/*
* @Author: Rumple
* @Email: ruipeng.wu@cyclone-robotics.com
* @DateTime: 2021/8/20 17:03
 */

package scaffold

type GraceServer struct {
}

func NewGraceServer() *GraceServer {
	return &GraceServer{}
}

func (g *GraceServer) Hallo() string {
	return "111"
}

// func (g *GraceServer) Start() {
// 	defer func() {
// 		db.Release()
// 	}()
//
// 	ctx, cancel := context.WithCancel(context.Background())
// 	wg := new(sync.WaitGroup)
//
// 	backgroundListener(ctx, wg)
//
// 	server := &http.Server{
// 		Addr:    fmt.Sprintf("0.0.0.0:%d", config.Config.ContainerPort),
// 		Handler: routers.InitEngine(),
// 	}
//
// 	go func() {
// 		logs.Info("Server Start ...")
// 		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			logs.FatalF("Listen Error : %s", err.Error())
// 		}
// 	}()
//
// 	quit := make(chan os.Signal)
// 	signal.Notify(
// 		quit,
// 		syscall.SIGHUP,
// 		syscall.SIGINT,
// 		syscall.SIGTERM,
// 		syscall.SIGQUIT,
// 		syscall.SIGKILL,
// 		// syscall.SIGTSTP,
// 	)
// 	<-quit
// 	logs.Info("Shutdown Server ...")
// 	// 退出后台的一些协程
// 	cancel()
// 	// 等待携程处理完毕
// 	wg.Wait()
//
// 	if err := server.Shutdown(context.TODO()); err != nil {
// 		logs.FatalF("Server Shutdown Error : %s ", err.Error())
// 	} else {
// 		logs.Info("Server Exiting ...")
// 	}
// }
