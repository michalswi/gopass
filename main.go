package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/michalswi/color"
)

func init() {
	flag.BoolVar(&generate, "g", false, "Generate 32-bytes secret key.")
	flag.BoolVar(&generate, "generate", false, "Generate 32-bytes secret key.")
	flag.StringVar(&encrypt, "e", "", "Encrypt the input file")
	flag.StringVar(&encrypt, "encrypt", "", "Encrypt the input file")
	flag.StringVar(&decrypt, "d", "", "Decrypt the input file")
	flag.StringVar(&decrypt, "decrypt", "", "Decrypt the input file")
	flag.StringVar(&encFile, "ef", "encrypted.dat", "Encrypted filename (+location)")
	flag.StringVar(&decFile, "df", "decrypted.dat", "Decrypted filename (+location)")

	flag.Usage = func() {
		h := []string{
			"",
			"Encrypt or decrypt a file using AES with a 256-bit key file.",
			"Secret key should be 32-bytes long.",
			"",
			"FLAGS:",
			"  -h, --help 						Show this help text",
			"  -g, --generate 					Generate 32-bytes secret key",
			"  -e, --encrypt <secret_key> <raw_data_file>		Encrypt data using secret key. Output to STDOUT",
			"  -d, --decrypt <secret_key> <encrypted_data_file>	Decrypt encrypted data using secret key. Output to STDOUT",
			"  -ef <encrypted_data_file> 				Encrypted filename (+location)",
			"  -df <decrypted_data_file>				Decrypted filename (+location)",
			"",
			"USAGE:",
			"  gopass [-h|--help]",
			"",
			"  > Generate 32-bytes secret key",
			"  $ gopass [-g|--generate]",
			"",
			"  > Encrypt data file using secret key.",
			"  Default encrypted data file - encrypted.dat",
			"  $ gopass [-e|--encrypt] <secret_key> <data_file>",
			"  $ gopass -e my32byteslongsecret data.txt",
			"",
			"  > Decrypt encrypted data file using secret key.",
			"  Default decrypted data file - decrypted.dat",
			"  $ gopass [-d|--decrypt] <secret_key> <encrypted_data_file>",
			"  $ gopass -d mysecret encrypted.dat",
			"",
			"  > Encrypt data file using secret key.",
			"  Provide location + filename where to save encrypted data file",
			"  $ gopass [-e|--encrypt] <secret_key> <data_file> -ef <encrypted_data_file>",
			"  $ gopass -e my32byteslongsecret data.txt -ef /tmp/encrypted.dat",
			"",
			"  > Decrypt encrypted data file (provide location + filename) using secret key.",
			"  Provide location + filename where to save decrypted data file",
			"  $ gopass [-d|--decrypt] <secret_key> <encrypted_data_file> -df <decrypted_data_file>",
			"  $ gopass -d mysecret /tmp/encrypted.dat -df /tmp/decrypted.dat",
			"",
		}
		ShowBanner()
		// fmt.Fprintf(os.Stderr, "%s", strings.Join(h, "\n"))
		println(color.Format(color.GREEN, strings.Join(h, "\n")))
	}
	flag.Parse()
	secretKey, dataFile, encFile, decFile = checkArgs()
}

func checkArgs() (string, string, string, string) {
	if len(os.Args) < 2 || len(os.Args) > 6 {
		fmt.Println(color.Format(color.RED, "Read the help text with -h or --help"))
		os.Exit(1)
	}

	if len(os.Args) == 2 && os.Args[1] == "-g" || os.Args[1] == "--generate" {
		return "", "", "", ""
	}

	// default encryption/decryption
	if len(os.Args) == 4 {
		if os.Args[1] == "-e" || os.Args[1] == "--encrypt" {
			return os.Args[2], os.Args[3], encFile, ""
		}
		if os.Args[1] == "-d" || os.Args[1] == "--decrypt" {
			return os.Args[2], encFile, "", decFile
		}
	}

	// specified encryption/decryption
	if len(os.Args) == 6 {
		if os.Args[1] == "-e" && os.Args[4] == "-ef" || os.Args[1] == "--encrypt" && os.Args[4] == "-ef" {
			return os.Args[2], os.Args[3], os.Args[5], ""
		}
		if os.Args[1] == "-d" && os.Args[4] == "-df" || os.Args[1] == "--decrypt" && os.Args[4] == "-df" {
			return os.Args[2], os.Args[3], "", os.Args[5]
		}
	}

	return "", "", "", ""
}

func main() {
	if generate {
		secret, err := Generate256BitSecret()
		if err != nil {
			logger.Println("Error:", err)
			return
		}
		ShowBanner()
		fmt.Println()
		fmt.Println(color.Format(color.GREEN, fmt.Sprintf("secret key: %x", base64.StdEncoding.EncodeToString(secret[:24]))))
		os.Exit(0)
	}

	secretKeyData := []byte(secretKey)
	// secretKeyData = []byte("passphrasewhichneedstobe32bytes!")

	fileData, err := os.ReadFile(dataFile)
	if err != nil {
		logger.Fatal("Unable to read file data contents.", err)
	}

	if encrypt != "" {
		ShowBanner()
		fmt.Println()
		fmt.Println(color.Format(color.GREEN, "Encrypting data."))
		// logger.Println(color.Format(color.GREEN, "Encrypting data."))
		cipherText, err := EncryptAES(secretKeyData, fileData)
		if err != nil {
			logger.Fatal("Error encrypting. ", err)
		}
		err = os.WriteFile(encFile, cipherText, 0644)
		if err != nil {
			logger.Fatal("Error write to file. ", err)
		}
		// fmt.Printf("%s", cipherText)
		fmt.Println(color.Format(color.GREEN, fmt.Sprintf("Encrypted file: '%s'.", encFile)))
		// logger.Println(color.Format(color.GREEN, fmt.Sprintf("Encrypted file: '%s'.", encFile)))
	}

	if decrypt != "" {
		ShowBanner()
		fmt.Println()
		fmt.Println(color.Format(color.GREEN, "Decrypting data."))
		// logger.Println(color.Format(color.GREEN, "Decrypting data."))
		message, err := DecryptAES(secretKeyData, fileData)
		if err != nil {
			logger.Fatal("Error decrypting. ", err)
		}
		err = os.WriteFile(decFile, message, 0644)
		if err != nil {
			logger.Fatal("Error write to file. ", err)
		}
		// fmt.Printf("%s", message)
		fmt.Println(color.Format(color.GREEN, fmt.Sprintf("Decrypted file: '%s'.", decFile)))
		// logger.Println(color.Format(color.GREEN, fmt.Sprintf("Decrypted file: '%s'.", decFile)))
	}
}
