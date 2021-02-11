# Logger

Self logger provides a layer on top of [zerolog](https://github.com/rs/zerolog) to initialize it and provide helper functions to easily manage your logging.


## Usage

```go
package main

import (
	"context"

	"github.com/joinself/logger"
)

func main() {
	logger.SetGlobalLevel(logger.DEBUG)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "session_id", "190827-30298-901283029-1232131")
	ctx = context.WithValue(ctx, "self_id", "1112223334")

	logger.Debug().Context(ctx).Msg("Yo!")
	// {"level":"debug","severity":100,"session_id":"190827-30298-901283029-1232131","self_id":"1112223334","timestamp":1613043647,"message":"Yo!"}

	logger.Info().Msg("Yo!")
	// {"level":"info","severity":200,"timestamp":1613043647,"message":"Yo!"}}

	logger.Info().Msg("Yo! %s", "Tom")
	// {"level":"info","severity":200,"timestamp":1613044425,"message":"Yo! Tom"}
}

```

