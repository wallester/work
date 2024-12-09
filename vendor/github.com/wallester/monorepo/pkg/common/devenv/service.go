package devenv

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/juju/errors"
	"github.com/wallester/monorepo/pkg/common/devenv/servicestate"
)

// Service represents a service in the monorepo; e.g. "api" and "account-service"
type Service struct {
	// ShortName is the name of the service used in the monorepo; e.g. "ams"
	ShortName string `yaml:"name"`
	// Directory is the directory (and executable!) name of the service in the monorepo; e.g. "account-service"
	Directory string `yaml:"directory"`
	// SkipManagement says that this service should not be managed by mu.
	SkipManagement bool `yaml:"skip"`
	// UsesInternalPackages says that this service uses internal packages.
	UsesInternalPackages bool `yaml:"internal"`
	// PCI says that this service is PCI-compliant.
	PCI bool `yaml:"pci"`
}

func (service *Service) String() string {
	return fmt.Sprintf("%s (%s)", service.ShortName, service.Directory)
}

// LogFile returns the path to the log file of the service.
func (service *Service) LogFile() string {
	return filepath.Join(service.FullPath(), fmt.Sprintf("%s.log", service.ShortName))
}

// FullPath returns the full path to the service directory.
func (service *Service) FullPath() string {
	return filepath.Join(MonorepoPath(), service.Directory)
}

func (service *Service) InternalPath() string {
	if service.UsesInternalPackages {
		return "internal"
	}

	return ""
}

func (service *Service) ConfigurationPath() string {
	return filepath.Join(MonorepoPath(), service.Directory, service.InternalPath(), "configuration")
}

// ExecutablePath returns the full path to the service executable.
func (service *Service) ExecutablePath() string {
	return filepath.Join(service.FullPath(), service.Directory)
}

// CmdExecutablePath returns monorepo/<service>/cmd/<service>
func (service *Service) CmdExecutablePath() string {
	return filepath.Join(service.FullPath(), "cmd", service.Directory)
}

func (service *Service) HasCmdExecutable() bool {
	return directoryExists(service.CmdExecutablePath())
}

// Start starts the service.
// If the service is already running, it does nothing.
// If the service is not running, it builds the service and runs it.
// cover is a flag that enables code coverage.
func (service *Service) Start(cover bool) (servicestate.ServiceState, error) {
	pid, err := service.PID()
	if err != nil {
		return "", errors.Annotate(err, "checking service state failed")
	}

	if pid != "" {
		return servicestate.Running, nil
	}

	if err := service.Build(cover); err != nil {
		return "", errors.Annotate(err, "building service failed")
	}

	if err := service.RunExecutable(cover); err != nil {
		return "", errors.Annotate(err, "running service failed")
	}

	pid, err = service.PID()
	if err != nil {
		return "", errors.Annotate(err, "checking service state failed")
	}

	if pid != "" {
		return servicestate.Started, nil
	}

	return servicestate.Unknown, nil
}

// Build builds the service.
// cover is a flag that enables code coverage.
func (service *Service) Build(cover bool) error {
	buildArgs := []string{"build"}

	if service.HasCmdExecutable() {
		// go build -C cmd/auth-service -o ../../
		buildArgs = append(buildArgs, "-C", filepath.Join("cmd", service.Directory), "-o", filepath.Join("..", ".."))
	}

	if cover {
		buildArgs = append(buildArgs, "-cover", "-covermode=atomic")
	}

	if !service.HasCmdExecutable() {
		// go build .
		buildArgs = append(buildArgs, ".")
	}

	cmd := exec.Command("go", buildArgs...)
	cmd.Dir = service.FullPath()
	res, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Annotatef(err, "building service failed: %s", string(res))
	}

	codesign, _ := strconv.ParseBool(os.Getenv("CODESIGN"))
	if runtime.GOOS == "darwin" && codesign {
		if err := service.CodeSign(); err != nil {
			return errors.Annotate(err, "codesigning service failed")
		}
	}

	return nil
}

// CodeSign signs the service with a developer certificate.
// This is only supported on macOS.
// It is not mandatory to sign the service.
func (service *Service) CodeSign() error {
	cmd := exec.Command("codesign", "-s", "Developer Certificate", service.Directory) //#nosec:G204
	cmd.Dir = service.FullPath()
	res, err := cmd.CombinedOutput()
	if err != nil {
		if !strings.Contains(string(res), "is already signed") {
			return errors.Annotatef(err, "signing service failed: %s", string(res))
		}
	}

	return nil
}

// Stop stops the service.
// If the service is not running, it does nothing.
func (service *Service) Stop() (servicestate.ServiceState, error) {
	pid, err := service.PID()
	if err != nil {
		return "", errors.Annotate(err, "checking service state failed")
	}

	if pid == "" {
		return servicestate.NotRunning, nil
	}

	killCmd := exec.Command("kill", pid) //#nosec:G204
	if err := killCmd.Run(); err != nil {
		return "", errors.Annotate(err, "stopping service failed")
	}

	pid, err = service.PID()
	if err != nil {
		return "", errors.Annotate(err, "checking service state failed")
	}

	if pid == "" {
		return servicestate.Stopped, nil
	}

	return servicestate.Unknown, nil
}

// PID returns the PID of the service.
func (service *Service) PID() (string, error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pgrep", "-x", service.Directory) //#nosec:G204
	case "linux":
		line := "ps -eo pid,cmd | awk '$2 == \"" + service.ExecutablePath() + "\" {print $1}'"
		cmd = exec.Command("bash", "-c", line) //#nosec:G204
	default:
		return "", errors.New("Unknown OS version")
	}

	output, err := cmd.Output()
	if err != nil {
		// Exit status 1 means that the process is not running.
		if IsExitStatus(err, 1) {
			return "", nil
		}

		return "", errors.Annotatef(err, "getting PID failed: %s", string(output))
	}

	return strings.TrimSpace(string(output)), nil
}

// RunExecutable runs the service executable.
// cover is a flag that enables code coverage.
func (service *Service) RunExecutable(cover bool) error {
	runCmd := filepath.Join(service.FullPath(), service.Directory)
	cmd := exec.Command(runCmd) //#nosec:G204

	// Optionally, enable code coverage.
	if cover {
		goCoverDir := IntegrationCoveragePath()
		if err := os.MkdirAll(goCoverDir, os.ModePerm); err != nil {
			return errors.Annotatef(err, "creating coverage directory %s failed", goCoverDir)
		}

		cmd.Env = append(os.Environ(), "GOCOVERDIR="+goCoverDir)
	}

	cmd.Dir = service.FullPath()
	if err := cmd.Start(); err != nil {
		return errors.Annotate(err, "starting service failed")
	}

	return nil
}

// private

func directoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return err == nil && info.IsDir()
}
