DROPHOSTS_VERSION := $(shell git describe --always)

test:
	go test . -cover

build:
	go build .

build-amd64:
	@env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-X main.appVersion=${DROPHOSTS_VERSION}" -o drophosts

release: clean build-amd64
	@zip drophosts_${DROPHOSTS_VERSION}_linux_amd64.zip drophosts
	@tar -cvzf drophosts_${DROPHOSTS_VERSION}_linux_amd64.tar.gz drophosts

docker: clean build-amd64 docker-image

docker-image: build-amd64
	@docker build -t qmxme/drophosts:${DROPHOSTS_VERSION} .

push: docker-image
	@docker push qmxme/drophosts:${DROPHOSTS_VERSION}
	@docker push qmxme/drophosts:latest

clean:
	@rm -f drophosts
	@rm -rf drophosts_*.zip
	@rm -rf drophosts_*.tar.gz
