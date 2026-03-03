# build:
# 	@echo "Compilando para Linux..."
# 	GOOS=linux GOARCH=amd64 go build -o build/ducky-linux
# 	@echo "Compilando para Windows..."
# 	GOOS=windows GOARCH=amd64 go build -o build/ducky.exe
# 	@echo "¡Listo! Revisa la carpeta /build"

	
BINARY_NAME=ducky
VERSION=0.1.0-beta
BUILD_DIR=dist
LDFLAGS=-ldflags="-s -w -X 'main.version=$(VERSION)'"

.PHONY: build clean

build: clean
	@mkdir -p $(BUILD_DIR)
	@echo "🏗️  Construyendo binarios para distribución (v$(VERSION))..."
	
	@echo "🐧 Linux amd64..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)
	
	@echo "🪟 Windows amd64..."
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME).exe
	
	@echo "✅ Archivos listos en la carpeta /$(BUILD_DIR)"

clean:
	@rm -rf $(BUILD_DIR)