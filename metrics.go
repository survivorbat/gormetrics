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

	metricAllTotal        = "all_total"
	metricAllDuration     = "all_duration"
	metricCreatesTotal    = "creates_total"
	metricCreatesDuration = "creates_duration"
	metricDeletesTotal    = "deletes_total"
	metricDeletesDuration = "deletes_duration"
	metricQueriesTotal    = "queries_total"
	metricQueriesDuration = "queries_duration"
	metricUpdatesTotal    = "updates_total"
	metricUpdatesDuration = "updates_duration"

	helpAllTotal        = `All queries requested`
	helpAllDuration     = `Duration of all queries requested in milliseconds`
	helpCreatesTotal    = `All create queries requested`
	helpCreatesDuration = `Duration of all create queries requested in milliseconds`
	helpDeletesTotal    = `All delete queries requested`
	helpDeletesDuration = `Duration of all delete queries requested in milliseconds`
	helpQueriesTotal    = `All select queries requested`
	helpQueriesDuration = `Duration of all select queries requested in milliseconds`
	helpUpdatesTotal    = `All update queries requested`
	helpUpdatesDuration = `Duration of all update queries requested in milliseconds`
)
