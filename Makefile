export GOPATH := $(shell pwd)


all: bin/js2png bin/serve

bin/serve: src/net.expobrain/serve/serve.go
	go install -v net.expobrain/serve

bin/js2png: src/net.expobrain/js2png/js2png.go
	go install -v net.expobrain/js2png

serve: bin/serve
	bin/serve

js2png: bin/js2png
	bin/js2png html/js/payload.js html/payload.png

clean:
	rm -rf bin/*
