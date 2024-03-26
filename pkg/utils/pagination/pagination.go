package pagination

func GeneratePaginationOffset(limit, offset int64) int64 {
	if offset == 0 {
		offset = 1
	}
	return (offset * limit) - limit
}
