package utils

import "github.com/Blessed0314/tru-test/api/internal/dtos"

func Paginate[T any](data []T, page, pageSize int) dtos.PaginationDTO[T] {
    total := len(data)

    if pageSize <= 0 {
        pageSize = 10
    }
    if page <= 0 {
        page = 1
    }

    start := (page - 1) * pageSize
    if start >= total {
        return dtos.PaginationDTO[T]{
            Total:    total,
            Page:     page,
            PageSize: pageSize,
            Count:    0,
            Stocks:     []T{},
        }
    }

    end := min(start + pageSize, total)

    return dtos.PaginationDTO[T]{
        Total:    total,
        Page:     page,
        PageSize: pageSize,
        Count:    end - start,
        Stocks:     data[start:end],
    }
}

