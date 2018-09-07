#
# Standard makefile for go projects
#

# standard definitions
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
GOFMT=$(GOCMD) fmt
BIN=bin

# project specific definitions
BASE_NAME=entity-id-ws
SRC_TREE=entityidws
RUNNER=scripts/entry.sh
PACKAGE=github.com/uvalib/$(BASE_NAME)

all: build-darwin build-linux

build-darwin:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -a -o $(BIN)/$(BASE_NAME).darwin $(PACKAGE)/$(SRC_TREE)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -a -installsuffix cgo -o $(BIN)/$(BASE_NAME).linux $(PACKAGE)/$(SRC_TREE)

test:
	$(GOCLEAN) -testcache $(PACKAGE)/...
	$(GOTEST) -v $(PACKAGE)/$(SRC_TREE)/tests $(if $(TEST),-run $(TEST),)

fmt:
	$(GOFMT) $(PACKAGE)/...

vet:
	$(GOVET) $(PACKAGE)/...

clean:
	$(GOCLEAN) $(PACKAGE)/...
	rm -f $(BIN)/$(BASE_NAME).*

run:
	cd $(BIN); rm -f $(BASE_NAME)
	cd $(BIN); ln -s $(BASE_NAME).darwin $(BASE_NAME)
	$(RUNNER)

#
# end of file
#
