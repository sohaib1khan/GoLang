package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io/ioutil"
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

//go:embed static/list-tools
var listTools []byte

//go:embed static/vimrc
var vimrc []byte

//go:embed static/welcome.sh
var welcomeScript []byte

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

	// Create a temporary working directory
	tempDir, err := ioutil.TempDir("", fmt.Sprintf("%s-", envName))
	if err != nil {
		fmt.Printf("Failed to create temporary directory: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tempDir)

	fmt.Printf("Working directory: %s\n", tempDir)

	// Write embedded files to the temporary directory
	writeFile(tempDir, "docker-compose.yml", dockerCompose)
	writeFile(tempDir, "Dockerfile", dockerfile)
	writeFile(tempDir, "list-tools", listTools)
	writeFile(tempDir, "vimrc", vimrc)
	writeFile(tempDir, "welcome.sh", welcomeScript)

	// Make the scripts executable
	makeExecutable(filepath.Join(tempDir, "list-tools"))
	makeExecutable(filepath.Join(tempDir, "welcome.sh"))

	// Launch the custom CLI with the host directory mounted
	runCustomCLI(tempDir, hostDir, envName)
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
	err := ioutil.WriteFile(path, content, 0644)
	if err != nil {
		fmt.Printf("Failed to write file %s: %v\n", filename, err)
		os.Exit(1)
	}
}

// makeExecutable sets the executable bit on a file
func makeExecutable(path string) {
	err := os.Chmod(path, 0755)
	if err != nil {
		fmt.Printf("Failed to make %s executable: %v\n", path, err)
		os.Exit(1)
	}
}

// runCustomCLI launches the custom CLI environment
func runCustomCLI(workDir, hostDir, envName string) {
	fmt.Println("Launching the custom CLI environment...")

	// Change to the working directory
	err := os.Chdir(workDir)
	if err != nil {
		fmt.Printf("Failed to change directory: %v\n", err)
		os.Exit(1)
	}

	// Execute the welcome script with the mounted directory
	cmd := exec.Command("/bin/bash", "./welcome.sh")
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("ENV_NAME=%s", envName),
		fmt.Sprintf("HOST_DIR=%s", hostDir),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to launch CLI: %v\n", err)
		os.Exit(1)
	}
}
