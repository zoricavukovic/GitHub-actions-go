.PHONY: default
default: build

all: clean get-deps build test

version := "0.1.0"

clean:
	rm -rf ./bin

build:
	mkdir -p bin
	go get .
	go build -o bin/service main.go

test: build
	go test -short -coverprofile=bin/cov.out `go list ./... | grep -v vendor/`
	go tool cover -func=bin/cov.out 


sonar: test
	sonar-scanner -Dsonar.projectVersion="$(version)"

start-sonar:
	docker run --name sonarqube -p 9000:9000 sonarqube