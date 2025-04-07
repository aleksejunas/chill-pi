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

verify-templates:
	@echo "📋 Sjekker at genererte templates kompilerer..."
	@cd $(PROJECT_ROOT) && go test -v ./handlers

verify-project: tidy verify lint verify-templates
	@echo "✅ Project verified"

run:
	@echo "🚀 Kjører server..."
	@cd $(PROJECT_ROOT) && go run main.go

dev:
	@echo "🌱 Utviklingsserver starter med logging til logs/server.log"
	@mkdir -p logs
	@cd $(PROJECT_ROOT) && go run main.go 2>&1 | tee -a logs/server.log

test:
	@echo "🧪 Kjører tester..."
	@cd $(PROJECT_ROOT) && go test ./... | tee logs/test.log

tidy:
	@echo "🔍 Rydder moduler..."
	@cd $(PROJECT_ROOT) && go mod tidy

clean:
	@echo "🧼 Viser .go-filer under 10 bytes:"
	@find . -type f -name "*.go" -size -10c -exec ls -lh {} \;
	@echo "❗ Slett dem manuelt hvis du er sikker."

clean-dry:
	@echo "🔍 Viser tomme .go-filer (ingen sletting)..."
	@find . -type f -name "*.go" -exec bash -c '[[ ! -s "{}" ]] && echo "{}"' \;

lint:
	@echo "🧽 Kjører linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		cd $(PROJECT_ROOT) && golangci-lint run ./... | tee logs/lint.log; \
	else \
		echo "🔧 golangci-lint ikke installert. Installer med:"; \
		echo "    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

scaffold-resource:
	@if [ -z "$(name)" ]; then \
		echo "❌ Du må oppgi navn: make scaffold-resource name=task"; exit 1; \
	fi
	@$(PROJECT_ROOT)/scripts/scaffold.sh $(name)

