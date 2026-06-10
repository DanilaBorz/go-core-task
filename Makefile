.PHONY: test coverage coverage-html clean

COVERAGE_DIR := coverage
COVERAGE_PROFILE := $(COVERAGE_DIR)/coverage.out
COVERAGE_HTML := $(COVERAGE_DIR)/coverage.html

test:
	go test ./...

coverage:
	mkdir -p $(COVERAGE_DIR)
	go test ./... -coverprofile=$(COVERAGE_PROFILE)
	go tool cover -func=$(COVERAGE_PROFILE)

coverage-html: coverage
	go tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	@echo "HTML coverage report: $(COVERAGE_HTML)"

clean:
	rm -rf $(COVERAGE_DIR)
