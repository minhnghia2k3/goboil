package main

import (
	"goboil/cmd"
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
