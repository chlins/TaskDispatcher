# A very simple package for generate random id

## Usage

```go
// default length 8
id, err := simpleid.New()

// specify length
id, err := simpleid.NewWithLength(8)

```