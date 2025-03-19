package dtos

type PaginationDTO[T any] struct {
    Total    int `json:"total"`    
    Page     int `json:"page"`     
    PageSize int `json:"pageSize"` 
    Count    int `json:"count"`    
    Stocks   []T `json:"stocks"`     
}