
DEPS := out/hello
DEPS += out/reader
DEPS += out/fibonacci

all: prefix $(DEPS)

prefix:
	mkdir -p out/
	echo $(DEPS)

clean:
	rm -rf out/

out/hello: hello.go
	go build -o out/hello hello.go

out/reader: reader.go
	go build -o out/reader reader.go

out/fibonacci: fibonacci.go
	go build -o out/fibonacci $<

test:
	echo "Test Hello"
	./out/hello
	echo "Test Arguments Reader and Open file"
	./out/reader hello.go
	echo Test Fibonacci Number Computing
	./out/fibonacci


