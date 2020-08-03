package main

import (
	"log"
	"os/exec"
)

func main() {

	devCd := exec.Command("cd", "D:/")
	_, err := devCd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}
