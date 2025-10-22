# Makefile para transcoding - Raspberry Pi 5 (ARM64)
# Federico Pereira <fpereira@cnsoluciones.com>

# Variables
IMAGE_NAME := transcoding-arm64
IMAGE_NAME_DEBIAN := transcoding-arm64-debian
CONTAINER_NAME := transcoding-builder
BIN_DIR := bin
LIB_DIR := lib

# Colores para output
GREEN := \033[0;32m
YELLOW := \033[1;33m
RED := \033[0;31m
NC := \033[0m # No Color

.PHONY: help build build-debian extract extract-debian clean all all-debian

# Target por defecto
help:
	@echo "$(GREEN)🎵 Transcoder Makefile para Raspberry Pi 5 (ARM64)$(NC)"
	@echo ""
	@echo "$(YELLOW)Comandos disponibles:$(NC)"
	@echo "  make build         - Construye la imagen Docker para ARM64 (Alpine/musl)"
	@echo "  make build-debian  - Construye la imagen Docker para ARM64 (Debian/glibc)"
	@echo "  make extract       - Extrae el binario y librerías del contenedor"
	@echo "  make extract-debian - Extrae el binario y librerías (Debian/glibc)"
	@echo "  make all           - Ejecuta build + extract (Alpine)"
	@echo "  make all-debian    - Ejecuta build + extract (Debian/glibc)"
	@echo "  make clean         - Limpia archivos generados"
	@echo "  make help          - Muestra esta ayuda"
	@echo ""
	@echo "$(YELLOW)Archivos generados:$(NC)"
	@echo "  $(BIN_DIR)/transcoding     - Binario para ARM64"
	@echo "  $(LIB_DIR)/libbcg729.so*   - Librerías nativas ARM64"

# Construir la imagen Docker para ARM64
build:
	@echo "$(GREEN)🔨 Construyendo imagen Docker para ARM64...$(NC)"
	docker build -f Dockerfile-for-raspberrypi5 -t $(IMAGE_NAME) .
	@echo "$(GREEN)✅ Imagen construida: $(IMAGE_NAME)$(NC)"

# Extraer binario y librerías del contenedor
extract: build
	@echo "$(GREEN)📦 Extrayendo binario y librerías...$(NC)"
	@mkdir -p $(BIN_DIR) $(LIB_DIR)
	@echo "$(YELLOW)Creando contenedor temporal...$(NC)"
	docker create --name $(CONTAINER_NAME) $(IMAGE_NAME)
	@echo "$(YELLOW)Extrayendo binario...$(NC)"
	docker cp $(CONTAINER_NAME):/usr/local/bin/transcoding $(BIN_DIR)/
	@echo "$(YELLOW)Extrayendo librerías...$(NC)"
	docker cp $(CONTAINER_NAME):/usr/local/lib/libbcg729.so $(LIB_DIR)/
	docker cp $(CONTAINER_NAME):/usr/local/lib/libbcg729.so.0 $(LIB_DIR)/
	@echo "$(YELLOW)Limpiando contenedor temporal...$(NC)"
	docker rm $(CONTAINER_NAME)
	@echo "$(GREEN)✅ Archivos extraídos:$(NC)"
	@echo "  📁 $(BIN_DIR)/transcoding"
	@ls -la $(BIN_DIR)/
	@echo "  📁 $(LIB_DIR)/libbcg729.so*"
	@ls -la $(LIB_DIR)/
	@echo ""
	@echo "$(GREEN)🎉 ¡Listo para Raspberry Pi 5!$(NC)"
	@echo "$(YELLOW)Para usar en la Raspberry Pi:$(NC)"
	@echo "  1. Copia el directorio $(BIN_DIR)/ y $(LIB_DIR)/ a tu Raspberry Pi"
	@echo "  2. Asegúrate de que las librerías estén en el PATH o usa LD_LIBRARY_PATH"
	@echo "  3. Ejecuta: ./$(BIN_DIR)/transcoding input.wav output.g729 --format g729"

# Construir la imagen Docker para ARM64 con Debian/glibc
build-debian:
	@echo "$(GREEN)🔨 Construyendo imagen Docker para ARM64 (Debian/glibc)...$(NC)"
	docker build -f Dockerfile-for-raspberrypi-debian -t $(IMAGE_NAME_DEBIAN) .
	@echo "$(GREEN)✅ Imagen construida: $(IMAGE_NAME_DEBIAN)$(NC)"

# Extraer binario y librerías del contenedor (Debian/glibc)
extract-debian: build-debian
	@echo "$(GREEN)📦 Extrayendo binario y librerías (Debian/glibc)...$(NC)"
	@mkdir -p $(BIN_DIR) $(LIB_DIR)
	@echo "$(YELLOW)Creando contenedor temporal...$(NC)"
	docker create --name $(CONTAINER_NAME) $(IMAGE_NAME_DEBIAN)
	@echo "$(YELLOW)Extrayendo binario...$(NC)"
	docker cp $(CONTAINER_NAME):/usr/local/bin/transcoding $(BIN_DIR)/
	@echo "$(YELLOW)Extrayendo librerías...$(NC)"
	docker cp $(CONTAINER_NAME):/usr/local/lib/libbcg729.so $(LIB_DIR)/
	docker cp $(CONTAINER_NAME):/usr/local/lib/libbcg729.so.0 $(LIB_DIR)/
	@echo "$(YELLOW)Limpiando contenedor temporal...$(NC)"
	docker rm $(CONTAINER_NAME)
	@echo "$(GREEN)✅ Archivos extraídos:$(NC)"
	@echo "  📁 $(BIN_DIR)/transcoding"
	@ls -la $(BIN_DIR)/
	@echo "  📁 $(LIB_DIR)/libbcg729.so*"
	@ls -la $(LIB_DIR)/
	@echo ""
	@echo "$(GREEN)🎉 ¡Listo para Raspberry Pi OS (Debian/glibc)!$(NC)"
	@echo "$(YELLOW)Para usar en la Raspberry Pi:$(NC)"
	@echo "  1. Copia el directorio $(BIN_DIR)/ y $(LIB_DIR)/ a tu Raspberry Pi"
	@echo "  2. Asegúrate de que las librerías estén en el PATH o usa LD_LIBRARY_PATH"
	@echo "  3. Ejecuta: ./$(BIN_DIR)/transcoding input.wav output.g729 --format g729"

# Ejecutar build + extract
all: extract

# Ejecutar build + extract (Debian/glibc)
all-debian: extract-debian

# Limpiar archivos generados
clean:
	@echo "$(YELLOW)🧹 Limpiando archivos generados...$(NC)"
	@rm -rf $(BIN_DIR) $(LIB_DIR)
	@echo "$(YELLOW)Limpiando imágenes Docker...$(NC)"
	@docker rmi $(IMAGE_NAME) 2>/dev/null || true
	@docker rmi $(IMAGE_NAME_DEBIAN) 2>/dev/null || true
	@echo "$(GREEN)✅ Limpieza completada$(NC)"

# Verificar que el binario funciona
test: extract
	@echo "$(GREEN)🧪 Probando binario extraído...$(NC)"
	@if [ -f "$(BIN_DIR)/transcoding" ]; then \
		echo "$(GREEN)✅ Binario encontrado$(NC)"; \
		file $(BIN_DIR)/transcoding; \
		echo "$(YELLOW)Verificando dependencias...$(NC)"; \
		ldd $(BIN_DIR)/transcoding 2>/dev/null || echo "ldd no disponible"; \
	else \
		echo "$(RED)❌ Binario no encontrado$(NC)"; \
		exit 1; \
	fi
