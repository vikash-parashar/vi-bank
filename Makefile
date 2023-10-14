build:
	go build -o main

run: build
	./main

clean:
	rm -f main
