.PHONY: install install-deps build-binary build-binary-arm install-binary test

BINARY_NAME=prometheus-crypto-exporter

build: install-deps build-binary

install: install-binary

install-deps:
	go get

build-binary:
	go build -o dist/${BINARY_NAME} main.go

build-binary-arm:
	GOOS=linux GOARCH=arm go build -o dist/${BINARY_NAME}_arm main.go

install-binary:
	mv dist/${BINARY_NAME} /usr/local/bin/${BINARY_NAME}
	chmod +x /usr/local/bin/${BINARY_NAME}
	@echo "Installed binary at '/usr/local/bin/${BINARY_NAME}'"