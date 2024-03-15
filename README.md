# gopass

![](https://img.shields.io/github/issues/michalswi/gopass)
![](https://img.shields.io/github/forks/michalswi/gopass)
![](https://img.shields.io/github/stars/michalswi/gopass)
![](https://img.shields.io/github/last-commit/michalswi/gopass)

Simple key password store.

### \# requirements

**go** in version **1.22**


### \# installation

```
todo

go install...
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
