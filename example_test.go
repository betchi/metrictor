package metrictor_test

import (
	"context"
	"time"

	"github.com/betchi/metrictor"
)

func ExampleSetExpVarString() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	metrictor.SetString(metrictor.OneTime, "stringKey", func() string {
		return "value"
	})
	metrictor.SetInt(metrictor.EachTime, "intKey", func() int64 {
		return int64(1)
	})
	metrictor.Run(ctx, time.Second*5)
}
