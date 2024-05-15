
.PHONY: tidy
# Run go mod tidy
tidy:
	go mod tidy -v -e

.PHONY: upgrade
# Upgrade packages
upgrade:
	go get -u -v -t ./...
	$(MAKE) tidy

.PHONY: upgrade2
# Upgrade packages by using go-mod-upgrade
upgrade2:
	goup -v && go-mod-upgrade -v

up-build: upgrade build

.PHONY: build
# Build package
build:
	go build ./...

.PHONY: test
# Run tests
test:
	@echo 'Testing'
	go test ./...

.PHONY: mod-list
# Run mod list
mod-list:
	go list -u -mod=mod -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' -m all

.PHONY: fix-udp
# Run fix UDP
fix-udp:
	sudo sysctl -w net.core.rmem_max=2500000
	sudo sysctl -w net.core.wmem_max=2500000

.PHONY: install-tools
# Run fix UDP
install-tools:
	go install github.com/blink-io/x/kratos/v2/cmd/protoc-gen-x-go-http@latest
	go install github.com/blink-io/x/connectrpc/cmd/protoc-gen-x-connect-go@latest