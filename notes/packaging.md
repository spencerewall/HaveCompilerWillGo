# Packaging in Go

Installing and distributing packages in Go turned out to be even simpler than I'd imagined. We picked up some basic packaging during the [go.org intro tutorial](https://go.dev/doc/code). 

## Basic packaging

After creating a project with a simple `mkdir`, we can create a module with:
```
> go mod init example/sewall/hello
go: creating new go.mod: module example/user/hello
```
This creates a file rooted in that directory
```
> cat go.mod
module example/user/hello

go 1.16

```
With your module, you can go and run the install command, which will both package it up and install it locally
```
> go install example/sewall/hello
```
The generated executable can be found at `$HOME/go/bin/hello` or `%USERPROFILE%\go\bin\hello.exe` on Windows.

## Controlling the Install Directory
The install directory is controlled by the `GOPATH` and `GOBIN` environment variables:
```
> go env -w GOBIN=/somewhere/else/bin
> go install example/sewall/hello
```
With GOBIN, the file is installed directly to the specified directory  `/somewhere/else/bin/hello.exe`. With `GOPATH`, however, it will be installed to a bin directory under the specified path:
```
> go env -w GOPATH=/somewhere/else
> go install example/sewall/hello
```
The above command installs to `/somewhere/else/bin/hello.exe`. You can also install to a specific directory with the `go install` command:
```
> go install <path>
```

## Including Dependencies

TODO - Including dependencies for distribution is an important part of any production build system. Learning out to manage and include them at distribution will be the topic of this section.