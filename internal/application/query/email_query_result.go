package query

import "github.com/Andres-MK/internal/application/common"

type EmailQueryResult struct {
	Result *common.EmailResult
}

type ListEmailsQueryResult struct {
	Result []*common.EmailResult
}