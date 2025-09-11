package gin

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

type Gin struct {
	*frameworks.BaseFramework
}

func New(moduleName string) frameworks.Template {
	return &Gin{
		BaseFramework: &frameworks.BaseFramework{
			ModuleName: moduleName,
			Name:       "Gin",
		},
	}
}

// Build builds project structure by initialize go module, making working directories, and writing files.
func (g *Gin) Build() error {
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

	return g.BuildProject(
		directoryPaths,
		files,
		"🔨 Creating Gin project structure",
		"🚀 Gin project structure built successfully!",
		g,
	)
}
