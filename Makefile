all: build

build:
	go build

test: build
	go test -v

deploy:
	go build -o signmsg.cgi
	scp signmsg.cgi $(LOC)
	rm -f signmsg.cgi

.PHONY: build test deploy
