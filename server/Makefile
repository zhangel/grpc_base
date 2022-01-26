BIN=grpc_base
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $(BIN) -ldflags="-r ." .
run:build
	./$(BIN) -config.file.path=./config/config_for_dev.ini
clean:
	rm -rvf $(BIN)
