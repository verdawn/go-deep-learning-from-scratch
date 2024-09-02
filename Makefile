
run-ch02: ch02
	go run ./cmd/ch02

ch02:
	go build -o ./build/ch02 ./cmd/ch02