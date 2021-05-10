// +build mage

package main

import (
	"fmt"
	// "os"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	// "github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	// cmd := exec.Command("go", "build", "-o", "MyApp", ".")
	cmd := exec.Command("go", "build", ".")
	return cmd.Run()
}

// Start Docker Container with Postgres DB.  This command is used for local
// development only
func RunDocker() error {
	fmt.Println("Mage starting postgres container...")
	cmd := exec.Command("docker", "compose", "up", "-d")
	return cmd.Run()
}

// Stop Running Postgres Docker Container.  under the hood running docker
// compose down
func StopDocker() error {
	fmt.Println("Mage stoping postgres container...")
	cmd := exec.Command("docker", "compose", "down")
	return cmd.Run()
}

// Start Dev Server. This command is used for local
// development and will start up your dev server... if docker postgres db is not
// started it will start that first and wait for your db to be up and running.
func StartAPI() error {
	mg.Deps(RunDocker)
	fmt.Println("Mage starting go_post_api api...")
	cmd := exec.Command("go", "run", ".")
	return cmd.Run()
}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	cmd := exec.Command("go", "install")
	return cmd.Run()
	// return os.Rename("./MyApp", "/usr/bin/MyApp")
	// return os.Rename("./MyApp", "/usr/bin/MyApp")
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Running go mod tidy...")
	// cmd := exec.Command("go", "get", "github.com/stretchr/piglatin")
	cmd := exec.Command("go", "mod", "tidy")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	// os.RemoveAll("MyApp")
}
