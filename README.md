# gopass

![](https://img.shields.io/github/issues/michalswi/gopass)
![](https://img.shields.io/github/forks/michalswi/gopass)
![](https://img.shields.io/github/stars/michalswi/gopass)
![](https://img.shields.io/github/last-commit/michalswi/gopass)
![](https://img.shields.io/github/release/michalswi/gopass)

Simple go data/password encryptor.


### \# requirements

**go** in version **1.23.2**


### \# installation

`gopass` requires go1.23.2 to install successfully.

```
go install -v github.com/michalswi/gopass@latest
```

You can also get `macos` or `linux` binary from [releases](https://github.com/michalswi/gopass/releases).


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

$ ./gopass -g

┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐
│ ┬│ │├─┘├─┤└─┐└─┐
└─┘└─┘┴  ┴ ┴└─┘└─┘
        v0.3.1 - @michalswi

secret: mZLbpOORzppLSOirl6KBnqUGQ56jJI1f


# encrypt data

$ ./gopass -e mZLbpOORzppLSOirl6KBnqUGQ56jJI1f data.txt

┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐
│ ┬│ │├─┘├─┤└─┐└─┐
└─┘└─┘┴  ┴ ┴└─┘└─┘
        v0.3.1 - @michalswi

Encrypting data.
Encrypted file: 'encrypted.dat'.


# decrypt data

$ ./gopass -d mZLbpOORzppLSOirl6KBnqUGQ56jJI1f encrypted.dat

┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐
│ ┬│ │├─┘├─┤└─┐└─┐
└─┘└─┘┴  ┴ ┴└─┘└─┘
        v0.3.1 - @michalswi

Decrypting data.
Decrypted file: 'decrypted.dat'.

$ cat decrypted.dat
Thu Mar 14 12:44:18 CET 2024
Thu Mar 14 12:44:19 CET 2024
Thu Mar 14 12:44:19 CET 2024
```


### \# warning

Keep in mind that secret key written in console can be later on read from:
```
$ cat ~/.*_history
```
