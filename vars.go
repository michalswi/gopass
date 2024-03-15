package main

import (
	"log"
	"os"
)

var (
	encFile, decFile    string
	encrypt, decrypt    string
	secretKey, dataFile string
	generate            bool
)

var logger = log.New(os.Stdout, "gopass ", log.LstdFlags|log.Lshortfile|log.Ltime|log.LUTC)