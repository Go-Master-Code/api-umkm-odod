package helper

func NormalizePagination(page *int, limit *int) {
	if *page < 1 {
		*page = 1
	}

	if *limit < 1 {
		*limit = 10
	}
}
