package plugins

import (
	"gorm.io/gorm"
)

var _ gorm.Plugin = (*GormError)(nil)

type GormError struct{}

func NewGormError() gorm.Plugin {
	return &GormError{}
}

func (p *GormError) Name() string {
	return "__gorm_error"
}

func (p *GormError) Initialize(db *gorm.DB) error {
	return nil
}

func (p *GormError) after(db *gorm.DB) {
}
