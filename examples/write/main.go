package main

import (
	"context"
	"log"
	"os/exec"
	"time"

	"github.com/azurity/go-conpty"
)

func main() {
	cpty, err := conpty.Start(exec.Command("cmd"))
	if err != nil {
		log.Fatalf("Failed to spawn a pty:  %v", err)
	}
	defer cpty.Close()

	cpty.Write([]byte("@echo off\r\n"))
	cpty.Write([]byte("echo hello\r\n"))
	cpty.Write([]byte("whoami\r\n"))
	time.Sleep(time.Second * 1)
	out := make([]byte, 1000)
	n, err := cpty.Read(out)
	log.Printf("ReadCount: %d, Error: %v", n, err)
	log.Printf("Read: %s", string(out[:n]))
	cpty.Write([]byte("exit 1234\r\n"))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	exitCode, err := cpty.Wait(ctx)
	log.Printf("ExitCode: %d, Error: %v", exitCode, err)
}
