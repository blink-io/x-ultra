
.PHONY: tidy
# Run go mod tidy
tidy:
	go mod tidy -v -e

.PHONY: upgrade
# Upgrade packages
upgrade:
	go get -u -v ./...
	$(MAKE) tidy

.PHONY: upgrade2
# Upgrade packages by using go-mod-upgrade
upgrade2:
	goup -v && go-mod-upgrade -v

.PHONY: build
# Build package
build:
	go build ./...

.PHONY: test
# Run tests
test:
	@echo 'Testing'
	go test ./...