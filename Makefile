.PHONY: clean all

all: mergo

mergo: main.go
	go build

clean:
	rm -f mergo
