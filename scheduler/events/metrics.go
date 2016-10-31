package events

import (
	"strings"

	xmetrics "code.uber.internal/infra/mesos-go/extras/metrics"
	"code.uber.internal/infra/mesos-go/scheduler"
)

func Metrics(harness xmetrics.Harness) Decorator {
	return func(h Handler) Handler {
		if h == nil {
			return h
		}
		return HandlerFunc(func(e *scheduler.Event) error {
			typename := strings.ToLower(e.GetType().String())
			return harness(func() error { return h.HandleEvent(e) }, typename)
		})
	}
}
