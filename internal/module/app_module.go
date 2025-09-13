package module

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var ServerModule = fx.Module("http",
	fx.Provide(func() *gin.Engine {
		r := gin.Default()

		r.Use(cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:    []string{"*"},
			ExposeHeaders:   []string{"*"},
			MaxAge:          12 * time.Hour,
		}))

		return r
	}),
	fx.Provide(func(r *gin.Engine) *http.Server {
		return &http.Server{
			Addr:    ":8080",
			Handler: r,
		}
	}),
	fx.Invoke(func(lc fx.Lifecycle, srv *http.Server) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
				defer cancel()
				return srv.Shutdown(ctx)
			},
		})
	}),
)
