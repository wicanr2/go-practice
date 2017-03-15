
DEPS := out/client
DEPS := out/server

all: prefix $(DEPS)

prefix:
	mkdir -p out/
	echo $(DEPS)

clean:
	rm -rf out/

out/client: client.go
	go build -o $@ $<

out/server: server.go
	go build -o $@ $<


