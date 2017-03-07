package dfu

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

type DfuFlash struct {
	LeftFirmware  string
	RightFirmware string
	LeftComplete  bool
	RightComplete bool
	Error         error
}

func execFlash(firmwarePath string) ([]byte, error) {
	cmdName := "dfu-util"
	cmdArgs := []string{"-D", firmwarePath}
	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return []byte{}, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(cmdOut))
	r := regexp.MustCompile("status(0) = No error condition is present")
	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			return cmdOut, nil
		}
	}

	return cmdOut, fmt.Errorf("unknown error: completion message not found")
}

func (df *DfuFlash) FlashLeft() (string, error) {
	_, err := execFlash(df.LeftFirmware)
	if err != nil {
		return "", err
	}

	df.LeftComplete = true
	return "", nil
}

func (df *DfuFlash) FlashRight() (string, error) {
	_, err := execFlash(df.RightFirmware)
	if err != nil {
		return "", err
	}

	df.RightComplete = true
	return "", nil
}

func NewDfuFlash(leftFirmwarePath string, rightFirmwarePath string) (*DfuFlash, error) {
	df := &DfuFlash{
		LeftFirmware:  leftFirmwarePath,
		RightFirmware: rightFirmwarePath,
		LeftComplete:  false,
		RightComplete: false,
	}

	return df, nil
}
