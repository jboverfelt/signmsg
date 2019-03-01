all: build

build:
	go build

deploy:
	GOOS=netbsd go build -o signmsg.cgi
	scp signmsg.cgi $(LOC)
	rm -f signmsg.cgi

.PHONY: build deploy
