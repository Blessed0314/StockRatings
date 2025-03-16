package utils

func Paginate[T any](data []T, page, pageSize int) ([]T, int, int) {
    total := len(data)
    
    if pageSize <= 0 {
        pageSize = 10
    }
    if page <= 0 {
        page = 1
    }

    // Calcular inicio y fin del slice paginado
    start := (page - 1) * pageSize
    if start >= total {
        return []T{}, total, 0
    }

    end := min(start + pageSize, total)

    return data[start:end], total, end - start
}

