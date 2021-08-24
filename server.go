/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/8/20 17:03
 */

package scaffold

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GraceServer struct {
	listen string
	engine *gin.Engine
	log    *zap.SugaredLogger

	// background tasks
	tasks      []func(ctx context.Context, wg *sync.WaitGroup)
	deferTasks []func()
	cancelFunc context.CancelFunc
}

func NewGraceServer(listen string, engine *gin.Engine, log *zap.SugaredLogger) *GraceServer {

	return &GraceServer{
		listen: listen,
		log:    log,
		engine: engine,
	}
}

// AddBackgroundTask append background goroutine and when server exit will wait then
// the task finish need call wg.Done()
func (g *GraceServer) AddBackgroundTask(tasks ...func(ctx context.Context, wg *sync.WaitGroup)) {
	g.tasks = append(g.tasks, tasks...)
}

func (g *GraceServer) AddDeferFunc(tasks ...func()) {
	g.deferTasks = append(g.deferTasks, tasks...)
}

// Start 没有stop方法 因为没必要
func (g *GraceServer) Start() {

	defer func() {

		if len(g.deferTasks) > 0 {
			for _, t := range g.deferTasks {
				t()
			}
		}
		_ = g.log.Sync()

	}()

	server := &http.Server{
		Addr:    g.listen,
		Handler: g.engine,
	}

	ctx, cancel := context.WithCancel(context.Background())
	g.cancelFunc = cancel

	wg := new(sync.WaitGroup)

	if len(g.tasks) > 0 {
		for _, task := range g.tasks {
			wg.Add(1)
			go task(ctx, wg)
		}
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

	if len(g.tasks) > 0 {
		// cancel background task
		g.cancelFunc()
		// and wait exit
		wg.Wait()
	}

	if err := server.Shutdown(context.TODO()); err != nil {
		g.log.Fatalf("Server Shutdown Error : %s ", err.Error())
	} else {
		g.log.Info("Server Exiting ...")
	}
}
