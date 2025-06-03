package models

import (
	"os"
	"strconv"
)

type Pagination struct {
	Limit      int    `json:"limit,omitempty,query:limit"`
	Page       int    `json:"page,omitempty,query:page"`
	Sort       string `json:"sort,omitempty,query:sort"`
	Order      string `json:"order,omitempty,query:order"`
	TotalRows  int64  `json:"total_rows,omitempty"`
	TotalPages int    `json:"total_pages,omitempty"`
}

func (objPage *Pagination) GetOffSet() int {
	return (objPage.GetPage() - 1) * objPage.GetLimit()
}

func (objPage *Pagination) GetLimit() int {
	strPerPage := os.Getenv("ITEMS_PER_PAGE")
	strMaxPage := os.Getenv("ITEMS_MAX_PAGE")

	maxPage, err := strconv.Atoi(strMaxPage)

	if err == nil || maxPage == 0 {
		maxPage = 999
	}

	perPage, err := strconv.Atoi(strPerPage)

	if err == nil && objPage.Limit == 0 {
		objPage.Limit = perPage
	}

	if objPage.Limit == 0 || objPage.Limit > maxPage {
		objPage.Limit = 50
	}

	return objPage.Limit
}

func (objPage *Pagination) GetPage() int {

	if objPage.Page == 0 {
		objPage.Page = 1
	}

	return objPage.Page
}

func (objPage *Pagination) GetSort() string {

	if objPage.Sort == "" {
		objPage.Sort = "id"
	}

	if objPage.Order == "" {
		objPage.Order = "desc"
	}

	return objPage.Sort + " " + objPage.Order
}
