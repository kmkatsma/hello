DEPS=go list -f '{{range .TestImports}}{{.}} {{end}}' ./...

update-deps:
	glide up

install-deps:
	glide install

fmt:
	bash -c 'go list ./... | grep -v vendor | xargs -n1 go fmt'

test:
	bash -c 'go list ./... | grep -v vendor | xargs -n1 go test -timeout=30s'