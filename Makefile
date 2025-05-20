DEBUG_FLAGS := RAS_DEBUG=true
MOCK_FLAGS := RAS_MOCK=true
RAS_FLAGS ?=

EXECUTABLE := ./build/srv

.PHONY: run debug mock build prepare install enable-service disable-service update-service test

prepare:
	# git pull --rebase

build: prepare 
	go build -o $(EXECUTABLE) ./cmd/server/main.go
	chmod 775 $(EXECUTABLE)
	
run: build
	$(RAS_FLAGS) ./build/srv


debug: RAS_FLAGS += $(DEBUG_FLAGS)

test: RAS_FLAGS += $(DEBUG_FLAGS)
test: RAS_FLAGS += $(MOCK_FLAGS)

test debug: build run

#=== SERVICE

SERVICE_NAME := ras.service

install: build
	mkdir -p /opt/ras/
	cp $(EXECUTABLE) /opt/ras/ras
	cp ./$(SERVICE_NAME) /lib/systemd/system/$(SERVICE_NAME)
	mkdir -p /etc/ras/
	cp ./config.yml /etc/ras/config.yml
	chmod 775 /etc/ras/config.yml

enable-service: install
	systemctl enable --now $(SERVICE_NAME) || journalctl -xeu $(SERVICE_NAME)

disable-service:
	systemctl disable --now $(SERVICE_NAME)

update-service: disable-service install enable-service
