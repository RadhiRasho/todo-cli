package build

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const binaryName = "todo"

var platformMap = map[string][]string{
	"darwin":   {"darwin", "amd64", "/usr/local/bin"},
	"linux":    {"linux", "amd64", "/usr/local/bin"},
	"windows":  {"windows", "amd64", filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "Programs", binaryName)},
}

func build() {
	platform, _, installDir := getPlatformInfo()
	if platform == "" {
		log.Fatal("Unsupported platform")
	}

	binName := strings.Join([]string{binaryName, getExecutableExtension(platform)}, "")

	buildCmd := exec.Command("go", "build", "todo.go", "-o", binName)
	output, err := buildCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Build failed: %s", output)
	}

	err = os.MkdirAll(installDir, 0755)
	if err != nil {
		log.Fatalf("Failed to create install directory: %v", err)
	}

	err = os.Rename(binName, installDir)
	if err != nil {
		log.Fatalf("Failed to move binary to install directory: %v", err)
	}

	fmt.Printf("Successfully installed %s to %s\n", binaryName, installDir)
}



func getPlatformInfo() (string, string, string) {
	platform := runtime.GOOS
	if info, ok := platformMap[platform]; ok {
		return info[0], info[1], info[2]
	}
	return "", "", ""
}

func getExecutableExtension(platform string) string {
	if platform == "windows" {
		return ".exe"
	}
	return ""
}