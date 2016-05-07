package util

import "log"

func LogFatalErr(err error, debug bool) {
	if err == nil || !debug {
		return
	}

	log.Fatal(err)
}

func LogFatalMsg(err error, msg string, debug bool) {
	if err == nil || !debug {
		return
	}

	log.Fatalln(msg)
}

func LogFatalfMsg(err error, msg string, debug bool) {
	if err == nil || !debug {
		return
	}

	log.Fatalf(msg, err)
}

func LogErr(err error, debug bool) {
	if err == nil || !debug {
		return
	}

	log.Println(err)
}

func LogErrMsg(err error, msg string, debug bool) {
	if err == nil || !debug {
		return
	}

	log.Println(msg)
}

func LogMsg(msg string, debug bool, fatal bool) {
	if !debug {
		return
	}

	if fatal {
		log.Fatalln(fatal)
	}

	log.Println(msg)
}
