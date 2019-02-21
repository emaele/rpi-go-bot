package commands

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

var (
	err error
)

//AvailableSpace returns free space of the sdcard in GB
func AvailableSpace() (string, error) {
	cmd := exec.Command("df", "--output=avail", "/")
	output, err := getOut(cmd)
	if err != nil {
		return "", err
	}

	msgSplit := strings.Split(output, "\n")
	if value, err := strconv.Atoi(msgSplit[1]); err == nil {
		return fmt.Sprintf("Available space %d GB 💾", value/1000000), nil
	}

	return "", err
}

func getOut(command *exec.Cmd) (output string, err error) {
	stdoutStderr, err := command.CombinedOutput()
	if err != nil {
		err = errors.New("Error")
	}

	output = string(stdoutStderr)

	return output, err
}
