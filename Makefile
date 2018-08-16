ARDUINO_IDE=/Applications/Arduino.app/Contents/MacOS/Arduino
FIRMATA_INO=./firmata/firmata.ino
GO_DIRS=publisher wemos

.PHONY: all upload build $(GO_DIRS)

upload:
	$(ARDUINO_IDE) $(FIRMATA_INO) &

build: $(GO_DIRS)

$(GO_DIRS):
	@echo "Building '$@'"
	@cd "$@" && go build
