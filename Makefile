PACKAGE_NAME = platformclientv2

build:
	env GOOS=windows GOARCH=amd64 go build ./${PACKAGE_NAME}
	env GOOS=linux GOARCH=amd64 go build ./${PACKAGE_NAME}
	env GOOS=darwin GOARCH=amd64 go build ./${PACKAGE_NAME}

test:
	go test ./${PACKAGE_NAME} -v -failfast
