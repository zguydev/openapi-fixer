# openapi-fixer
> A powerful tool to fix OpenAPI spec to ensure compatibility with various code generators and tools

[![Go Reference](https://pkg.go.dev/badge/github.com/zguydev/openapi-fixer.svg)](https://pkg.go.dev/github.com/zguydev/openapi-fixer)
[![Go Report Card](https://goreportcard.com/badge/github.com/zguydev/openapi-fixer?style=flat-square)](https://goreportcard.com/report/github.com/zguydev/openapi-fixer)

## Introduction
`openapi-fixer` is a Go-based tool designed to help developers automatically apply fixups to OpenAPI specification files. It provides tooling to modify OpenAPI 3.0 specification file to ensure it's compatibility with various code generators and tools while maintaining the integrity of the API specification.

## Install
```shell
go install github.com/zguydev/openapi-fixer@latest
```

### Usage
```go
//go:generate go run github.com/zguydev/openapi-fixer openapi.yaml fixed.openapi.yaml --fixups ./fixups/ --config .openapi-fixer.yaml
```

## Examples
Explore ready-to-use examples:

| Example Name         | Description                   | Path                                    |
| -------------------- | ----------------------------- | --------------------------------------- |
| 🤖 **OpenAI Example** | Fixups for the OpenAI API schema on example of `POST /chat/completions` endpoint | [`examples/OpenAI`](./examples/OpenAI/) |

## License
This project is licensed under the terms of the Apache License 2.0. See the [LICENSE](./LICENSE) file for details.
