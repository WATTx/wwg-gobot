ARDUINO_IDE=/Applications/Arduino.app/Contents/MacOS/Arduino
FIRMATA_INO=./firmata/firmata.ino
GO_DIRS=publisher wemos

.PHONY: upload build $(GO_DIRS)

build: $(GO_DIRS)

$(GO_DIRS):
	@echo "Building '$@'"
	@cd "$@" && go build

upload:
	$(ARDUINO_IDE) $(FIRMATA_INO) &
