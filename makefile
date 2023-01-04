start:
	clear && ./lexb hello.lexb

run:
	clear && go run lexb.go hello.lexb

build:
	clear && go build lexb.go
	clear && ./lexb hello.lexb
