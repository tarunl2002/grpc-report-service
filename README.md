# gRPC Report Generation Service (Task Submission)

This is a Go project I built as part of a task. The service uses gRPC and a cron job to generate reports for some users every 10 seconds. Everything is stored in-memory, and there's also a basic health check endpoint.

## What It Does

- gRPC endpoint: `GenerateReport(UserID)` returns a report ID.
- Every 10 seconds, a few hardcoded user IDs get their reports generated automatically (via cron).
- Reports are stored in a map (so it's all in memory).
- All actions are logged with timestamps.
- There's a `HealthCheck()` gRPC endpoint too.

## Tech Used

- Go 1.24.4
- gRPC
- Protocol Buffers
- robfig/cron for scheduling
- Go’s built-in `log` package

## How to Run

1. Clone the repo.
2. Run `go mod tidy`
3. Generate the gRPC code:
4. Start the server:

## Testing the gRPC Endpoints

You can use `grpcurl` to test:

Generate report:

Health check:
proto/ - protobuf file
server/ - gRPC logic
cron/ - cron job code
main.go - app entry point
Dockerfile - for container build

## Deployment

I also deployed it using Render (Go 1.24 image). It's running fine and you can test gRPC endpoints if needed.

## Notes

- No database used, just in-memory maps.
- Logs print to the console.

