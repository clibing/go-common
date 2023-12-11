package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/clibing/go-common/api"
	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"
	"github.com/clibing/go-common/internal/pkg/config"
	"github.com/clibing/go-common/internal/pkg/log"
	"github.com/clibing/go-common/internal/pkg/model"
	"github.com/urfave/cli/v2"
)

func server() *cli.Command {

	return &cli.Command{
		Name: "server",
		Flags: []cli.Flag{
			configFlag(),
		},

		Action: func(c *cli.Context) error {
			input := c.String("config")

			cfg, err := config.LoadConfig(input)
			if err != nil {
				fmt.Println("加载配置异常: ", input)
				return nil
			}

			// logger
			logger := log.Init(cfg.Log.Level)

			// cli.ShowAppHelp(c)
			router := gin.New()
			router.Use(ginlogrus.Logger(logger), gin.Recovery())
			// Set a lower memory limit for multipart forms (default is 32 MiB)
			router.MaxMultipartMemory = 8 << 20 // 8 MiB

			err = api.BindHandle(&model.Context{
				Router: router,
				Logger: logger,
			})
			if err != nil {
				logger.Error("gin bind hander error", err)
				return err
			}

			srv := &http.Server{
				Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
				Handler: router,
			}

			go func() {
				// service connections
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Fatalf("listen: %s\n", err)
				}
			}()

			// Wait for interrupt signal to gracefully shutdown the server with
			// a timeout of 3 seconds.
			// the channel used with signal.Notify should be buffered
			// fix https://blog.wu-boy.com/2021/03/why-use-buffered-channel-in-signal-notify/
			quit := make(chan os.Signal, 1)
			// kill (no param) default send syscanll.SIGTERM
			// kill -2 is syscall.SIGINT
			// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			logger.Info("Shutdown Server ...")

			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				logger.Fatalf("Server Shutdown: %s", err)
			}
			// catching ctx.Done(). timeout of 5 seconds.
			select {
			case <-ctx.Done():
				logger.Info("timeout of 3 seconds.")
			case <-time.After(10 * time.Millisecond):
				logger.Info("waiting & receive signal")
			}

			logger.Info("Server exit.")


			return nil
		},
		Subcommands: []*cli.Command{},
	}
}
