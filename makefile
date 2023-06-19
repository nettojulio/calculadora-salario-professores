build:
	@make build-linux && make build-windows && make build-mac-m1

build-linux:
	@cd cmd && GOOS=linux GOARCH=amd64 go build -o ../calculadora-linux

build-windows:
	@cd cmd && GOOS=windows GOARCH=amd64 go build -o ../calculadora-windows.exe

build-mac-m1:
	@cd cmd && GOOS=darwin GOARCH=arm64 go build -o ../calculadora-mac-m1

run:
	@go run cmd/main.go
