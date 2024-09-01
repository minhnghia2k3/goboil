package fiber

import (
	_ "embed"
	"github.com/minhnghia2k3/goboil/frameworks"
	"github.com/minhnghia2k3/goboil/helpers"
	"os/exec"
)

//go:embed templates/main.tmpl
var main []byte

//go:embed templates/config.tmpl
var config []byte

//go:embed templates/auth_middleware.tmpl
var auth []byte

//go:embed templates/routes.tmpl
var routes []byte

//go:embed templates/user.tmpl
var user []byte

//go:embed templates/user_controller.tmpl
var userController []byte

//go:embed templates/.env.tmpl
var env []byte

type Fiber struct {
	ModuleName string
}

func New(moduleName string) frameworks.Template {
	return &Fiber{moduleName}
}

// Build builds project structure by initialize go module, making working directories, and writing files.
func (f *Fiber) Build() error {
	directoryPaths := []string{
		"./cmd",
		"./config",
		"./controllers",
		"./middlewares",
		"./models",
		"./routes",
	}

	files := map[string][]byte{
		"./cmd/main.go":                     main,
		"./config/config.go":                config,
		"./controllers/user_controllers.go": userController,
		"./middlewares/auth_middleware.go":  auth,
		"./models/user.go":                  user,
		"./routes/routes.go":                routes,
		".env":                              env,
	}

	done := make(chan bool)

	go helpers.Loading("Creating Fiber project structure",
		`🚀 Project structure built successfully!
🚀 To start the application run: $ go run cmd/main.go`, done)

	// Init go.mod
	err := helpers.InitModule(f.ModuleName)
	if err != nil {
		done <- true
		return err
	}

	// Creating directories
	for _, path := range directoryPaths {
		err = helpers.CreateDir(path)
		if err != nil {
			done <- true
			return err
		}
	}

	// Creating files
	for path, content := range files {
		err = helpers.WriteFileFromTemplate(path, content, f)
		if err != nil {
			done <- true
			return err
		}
	}

	// Tidying module
	_ = exec.Command("go", "mod", "tidy").Run()

	done <- true
	return nil
}
