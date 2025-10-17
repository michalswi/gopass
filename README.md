# gopass

![](https://img.shields.io/github/issues/michalswi/gopass)
![](https://img.shields.io/github/forks/michalswi/gopass)
![](https://img.shields.io/github/stars/michalswi/gopass)
![](https://img.shields.io/github/last-commit/michalswi/gopass)
![](https://img.shields.io/github/release/michalswi/gopass)

GO data/password encryptor.


### \# requirements

**go** in version **1.25.3**


### \# installation

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
Fri Oct 17 18:26:48 CEST 2025
Fri Oct 17 18:26:49 CEST 2025
Fri Oct 17 18:26:49 CEST 2025


# generate secret key

$ ./gopass -g

┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐
│ ┬│ │├─┘├─┤└─┐└─┐
└─┘└─┘┴  ┴ ┴└─┘└─┘
        v0.4.1 - @michalswi

secret: mZLbpOORzppLSOirl6KBnqUGQ56jJI1f


# encrypt data

$ ./gopass -e mZLbpOORzppLSOirl6KBnqUGQ56jJI1f data.txt

┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐
│ ┬│ │├─┘├─┤└─┐└─┐
└─┘└─┘┴  ┴ ┴└─┘└─┘
        v0.4.1 - @michalswi

Encrypting data.
Encrypted file: 'encrypted.dat'.


# decrypt data

$ ./gopass -d mZLbpOORzppLSOirl6KBnqUGQ56jJI1f encrypted.dat

┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐
│ ┬│ │├─┘├─┤└─┐└─┐
└─┘└─┘┴  ┴ ┴└─┘└─┘
        v0.4.1 - @michalswi

Decrypting data.
Decrypted file: 'decrypted.dat'.

$ cat decrypted.dat
Fri Oct 17 18:26:48 CEST 2025
Fri Oct 17 18:26:49 CEST 2025
Fri Oct 17 18:26:49 CEST 2025
```


### \# warning

Keep in mind that secret key written in console can be later on read from:
```
$ cat ~/.*_history
```
