package models

import (
	"github.com/igorfranzoi/golibfunctions/config"
)

type Pagination struct {
	Limit      int    `json:"limit,omitempty" form:"limit"`
	Page       int    `json:"page,omitempty" form:"page"`
	Sort       string `json:"sort,omitempty" form:"sort"`
	Order      string `json:"order,omitempty" form:"order"`
	TotalRows  int64  `json:"total_rows,omitempty"`
	TotalPages int    `json:"total_pages,omitempty"`
}

var DefaultPagination = config.DefaultConfig

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetLimit(cfg *config.ConfigPagination) int {
	if p.Limit == 0 {
		p.Limit = cfg.DefaultLimit
	}
	if p.Limit > cfg.MaxLimit {
		p.Limit = cfg.MaxLimit
	}
	return p.Limit
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit(&DefaultPagination)
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id"
	}
	if p.Order == "" {
		p.Order = "desc"
	}
	return p.Sort + " " + p.Order
}
