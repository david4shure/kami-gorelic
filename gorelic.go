package kami_gorelic

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/context"

	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/gorelic"
)

var agent *gorelic.Agent

func Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) context.Context {
	startTime := time.Now()
	if agent != nil {
		agent.HTTPTimer.UpdateSince(startTime)
	}
	return ctx
}

func InitNewrelicAgent(license string, appname string, verbose bool) error {

	if license == "" {
		return fmt.Errorf("Please specify NewRelic license")
	}

	agent = gorelic.NewAgent()
	agent.NewrelicLicense = license

	agent.HTTPTimer = metrics.NewTimer()
	agent.CollectHTTPStat = true
	agent.Verbose = verbose

	agent.NewrelicName = appname
	agent.Run()
	return nil
}
