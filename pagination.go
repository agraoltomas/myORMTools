package myORMTools

import (
	"github.com/Masterminds/squirrel"
	"net/http"
	"strconv"
)

const PaginationDefaultSize = 10
const PaginationStartingPage = 1

type PaginationData struct {
	page uint64
	size uint64
}

func GetPaginationData(r *http.Request) PaginationData {
	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	if err != nil {
		page = PaginationStartingPage
	}
	size, err := strconv.ParseInt(r.URL.Query().Get("size"), 10, 64)
	if err != nil {
		size = PaginationDefaultSize
	}
	return PaginationData{
		uint64(page),
		uint64(size),
	}
}

func Paginate(sq squirrel.SelectBuilder, data *PaginationData) squirrel.SelectBuilder {
	sq = sq.Limit(data.size)
	if data.page > 1 {
		sq = sq.Offset(data.page * data.size)
	}
	return sq
}
