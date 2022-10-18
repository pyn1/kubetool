GOPATH=$(shell go env GOPATH)

all:
	go build -o kubetool -buildvcs=false
	chmod ugo+x kubetool

clean:
	rm -rf kubetool
