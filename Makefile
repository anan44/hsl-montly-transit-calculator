guide:
	$(info "use specific make commands and export AWS_PROFILE before deploying")

test:
	go test ./api
	go test ./hsl

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
	zip function.zip main

deploy: | build
	npm run cdk:deploy --prefix infra

