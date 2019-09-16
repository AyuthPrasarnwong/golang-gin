package newrelic

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent"
)

const (
	// NewRelicTxnKey is the key used to retrieve the NewRelic Transaction from the context
	NewRelicTxnKey = "NewRelicTxnKey"
)

func Setup(router *gin.Engine)  {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := newrelic.NewConfig(os.Getenv("NEW_RELIC_APP_NAME"), os.Getenv("NEW_RELIC_LICENSE_KEY"))
	config.HostDisplayName = "tcg-api-bs.eggsmartpos.com"
	config.Labels = map[string]string{
		"Project": "tcg",
		"Environment": os.Getenv("APP_ENV"),
		"Server": "on-premise",
		"Tier": "backend",
		"Language": "golang",
	}

	app, err := newrelic.NewApplication(config)
	if err != nil {
		log.Printf("failed to make new_relic app: %v", err)
	} else {
		router.Use(NewRelicMonitoring(app))
	}
}

// NewRelicMonitoring is a middleware that starts a newrelic transaction, stores it in the context, then calls the next handler
func NewRelicMonitoring(app newrelic.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		txn := app.StartTransaction(ctx.Request.URL.Path, ctx.Writer, ctx.Request)
		defer txn.End()
		ctx.Set(NewRelicTxnKey, txn)
		ctx.Next()
	}
}
