package metrictor

import (
	"context"
	"expvar"
	"runtime"
	"time"
)

var (
	startupTime time.Time

	metrics            = expvar.NewMap("metrics")
	startupTimeRFC3339 = new(expvar.String)
	upTime             = new(expvar.Int)
	goVersion          = new(expvar.String)
	goOS               = new(expvar.String)
	goArch             = new(expvar.String)
	numGoroutine       = new(expvar.Int)
	numCPU             = new(expvar.Int)
	numCgoCall         = new(expvar.Int)
)

// Run runs metrics collector
func Run(ctx context.Context) {
	metrics.Set("goVersion", goVersion)
	metrics.Set("goOs", goOS)
	metrics.Set("goArch", goArch)
	metrics.Set("numGoroutine", numGoroutine)
	metrics.Set("numCpu", numCPU)
	metrics.Set("numCgoCall", numCgoCall)
	metrics.Set("startup", startupTimeRFC3339)
	metrics.Set("uptime", upTime)

	startupTime = time.Now()
	startupTimeRFC3339.Set(startupTime.Format(time.RFC3339))

	collect()
	ticker := time.Tick(time.Second * 5)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker:
				collect()
			}
		}
	}()
}

func collect() {
	goVersion.Set(runtime.Version())
	goOS.Set(runtime.GOOS)
	goArch.Set(runtime.GOARCH)
	numGoroutine.Set(int64(runtime.NumGoroutine()))
	numCPU.Set(int64(runtime.NumCPU()))
	numCgoCall.Set(int64(runtime.NumCgoCall()))
	upTime.Set(int64(time.Now().Sub(startupTime).Seconds()))
}
