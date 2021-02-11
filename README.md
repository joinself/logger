# Logger

Self logger provides a layer on top of [zerolog](https://github.com/rs/zerolog) to initialize it and provide helper functions to easily manage your logging.


## Usage

```go
package main

import (
	"github.com/joinself/logger"
)

func main() {
	logger.SetGlobalLevel(logger.DEBUG)

	logger.Debug(logger.L{"Yo!", "190827-30298-901283029-1232131", "1112223334"})
	// {"severity":100,"session_id":"190827-30298-901283029-1232131","self_id":"1112223334","timestamp":1612966902,"message":"Yo!"}

	logger.Info(logger.L{Body: "Yo!"})
	// {"severity":100,"timestamp":1612966919,"message":"Yo!"}
}
```

