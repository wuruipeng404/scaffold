# Scaffold

- Scaffolding for quickly building web services based on gin and zap

### install

```shell
go get -u github.com/wuruipeng404/scaffold@v1.5.6
```

### Usage

```golang

package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wuruipeng404/scaffold"
	"github.com/wuruipeng404/scaffold/logger"
	"github.com/wuruipeng404/scaffold/orm"
	"sync"
)

func init() {
	logger.InitLogger() // only os.stdout

	logger.InitPersistenceLogger("./some_logfile")
}

func main() {
	server := scaffold.NewGraceServer(":8000", gin.Default())

	server.AddBackgroundTask(func(ctx context.Context, wg *sync.WaitGroup) {
		// some background task like cronjob
	})

	server.AddDeferFunc(func() {
		// resource recycle
		logger.Debugf("exit")
	})
	server.Start()
}
```
