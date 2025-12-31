build:
	go build -o main acolyte.go pretre.go main.go

run: build
	./main

clean:
	rm -f main
