GO_TEST = go tool gotest.tools/gotestsum --format pkgname
GO_LINT = go tool github.com/golangci/golangci-lint/v2/cmd/golangci-lint
GO_FUMPT = go tool mvdan.cc/gofumpt

#   🔨 TOOLS       #
##@ Tools

prep: prep/tools

prep/tools:
	@if ! command -v copywrite >/dev/null 2>&1; then \
		echo "copywrite is not installed. Installing copywrite..."; \
		go install github.com/hashicorp/copywrite@latest; \
	fi

	@if [ ! -f .lefthook.yml ]; then \
		echo "Installing lefthook configuration..."; \
		lefthook install; \
	fi

#   🧹 Formatting   #
##@ Formatting

fmt:
	$(GO_FUMPT) -w .

fmt/check:
	$(GO_FUMPT) -d .

#   🔍 Linting     #
##@ Linting

lint:
	$(GO_LINT) run ./...

#   ⛹🏽‍ License   #
##@ License

license: license/headers/check

license/headers/check:
	copywrite headers --plan

test/unit:
	mkdir -p build/reports
	$(GO_TEST) --junitfile build/reports/test-unit.xml -- -race ./... -count=1 -short -cover -coverprofile build/reports/unit-test-coverage.out
