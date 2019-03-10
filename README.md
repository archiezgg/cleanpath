# cleanpath
cleanpath is a lightweight Golang Command-line application to show the duplicates in the UNIX system's PATH variable, if any. With the cli flags the user can invoke other mechanisms on PATH as well.

## Installing Go
First of all, in need of using cleanpath, you need to compile the source code to get an executable binary of it. In order to do that, you'll need Go installed on your computer. Please follow the the official [Google documentation](https://golang.org/doc/install). 

## Creating an executable
Download the source code, and in the directory run `go build -o bin/cleanpath`. This will create a folder named *bin* with an executable binary in it called *cleanpath*. Place this binary file somewhere in your *PATH* so you can call it by simply typing `cleanpath`.

## Usage and flags
`cleanpath [OPTIONS]`


Just by simply calling `cleanpath` you get the help page, which is:

```
Welcome to cleanpath, a lightweight Go CLI application to show your duplicates in your PATH variable!

Usage: cleanpath [OPTIONS]

Options:
	
  -h,		Shows help page.
	
  -g,		Shows your current PATH.
	
  -s,		Prints status about the number of duplicates and the duplicates itself.
	
  -p,		Prints what would your clean path look like.
```
## Cleaning your PATH
As you can see `cleanpath -p` returns the "cleaned" path to standard output, sou you can save it to your PATH by running 

`export PATH=$(cleanpath -p)`

## Saving the changes
The above command will only clean your path for the time as long as your shell is running. To save the changes, save the `export PATH=$(cleanpath -p)` command to your configuration file which gets loaded everytime you run a shell (.bashrc, .zshrc, .profile etc. - whichever you use).

# IMPORTANT
In order to avoid bugs, **save the above command at the _end of your config file_**. Since `cleanpath` is working with the current PATH when it is ran, it's important to have all the PATH edited and setup in the configuration file **before** the cleaning.
