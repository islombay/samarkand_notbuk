package models

type Pagination struct {
	Limit      int    `form:"limit"`
	Page       int    `form:"page"`
	Query      string `form:"q"`
	CategoryID string `form:"category_id" binding:"omitempty,uuid" json:"category_id"`
	BrandID    string `form:"brand_id" binding:"omitempty,uuid" json:"brand_id"`

	Data   interface{} `form:"-" json:"-"`
	Count  int64       `form:"-" json:"-"`
	Offset int         `form:"-" json:"-"`
}

func (p *Pagination) Fix() {
	if p.Limit <= 0 {
		p.Limit = 10
	}

	if p.Page <= 0 {
		p.Page = 1
	}

	p.Count = 0

	p.Offset = (p.Page - 1) * p.Limit
}
