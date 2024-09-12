all: ezjwt

ezjwt:
	go build

test:
	go test

clean:
	- go clean
	- rm -f ezjwt-*

cross:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ezjwt-linux-amd64
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ezjwt-windows-amd64.exe
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ezjwt-darwin-amd64

