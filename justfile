run:
  # Start the app
  go run main.go

lint:
  # Run lint.
  golangci-lint run

test:
  # Run testing.
  go test ./...

v-test:
  # Run verbose testing.
  go test ./... -v
