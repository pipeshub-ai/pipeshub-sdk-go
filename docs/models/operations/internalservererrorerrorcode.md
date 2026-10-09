# InternalServerErrorErrorCode

Machine-readable error code.

- `HTTP_INTERNAL_SERVER_ERROR` — explicit
  `InternalServerError` raised by the handler.
- `INTERNAL_ERROR` — unhandled exception
  caught by the global error middleware.
- `MIDDLEWARE_ERROR` — the error middleware
  itself failed while serializing the
  response.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.InternalServerErrorErrorCodeHTTPInternalServerError

// Open enum: custom values can be created with a direct type cast
custom := operations.InternalServerErrorErrorCode("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `InternalServerErrorErrorCodeHTTPInternalServerError` | HTTP_INTERNAL_SERVER_ERROR                            |
| `InternalServerErrorErrorCodeInternalError`           | INTERNAL_ERROR                                        |
| `InternalServerErrorErrorCodeMiddlewareError`         | MIDDLEWARE_ERROR                                      |