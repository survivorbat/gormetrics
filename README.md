# Gormetrics

[![Go Report Card](https://goreportcard.com/badge/github.com/survivorbat/gormetrics)](https://goreportcard.com/report/github.com/survivorbat/gormetrics)
[![GoDoc](https://godoc.org/github.com/survivorbat/gormetrics?status.svg)](http://godoc.org/github.com/survivorbat/gormetrics)

_Forked from [profects/gormetrics](https://github.com/profects/gormetrics)._

A plugin for GORM providing metrics using Prometheus.

Warning: this plugin is still in an early stage of development. APIs may change.

## Usage

```go
import "github.com/survivorbat/gormetrics"

if err := gormetrics.Register(db, "my_database"); err != nil {
	// handle the error
}
```

Gormetrics does not expose the metrics endpoint using promhttp, you have to do this yourself.
You can use the following snippet for exposing metrics on port 2112 at `/metrics`:

```go
go func() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":2112", nil))
}()
```

## Exported metrics

| Type      | Metric                        | Purpose                                               |
|-----------|-------------------------------|-------------------------------------------------------|
| Counter   | gormetrics_all_total          | Counts how many queries have been performed           |
| Counter   | gormetrics_creates_total      | Counts how many create-queries have been performed    |
| Counter   | gormetrics_deletes_total      | Counts how many delete-queries have been performed    |
| Counter   | gormetrics_updates_total      | Counts how many update-queries have been performed    |
| Counter   | gormetrics_queries_total      | Counts how many select-queries have been performed    |
| Histogram | gormetrics_all_duration       | A histogram of all query durations in milliseconds    |
| Histogram | gormetrics_creates_duration   | A histogram of create-query durations in milliseconds |
| Histogram | gormetrics_deletes_duration   | A histogram of delete-query durations in milliseconds |
| Histogram | gormetrics_updates_duration   | A histogram of update-query durations in milliseconds |
| Histogram | gormetrics_queries_duration   | A histogram of select-query durations in milliseconds |
| Gauge     | gormetrics_connections_idle   | Amount of idle connections                            |
| Gauge     | gormetrics_connections_in_use | Amount of in-use connections                          |
| Gauge     | gormetrics_connections_open   | Amount of open connections                            |

These all have the following labels:

- `database`: the name of the database
- `driver`: the driver for the database (e.g. pq)
- `status`: fail or success (only for query-related metrics)
