.PHONY: help generate clean install-tools release release-patch release-minor release-major version

GOPATH_BIN := $(shell go env GOPATH)/bin

help:
	@echo "Доступные команды:"
	@echo "  make generate      - Генерация Go кода из proto файлов"
	@echo "  make clean         - Удаление сгенерированного кода"
	@echo "  make install-tools - Установка необходимых инструментов"
	@echo "  make version       - Показать текущую версию (из git tag)"
	@echo "  make release       - Автоматический релиз (patch версия: v1.0.0 -> v1.0.1)"
	@echo "  make release-patch - Релиз patch версии (v1.0.0 -> v1.0.1)"
	@echo "  make release-minor - Релиз minor версии (v1.0.0 -> v1.1.0)"
	@echo "  make release-major - Релиз major версии (v1.0.0 -> v2.0.0)"
	@echo "  make release VER=v1.0.0 - Релиз конкретной версии"

# Установка инструментов для генерации proto
install-tools:
	@echo "Установка protoc-gen-go и protoc-gen-go-grpc..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "Установлено"

# Генерация Go кода из всех proto файлов
generate:
	@echo "Генерация proto файлов..."
	@mkdir -p generated
	
	@echo "Генерация common types..."
	PATH=$(GOPATH_BIN):$$PATH protoc --go_out=generated --go_opt=paths=source_relative \
		--proto_path=proto \
		proto/common/types.proto
	
	@echo "Генерация user models..."
	PATH=$(GOPATH_BIN):$$PATH protoc --go_out=generated --go_opt=paths=source_relative \
		--proto_path=proto \
		proto/users_service/v1/user.proto
	
	@echo "Генерация auth service..."
	PATH=$(GOPATH_BIN):$$PATH protoc --go_out=generated --go_opt=paths=source_relative \
		--go-grpc_out=generated --go-grpc_opt=paths=source_relative \
		--proto_path=proto \
		proto/users_service/v1/auth_service.proto
	
	@echo "Генерация user service..."
	PATH=$(GOPATH_BIN):$$PATH protoc --go_out=generated --go_opt=paths=source_relative \
		--go-grpc_out=generated --go-grpc_opt=paths=source_relative \
		--proto_path=proto \
		proto/users_service/v1/user_service.proto
	
	@echo "Генерация admin service..."
	PATH=$(GOPATH_BIN):$$PATH protoc --go_out=generated --go_opt=paths=source_relative \
		--go-grpc_out=generated --go-grpc_opt=paths=source_relative \
		--proto_path=proto \
		proto/users_service/v1/admin_service.proto

	@echo "Генерация test service..."
	PATH=$(GOPATH_BIN):$$PATH protoc --go_out=generated --go_opt=paths=source_relative \
		--go-grpc_out=generated --go-grpc_opt=paths=source_relative \
		--proto_path=proto \
		proto/users_service/v1/test.proto
	
	@echo "✅ Генерация завершена! Файлы в generated/"

# Очистка сгенерированного кода
clean:
	@echo "Удаление сгенерированных файлов..."
	rm -rf generated/
	@echo "✅ Очистка завершена!"

# Показать текущую версию
version:
	@git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0"

# Получить следующую patch версию (v1.0.0 -> v1.0.1)
get-next-patch:
	@LAST_TAG=$$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0"); \
	VERSION=$$(echo $$LAST_TAG | sed 's/v//'); \
	MAJOR=$$(echo $$VERSION | cut -d. -f1); \
	MINOR=$$(echo $$VERSION | cut -d. -f2); \
	PATCH=$$(echo $$VERSION | cut -d. -f3); \
	NEXT_PATCH=$$((PATCH + 1)); \
	echo "v$${MAJOR}.$${MINOR}.$${NEXT_PATCH}"

# Получить следующую minor версию (v1.0.0 -> v1.1.0)
get-next-minor:
	@LAST_TAG=$$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0"); \
	VERSION=$$(echo $$LAST_TAG | sed 's/v//'); \
	MAJOR=$$(echo $$VERSION | cut -d. -f1); \
	MINOR=$$(echo $$VERSION | cut -d. -f2); \
	NEXT_MINOR=$$((MINOR + 1)); \
	echo "v$${MAJOR}.$${NEXT_MINOR}.0"

# Получить следующую major версию (v1.0.0 -> v2.0.0)
get-next-major:
	@LAST_TAG=$$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0"); \
	VERSION=$$(echo $$LAST_TAG | sed 's/v//'); \
	MAJOR=$$(echo $$VERSION | cut -d. -f1); \
	NEXT_MAJOR=$$((MAJOR + 1)); \
	echo "v$${NEXT_MAJOR}.0.0"

# Автоматический релиз (patch версия по умолчанию)
release: release-patch

# Релиз patch версии (v1.0.0 -> v1.0.1)
release-patch:
	@VER=$$($(MAKE) -s get-next-patch); \
	$(MAKE) do-release VER=$$VER

# Релиз minor версии (v1.0.0 -> v1.1.0)
release-minor:
	@VER=$$($(MAKE) -s get-next-minor); \
	$(MAKE) do-release VER=$$VER

# Релиз major версии (v1.0.0 -> v2.0.0)
release-major:
	@VER=$$($(MAKE) -s get-next-major); \
	$(MAKE) do-release VER=$$VER

# Внутренняя команда для выполнения релиза
do-release:
	@if [ -z "$(VER)" ]; then \
		echo "❌ Ошибка: версия не указана"; \
		exit 1; \
	fi
	@echo "🚀 Создание релиза $(VER)..."
	@echo "1. Генерация кода..."
	@$(MAKE) generate
	@echo "2. Проверка изменений..."
	@if git diff --quiet 2>/dev/null && git diff --cached --quiet 2>/dev/null; then \
		echo "   ⚠️  Нет изменений для коммита"; \
	else \
		echo "3. Добавление файлов в git..."; \
		git add .; \
		echo "4. Создание коммита..."; \
		git commit -m "Release $(VER)" || true; \
		echo "5. Создание тега..."; \
		git tag -a $(VER) -m "Release $(VER)" || true; \
		echo "6. Отправка на удаленный репозиторий..."; \
		git push origin main --tags || echo "   ⚠️  Не удалось отправить (проверьте настройки git)"; \
		echo "✅ Релиз $(VER) создан и отправлен!"; \
	fi