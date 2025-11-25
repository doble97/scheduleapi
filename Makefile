# Makefile

# Comandos de Go
GO = go
GOCMD = $(GO)
GOLINT = golangci-lint
# Variables de carpetas
CMD_DIR = cmd/server/main.go
# Directorio de los tests de todo el proyecto (incluye services, handlers, y repositorios)
TEST_ALL_DIR = ./internal/... 

# Directorio de los tests unitarios (solo Servicios/Casos de Uso)
TEST_SERVICE_DIR = ./internal/core/services/... 

# --- TAREAS  ---
.PHONY: start
start: ## Corre la aplicación principal.
	@echo "🔥 Iniciando la aplicación..."
	$(GOCMD) run $(CMD_DIR)

.PHONY: test	
test: ## Ejecuta todos los tests en todo el proyecto.
	@echo "🧪 Ejecutando todos los tests..."
	go test $(TEST_ALL_DIR) -v

.PHONY: unit-test
unit-test: ## Ejecuta solo los tests de la capa de Servicios (lógica de negocio).
	@echo "🧪 Ejecutando tests unitarios (Servicios)..."
	go test $(TEST_SERVICE_DIR) -v

.PHONY: coverage
coverage: ## Genera el reporte de cobertura de los Servicios y lo abre en HTML.
	@echo "📊 Calculando cobertura de la lógica de negocio..."
	# 1. Ejecuta tests y genera el archivo de perfil
	go test $(TEST_SERVICE_DIR) -coverprofile=coverage.out
	
	# 2. Muestra el porcentaje de cobertura en la consola
	go tool cover -func=coverage.out
	
	# 3. Abre la vista HTML interactiva en el navegador
	go tool cover -html=coverage.out
	
	# 4. Limpia el archivo de perfil generado
	@rm coverage.out