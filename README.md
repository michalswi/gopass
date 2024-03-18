# gopass

![](https://img.shields.io/github/issues/michalswi/gopass)
![](https://img.shields.io/github/forks/michalswi/gopass)
![](https://img.shields.io/github/stars/michalswi/gopass)
![](https://img.shields.io/github/last-commit/michalswi/gopass)

Simple key password store.

### \# requirements

**go** in version **1.22**


### \# installation

`gopass` requires go1.22 to install successfully.

```
go install -v github.com/michalswi/gopass@latest
```

### \# usage

```
gopass -h
```
This will display help for the tool. It supports:
```
Encrypt or decrypt a file using AES with a 256-bit key file.
Secret key should be 32-bytes long.

FLAGS:
  -h, --help 						Show this help text
  -g, --generate 					Generate 32-bytes secret key
  -e, --encrypt <secret_key> <raw_data_file>		Encrypt data using secret key. Output to STDOUT
  -d, --decrypt <secret_key> <encrypted_data_file>	Decrypt encrypted data using secret key. Output to STDOUT
  -ef <encrypted_data_file> 				Encrypted filename (+location)
  -df <decrypted_data_file>				    Decrypted filename (+location)

USAGE:
  gopass [-h|--help]

  > Generate 32-bytes secret key
  $ gopass [-g|--generate]

  > Encrypt data file using secret key.
  Default encrypted data file - encrypted.dat
  $ gopass [-e|--encrypt] <secret_key> <data_file>
  $ gopass -e my32byteslongsecret data.txt

  > Decrypt encrypted data file using secret key.
  Default decrypted data file - decrypted.dat
  $ gopass [-d|--decrypt] <secret_key> <encrypted_data_file>
  $ gopass -d mysecret encrypted.dat

  > Encrypt data file using secret key.
  Provide location + filename where to save encrypted data file
  $ gopass [-e|--encrypt] <secret_key> <data_file> -ef <encrypted_data_file>
  $ gopass -e my32byteslongsecret data.txt -ef /tmp/encrypted.dat

  > Decrypt encrypted data file (provide location + filename) using secret key.
  Provide location + filename where to save decrypted data file
  $ gopass [-d|--decrypt] <secret_key> <encrypted_data_file> -df <decrypted_data_file>
  $ gopass -d mysecret /tmp/encrypted.dat -df /tmp/decrypted.dat
```
### \# example
```
# create dummy data file

$ date >> data.txt
$ cat data.txt
Thu Mar 14 12:44:18 CET 2024
Thu Mar 14 12:44:19 CET 2024
Thu Mar 14 12:44:19 CET 2024


# generate secret key

$ ./gopass -g

┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐
│ ┬│ │├─┘├─┤└─┐└─┐
└─┘└─┘┴  ┴ ┴└─┘└─┘
	v0.1.0 - @michalswi

secret key: rlJnWAUs/r+39mvTzewehlJJ+hDUEwV9


# encrypt data

todo


# decrypt data

todo
```
