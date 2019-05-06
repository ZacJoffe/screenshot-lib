package screenshot

import (
	"os"
	"os/exec"
)

// dir of saved image
const outputDir = "/tmp/screenshot.png"

// checkCommand looks sees is a given command exists on the user's system, if the returned error is nil then it exists
func checkCommand(command string) error {
	_, err := exec.LookPath(command)
	return err
}

// Select creates an image file with the selected region using the "import" command, then returns a pointer to that file
func Select() (*os.File, error) {
	// check that import command exists
	err := checkCommand("maim")
	if err != nil {
		return nil, err
	}

	// create the command "import /tmp/screenshot.png" to save a selection screenshot to that dir
	cmd := exec.Command("maim", "-s", outputDir)

	// run the command, check if successful
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	// open the new image file, defer it's closing
	file, err := os.Open(outputDir)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// return the file
	return file, nil
}

// Screen creates an image file with of the whole screen using the "import -window root" command, then returns a pointer to that file
func Screen() (*os.File, error) {
	// check that import command exists
	err := checkCommand("maim")
	if err != nil {
		return nil, err
	}

	// create the command "import -window root /tmp/screenshot.png" to save screenshot of whole screen to that dir
	cmd := exec.Command("maim", outputDir)

	// run the command, check if successful
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	// open the new image file, defer it's closing
	file, err := os.Open(outputDir)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// return the file
	return file, nil
}
