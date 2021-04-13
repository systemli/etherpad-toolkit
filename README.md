# etherpad-toolkit

[![Integration](https://github.com/systemli/etherpad-toolkit/actions/workflows/integration.yml/badge.svg)](https://github.com/systemli/etherpad-toolkit/actions/workflows/integration.yml) [![Quality](https://github.com/systemli/etherpad-toolkit/actions/workflows/quality.yml/badge.svg)](https://github.com/systemli/etherpad-toolkit/actions/workflows/quality.yml)

**Etherpad Toolkit is a collection for most common [Etherpad](https://github.com/ether/etherpad-lite) maintenance tasks.**

```
Usage:
  etherpad-toolkit [command]

Available Commands:
  delete-pad  Removes a single Pad
  help        Help about any command
  metrics     Serves Pad related metrics
  move-pad    Moves a single Pad
  purge       Removes old Pads entirely from Etherpad

Flags:
      --etherpad.apikey string   API Key for Etherpad
      --etherpad.url string      URL to access Etherpad (default "http://localhost:9001")
  -h, --help                     help for etherpad-toolkit
      --log.format string        Format for log output (default "text")
      --log.level string         Log level (default "info")

Use "etherpad-toolkit [command] --help" for more information about a command.
```
