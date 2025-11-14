.PHONY: help generate clean install-tools release version

GOPATH_BIN := $(shell go env GOPATH)/bin

help:
	@echo "Доступные команды:"
	@echo "  make generate      - Генерация Go кода из proto файлов"
	@echo "  make clean         - Удаление сгенерированного кода"
	@echo "  make install-tools - Установка необходимых инструментов"
	@echo "  make version       - Показать текущую версию (из git tag)"
	@echo "  make release VER=v1.0.0 - Создать релиз (генерация + коммит + тег)"

# Установка инструментов для генерации proto
install-tools:
	@echo "Установка protoc-gen-go и protoc-gen-go-grpc..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "Установлено"

# Генерация Go кода из всех proto файлов
generate:
	@echo "Генерация proto файлов..."
	@mkdir -p generated/go
	
	@echo "Генерация common types..."
	PATH=$(GOPATH_BIN):$$PATH protoc --go_out=generated/go --go_opt=paths=source_relative \
		--proto_path=proto \
		proto/common/types.proto
	
	@echo "Генерация auth service..."
	PATH=$(GOPATH_BIN):$$PATH protoc --go_out=generated/go --go_opt=paths=source_relative \
		--go-grpc_out=generated/go --go-grpc_opt=paths=source_relative \
		--proto_path=proto \
		proto/auth/v1/auth.proto
	
	@echo "✅ Генерация завершена! Файлы в generated/go/"

# Очистка сгенерированного кода
clean:
	@echo "Удаление сгенерированных файлов..."
	rm -rf generated/
	@echo "✅ Очистка завершена!"

# Показать текущую версию
version:
	@git describe --tags --always --dirty 2>/dev/null || echo "Версия не определена (нет git тегов)"

# Создать новый релиз
# Использование: make release VER=v1.0.0
release:
	@if [ -z "$(VER)" ]; then \
		echo "❌ Ошибка: укажите версию. Пример: make release VER=v1.0.0"; \
		exit 1; \
	fi
	@echo "🚀 Создание релиза $(VER)..."
	@echo "1. Генерация кода..."
	@$(MAKE) generate
	@echo "2. Проверка git статуса..."
	@if ! git diff --quiet generated/ 2>/dev/null; then \
		echo "   Обнаружены изменения в generated/"; \
	fi
	@echo "3. Создание git тега $(VER)..."
	@echo "   Выполните вручную:"
	@echo "   git add ."
	@echo "   git commit -m 'Release $(VER)'"
	@echo "   git tag -a $(VER) -m 'Release $(VER)'"
	@echo "   git push origin main --tags"
	@echo "✅ Готово к релизу! Выполните команды выше."


