.PHONY: build clean

build:
	GOPATH=$(shell pwd) /usr/local/go/bin/go build -o bin/render

clean:
	rm -rf ./bin

install-deps:
	sudo apt install xfonts-75dpi wkhtmltopdf imagemagick
	wget https://golang.org/dl/go1.15.2.linux-armv6l.tar.gz
	tar zxvf go1.15.2.linux-armv6l.tar.gz
	sudo cp go /usr/local/go
	wget https://www.waveshare.com/w/upload/1/15/IT8951.tar.gz
	tar zxvf IT8951.tar.gz
	cd IT8951 && make clean
	cd IT8951 && make
