run:
  # Start the app
  just esbuild
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

gen:
  # Run code generation.
  just esbuild
  tygo generate

esbuild:
  esbuild ./frontend/src/main.ts --bundle --minify --outfile=./frontend/js/index.min.js
