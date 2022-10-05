package common

import "log"

func CheckError(err error, msg string) {
	if err != nil {
		log.Printf("ERR: %s: %s\n", msg, err.Error())
	}
}

func CheckFatalError(err error, msg string) {
	if err != nil {
		log.Fatalf("FAILED: %s: %s\n", msg, err.Error())
	}
}
