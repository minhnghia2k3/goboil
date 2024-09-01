package cmd

import (
	"github.com/spf13/cobra"
	"goboil/frameworks"
	"goboil/frameworks/fiber"
	"goboil/frameworks/gfly"
	"goboil/frameworks/gin"
	"goboil/helpers"
	"log"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goboil",
	Short: "A boilerplate support building application using famous web frameworks in Go.",
	Run: func(cmd *cobra.Command, args []string) {
		// Send a greeting
		helpers.Greeting()

		// Select a framework template
		index, err := helpers.SelectTemplates()
		if err != nil {
			log.Fatal("Error selecting templates: ", err)
		}

		// Prompt module name
		module, err := helpers.PromptModulePath()
		if err != nil {
			log.Fatal("Error selecting module name: ", err)
		}

		// Switch case build project structure for each template
		var template frameworks.Template
		switch index {
		case 0:
			template = gin.New(module)
		case 1:
			template = fiber.New(module)
		case 2:
			template = gfly.New(module)
		default:
			panic("Something gone wrong!")
		}

		// Build the project structure
		if err := template.Build(); err != nil {
			log.Fatal("Error building template: ", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	err := helpers.ClearTerminal()
	if err != nil {
		panic(err)
	}
}
