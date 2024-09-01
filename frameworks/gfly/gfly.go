package gfly

import (
	"fmt"
	"goboil/frameworks"
)

type Gfly struct {
	module string
}

func New(moduleName string) frameworks.Template {
	return &Gfly{moduleName}
}

func (f *Gfly) Build() error {
	fmt.Println("Gfly is coming soon, for more detail access: https://gfly.dev/")
	return nil
}
