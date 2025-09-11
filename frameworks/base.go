package frameworks

import (
	"github.com/minhnghia2k3/goboil/helpers"
	"os/exec"
)

// BaseFramework provides common functionality for framework implementations
type BaseFramework struct {
	ModuleName string
	Name       string
}

// BuildProject builds the project structure with the given directories, files and templates
func (b *BaseFramework) BuildProject(directoryPaths []string, files map[string][]byte, loadingMsg, successMsg string, data interface{}) error {
	done := make(chan bool)
	
	go helpers.Loading(loadingMsg, successMsg, done)

	// Init go.mod
	if err := helpers.InitModule(b.ModuleName); err != nil {
		done <- true
		return err
	}

	// Create directories
	for _, path := range directoryPaths {
		if err := helpers.CreateDir(path); err != nil {
			done <- true
			return err
		}
	}

	// Create files from templates
	for path, content := range files {
		if err := helpers.WriteFileFromTemplate(path, content, data); err != nil {
			done <- true
			return err
		}
	}

	// Tidy module
	_ = exec.Command("go", "mod", "tidy").Run()

	done <- true
	return nil
}

// GetStandardDirectories returns the standard directory structure for web frameworks
func GetStandardDirectories() []string {
	return []string{
		"./cmd",
		"./config",
		"./controllers",
		"./middlewares",
		"./models",
		"./routes",
	}
}

// GetStandardFiles returns the standard file paths for web frameworks
func GetStandardFileMapping() map[string]string {
	return map[string]string{
		"./cmd/main.go":                     "main",
		"./config/config.go":                "config",
		"./controllers/user_controllers.go": "userController",
		"./middlewares/auth_middleware.go":  "auth",
		"./models/user.go":                  "user",
		"./routes/routes.go":                "routes",
		".env":                              "env",
	}
}