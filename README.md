# gopass

![](https://img.shields.io/github/issues/michalswi/gopass)
![](https://img.shields.io/github/forks/michalswi/gopass)
![](https://img.shields.io/github/stars/michalswi/gopass)
![](https://img.shields.io/github/last-commit/michalswi/gopass)
![](https://img.shields.io/github/release/michalswi/gopass)

Simple go data/password encryptor.

### \# requirements

**go** in version **1.22**


### \# installation

`gopass` requires go1.22 to install successfully.

```
go install -v github.com/michalswi/gopass@latest
```

### \# usage

This will display help for the tool.

```
gopass -h
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

$ gopass -g

┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐
│ ┬│ │├─┘├─┤└─┐└─┐
└─┘└─┘┴  ┴ ┴└─┘└─┘
	v0.1.0 - @michalswi

secret: r10jNonOY5gHOESvbig1Wpkb5baSoKyZ


# encrypt data

$ gopass -e r10jNonOY5gHOESvbig1Wpkb5baSoKyZ data.txt

┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐
│ ┬│ │├─┘├─┤└─┐└─┐
└─┘└─┘┴  ┴ ┴└─┘└─┘
	v0.1.0 - @michalswi

Encrypting data.
Encrypted file: 'encrypted.dat'.


# decrypt data

$ gopass -d r10jNonOY5gHOESvbig1Wpkb5baSoKyZ encrypted.dat

┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐
│ ┬│ │├─┘├─┤└─┐└─┐
└─┘└─┘┴  ┴ ┴└─┘└─┘
	v0.1.0 - @michalswi

Decrypting data.
Decrypted file: 'decrypted.dat'.

$ cat decrypted.dat
Thu Mar 14 12:44:18 CET 2024
Thu Mar 14 12:44:19 CET 2024
Thu Mar 14 12:44:19 CET 2024
```
