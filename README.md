# Line In File

## Usage
```
Usage: lineinfile <filename> '<line>'

Lineinfile is a simple command line utility to ensure
that a file contains a given line of text.

Examples:
     lineinfile ~/.bashrc alias cdg="cd ~/git-repos"
```

## Install
```
go build lineinfile.go
sudo cp ./lineinfile /usr/local/bin/lineinfile
```
