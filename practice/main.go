package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func killServer(pidfile string) error {
	file, err := os.Open(pidfile)
	if err != nil {
		return err
	}

	defer file.Close()

	var pid int
	if _,err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return errors.
	}

	fmt.Printf("Killing server with pid=%d\n", pid)
	return nil
}


func main() {}