.PHONY: test
.PHONY: swagger
.PHONY: static

build:
	@go mod tidy -go="${GOLANG_VERSION}"
	@if [ "$(SKIP_UNIT_TEST)" = "false" ]; then $(MAKE) test; else echo "Skipped unit tests"; fi
	@$(MAKE) swagger
	@echo "building icx dashboard service"
	@go build -o build/icx_dashboard -a cmd/icx_dashboard/icx_dashboard.go
	@echo "build  icx dashboard service"
	@$(MAKE) static
swagger:
	@echo "initializing swagger"
	@swag init -g cmd/icx_dashboard/icx_dashboard.go --output docs/ --parseDependency
static:
	@echo "generating static pages"
	@go run ./cmd/icx_dashboard_doc/icx_dashboard_doc.go
	@blackfriday-tool -css="https://unpkg.com/sakura.css/css/sakura.css" ./docs/Static_HTML_Doc.md ./docs/static.html
	@cp ./docs/document.html build/
	@cp ./docs/static.html build/

test:
	@echo "running unit tests"
	@go test -v -coverpkg=./... -coverprofile=profile.cov ./test/...

coverage:
	@echo "generating coverage"
	@go tool cover -html=profile.cov -o coverage.html

lint:
	@echo "running linter fix"
	@golangci-lint run --fix