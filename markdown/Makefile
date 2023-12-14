BINARY_NAME=MarkDown.App
APP_NAME=MarkDown
VERSION=0.0.1

## build: build binary and package app
build:
	rm -rf ${BINARY_NAME}
	rm -rf fyne-md
	fyne package -appVersion ${VERSION} -name ${APP_NAME} -release

## clean: runs go clean and deletes binary
clean:
	@go clean
	@rm -rf ${BINARY_NAME}

test:
	go test -v ./...