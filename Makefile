#
# afto
#
# http://github.com/hako/afto
#
# @author Wesley Hill (@hakobyte) <wesley@hakobaito.co.uk>
# @license http://opensource.org/licenses/MIT
#

build:
	rm -rf bin
	mkdir bin
	go build  -v -o afto cmd/afto/main.go cmd/afto/scripts.go
	mv afto bin

test:
	go test ./...

test-verbose:
	go test -v ./...

test-coverage:
	go test -v -cover ./...

scrutinise:
	go vet ./...
	golint ./...

scrutinise-verbose: test-coverage scrutinise

clean:
	rm -rf bin
