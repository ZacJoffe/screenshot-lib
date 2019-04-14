package screenshot

import (
	"os/exec"
)

// dir of saved image
const outputDir = "/tmp/screenshot.png"

// checkCommand looks sees is a given command exists on the user's system, if the returned error is nil then it exists
func checkCommand(command string) error {
	_, err := exec.LookPath(command)
	return err
}

// Select creates an image file with the selected region using the "import" command
func Select() error {
	// check that import command exists
	err := checkCommand("import")
	if err != nil {
		return err
	}

	// create the command "import /tmp/screenshot.png" to save a selection screenshot to that dir
	cmd := exec.Command("import", outputDir)

	// run the command, will return nil if successful
	return cmd.Run()
}

// Screen creates an image file with of the whole screen using the "import -window root" command
func Screen() error {
	// check that import command exists
	err := checkCommand("import")
	if err != nil {
		return err
	}

	// create the command "import -window root /tmp/screenshot.png" to save screenshot of whole screen to that dir
	cmd := exec.Command("import", "-window", "root", outputDir)

	// run the command, will return nil if successful
	return cmd.Run()
}
