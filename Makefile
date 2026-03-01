.PHONY: tidy test latest

tidy:
	go mod tidy

test:
	go test -v ./...

latest:
	@git tag --sort=-v:refname | head -n1