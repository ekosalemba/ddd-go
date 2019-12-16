# DDD go
Domain Driven Design written in Go

# Installation

Install GVM (Go Version Manager) on your machine using this guideline https://github.com/moovweb/gvm.

```
gvm install go1.13.5
gvm use go1.13.5
mdkir $GOPATH/src
git clone git@github.com:ekosalemba/ddd-go.git $GOPATH/src/ddd-go
cd $GOPATH/src/ddd-go
```

Install dependency Manager https://github.com/golang/dep.

```
cd $GOPATH/src/ddd-go
dep ensure
```


# How to run
1. copy file at presentation/config/config.json.sample and rename to rename to config.json
2. move to presentation directory
3. run command `go run main.go`
