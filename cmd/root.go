package cmd

import (
	"fmt"
	"github.com/minhnghia2k3/goboil/frameworks"
	"github.com/minhnghia2k3/goboil/frameworks/fiber"
	"github.com/minhnghia2k3/goboil/frameworks/gfly"
	"github.com/minhnghia2k3/goboil/frameworks/gin"
	"github.com/minhnghia2k3/goboil/helpers"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goboil",
	Short: "A boilerplate support building application using famous web frameworks in Go.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Send a greeting
		helpers.Greeting()

		// Select a framework template
		index, err := helpers.SelectTemplates()
		if err != nil {
			return fmt.Errorf("error selecting templates: %w", err)
		}

		// Prompt module name
		module, err := helpers.PromptModulePath()
		if err != nil {
			return fmt.Errorf("error getting module name: %w", err)
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
			return fmt.Errorf("invalid template index: %d", index)
		}

		// Build the project structure
		if err := template.Build(); err != nil {
			return fmt.Errorf("error building template: %w", err)
		}

		return nil
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
	if err := helpers.ClearTerminal(); err != nil {
		// If terminal clearing fails, continue without panicking
		fmt.Fprintf(os.Stderr, "Warning: could not clear terminal: %v\n", err)
	}
}
