package main

import (
	"github.com/minhnghia2k3/goboil/cmd"
	"log"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Panic occurred:", err)
		}
	}()
	cmd.Execute()
}
