package gin

import (
	_ "embed"
	"goboil/frameworks"
	"goboil/helpers"
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

type Gin struct {
	ModuleName string
}

func New(moduleName string) frameworks.Template {
	return &Gin{moduleName}
}

// Build builds project structure by initialize go module, making working directories, and writing files.
func (f *Gin) Build() error {
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
	go helpers.Loading("🔨 Creating project structure",
		"🚀 Project structure built successfully", done)

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

	// Writing files
	for file, content := range files {
		err = helpers.WriteFileFromTemplate(file, content, f)
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
