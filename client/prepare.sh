#!/usr/bin/env bash

set -euo pipefail

command -v arduino-cli > /dev/null || (echo "Error: arduino-cli not in PATH" && exit 1)

arduino-cli core update-index --additional-urls https://arduino.esp8266.com/stable/package_esp8266com_index.json
arduino-cli core install esp8266:esp8266

echo "All done."
