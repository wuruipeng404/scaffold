/*
* @Author: Rumple
* @Email: ruipeng.wu@cyclone-robotics.com
* @DateTime: 2021/8/20 17:03
 */

package scaffold

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/wuruipeng404/scaffold/logger"
	"go.uber.org/zap"
)

type GraceServer struct {
	listen string
	engine *gin.Engine
	log    *zap.SugaredLogger
}

func NewGraceServer(listen string, engine *gin.Engine, log *zap.SugaredLogger) *GraceServer {

	if log == nil {
		log = logger.DefaultLogger()
	}

	return &GraceServer{
		listen: listen,
		log:    log,
		engine: engine,
	}
}

// Start 没有stop方法 因为没必要
func (g *GraceServer) Start() {

	defer func() {
		_ = g.log.Sync()
	}()

	server := &http.Server{
		Addr:    g.listen,
		Handler: g.engine,
	}

	go func() {
		g.log.Info("Server Start ...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			g.log.Fatalf("listen Error : %s", err.Error())
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(
		quit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL,
		// syscall.SIGTSTP,
	)
	<-quit
	g.log.Info("Shutdown Server ...")

	if err := server.Shutdown(context.TODO()); err != nil {
		g.log.Fatalf("Server Shutdown Error : %s ", err.Error())
	} else {
		g.log.Info("Server Exiting ...")
	}
}

// func (g *GraceServer) Stop() {
// 	g.signChan <- syscall.SIGTERM // same with docker stop
// }
