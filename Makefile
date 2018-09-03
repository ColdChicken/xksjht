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

clean:
	rm -rf bin
	rm -rf pkg
