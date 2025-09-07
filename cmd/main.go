package main

import (
	"context"
	"log"

	"github.com/TruongHoang2004/ngoclam-zmp-backend/internal/module"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.NopLogger, // Disable Fx's own logging
		module.ConfigModule(),
		module.DatabaseModule,
		module.ServerModule,
		module.SwaggerModule,
		module.CategoryModule,
		module.ProductModule,
		module.ImageModule,
		module.RouterModule,

		// Hook cho log start/stop toàn app
		fx.Invoke(func(lc fx.Lifecycle) {
			lc.Append(fx.Hook{
				OnStart: OnStart,
				OnStop:  OnStop,
			})
		}),
	)

	// Đây sẽ block cho đến khi nhận tín hiệu OS (Ctrl+C, docker stop)
	app.Run()
}

func OnStart(ctx context.Context) error {
	log.Println("Application is starting...")
	log.Println("Swagger documentation available at: http://localhost:8080/swagger/index.html")
	return nil
}

func OnStop(ctx context.Context) error {
	log.Println("Application is stopping...")
	return nil
}
