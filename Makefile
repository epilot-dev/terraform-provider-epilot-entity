.PHONY: *
all: speakeasy docs

speakeasy:
	speakeasy generate sdk --lang terraform -o . -s entity.yaml

docs:
	go generate ./...

