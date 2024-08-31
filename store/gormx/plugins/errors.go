package plugins

import (
	"github.com/go-bamboo/pkg/store/gormx/ecode"
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

const errorAfterName = "error:after"

var _ gorm.Plugin = (*GormError)(nil)

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
		return ecode.GormErrRecordNotFound(err)
	}
	if errors.Is(err, gorm.ErrInvalidTransaction) {
		return ecode.ErrInvalidTransaction(err)
	}
	if errors.Is(err, gorm.ErrNotImplemented) {
		return ecode.ErrNotImplemented(err)
	}
	if errors.Is(err, gorm.ErrMissingWhereClause) {
		return ecode.ErrMissingWhereClause(err)
	}
	return err
}
