package commands

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

const (
	b  = 1
	kb = 1024 * b
	mb = 1024 * kb
	gb = 1024 * mb
)

//AvailableSpace returns free space of the sdcard in GB
func AvailableSpace() (string, error) {

	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/", &fs)
	if err != nil {
		return "", err
	}

	free := fs.Bavail * uint64(fs.Bsize)
	return fmt.Sprintf("Free space: %.2f GB 💾", float64(free)/float64(gb)), nil
}

// GetTemp gets the actual temperature of your rpi's CPU
func GetTemp() (temp string) {

	cmd := exec.Command("/opt/vc/bin/vcgencmd", "measure_temp")
	if stdoutStderr, err := cmd.CombinedOutput(); err == nil {
		log := string(stdoutStderr)
		temp = strings.Trim(log, "temp='C\n")
	}

	return
}
