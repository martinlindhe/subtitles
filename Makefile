install:
	go install ./cmd/subber/subber.go
	go install ./cmd/ssa2srt/subber.go

test:
	goreleaser check
	go test -v ./...
