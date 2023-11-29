package helpers

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

func GeneratePaginationFromRequest(r *http.Request) Pagination {
	limit := 10
	page := 1
	sort := "created_at asc"
	query := r.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "sort":
			sort = queryValue
		}
	}
	return Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

}
