
tidy:
	go mod tidy
.PHONY: tidy

upgrade:
	go get -u -v ./...
	$(MAKE) tidy
.PHONY: upgrade