package one_to_many

import (
	"CoolGoPkg/apply_gorm/v1/model"
	"fmt"
)

func CreateIndex() error {

	indexs := []*model.Index{
		{
			Code: "000004.SS",
			Name: "创业指数",
			Funds: []*model.Fund{
				{Code: "400001.SZ", Name: "4a", Desc: "最大基金"},
				{Code: "400002.SZ", Name: "4b", Desc: "第二大基金"},
			},
		},
		{
			Code: "000300.SZ",
			Name: "沪深指数",
			Funds: []*model.Fund{
				{Code: "300003.SZ", Name: "e", Desc: "eee"},
				{Code: "300002.SZ", Name: "d", Desc: "ddd"},
			},
		},
	}
	for _, index := range indexs {
		err := model.DB().Debug().FirstOrCreate(index).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func GetIndex() error {
	indxes := []*model.Index{}
	err := model.DB().Find(&indxes).Error
	if err != nil {
		return nil
	}

	for _, index := range indxes {
		fmt.Println("index : ", index)
		for _, v := range index.Funds {
			fmt.Println("fund : ", v)
		}
	}
	return nil
}

func GetIndexPreload() error {
	indxes := []*model.Index{}
	err := model.DB().Preload("Funds").Debug().Find(&indxes).Error
	if err != nil {
		return nil
	}

	for _, index := range indxes {
		fmt.Println("index : ", index)
		for _, v := range index.Funds {
			fmt.Println("fund : ", v)
		}
	}
	return nil
}
func GetFundPreload() error {
	funds := []*model.Fund{}
	err := model.DB().Preload("Index").Debug().Find(&funds).Error
	if err != nil {
		return nil
	}

	for _, fund := range funds {
		fmt.Println("fund : ", fund)
		fmt.Println("idx : ", fund.Index)
	}
	return nil
}

func GetIndexAssociation(code string) error {
	funds := []*model.Fund{}
	err := model.DB().Model(&model.Index{Code: code}).Association("Funds").Find(&funds).Error
	if err != nil {
		return nil
	}
	for _, v := range funds {
		fmt.Println("fund : ", v)
	}
	return nil
}

func AppendFundOfIndex(indexCode string, newFunds []model.Fund) error {
	return model.DB().Model(&model.Index{Code: indexCode}).Association("Funds").Append(newFunds).Error
}

func AppendIndexOfFund(fundCode string, newIndex *model.Index) error {
	return nil
}

func ReplaceFundOfIndex(indexCode string, newFunds []*model.Fund) error {
	return model.DB().Model(&model.Index{Code: indexCode}).Association("Funds").Replace(newFunds).Error

}
