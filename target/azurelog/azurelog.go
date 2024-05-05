package azurelog

// FormatMode switches between Table, Time series, or Heatmap. Table will only work
// in the Table panel. Heatmap is suitable for displaying metrics of the
// Histogram type on a Heatmap panel. Under the hood, it converts cumulative
// histograms to regular ones and sorts series by the bucket bound.
type FormatMode string

const (
	FormatTable      FormatMode = "table"
	FormatHeatmap    FormatMode = "heatmap"
	FormatTimeSeries FormatMode = "time_series"
)

// Option represents an option that can be used to configure a AzureLog query.
type Option func(target *AzureLog)

// AzureLog represents a AzureLog query.
type AzureLog struct {
	Ref            string
	Hidden         bool
	Expr           string
	IntervalFactor int
	Interval       string
	Step           int
	LegendFormat   string
	Instant        bool
	Format         string
}

// New creates a new AzureLog query.
func New(query string, options ...Option) *AzureLog {
	azureLog := &AzureLog{
		Expr:   query,
		Format: string(FormatTimeSeries),
	}

	for _, opt := range options {
		opt(azureLog)
	}

	return azureLog
}

// Legend sets the legend format.
func Legend(legend string) Option {
	return func(AzureLog *AzureLog) {
		AzureLog.LegendFormat = legend
	}
}

// Ref sets the reference ID for this query.
func Ref(ref string) Option {
	return func(AzureLog *AzureLog) {
		AzureLog.Ref = ref
	}
}

// Hide the query. Grafana does not send hidden queries to the data source,
// but they can still be referenced in alerts.
func Hide() Option {
	return func(AzureLog *AzureLog) {
		AzureLog.Hidden = true
	}
}

// Instant marks the query as "instant, which means AzureLog will only return the latest scrapped value.
func Instant() Option {
	return func(AzureLog *AzureLog) {
		AzureLog.Instant = true
	}
}

// Format indicates how the data should be returned.
func Format(format FormatMode) Option {
	return func(AzureLog *AzureLog) {
		AzureLog.Format = string(format)
	}
}

// IntervalFactor sets the resolution factor.
func IntervalFactor(factor int) Option {
	return func(AzureLog *AzureLog) {
		AzureLog.IntervalFactor = factor
	}
}
