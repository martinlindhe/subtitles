install:
	go install -v cmd/subber/subber.go

release:
	goreleaser --rm-dist

update-deps:
	rm -rf vendor
	dep ensure -update

test:
	go test -v ./...
