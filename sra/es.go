package sra

import (
	"github.com/elastic/go-elasticsearch/v7/estransport"
	"github.com/go-kratos/kratos/v2/errors"
)

type EsSelector struct {
}

func (s *EsSelector) Select([]*estransport.Connection) (*estransport.Connection, error) {
	return nil, errors.ServiceUnavailable("Unavailable", "not impl")
}
