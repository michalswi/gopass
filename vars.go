package main

import (
	"log"
	"os"
)

var version = "v0.3.1"

var (
	generate            bool
	appVersion          bool
	encFile, decFile    string
	encrypt, decrypt    string
	secretKey, dataFile string
)

var logger = log.New(os.Stdout, "gopass ", log.LstdFlags|log.Lshortfile|log.Ltime|log.LUTC)
