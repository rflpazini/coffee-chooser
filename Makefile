BINARY_NAME=bin/coffee-choose

# Main
build:
	make deps-upgrade
	make easyjson
	go build -o ${BINARY_NAME} ./cmd

run:
	go build -o ${BINARY_NAME} ./cmd
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test -cover ./...

# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get github.com/mailru/easyjson && go install github.com/mailru/easyjson/...@latest
	go get -u -t -d -v ./...
	go mod tidy

deps-cleancache:
	go clean -modcache

# ==============================================================================
# Tools commands

easyjson:
	go generate ./pkg/api/...
	go generate ./pkg/config/...
