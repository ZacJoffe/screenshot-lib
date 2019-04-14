package main

import (
	"log"
	"os/exec"
)

// checkCommand looks sees is a given command exists on the user's system, if the returned error is nil then it exists
func checkCommand(command string) error {
	_, err := exec.LookPath(command)
	return err
	/*
		if err != nil {
			return err
		}

		return nil
	*/
}

func main() {
	err := checkCommand("import")
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("import", "screenshot.png")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
