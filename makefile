test:
	go test ./... -v

test_coverage:
	go test ./... --cover

dep:
	go mod download

vet:
	go vet
