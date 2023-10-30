package common

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Paging struct {
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
	Limit int `json:"limit"`
}

func GetPagingItem(s string, r *gin.Context, _default int) (int, error){
	item := r.Query(s)
	if item == "" {
		return _default, nil
	} else {
		result, err := strconv.Atoi(item)
		if err != nil {
			return 0, err
		}
		return result, nil
	}
}

func GetPaging(p *Paging, r *gin.Context) error{
	var err error

	p.Page, err = GetPagingItem("page", r, 1)
	p.Limit, err = GetPagingItem("limit", r, 10)

	if err!=nil{
		return err
	}
	return nil
}
