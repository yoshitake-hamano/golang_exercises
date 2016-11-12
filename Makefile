SOURCES:=$(wildcard *.go)
TARGETS:=$(patsubst %.go,%,$(SOURCES))

all: $(TARGETS)

hello: hello.go

%: %.go
	go build $^

clean:
	rm -f $(TARGETS)
