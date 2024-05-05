package timeseries

import (
	"github.com/anil-appface/grabana/target/azurelog"
	"github.com/anil-appface/grabana/target/graphite"
	"github.com/anil-appface/grabana/target/influxdb"
	"github.com/anil-appface/grabana/target/loki"
	"github.com/anil-appface/grabana/target/prometheus"
	"github.com/anil-appface/grabana/target/stackdriver"
	"github.com/anil-appface/sdk"
)

// WithAzureLogTarget adds a azure log query to the graph.
func WithAzureLogTarget(query, resource, timeColumn string, options ...azurelog.Option) Option {
	target := azurelog.New(query, options...)

	return func(graph *TimeSeries) error {
		graph.Builder.AddTarget(&sdk.Target{
			RefID:          target.Ref,
			Hide:           target.Hidden,
			Expr:           target.Expr,
			IntervalFactor: target.IntervalFactor,
			Interval:       target.Interval,
			Step:           target.Step,
			LegendFormat:   target.LegendFormat,
			Instant:        target.Instant,
			Format:         target.Format,
			AzureLogAnalytics: sdk.AzureLogAnalytics{
				TimeColumn:    timeColumn,
				DashboardTime: true,
				Query:         query,
				ResultFormat:  "time_series",
				Resources: []string{
					resource,
				},
			},
			QueryType: "Azure Log Analytics",
		})

		return nil
	}
}

// WithPrometheusTarget adds a prometheus query to the graph.
func WithPrometheusTarget(query string, options ...prometheus.Option) Option {
	target := prometheus.New(query, options...)

	return func(graph *TimeSeries) error {
		graph.Builder.AddTarget(&sdk.Target{
			RefID:          target.Ref,
			Hide:           target.Hidden,
			Expr:           target.Expr,
			IntervalFactor: target.IntervalFactor,
			Interval:       target.Interval,
			Step:           target.Step,
			LegendFormat:   target.LegendFormat,
			Instant:        target.Instant,
			Format:         target.Format,
		})

		return nil
	}
}

// WithGraphiteTarget adds a Graphite target to the table.
func WithGraphiteTarget(query string, options ...graphite.Option) Option {
	target := graphite.New(query, options...)

	return func(graph *TimeSeries) error {
		graph.Builder.AddTarget(target.Builder)

		return nil
	}
}

// WithInfluxDBTarget adds an InfluxDB target to the graph.
func WithInfluxDBTarget(query string, options ...influxdb.Option) Option {
	target := influxdb.New(query, options...)

	return func(graph *TimeSeries) error {
		graph.Builder.AddTarget(target.Builder)

		return nil
	}
}

// WithStackdriverTarget adds a stackdriver query to the graph.
func WithStackdriverTarget(target *stackdriver.Stackdriver) Option {
	return func(graph *TimeSeries) error {
		graph.Builder.AddTarget(target.Builder)

		return nil
	}
}

// WithLokiTarget adds a loki query to the graph.
func WithLokiTarget(query string, options ...loki.Option) Option {
	target := loki.New(query, options...)

	return func(graph *TimeSeries) error {
		graph.Builder.AddTarget(&sdk.Target{
			Hide:         target.Hidden,
			Expr:         target.Expr,
			LegendFormat: target.LegendFormat,
		})

		return nil
	}
}
