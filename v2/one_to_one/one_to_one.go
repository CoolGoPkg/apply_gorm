package one_to_one

import (
	"CoolGoPkg/apply_gorm/v2/model"
)

func InsertStockInfo(param *model.Stock) error {
	return model.DB.Create(param).Error
}

func GetStock(id int64) (model.Stock, error) {
	stock := model.Stock{}
	err := model.DB.Model(&model.Stock{}).Where("fiances.id=?", id).Preload("Fiance").Find(&stock).Error
	if err != nil {
		return stock, err
	}

	return stock, nil
}

func GetFiance(id int64) (model.FianceWithStock, error) {
	fiance := model.FianceWithStock{}
	err := model.DB.Model(&model.Fiance{}).Where("id=?", id).Preload("Stock").Find(&fiance).Error
	if err != nil {
		return fiance, err
	}

	return fiance, nil

}
