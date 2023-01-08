# isatty

[![Go Reference](https://pkg.go.dev/badge/gopkg.ilharper.com/x/isatty.svg)](https://pkg.go.dev/gopkg.ilharper.com/x/isatty)
[![Go Report Card](https://goreportcard.com/badge/github.com/ifrstr/isatty)](https://goreportcard.com/report/github.com/ifrstr/isatty)

Determine whether the stream is a tty.

## Features

- Supports most platforms

- Supports force override using environment variable `FORCE_TTY`

## Install

```sh
go get gopkg.ilharper.com/x/isatty
```

## Usage

```go
import "gopkg.ilharper.com/x/isatty"

if isatty.Isatty(os.Stdout.Fd()) {
	// TTY
} else {
	// Not TTY
}
```

[![Go Reference](https://pkg.go.dev/badge/gopkg.ilharper.com/x/isatty.svg)](https://pkg.go.dev/gopkg.ilharper.com/x/isatty)

## Using `FORCE_TTY`

User can override the result using the environment variable `FORCE_TTY`.

```ini
# Force TTY
FORCE_TTY=
FORCE_TTY=1
FORCE_TTY=true

# Force not TTY
FORCE_TTY=0
FORCE_TTY=false
```

## Test

To determine whether the current `stdout` is a tty:

```sh
go test
```

## Defects on Windows

See #3.

## LICENSE

[MIT](https://github.com/ifrstr/isatty/blob/master/LICENSE)
