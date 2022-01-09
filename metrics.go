// Copyright 2019 Profects Group B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gormetrics

const (
	labelStatus   = "status"
	labelDatabase = "database"
	labelDriver   = "driver"

	// Statuses for metrics (values of labelStatus).
	metricStatusFail    = "fail"
	metricStatusSuccess = "success"

	metricOpenConnections  = "connections_open"
	metricIdleConnections  = "connections_idle"
	metricInUseConnections = "connections_in_use"

	helpOpenConnections  = `Currently open connections to the database`
	helpIdleConnections  = `Currently idle connections to the database`
	helpInUseConnections = `Currently in use connections`

	metricAllTotal           = "all_total"
	metricAllAverageTime     = "all_total_average_time"
	metricCreatesTotal       = "creates_total"
	metricCreatesAverageTime = "creates_average_time"
	metricDeletesTotal       = "deletes_total"
	metricDeletesAverageTime = "deletes_average_time"
	metricQueriesTotal       = "queries_total"
	metricQueriesAverageTime = "queries_average_time"
	metricUpdatesTotal       = "updates_total"
	metricUpdatesAverageTime = "updates_average_time"

	helpAllTotal           = `All queries requested`
	helpAllAverageTime     = `Average time of all queries requested`
	helpCreatesTotal       = `All create queries requested`
	helpCreatesAverageTime = `Average time of allcreate queries requested`
	helpDeletesTotal       = `All delete queries requested`
	helpDeletesAverageTime = `Average time of all delete queries requested`
	helpQueriesTotal       = `All select queries requested`
	helpQueriesAverageTime = `Average time of all select queries requested`
	helpUpdatesTotal       = `All update queries requested`
	helpUpdatesAverageTime = `Average time of all update queries requested`
)
