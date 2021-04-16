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

The command copies a pad with full history and chat. If force is true and the destination pad exists, it will be overwritten.

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
      --listen.addr string   Address on which to expose metrics. (default ":9012")
      --suffixes string      Suffixes to group the pads. (default "keep,temp")
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

The command checks every Pad for it’s last edited date. If it is older than the defined limit, the pad will be deleted.

Pads without any changes (revisions) will be deleted. This can happen when no content was changed in the pad
(e.g. a person misspelles a pad).
Pads will grouped by the pre-defined suffixes. Every suffix has a defined expiration time. If the pad is older than the
defined expiration time, the pad will be deleted.

Example:

`etherpad-toolkit purge --expiration "default:720h,temp:24h,keep:8760h"`

This configuration will group the pads in three clusters: default (expiration: 30 days, suffix is required!),
temp (expiration: 24 hours), keep (expiration: 365 days). If pads in the clusters older than the given expiration the
pads will be deleted.

```
Usage:
  etherpad-toolkit purge [flags]

Flags:
      --concurrency int     Concurrency for the purge process (default 4)
      --dry-run             Enable dry-run
      --expiration string   Configuration for pad expiration duration. Example: "default:720h,temp:24h,keep:8760h"
  -h, --help                help for purge
```
