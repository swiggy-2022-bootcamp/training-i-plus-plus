package config

import (
	"fmt"
	"time"

	"mini-project/bff/service"

	"github.com/go-kit/kit/metrics"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           service.StationService
}

func (mw InstrumentingMiddleware) AddStation(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "addStation", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.AddStation(s)
	return
}

func (mw InstrumentingMiddleware) UpdateStation(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "updateStation", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.UpdateStation(s)
	return
}

func (mw InstrumentingMiddleware) RetrieveStation(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "retrieveStation", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.RetrieveStation(s)
	return
}

func (mw InstrumentingMiddleware) DeleteStation(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "deleteStation", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.DeleteStation(s)
	return
}
