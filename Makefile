
.PHONY: tidy
# Run go mod tidy
tidy:
	go mod tidy -v -e

.PHONY: upgrade
# Upgrade packages
upgrade:
	go get -u -v ./...
	$(MAKE) tidy

.PHONY: build
# Build package
build:
	go build ./...
