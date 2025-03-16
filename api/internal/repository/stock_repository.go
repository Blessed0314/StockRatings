package repository

import (
    "github.com/Blessed0314/tru-test/api/internal/models"
    "github.com/Blessed0314/tru-test/api/pkg/db"
    "gorm.io/gorm"
)

type StockRepository struct {
    DB *gorm.DB
}

func NewStockRepository() *StockRepository {
    return &StockRepository{
        DB: db.DB,
    }
}

func (r *StockRepository) Save(stock *models.StockRating) error {
    if err := r.DB.Create(stock).Error; err != nil {
        return err
    }
    return nil
}

func (r *StockRepository) Update(stock *models.StockRating) error {
    if err := r.DB.Save(stock).Error; err != nil {
        return err
    }
    return nil
}

func (r *StockRepository) GetAll() ([]models.StockRating, error) {
    var stocks []models.StockRating
    if err := r.DB.Find(&stocks).Error; err != nil {
        return nil, err
    }
    return stocks, nil
}

func (r *StockRepository) GetByTicker(ticker string) (*models.StockRating, error) {
    var stock models.StockRating
    if err := r.DB.Where("ticker = ?", ticker).First(&stock).Error; err != nil {
        return nil, err
    }
    return &stock, nil
}

