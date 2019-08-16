package gormetrics

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/profects/gormetrics/gormi"
	"github.com/profects/gormetrics/gormi/adapter/unforked"
)

// Register gormetrics. Options (opts) can be used to configure the Prometheus
// namespace and GORM plugin scope.
func Register(db *gorm.DB, dbName string, opts ...RegisterOpt) error {
	return RegisterInterface(unforked.New(db), dbName, opts...)
}

// RegisterInterface registers gormetrics with a gormi.DB interface, which can
// be created using one of the adapters in gormi/adapter. This can be useful if
// you use a forked version of GORM.
// Options (opts) can be used to configure the Prometheus namespace and
// GORM plugin scope.
func RegisterInterface(db gormi.DB, dbName string, opts ...RegisterOpt) error {
	if db == nil {
		return ErrDbIsNil
	}

	driverName := sqlDriverToDriverName(db.DB().Driver())
	handlerOpts := getOpts(opts)
	info := extraInfo{
		dbName:     dbName,
		driverName: driverName,
	}

	handler, err := newCallbackHandler(info, handlerOpts)
	if err != nil {
		return errors.Wrap(err, "could not create callback handler")
	}
	handler.registerCallback(db.Callback())

	dbInterface := databaseFrom(info, db.DB())
	dbMetrics, err := newDatabaseMetrics(dbInterface, handlerOpts)
	if err != nil {
		return errors.Wrap(err, "could not create database metrics exporter")
	}
	go dbMetrics.maintain()

	return nil
}
