# Command to create symlink for subpackages: ln -s ../Makefile Makefile

export GOFLAGS=-mod=vendor
export GOPRIVATE=github.com/wallester

install-go:
	@go install ./...

install: install-go

mod-vendor:
	go mod tidy && go mod vendor