# golang-programs

# Official Documentation :
    https://go.dev/doc/
    https://pkg.go.dev/

# 1. Initialise specific package in golang :
    go mod init package_name

# 2. Run the golang program :
    go run main.go

# 3. gopath and godocs :
    go help
    go help gopath

# 4. Some important packages in golang :
    fmt
    strings
    bufio
    os
    strconv
    sort
    math/rand

# 5. Set the golang environment and build for different os :
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

