# wcnexus-venture-server
The server part of WCNexus Venture service

## Development and Compilation

### Fetch package

``` shell
  $ go build
  $ go mod tidy
  $ go mod vendor
```

### Get the dist folder ready

- Create the `dist` folder

### Run

You should always start your program at `dist` folder after compilation.

``` shell
  # Linux
  $ cd dist
  $ go build -o ./server ../src/server.go && ./server
```

``` powershell
  # Windows
  PS> cd dist
  PS> go build -o .\server.exe ..\src\server.go ; .\server.exe
```