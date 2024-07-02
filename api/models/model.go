package models

type Pagination struct {
	Limit  int         `form:"limit"`
	Page   int         `form:"page"`
	Offset int         `form:"-" json:"-"`
	Query  string      `form:"q"`
	Data   interface{} `form:"-" json:"-"`
	Count  int64       `form:"-" json:"-"`
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
