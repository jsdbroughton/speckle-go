# Speckle Go

This is the Go implementation of the Speckle SDK.

## Project Structure

- `cmd/speckle`: Contains the main application entry point
- `internal`: Private application and library code
  - `api`: API-related code
  - `core`: Core functionality
  - `logging`: Logging utilities
  - `objects`: Object definitions
  - `serialization`: Serialization utilities
  - `transports`: Transport layer implementations
- `pkg`: Library code that's ok to use by external applications
- `examples`: Example code
- `tests`: Additional external tests
