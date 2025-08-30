package dbx

func GetOffset(page, size int) (int, int) {
	if page != 0 {
		page -= 1
	}
	return page * size, size
}
