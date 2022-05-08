package main

import (
	"context"
	"github.com/itsfunny/go-cell/application"
	"github.com/itsfunny/go-cell/extension/http"
	"github.com/itsfunny/go-cell/extension/oec"
	"github.com/itsfunny/go-cell/extension/swagger"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"os"
)

func main() {
	logsdk.SetGlobalLogLevel(logsdk.DebugLevel)
	app := application.New(context.Background(),
		http.HttpModule,
		swagger.SwaggerModule,
		oec.OecModule,
	)
	app.Run(os.Args)
}
