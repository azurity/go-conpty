package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/azurity/go-conpty"
)

func main() {
	cpty, err := conpty.Start(exec.Command("cmd"))
	if err != nil {
		log.Fatalf("Failed to spawn a pty:  %v", err)
	}
	defer cpty.Close()

	go func() {
		go io.Copy(os.Stdout, cpty)
		io.Copy(cpty, os.Stdin)
	}()

	exitCode, err := cpty.Wait(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("ExitCode: %d", exitCode)
}
