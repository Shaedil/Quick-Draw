PORT ?=

all:
	arduino-cli compile --fqbn esp8266:esp8266:nodemcu ./NodeMCU

clean:
	@rm -vrf ./NodeMCU/build

flash: all
ifndef PORT
	$(error make variable: [PORT] (not defined))
endif
	arduino-cli upload --verify --port $(PORT) --fqbn esp8266:esp8266:nodemcu ./NodeMCU

.PHONY: clean flash
