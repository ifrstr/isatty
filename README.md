# isatty

[![Go Reference](https://pkg.go.dev/badge/gopkg.ilharper.com/x/isatty.svg)](https://pkg.go.dev/gopkg.ilharper.com/x/isatty)
[![Go Report Card](https://goreportcard.com/badge/github.com/ifrstr/isatty)](https://goreportcard.com/report/github.com/ifrstr/isatty)

Determine whether the stream is a tty.

## Install

```sh
go get gopkg.ilharper.com/x/isatty
```

## Usage

```go
import "gopkg.ilharper.com/x/isatty"

if isatty.Isatty(os.Stdout.Fd()) {
    // Is a tty
}
else {
    // Isn't a tty
}
```

[![Go Reference](https://pkg.go.dev/badge/gopkg.ilharper.com/x/isatty.svg)](https://pkg.go.dev/gopkg.ilharper.com/x/isatty)

## LICENSE

[MIT](https://github.com/ifrstr/isatty/blob/master/LICENSE)
