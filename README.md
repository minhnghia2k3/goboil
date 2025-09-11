<img alt="goboil.svg" height="500" src="https://raw.githubusercontent.com/egonelbre/gophers/63b1f5a9f334f9e23735c6e09ac003479ffe5df5/vector/superhero/zorro.svg" width="500"/>

# 🚀 Goboil

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.22-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/license-Apache%202.0-green)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/minhnghia2k3/goboil)](https://goreportcard.com/report/github.com/minhnghia2k3/goboil)

A powerful CLI application that helps developers quickly generate project structure for Go web applications using popular frameworks like **Gin**, **Fiber**, and **gFly**. It automates the setup process, creating directories, configuration files, and boilerplate code, so you can start coding right away.

## ✨ Features

- 🚀 **Quick Setup**: Generate complete project structure in seconds
- 🎯 **Multiple Frameworks**: Support for Gin, Fiber, and gFly (coming soon)
- 📁 **Standard Structure**: Creates industry-standard project layout
- 🔧 **Auto Configuration**: Includes middleware, routes, and configuration setup
- 🎨 **Interactive CLI**: Beautiful and intuitive command-line interface
- ⚡ **Performance Optimized**: Concurrent API calls for fast framework information fetching

## 🎬 Demo

![demo](https://s11.gifyu.com/images/Soypn.gif)

## 📦 Installation

### Option 1: Go Install (Recommended)

Make sure you have Go 1.22+ installed, then run:

```bash
go install -v github.com/minhnghia2k3/goboil/cmd/goboil@v0.1.3
```

### Option 2: Download Binary

Visit the [releases page](https://github.com/minhnghia2k3/goboil/releases) and download the binary for your platform.

### Option 3: Build from Source

```bash
git clone https://github.com/minhnghia2k3/goboil.git
cd goboil
make gobuild
```

## 🚀 Quick Start

1. **Run goboil**:
   ```bash
   goboil
   ```

2. **Select a framework** from the interactive menu:
   - **Gin** - High-performance HTTP web framework
   - **Fiber** - Express-inspired web framework
   - **gFly** - Coming soon!

3. **Enter your module path**:
   ```
   Example: github.com/yourusername/my-awesome-api
   ```

4. **Start coding**! Your project structure is ready:
   ```bash
   cd your-project
   go run cmd/main.go
   ```

## 📂 Generated Project Structure

```
your-project/
├── cmd/
│   └── main.go              # Application entry point
├── config/
│   └── config.go            # Configuration management
├── controllers/
│   └── user_controllers.go  # Request handlers
├── middlewares/
│   └── auth_middleware.go   # Authentication middleware
├── models/
│   └── user.go              # Data models
├── routes/
│   └── routes.go            # Route definitions
├── .env                     # Environment variables
├── go.mod                   # Go module file
└── go.sum                   # Go module checksums
```

## 🛠️ Framework Support

| Framework | Status | Description |
|-----------|---------|-------------|
| [Gin](https://github.com/gin-gonic/gin) | ✅ Supported | High-performance HTTP web framework |
| [Fiber](https://github.com/gofiber/fiber) | ✅ Supported | Express-inspired web framework built on Fasthttp |
| [gFly](https://gfly.dev/) | 🔄 Coming Soon | Modern Go web framework |

## 💡 Usage Examples

### Creating a Gin Project

```bash
$ goboil
# Select Gin from the menu
# Enter module path: github.com/myusername/gin-api
# Project generated successfully!

$ cd gin-api
$ go run cmd/main.go
# Server starts on http://localhost:8080
```

### Creating a Fiber Project

```bash
$ goboil
# Select Fiber from the menu  
# Enter module path: github.com/myusername/fiber-api
# Project generated successfully!

$ cd fiber-api
$ go run cmd/main.go
# Server starts with Fiber configuration
```

## ⚙️ System Requirements

- **Go**: Version 1.22 or higher
- **OS**: Linux, macOS, Windows
- **Memory**: Minimum 512MB RAM
- **Disk**: 50MB free space

## 🔧 Development

### Prerequisites

- Go 1.22+
- Make (optional)

### Building from Source

```bash
# Clone the repository
git clone https://github.com/minhnghia2k3/goboil.git
cd goboil

# Install dependencies
go mod tidy

# Build the application
make gobuild
# or
go build -o build/goboil ./cmd/goboil

# Run tests
go test ./...

# Run linters
make golint
```

### Available Make Commands

- `make gobuild` - Build the application
- `make golint` - Run golangci-lint
- `make gosec` - Run security analysis
- `make gocritic` - Run code critic analysis
- `make gotidy` - Tidy Go modules

## 🤝 Contributing

We welcome contributions! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

### Development Process

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests and linters (`make golint && go test ./...`)
5. Commit your changes (`git commit -m 'Add some amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

### Code Style

- Follow standard Go formatting (`go fmt`)
- Add comments for exported functions
- Write tests for new functionality
- Keep functions small and focused

## 🐛 Troubleshooting

### Common Issues

**Issue: `command not found: goboil`**
- Solution: Make sure `$GOPATH/bin` is in your PATH, or run `go env GOPATH` to check your Go workspace

**Issue: `too many requests to GitHub API`**
- Solution: Wait a few minutes before retrying. The app fetches framework information from GitHub API

**Issue: `module path validation error`**
- Solution: Use a valid module path format like `github.com/username/project` or `example.com/myproject`

**Issue: Build fails with `go mod tidy`**
- Solution: Ensure you have a stable internet connection and Go modules proxy access

### Getting Help

- 📋 [Create an issue](https://github.com/minhnghia2k3/goboil/issues) for bugs
- 💬 [Start a discussion](https://github.com/minhnghia2k3/goboil/discussions) for questions
- 📧 Contact: [minhnghia2k3](https://github.com/minhnghia2k3)

## 📝 Changelog

See [RELEASES](https://github.com/minhnghia2k3/goboil/releases) for detailed changelog.

## 📄 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Thanks to the [Gin](https://github.com/gin-gonic/gin) and [Fiber](https://github.com/gofiber/fiber) communities
- Gopher artwork by [Egon Elbre](https://github.com/egonelbre/gophers)
- Built with ❤️ by [minhnghia2k3](https://github.com/minhnghia2k3)

---

⭐ If you find this project helpful, please give it a star!
