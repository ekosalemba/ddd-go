package utils

import "os/exec"
import "log"

func GenerateGuid() []byte {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}
