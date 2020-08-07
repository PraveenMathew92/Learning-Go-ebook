# CHAPTER 4 PACKAGES

Go build creates binary which unlike JVM languages does not run on a virtual machine.

## Running tests in GO

The below command is used to run test

``` go test ```

The command checks for file with the name *_test.go in the GOPATH and the GOROOT

#### GOPATH
is an environment variable.
To run the tests in this directory set the GOPATH as

``` export GOPATH=$(pwd) ```

in the _learning-go-ebook/ch04/exercise_ directory.

Once the GOPATH is set run

```
go test stack
```
