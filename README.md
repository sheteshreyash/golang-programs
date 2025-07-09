# Official Documentation

* [The Go Programming Language Documentation](https://go.dev/doc/)
* [pkg.go.dev (Go Package Index)](https://pkg.go.dev/)

---

## 1. Display all packages, initialize a module, and build

```bash
# Initialize a new module in the current directory
go mod init <module_name>

# Build the module in the current directory
go build .

# List all required modules & versions
go list -m all
```

---

## 2. Run the Go program

```bash
# run main.go file
go run main.go

# run main.go file with race flag
go run --race main.go
```

---

## 3. GOPATH & Go help

```bash
# Show general help
go help

# Show help on GOPATH
go help gopath
```

---

## 4. Commonly Used Standard Packages

```go
import (
    "bufio"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "net/url"
    "os"
    "sort"
    "strconv"
    "strings"
    "time"
    "sync"
)
```

---

## 5. Environment & Cross‑Compiling

```bash
# Show your Go environment
go env
```

### Unix / Linux / macOS

```bash
GOOS="windows" go build   # Build for Windows
GOOS="linux"   go build   # Build for Linux
GOOS="darwin"  go build   # Build for macOS
```

### Windows Command Prompt (cmd.exe)

```cmd
set GOOS=windows
go build

set GOOS=linux
go build

set GOOS=darwin
go build
```

### Windows PowerShell

```powershell
$env:GOOS = "windows"; go build
$env:GOOS = "linux";   go build
$env:GOOS = "darwin";  go build
```

---

## 6. Go Modules

```bash
# Initialize a module with a repository path
go mod init github.com/yourusername/yourmodule

# Add or update a dependency
go get -u github.com/gorilla/mux

# Remove unused dependencies
go mod tidy

# Verify dependencies have not changed
go mod verify

# Why a module is needed
go mod why github.com/gorilla/mux

# Show module requirement graph
go mod graph

# Edit Go version in go.mod
go mod edit -go=1.20

# Change module path in go.mod
go mod edit -module=github.com/yourusername/newmodule

# Create a vendor directory
go mod vendor
```

---

## 7. Start the Node.js Webserver

```bash
cd Icowebserver
npm install
npm start
```

---

## 8. Gorilla Mux (HTTP Router)

```bash
go get -u github.com/gorilla/mux
```

---

## 9. MongoDB Go Driver

```bash
go get go.mongodb.org/mongo-driver/mongo
```

---

## 10. MongoDB Connection String

```bash
mongodb://localhost:27017/golangmongobasics
```

---

## 11. Initialise module with github repository name

```bash
go mod init github.com/sheteshreyash/module_name
```
