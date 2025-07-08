# Official Documentation

    https://go.dev/doc/
    https://pkg.go.dev/

## 1. Display all packages or lists and initialise a specific package in golang and build

    go mod init package_name
    go build .
    go list -m all

## 2. Run the golang program

    go run main.go

## 3. gopath and godocs

    go help
    go help gopath

## 4. Some important packages in golang

    fmt
    strings
    bufio
    os
    strconv
    sort
    math/rand
    io/ioutil
    net/http
    net/url
    encoding/json
    log
    time

## 5. Set the golang environment and build for different os

    go env
    1. For Unix terminal :
        GOOS="windows" go build
        GOOS="linux" go build
        GOOS="darwin" go build
    2. For Windows command prompt :
        set GOOS = windows go build
        set GOOS = linux go build
        set GOOS = darwin go build
    3. For Windows powershell :
        $env:GOOS = "windows" go build
        $env:GOOS = "linux" go build
        $env:GOOS = "darwin" go build

## 6. Modules in golang

    go mod init github.com/sheteshreyash/modules_name
    go get -u github.com/gorilla/mux
    go mod tidy (expensive)
    go mod verify
    go mod why github.com/gorilla/mux (expensive)
    go mod graph (expensive)
    go mod edit -go go_version
    go mod edit -module new_module_name
    go mod vendor

## 7. Start the node js webserver in the Icowebserver directory

    npm install
    npm start

## 8. Other requirements

    go get -u github.com/gorilla/mux
