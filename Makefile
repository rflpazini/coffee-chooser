SERVICE_NAME=coffee-chooser
BINARY_NAME=bin/$(SERVICE_NAME)

# Main
build: easyjson
	@echo Building $(BINARY_NAME)
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

deps:
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
	go get github.com/mailru/easyjson && go install github.com/mailru/easyjson/...@latest
	go install golang.org/x/tools/...@latest
	go mod download

deps-reset:
	git checkout -- go.mod
	go mod tidy

deps-upgrade:
	go get github.com/mailru/easyjson && go install github.com/mailru/easyjson/...@latest
	go get -u -t -d -v ./...
	go mod tidy

deps-cleancache:
	go clean -modcache

# ==============================================================================
# Tools commands
gen: easyjson-clean \
easyjson

easyjson:
	@echo Generating easyjson files...
	go generate ./pkg/api/...
	go generate ./pkg/service/...
	go generate ./pkg/config/...

easyjson-clean:
	@echo Cleaning easyjson files...
	find . -name "*_easyjson.go" -exec rm -f {} \;