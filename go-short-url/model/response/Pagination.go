package response

type Pagination struct {
	PageSize    int         `json:"pageSize,omitempty" form:"pageSize"`   //每页几条
	CurrentPage int         `json:"pageIndex,omitempty" form:"pageIndex"` // 当前第几页
	Sort        string      `json:"sort,omitempty" form:"sort"`           //排序
	TotalRows   int         `json:"totalRows"`                            //共多少行
	TotalPages  int         `json:"totalPages"`                           // 共多少业
	Rows        interface{} `json:"rows"`                                 // 数据
}

func (p *Pagination) SetRows(data interface{}) {
	p.Rows = data
}

func NewPagination(PageSize int, PageIndex int, Sort string, TotalRows int) *Pagination {
	p := &Pagination{
		PageSize:    PageSize,
		CurrentPage: PageIndex,
		Sort:        Sort,
		TotalRows:   TotalRows,
	}
	p.TotalPages = p.GetTotalPages()
	return p
}

func (p *Pagination) GetOffset() int {
	return (p.GetCurrentPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.PageSize == 0 {
		p.PageSize = 10
	}
	return p.PageSize
}

func (p *Pagination) GetCurrentPage() int {
	return p.CurrentPage
}

func (p *Pagination) GetTotalPages() int {
	if p.TotalRows == 0 {
		return 0
	}
	if p.TotalRows < p.PageSize {
		return 1
	}
	if p.TotalRows%p.PageSize == 0 {
		return p.TotalRows / p.PageSize
	}
	return (p.TotalRows / p.PageSize) + 1
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}
