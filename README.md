# etherpad-toolkit

[![Integration](https://github.com/systemli/etherpad-toolkit/actions/workflows/integration.yml/badge.svg)](https://github.com/systemli/etherpad-toolkit/actions/workflows/integration.yml) [![Quality](https://github.com/systemli/etherpad-toolkit/actions/workflows/quality.yml/badge.svg)](https://github.com/systemli/etherpad-toolkit/actions/workflows/quality.yml) [![Release](https://github.com/systemli/etherpad-toolkit/actions/workflows/release.yml/badge.svg)](https://github.com/systemli/etherpad-toolkit/actions/workflows/release.yml)

**Etherpad Toolkit is a collection for most common [Etherpad](https://github.com/ether/etherpad-lite) maintenance tasks.**

```
Usage:
  etherpad-toolkit [command]

Available Commands:
  copy-pad    Copies a single Pad
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

## Docker

You can run the etherpad-toolkit with Docker

```
docker run systemli/etherpad-toolkit:latest --help
```

## Commands

### Copy Pad

The command copies a pad with full history and chat.. If force is true and the destination pad exists, it will be overwritten.

```
Usage:
  etherpad-toolkit copy-pad [sourceID] [destinationID] [flags]

Flags:
      --force   If set and the destination pad exists, it will be overwritten.
  -h, --help    help for copy-pad

```

### Delete Pad

The command removes a single pad entirely from Etherpad.

```
Usage:
  etherpad-toolkit delete-pad [pad] [flags]

Flags:
  -h, --help   help for delete-pad

```

### Metrics

The Command serves the count of pads grouped by suffix in Prometheus format.

```
Usage:
  etherpad-toolkit metrics [flags]

Flags:
  -h, --help                 help for metrics
      --listen.addr string    (default ":9012")
```

### Move Pad

The command moves a single pad. If force is true and the destination pad exists, it will be overwritten.

```
Usage:
  etherpad-toolkit move-pad [sourceID] [destinationID] [flags]

Flags:
      --force   If set and the destination pad exists, it will be overwritten.
  -h, --help    help for move-pad
```

### Purge

The command checks every Pad if the last edited date is older than the defined limit. Older Pads will be deleted.

Pads without a suffix will be deleted after 30 days of inactivity.  
Pads with the suffix "-temp" will be deleted after 24 hours of inactivity.  
Pads with the suffix "-keep" will be deleted after 365 days of inactivity.  

```
Usage:
  etherpad-toolkit purge [flags]

Flags:
      --concurrency int   Concurrency for the purge process (default 4)
      --dry-run           Enable dry-run
  -h, --help              help for purge
```
