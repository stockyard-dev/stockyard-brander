# Stockyard Brander

**Email signature manager — define a template once, everyone gets the same footer**

Part of the [Stockyard](https://stockyard.dev) family of self-hosted developer tools.

## Quick Start

```bash
docker run -p 9180:9180 -v brander_data:/data ghcr.io/stockyard-dev/stockyard-brander
```

Or with docker-compose:

```bash
docker-compose up -d
```

Open `http://localhost:9180` in your browser.

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `9180` | HTTP port |
| `DATA_DIR` | `./data` | SQLite database directory |
| `BRANDER_LICENSE_KEY` | *(empty)* | Pro license key |

## Free vs Pro

| | Free | Pro |
|-|------|-----|
| Limits | 1 template, 10 members | Unlimited templates and members |
| Price | Free | $1.99/mo |

Get a Pro license at [stockyard.dev/tools/](https://stockyard.dev/tools/).

## Category

Operations & Teams

## License

Apache 2.0
