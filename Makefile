gobuild: gotidy gosec gocritic golint
	go build -o "./build/goboil" ./cmd/goboil

gosec:
	gosec ./...

gocritic:
	gocritic check -enableAll ./...

golint:
	golangci-lint run ./...

gotidy:
	go mod tidy
PHONY: gobuild gosec gocritic golint


