all: build-windows build-linux build

build-windows:
	GOOS=windows go build -o ./bin/win/spritesheetgen ./internal/main.go
build-linux:
	GOOS=linux go build -o ./bin/linux/spritesheetgen ./internal/main.go
build:
	go build -o ./bin/macOs/spritesheetgen ./internal/main.go
