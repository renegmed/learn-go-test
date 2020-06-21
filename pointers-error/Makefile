testv1:
	go test v1/* -v
.PHONY: testv1

install-errorcheck:
	go get -u github.com/kisielk/errorcheck
.PHONY: install-errorcheck

error-check-v1:
	errcheck ./v1/*
.PHONY: error-check-v1

