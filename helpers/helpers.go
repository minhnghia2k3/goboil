package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/manifoldco/promptui"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"time"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

const (
	Gin   = "gin"
	Fiber = "fiber"
	Gfly  = "gFly"
)

type framework struct {
	Name        string `json:"name"`
	Stars       int    `json:"watchers_count"`
	Description string `json:"description"`
}

type repository struct {
	Owner string
	Repo  string
}

var cls map[string]func()

func Greeting() {
	fmt.Println(Cyan + `
.-------------------------------------------------------------.
|                                                             |
|  ________  ________  ________  ________  ___  ___           |
| |\   ____\|\   __  \|\   __  \|\   __  \|\  \|\  \          |
| \ \  \___|\ \  \|\  \ \  \|\ /\ \  \|\  \ \  \ \  \         |
|  \ \  \  __\ \  \\\  \ \   __  \ \  \\\  \ \  \ \  \        |
|   \ \  \|\  \ \  \\\  \ \  \|\  \ \  \\\  \ \  \ \  \____   |
|    \ \_______\ \_______\ \_______\ \_______\ \__\ \_______\ |
|     \|_______|\|_______|\|_______|\|_______|\|__|\|_______| |
|                                                             |
'-------------------------------------------------------------'` + Reset)
}

// ClearTerminal clear the terminal screen based on os.
func ClearTerminal() error {
	cls = make(map[string]func())
	cls["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	cls["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	cls["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}

	value, ok := cls[runtime.GOOS] // runtime.GOOS -> linux, windows, darwin etc.
	if ok {
		value() // Execute cls function
	} else {
		return fmt.Errorf("your platform is unsupported! I can't clear terminal screen :(")
	}
	return nil
}

// SelectTemplates allows client select a specific framework then returns the framework's index.
func SelectTemplates() (int, error) {
	// Fetch frameworks info
	frameworks, err := FetchInfo()
	if err != nil {
		return 0, err
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "➙ {{ .Name | magenta | italic }}",
		Inactive: "  {{ .Name | white }}",
		Selected: "🚀 You've choose: {{ .Name | magenta | italic }}",
		Details: `--------- Framework Details ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Stars:" | faint }}	{{ .Stars | yellow }}
{{ "Description:" | faint }}	{{ .Description }}`,
	}

	// Search function
	searcher := func(input string, index int) bool {
		fw := frameworks[index]
		name := strings.ReplaceAll(strings.ToLower(fw.Name), " ", "")
		input = strings.ReplaceAll(strings.ToLower(input), " ", "")

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Select a web framework:",
		Items:     frameworks,
		Templates: templates,
		Size:      5,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return 0, err
	}

	return i, nil
}

// PromptModulePath allows client type in module path.
func PromptModulePath() (string, error) {
	validate := func(input string) error {
		if input == "" {
			return fmt.Errorf("please enter a valid module path")
		}

		if strings.Contains(input, " ") {
			return fmt.Errorf("module path should not contain spaces")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . }} ",
	}

	prompt := promptui.Prompt{
		Label:     "Your module path (e.g. example.com/my-project)",
		Validate:  validate,
		Templates: templates,
	}

	module, err := prompt.Run()

	if err != nil {
		return "", fmt.Errorf("Prompt failed %v\n", err)
	}

	return module, nil
}

// CreateDir creates directory named path
func CreateDir(path string) error {
	if err := os.MkdirAll(path, 0o750); err != nil {
		return fmt.Errorf("error creating directory %s: %v", path, err)
	}
	return nil
}

// CreateFile creates file to path, executes provided data into template file
func CreateFile(path string, tmpl *template.Template, data interface{}) error {
	f, err := os.Create(filepath.Clean(path))
	if err != nil {
		return fmt.Errorf("error creating file %v\n", err)
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		return fmt.Errorf("error executing %v\n", err)
	}
	return nil
}

// Loading displays a loading effect until the done channel is closed.
func Loading(msg, finish string, done chan bool) {
	fmt.Print(msg + " ")

	// Spinner characters
	spinner := []rune{'|', '/', '-', '\\'}

	i := 0
	for {
		select {
		case <-done:
			fmt.Println("\n" + finish)
		default:
			fmt.Printf("\b%s", string(spinner[i%len(spinner)]))
			i++
			time.Sleep(200 * time.Millisecond) // Adjust the speed as needed
		}
	}
}

// InitModule executes `go mod init` to init module by provided moduleName.
func InitModule(moduleName string) error {
	var (
		out    bytes.Buffer
		stderr bytes.Buffer
	)

	command := exec.Command("go", "mod", "init", moduleName)
	command.Stdout = &out
	command.Stderr = &stderr
	err := command.Run()
	if err != nil {
		return fmt.Errorf("Error initialize go modules: %v\n", command.Stderr)
	}
	return nil
}

// WriteFileFromTemplate parses template content for given file path
func WriteFileFromTemplate(path string, tmplContent []byte, data interface{}) error {
	tmpl, err := template.New(filepath.Base(path)).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("error passing template %v\n", err)
	}
	return CreateFile(path, tmpl, data)
}

func FetchInfo() ([]framework, error) {
	// curl -L   -H "Accept: application/vnd.github+json"
	// https://api.github.com/repos/minhnghia2k3/green-light | jq '. | {stars: .watchers_count, description: .description}'
	/*
		[
			gin: {
				stars: int
				description: string
			}
		]
	*/
	var result []framework
	frameworks := []repository{
		{
			Owner: "gin-gonic",
			Repo:  Gin,
		},
		{
			Owner: "gofiber",
			Repo:  Fiber,
		},
		{
			Owner: "JiveIO",
			Repo:  Gfly,
		},
	}

	for _, fr := range frameworks {
		res, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s", fr.Owner, fr.Repo))
		if err != nil {
			return nil, fmt.Errorf("Error fetching repo %v\n", err)
		}
		if res.StatusCode != http.StatusOK {
			switch res.StatusCode {
			case http.StatusForbidden:
				return nil, fmt.Errorf("too many request to this resource, try again later")
			case http.StatusNotFound:
				return nil, fmt.Errorf("resource not found")
			default:
				return nil, fmt.Errorf("unexpected status code %v", res.StatusCode)
			}
		}

		// Read response body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("Error fetching framework %s: %v\n", fr.Repo, err)
		}

		// Unmarshal Golang struct
		var info framework
		err = json.Unmarshal(body, &info)
		if err != nil {
			return nil, fmt.Errorf("Error fetching framework %s: %v\n", fr.Repo, err)
		}

		result = append(result, info)
		_ = res.Body.Close()
	}

	return result, nil
}
