.PHONY: all be clean

GOPATH :=
ifeq ($(OS),Windows_NT)
	GOPATH := $(CURDIR)/_vender;$(CURDIR)
else
	GOPATH := $(CURDIR)/_vender:$(CURDIR)
endif

export GOPATH

all: be 

be:
	go install be/cmd/xksjht
	go install be/cmd/parser

clean:
	rm -rf bin
	rm -rf pkg
