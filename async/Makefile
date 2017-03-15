
DEPS := out/gopp

all: prefix $(DEPS)

prefix:
	mkdir -p out/
	echo $(DEPS)

clean:
	rm -rf out/

out/gopp: gochan.go
	go build -o $@ $<

test:
	echo "Test PING-PONG"
	./out/gopp


