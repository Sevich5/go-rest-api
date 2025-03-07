package requestdto

type PaginatedDto struct {
	Limit  int `form:"limit" binding:"required"`
	Offset int `form:"offset"`
}
