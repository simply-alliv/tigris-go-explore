# tigris-go-explore

# Local Development

## 1. Configure environment variables

- Run `cp .env.example .env` to create a configuration file for your environment variables.
- Set `TIGRIS_CLIENT_ID`, `TIGRIS_CLIENT_SECRET`, and `TIGRIS_PROJECT` environment variables with your Tigris credentials for the Go SDK.

## 2. Seed Tigris database with test data

- Set `SEED_DATA` to true.

```
go run main.go
```

## 3. Start local server

- Set `SEED_DATA` to false, or leave it empty.

```
go run main.go
```