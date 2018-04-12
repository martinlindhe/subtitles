install:
	go install -v ./...

release:
	goreleaser --rm-dist

update-deps:
	rm -rf vendor
	dep ensure -update

test:
	go test -v ./...
