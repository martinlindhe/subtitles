install:
	go install ./cmd/subber/subber.go

test:
	goreleaser check
	go test -v ./...
