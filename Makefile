DEBUG_FLAGS := RAS_DEBUG=true

.PHONY: debug

debug:
	$(DEBUG_FLAGS) go run ./cmd/server/main.go
