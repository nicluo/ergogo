package dfu

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

func Scan() (bool, error) {
	dfuCount := 0
	cmdName := "dfu-util"
	cmdArgs := []string{"-l"}
	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return false, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(cmdOut))
	r := regexp.MustCompile("Found DFU:")
	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			log.Println(scanner.Text())
			dfuCount++
		}
	}

	switch dfuCount {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, fmt.Errorf("too many dfu found (expected 1): %d", dfuCount)
	}
}
