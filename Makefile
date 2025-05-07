DEBUG_FLAGS := RAS_DEBUG=true
MOCK_FLAGS := RAS_MOCK=true
RAS_FLAGS ?=

EXECUTABLE := ./build/srv

.PHONY: run debug mock build prepare install

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


install: build
	mkdir -p /opt/ras/
	cp $(EXECUTABLE) /opt/ras/ras
	cp ./ras.service /lib/systemd/system/ras.service
	mkdir -p /etc/ras/
	cp ./config.yml /etc/ras/config.yml
	chmod 775 /etc/ras/config.yml

run-service: install
	systemctl start ras.service || journalctl -xeu ras.service
