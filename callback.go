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

import (
	"fmt"
	"gorm.io/gorm"
	"time"

	"github.com/pkg/errors"

	"github.com/prometheus/client_golang/prometheus"
)

// callbackHandler manages gorm query callback handling so query
// statistics are always up to date.
type callbackHandler struct {
	opts          *pluginOpts
	counters      *queryCounters
	defaultLabels map[string]string
}

func (h *callbackHandler) registerCallback(db *gorm.DB) {
	cb := db.Callback()

	cb.Create().Before("gorm:create").Register(
		h.opts.callbackName("before_create"),
		h.setStartTime,
	)

	cb.Create().After("gorm:after_create").Register(
		h.opts.callbackName("after_create"),
		h.afterCreate,
	)

	cb.Delete().Before("gorm:delete").Register(
		h.opts.callbackName("before_delete"),
		h.setStartTime,
	)

	cb.Delete().After("gorm:after_delete").Register(
		h.opts.callbackName("after_delete"),
		h.afterDelete,
	)

	cb.Query().Before("gorm:query").Register(
		h.opts.callbackName("before_query"),
		h.setStartTime,
	)

	cb.Query().After("gorm:after_query").Register(
		h.opts.callbackName("after_query"),
		h.afterQuery,
	)

	cb.Update().Before("gorm:update").Register(
		h.opts.callbackName("before_update"),
		h.setStartTime,
	)

	cb.Update().After("gorm:after_update").Register(
		h.opts.callbackName("after_update"),
		h.afterUpdate,
	)
}

func (h *callbackHandler) setStartTime(db *gorm.DB) {
	db.Set("timeStart", time.Now())
}

func (h *callbackHandler) afterCreate(db *gorm.DB) {
	h.updateCounterVectors(db, h.counters.creates)
	h.updateHistogramVectors(db, h.counters.createsDuration)
}

func (h *callbackHandler) afterDelete(db *gorm.DB) {
	h.updateCounterVectors(db, h.counters.deletes)
	h.updateHistogramVectors(db, h.counters.deletesDuration)
}

func (h *callbackHandler) afterQuery(db *gorm.DB) {
	h.updateCounterVectors(db, h.counters.queries)
	h.updateHistogramVectors(db, h.counters.queriesDuration)
}

func (h *callbackHandler) afterUpdate(db *gorm.DB) {
	h.updateCounterVectors(db, h.counters.updates)
	h.updateHistogramVectors(db, h.counters.updatesDuration)
}

// updateCounterVectors registers one or more of prometheus.CounterVec to increment
// with the status in db (any type of query). If any errors are in
// db.DB().GetErrors(), a status "fail" will be assigned to the increment.
// Otherwise, a status "success" will be assigned.
// Increments h.counters.all (gormetrics_all_total) by default.
func (h *callbackHandler) updateCounterVectors(db *gorm.DB, vectors ...*prometheus.CounterVec) {
	vectors = append(vectors, h.counters.all)

	_, err := db.DB()
	status := metricStatusFail
	if err == nil {
		status = metricStatusSuccess
	}

	labels := mergeLabels(prometheus.Labels{
		labelStatus: status,
	}, h.defaultLabels)

	for _, counter := range vectors {
		if counter == nil {
			continue
		}

		counter.With(labels).Add(1)
	}
}

// updateHistogramVectors registers one or more of prometheus.HistogramVec to add observations to
// with the status in db (any type of query). If any errors are in
// db.DB().GetErrors(), a status "fail" will be assigned to the increment.
// Otherwise, a status "success" will be assigned.
// Increments h.counters.allDuration (gormetrics_all_total) by default.
func (h *callbackHandler) updateHistogramVectors(db *gorm.DB, vectors ...*prometheus.HistogramVec) {
	vectors = append(vectors, h.counters.allDuration)

	_, err := db.DB()
	status := metricStatusFail
	if err == nil {
		status = metricStatusSuccess
	}

	labels := mergeLabels(prometheus.Labels{
		labelStatus: status,
	}, h.defaultLabels)

	for _, histogram := range vectors {
		if histogram == nil {
			continue
		}

		startTime, ok := db.Get("timeStart")

		if !ok {
			continue
		}

		elapsed := time.Since(startTime.(time.Time)).Milliseconds()
		histogram.With(labels).Observe(float64(elapsed))
	}
}

// extraInfo contains information for filtering the provided metrics.
type extraInfo struct {
	// The name of the database in use.
	dbName string

	// The name of the driver powering database/sql (underlying database for GORM).
	driverName string
}

// newCallbackHandler creates a new callback handler configured with info and opts.
// info does not contain any mandatory information for the functioning of the
// function, but sets label values which can be useful in the usage of
// the provided metrics (driver, database, connection).
// Automatically registers metrics.
func newCallbackHandler(info extraInfo, opts *pluginOpts) (*callbackHandler, error) {
	counters, err := newQueryCounters(opts.prometheusNamespace)
	if err != nil {
		return nil, errors.Wrap(err, "could not create query gauges")
	}

	return &callbackHandler{
		opts:     opts,
		counters: counters,
		defaultLabels: prometheus.Labels{
			labelDriver:   info.driverName,
			labelDatabase: info.dbName,
		},
	}, nil
}

// callbackName creates a GORM callback name based on the configured plugin
// db and callback name.
func (c *pluginOpts) callbackName(callback string) string {
	return fmt.Sprintf("%v:%v", c.gormPluginScope, callback)
}

// Merges maps a and b. a is returned with extra values from b. Existing items
// in a with a matching key in b will not get overwritten.
func mergeLabels(a, b prometheus.Labels) prometheus.Labels {
	for k, v := range b {
		if _, exists := a[k]; !exists {
			a[k] = v
		}
	}
	return a
}
