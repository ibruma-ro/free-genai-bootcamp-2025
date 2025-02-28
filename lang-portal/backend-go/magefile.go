package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Build compiles the application
func Build() error {
	fmt.Println("Building the application...")
	cmd := exec.Command("go", "build", "-o", "myapp", "./cmd/server")
	return cmd.Run()
}

// Run starts the application
func Run() error {
	fmt.Println("Running the application...")
	cmd := exec.Command("./myapp")
	return cmd.Run()
}

// RunMigrations executes database migrations
func RunMigrations() error {
	fmt.Println("Running database migrations...")
	cmd := exec.Command("go", "run", "./cmd/seeder.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// InitDB initializes the database with seed data
func InitDB() error {
	fmt.Println("Initializing database with seed data...")
	cmd := exec.Command("go", "run", "./cmd/seeder.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
