DROPHOSTS_VERSION ?= latest

test:
	go test . -cover

build:
	go build .

build-amd64:
	@env GOOS=linux GOARCH=amd64 go build -ldflags="-X main.appVersion=${DROPHOSTS_VERSION}" -o drophosts

build-i386:
	@env GOOS=linux GOARCH=386 go build -ldflags="-X main.appVersion=${DROPHOSTS_VERSION}" -o drophosts_i386

release: build-amd64 build-i386
	@zip drophosts_${DROPHOSTS_VERSION}_linux_amd64.zip drophosts
	@tar -cvzf drophosts_${DROPHOSTS_VERSION}_linux_amd64.tar.gz drophosts
	@rm drophosts

	@mv drophosts_i386 drophosts
	@zip drophosts_${DROPHOSTS_VERSION}_linux_386.zip drophosts
	@tar -cvzf drophosts_${DROPHOSTS_VERSION}_linux_386.tar.gz drophosts
	@rm drophosts

docker: build-amd64 docker-image clean

docker-image:
	@docker build -t qmxme/drophosts:${DROPHOSTS_VERSION} .

clean:
	@rm -f drophosts
	@rm -rf drophosts_*.zip
	@rm -rf drophosts_*.tar.gz
