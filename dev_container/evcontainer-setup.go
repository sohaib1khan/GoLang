package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Embed the files from the `static/` directory
//go:embed static/docker-compose.yml
var dockerCompose []byte

//go:embed static/Dockerfile
var dockerfile []byte

func main() {
	fmt.Println("Welcome to the custom CLI environment setup!")
	fmt.Println("--------------------------------------------")

	// Prompt user for the environment name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a name for your environment (e.g., devbox): ")
	envName, _ := reader.ReadString('\n')
	envName = strings.TrimSpace(envName)

	if envName == "" {
		envName = "devbox" // Default name
	}

	fmt.Printf("Configuring environment: %s\n", envName)
	showLoadingAnimation("Setting up your environment", 3)

	// Get the directory where the script is running
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get current directory: %v\n", err)
		os.Exit(1)
	}

	// Create a persistent directory in the same directory as the script
	hostDir := filepath.Join(currentDir, fmt.Sprintf("%s-data", envName))
	err = os.MkdirAll(hostDir, 0755)
	if err != nil {
		fmt.Printf("Failed to create host directory: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Host directory created: %s\n", hostDir)

	// Write embedded Dockerfile and Compose files
	writeFile(currentDir, "Dockerfile", dockerfile)
	writeFile(currentDir, "docker-compose.yml", dockerCompose)

	// Build Docker image
	imageName := fmt.Sprintf("%s-image", envName)
	buildDockerImage(currentDir, imageName)

	// Run Docker container
	runDockerContainer(imageName, hostDir)
}

// showLoadingAnimation displays a loading animation for a specified duration
func showLoadingAnimation(message string, durationSeconds int) {
	loadingChars := []rune{'|', '/', '-', '\\'}
	fmt.Printf("%s", message)

	for i := 0; i < durationSeconds*10; i++ {
		fmt.Printf("\r%s %c", message, loadingChars[i%len(loadingChars)])
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("\rDone!                               ")
}

// writeFile writes the embedded content to a file
func writeFile(dir, filename string, content []byte) {
	path := filepath.Join(dir, filename)
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		fmt.Printf("Failed to write file %s: %v\n", filename, err)
		os.Exit(1)
	}
}

// buildDockerImage builds a Docker image from the provided directory
func buildDockerImage(buildDir, imageName string) {
	fmt.Println("Building Docker image...")
	cmd := exec.Command("docker", "build", "-t", imageName, buildDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to build Docker image: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Docker image built successfully.")
}

// runDockerContainer starts a Docker container with the given image and mounts the host directory
func runDockerContainer(imageName, hostDir string) {
	fmt.Println("Starting Docker container...")
	cmd := exec.Command("docker", "run", "-it", "--rm",
		"-v", fmt.Sprintf("%s:/data", hostDir), // Bind-mount host directory
		imageName,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to start Docker container: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Docker container exited.")
}
