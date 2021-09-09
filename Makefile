install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.39.0
	go get golang.org/x/tools/cmd/goimports
	# go get github.com/golang/protobuf/protoc-gen-go@v1.3.2

lint:
	go mod tidy
	goimports -local github.com/Laisky/gorm -w .
	gofmt -s -w .
	golangci-lint run -E golint,depguard,gocognit,goconst,gofmt,misspell,exportloopref,nilerr #,gosec,lll

changelog:
	./_deployment/generate_changelog.sh
