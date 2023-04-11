.PHONY: all
all: docs

speakeasy:
	speakeasy generate sdk --lang terraform -o . -s entity.yaml

docs: speakeasy
	go generate ./...

