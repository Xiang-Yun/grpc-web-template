# CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o bin/api cmd/api/*.go
# CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o bin/web cmd/web/*.go

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o bin/api cmd/api/*.go
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o bin/web cmd/web/*.go


