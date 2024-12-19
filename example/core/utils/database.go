package utils

import (
	"github.com/dreamph/dbre/example/core/models"
	"github.com/dreamph/dbre/query"
)

func ToQueryLimit(pageLimit *models.PageLimit) *query.Limit {
	if pageLimit == nil {
		return nil
	}
	limit := &query.Limit{}
	limit.PageSize = pageLimit.PageSize
	limit.Offset = (pageLimit.PageNumber - 1) * pageLimit.PageSize
	return limit
}
