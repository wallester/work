package env

import (
	"os"
	"path/filepath"

	"github.com/juju/errors"
	"github.com/wallester/godotenv"
	"github.com/wallester/monorepo/pkg/common/devenv"
)

func LoadDotenv() error {
	if err := loadDotenv(""); err != nil {
		return errors.Annotate(err, "loading dotenv failed")
	}

	return nil
}

func LoadDotenvByPath(path string) error {
	if err := loadDotenv(path); err != nil {
		return errors.Annotate(err, "loading dotenv by path failed")
	}

	return nil
}

// private

func loadDotenv(path string) error {
	globalEnvPath := filepath.Join(path, ".env.global")

	if _, err := os.Stat(globalEnvPath); err != nil {
		globalEnvPath = getGlobalEnvPath(path)
	}

	if _, err := os.Stat(globalEnvPath); err == nil {
		if err := godotenv.Load(globalEnvPath); err != nil {
			return errors.Annotate(err, "loading .env.global file failed")
		}
	}

	environment := os.Getenv("GLOBAL_ENVIRONMENT")
	if environment == "" {
		return errors.New("loading GLOBAL_ENVIRONMENT failed")
	}

	envPath := filepath.Join(path, ".env")
	if _, err := os.Stat(envPath); err == nil {
		if err := godotenv.Load(envPath); err != nil {
			return errors.Annotate(err, "loading .env file failed")
		}
	}

	if err := godotenv.Load(filepath.Join(path, ".env."+environment), filepath.Join(path, ".env.shared")); err != nil {
		return errors.Annotate(err, "loading .env files failed")
	}

	return nil
}

func getGlobalEnvPath(path string) string {
	switch path {
	case emptyPath, singleDotPath:
		return filepath.Join("..", "pkg", "common", ".env.global")
	case doubleDotPath:
		return filepath.Join("..", "..", "pkg", "common", ".env.global")
	case doubleDoubleDotPath:
		return filepath.Join("..", "..", "..", "pkg", "common", ".env.global")
	case tripleDoubleDotPath:
		return filepath.Join("..", "..", "..", "..", "pkg", "common", ".env.global")
	default:
		return filepath.Join(devenv.MonorepoPath(), "pkg", "common", ".env.global")
	}
}

const (
	emptyPath           = ""
	singleDotPath       = "."
	doubleDotPath       = ".."
	doubleDoubleDotPath = "../.."
	tripleDoubleDotPath = "../../.."
)
