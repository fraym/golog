# golog

This package provides an abstraction layer for logging in your application.

You can use the `Logger` interface whereever you need logging. The used implementation is simple to be changed. Just instantiate one of the built in loggers or create your own implementation of the `Logger` interface:

- `NewZerologLogger`: Instantiates a [zerolog](https://github.com/rs/zerolog) logger
- `NewLogrusLogger`: Instantiates a [logrus](https://github.com/sirupsen/logrus) logger

## usage

```go
logger := golog.NewZerologLogger()

logger.Debug().Write("your log message")

logger.Debug().WithFields(map[string]any{
	"field": "value",
}).Write("your log message")

logger.Debug().WithError(fmt.Errorf("your error")).Write("your log message")
```
