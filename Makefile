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
	@mkdir -p $(PROJECT_ROOT)/handlers $(PROJECT_ROOT)/models
	@RESOURCE_STRUCT=$$(echo $(name) | sed 's/.*/\u&/g'); \
	FILE=$(PROJECT_ROOT)/handlers/$(name)_resource.go; \
	MODEL=$(PROJECT_ROOT)/models/$(name).go; \
	if [ -f $$FILE ]; then \
		echo "⚠️  $$FILE finnes allerede"; \
	else \
		echo "🛠 Lager $$FILE..."; \
		cat > $$FILE <<EOF
package handlers

import (
	"net/http"
	"chill-pi/database"
	"chill-pi/models"
	"chill-pi/utils"
	"github.com/gin-gonic/gin"
)

type $$RESOURCE_STRUCT Resource struct{}

func ($$RESOURCE_STRUCT Resource) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GET all $(name)s"})
}

func ($$RESOURCE_STRUCT Resource) GetOne(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GET one $(name)"})
}

func ($$RESOURCE_STRUCT Resource) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "CREATE $(name)"})
}

func ($$RESOURCE_STRUCT Resource) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UPDATE $(name)"})
}

func ($$RESOURCE_STRUCT Resource) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DELETE $(name)"})
}

func ($$RESOURCE_STRUCT Resource) Patch(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "PATCH ikke implementert"})
}
EOF; \
	fi; \
	if [ -f $$MODEL ]; then \
		echo "⚠️  $$MODEL finnes allerede"; \
	else \
		echo "📦 Lager modell: $$MODEL"; \
		cat > $$MODEL <<EOF
package models

type $$RESOURCE_STRUCT struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
EOF; \
	fi; \
	echo "✅ Ferdig. Husk å registrere route i routes/routes.go:"
	@echo "    RegisterResourceRoutes(r, \"/$(name)s\", handlers.$$RESOURCE_STRUCT Resource{})"
