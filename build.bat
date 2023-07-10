SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o bin/hash-darwin-amd64

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o bin/hash-linux-amd64

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o bin/hash-windows-amd64.exe