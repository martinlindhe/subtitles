install:
	go install -v ./...

release:
	goreleaser --rm-dist
