PROJECT_ROOT := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

verify:
	@echo "🔍 Verifiserer prosjektet..."

	@echo "📦 Kjører go build..."
	@go build -o /dev/null ./... || { echo '❌ Build feilet'; exit 1; }

	@echo "🧠 Kjører go vet..."
	@go vet ./... || { echo '❌ go vet fant problemer'; exit 1; }

	@echo "🧹 Ser etter trivielt tomme Go-filer..."
	@go run tools/scanempty/main.go || echo "⚠️  scanempty feilet (OK hvis du ikke har laget den ennå)"

	@echo "✅ Verifisering ferdig"

# 🚀 Kjør main.go
run:
	@echo "🚀 Kjører server..."
	@cd $(PROJECT_ROOT) && go run main.go

# 🌱 Dev-server med logging til server.log og terminal
dev:
	@echo "🌱 Utviklingsserver starter med logging til logs/server.log"
	@mkdir -p logs
	@cd $(PROJECT_ROOT) && go run main.go 2>&1 | tee -a logs/server.log

# 🧪 Test og lagre til test.log
test:
	@echo "🧪 Kjører tester..."
	@cd $(PROJECT_ROOT) && go test ./... | tee logs/test.log

# 🧼 Rydd go.mod og go.sum
tidy:
	@echo "🔍 Rydder moduler..."
	@cd $(PROJECT_ROOT) && go mod tidy

# 🧹 Slett tomme .go-filer
clean:
	@echo "🧼 Viser .go-filer under 10 bytes:"
	@find . -type f -name "*.go" -size -10c -exec ls -lh {} \;
	@echo "❗ Slett dem manuelt hvis du er sikker."

# 🔍 Liste over tomme .go-filer
clean-dry:
	@echo "🔍 Viser tomme .go-filer (ingen sletting)..."
	@find . -type f -name "*.go" -exec bash -c '[[ ! -s "{}" ]] && echo "{}"' \;

# 🧽 Linter (krever golangci-lint installert)
lint:
	@echo "🧽 Kjører linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		cd $(PROJECT_ROOT) && golangci-lint run ./... | tee logs/lint.log; \
	else \
		echo "🔧 golangci-lint ikke installert. Installer med:"; \
		echo "    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi
