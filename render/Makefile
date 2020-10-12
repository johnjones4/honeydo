.PHONY: build clean

build:
	GOPATH=$(shell pwd) /usr/local/go/bin/go build -o bin/render

clean:
	rm -rf ./bin

install-deps:
	sudo apt install xfonts-75dpi wkhtmltopdf imagemagick
	wget http://www.airspayce.com/mikem/bcm2835/bcm2835-1.68.tar.gz
	tar zxvf bcm2835-1.68.tar.gz
	rm bcm2835-1.68.tar.gz
	cd bcm2835-1.68 && ./configure
	cd bcm2835-1.68 && sudo make check
	cd bcm2835-1.68 && sudo make install
	rm -rf bcm2835-1.68
	wget https://golang.org/dl/go1.15.2.linux-armv6l.tar.gz
	tar zxvf go1.15.2.linux-armv6l.tar.gz
	sudo mv go /usr/local/go
	wget https://www.waveshare.com/w/upload/1/15/IT8951.tar.gz
	tar zxvf IT8951.tar.gz
	rm IT8951.tar.gz
	cd IT8951 && make clean
	cd IT8951 && make
