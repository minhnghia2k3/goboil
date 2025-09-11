package fiber

import (
	_ "embed"
	"github.com/minhnghia2k3/goboil/frameworks"
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
	*frameworks.BaseFramework
}

func New(moduleName string) frameworks.Template {
	return &Fiber{
		BaseFramework: &frameworks.BaseFramework{
			ModuleName: moduleName,
			Name:       "Fiber",
		},
	}
}

// Build builds project structure by initialize go module, making working directories, and writing files.
func (f *Fiber) Build() error {
	directoryPaths := frameworks.GetStandardDirectories()

	files := map[string][]byte{
		"./cmd/main.go":                     main,
		"./config/config.go":                config,
		"./controllers/user_controllers.go": userController,
		"./middlewares/auth_middleware.go":  auth,
		"./models/user.go":                  user,
		"./routes/routes.go":                routes,
		".env":                              env,
	}

	return f.BuildProject(
		directoryPaths,
		files,
		"🔨 Creating Fiber project structure",
		`🚀 Fiber project structure built successfully!
🚀 To start the application run: $ go run cmd/main.go`,
		f,
	)
}
