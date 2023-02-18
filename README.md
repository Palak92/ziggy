# ziggy
RPC API service  to serve up Crypto price information.

## Prerequites (MacOs)
1. Install go
2. Intall protobuf

```
$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+
```

3. Install the protocol compiler plugins for Go.

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

4. Update your PATH so that the protoc compiler can find the plugins.

```
$ export PATH="$PATH:$(go env GOPATH)/bin"

```