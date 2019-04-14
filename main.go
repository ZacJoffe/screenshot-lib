package main

import (
	"log"
	"os/exec"
)

const outputDir = "/tmp/screenshot.png"

// checkCommand looks sees is a given command exists on the user's system, if the returned error is nil then it exists
func checkCommand(command string) error {
	_, err := exec.LookPath(command)
	return err
}

func ScreenshotSelect() error {
	err := checkCommand("import")
	if err != nil {
		return err
	}

	cmd := exec.Command("import", outputDir)
	return cmd.Run()
}

func ScreenshotScreen() error {
	err := checkCommand("import")
	if err != nil {
		return err
	}

	cmd := exec.Command("import", "-window", "root", outputDir)
	return cmd.Run()
}

func main() {
	/*
		err := checkCommand("import")
		if err != nil {
			log.Fatal(err)
		}

		cmd := exec.Command("import", outputDir)
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	*/
	err := ScreenshotScreen()
	if err != nil {
		log.Fatal(err)
	}
}
