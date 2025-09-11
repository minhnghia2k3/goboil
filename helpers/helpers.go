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
	"sync"
	"text/template"
	"time"
)

// ANSI Color codes
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

// Framework names
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
		input = strings.TrimSpace(input)
		if input == "" {
			return fmt.Errorf("module path cannot be empty")
		}

		if strings.Contains(input, " ") {
			return fmt.Errorf("module path cannot contain spaces")
		}

		// Check for invalid characters
		invalidChars := []string{"\\", "\"", "'", "<", ">", "|", "?", "*"}
		for _, char := range invalidChars {
			if strings.Contains(input, char) {
				return fmt.Errorf("module path cannot contain '%s'", char)
			}
		}

		// Should have at least one dot for domain
		if !strings.Contains(input, ".") && !strings.HasPrefix(input, "github.com/") {
			return fmt.Errorf("module path should be a valid domain/path (e.g., example.com/my-project)")
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
		Label:     "Your module path (e.g. github.com/username/my-project)",
		Validate:  validate,
		Templates: templates,
	}

	module, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("failed to get module path: %w", err)
	}

	return strings.TrimSpace(module), nil
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
			return
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
	frameworks := []repository{
		{Owner: "gin-gonic", Repo: Gin},
		{Owner: "gofiber", Repo: Fiber},
		{Owner: "JiveIO", Repo: Gfly},
	}

	result := make([]framework, len(frameworks))
	var wg sync.WaitGroup
	errChan := make(chan error, len(frameworks))

	for i, fr := range frameworks {
		wg.Add(1)
		go func(index int, repo repository) {
			defer wg.Done()
			
			res, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s", repo.Owner, repo.Repo))
			if err != nil {
				errChan <- fmt.Errorf("error fetching repo %s/%s: %w", repo.Owner, repo.Repo, err)
				return
			}
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				switch res.StatusCode {
				case http.StatusForbidden:
					errChan <- fmt.Errorf("too many requests to GitHub API, try again later")
				case http.StatusNotFound:
					errChan <- fmt.Errorf("repository %s/%s not found", repo.Owner, repo.Repo)
				default:
					errChan <- fmt.Errorf("unexpected status code %d for %s/%s", res.StatusCode, repo.Owner, repo.Repo)
				}
				return
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				errChan <- fmt.Errorf("error reading response body for %s: %w", repo.Repo, err)
				return
			}

			var info framework
			if err := json.Unmarshal(body, &info); err != nil {
				errChan <- fmt.Errorf("error unmarshaling JSON for %s: %w", repo.Repo, err)
				return
			}

			result[index] = info
		}(i, fr)
	}

	wg.Wait()
	close(errChan)

	// Check if any errors occurred
	if err := <-errChan; err != nil {
		return nil, err
	}

	return result, nil
}
