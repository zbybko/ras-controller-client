DEBUG_FLAGS := RAS_DEBUG=true
MOCK_FLAGS := RAS_MOCK=true
RAS_FLAGS ?=

.PHONY: run debug mock
	
run:
	$(RAS_FLAGS) go run ./cmd/server/main.go


debug: RAS_FLAGS += $(DEBUG_FLAGS)

test: RAS_FLAGS += $(DEBUG_FLAGS)
test: RAS_FLAGS += $(MOCK_FLAGS)

test debug: run
