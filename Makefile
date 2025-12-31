build:
	go build -o main utils.go acolyte.go pretre.go main.go

run: build
	./main

clean:
	rm -f main
