GOPATH:=$(CURDIR)
export GOPATH

all: bld

bld: tcpIpServer

tcpIpServer:
	go build -o bin/tcpIpServer tcpIpServer

clean:
	@rm -f bin/tcpIpServer
	@rm -rf ./pkg
	@rm -rf status
	@rm -f  log/*log*
	@rm -rf ./output

cleanlog:
	@rm -f log/*log*
