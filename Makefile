TEST_POSTFIX  :=_test.go
SOURCE_POSTFIX:=.go
SOURCES       :=$(wildcard *$(SOURCE_POSTFIX))
SOURCES       :=$(filter-out $(wildcard *$(TEST_POSTFIX)), $(SOURCES))
TARGETS       :=$(patsubst %$(SOURCE_POSTFIX),%,$(SOURCES))
TEST_SOURCES  :=$(wildcard *$(TEST_POSTFIX))
TEST_TARGETS  :=$(patsubst %$(SOURCE_POSTFIX),%,$(TEST_SOURCES))

all: $(TARGETS)

test: $(TEST_TARGETS)

%_test: %_test.go
	go test -bench . $^ $(patsubst %$(TEST_POSTFIX),%,$^).go

%: %.go
	go build $^

clean:
	rm -f $(TARGETS)
