package gormx

import (
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

const errorAfterName = "error:after"

type GormError struct{}

func NewGormError() gorm.Plugin {
	return &GormError{}
}

func (p *GormError) Name() string {
	return "__gorm_error"
}

func (p *GormError) Initialize(db *gorm.DB) error {
	// 结束后
	db.Callback().Create().After("gorm:error_create").Register(errorAfterName, p.after)
	db.Callback().Query().After("gorm:error_query").Register(errorAfterName, p.after)
	db.Callback().Delete().After("gorm:error_delete").Register(errorAfterName, p.after)
	db.Callback().Update().After("gorm:error_update").Register(errorAfterName, p.after)
	db.Callback().Row().After("gorm:error_row").Register(errorAfterName, p.after)
	db.Callback().Raw().After("gorm:error_raw").Register(errorAfterName, p.after)
	return nil
}

func (p *GormError) after(db *gorm.DB) {
	if db.Error != nil {
		db.Error = WrapGormError(db.Error)
	}
}

// WrapGormError 转换错误
func WrapGormError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return GormErrRecordNotFound(err)
	}
	if errors.Is(err, gorm.ErrInvalidTransaction) {
		return ErrInvalidTransaction(err)
	}
	if errors.Is(err, gorm.ErrNotImplemented) {
		return ErrNotImplemented(err)
	}
	if errors.Is(err, gorm.ErrMissingWhereClause) {
		return ErrMissingWhereClause(err)
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////gorm

// GormErrRecordNotFound 找不到记录
func GormErrRecordNotFound(err error) error {
	return errors.InternalServer("GormErrRecordNotFound", err.Error())
}

func IsGormErrRecordNotFound(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "GormErrRecordNotFound" && se.Code == 500
}

func ErrInvalidTransaction(err error) error {
	return errors.InternalServer("ErrInvalidTransaction", err.Error())
}

func IsErrInvalidTransaction(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "ErrInvalidTransaction" && se.Code == 500
}

// ErrNotImplemented not implemented
func ErrNotImplemented(err error) error {
	return errors.InternalServer("ErrNotImplemented", err.Error())
}

// ErrMissingWhereClause missing where clause
func ErrMissingWhereClause(err error) error {
	return errors.InternalServer("ErrMissingWhereClause", err.Error())
}

// ErrUnsupportedRelation unsupported relations
func ErrUnsupportedRelation(err error) error {
	return errors.InternalServer("ErrUnsupportedRelation", err.Error())
}

// ErrPrimaryKeyRequired primary keys required
func ErrPrimaryKeyRequired(err error) error {
	return errors.InternalServer("ErrPrimaryKeyRequired", err.Error())
}

// ErrModelValueRequired model value required
func ErrModelValueRequired(err error) error {
	return errors.InternalServer("ErrModelValueRequired", err.Error())
}

// ErrInvalidData unsupported data
func ErrInvalidData(err error) error {
	return errors.InternalServer("ErrInvalidData", err.Error())
}

// ErrUnsupportedDriver unsupported driver
func ErrUnsupportedDriver(err error) error {
	return errors.InternalServer("ErrUnsupportedDriver", err.Error())
}

// ErrRegistered registered
func ErrRegistered(err error) error {
	return errors.InternalServer("ErrRegistered", err.Error())
}

// ErrInvalidField invalid field
func ErrInvalidField(err error) error {
	return errors.InternalServer("ErrInvalidField", err.Error())
}

// ErrEmptySlice empty slice found
func ErrEmptySlice(err error) error {
	return errors.InternalServer("ErrEmptySlice", err.Error())
}

// ErrDryRunModeUnsupported dry run mode unsupported
func ErrDryRunModeUnsupported(err error) error {
	return errors.InternalServer("ErrDryRunModeUnsupported", err.Error())
}

// ErrInvalidDB invalid db
func ErrInvalidDB(err error) error {
	return errors.InternalServer("ErrInvalidDB", err.Error())
}

// ErrInvalidValue invalid value
func ErrInvalidValue(err error) error {
	return errors.InternalServer("ErrInvalidValue", err.Error())
}

// ErrInvalidValueOfLength invalid values do not match length
func ErrInvalidValueOfLength(err error) error {
	return errors.InternalServer("ErrInvalidValueOfLength", err.Error())
}
