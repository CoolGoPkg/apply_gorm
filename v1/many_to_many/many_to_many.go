package many_to_many

import (
	"CoolGoPkg/apply_gorm/v1/model"
	"fmt"
)

func CreateFundAndTag() error {
	idxAndTags := []*model.Index{
		{Code: "xxxxxxx.SS", Name: "xxxxxxxxxx", Tags: []*model.Tag{{
			Name:   "testTag8",
			Level:  0,
			Father: 0,
		}, {
			Name:   "testTag9",
			Level:  0,
			Father: 0,
		}},
			Funds: []*model.Fund{{
				Code: "mmmmmmm.ZZ",
				Name: "mmmmm",
				Desc: "mfund",
			}, {
				Code: "mmmmmmm2.ZZ",
				Name: "mmmmm2",
				Desc: "m2fund",
			}}},
		{Code: "nnnnnnnn.SS", Name: "mmmmmmmmmm", Tags: []*model.Tag{{
			Name:   "testTag7",
			Level:  0,
			Father: 0,
		}}},
	}

	for _, idxTag := range idxAndTags {
		err := model.DB().Create(idxTag).Error
		if err != nil {
			return err
		}
	}

	return nil

}

func GetFunds() error {
	indexes := []*model.Index{}
	err := model.DB().Preload("Tags").Find(&indexes).Error
	if err != nil {
		return err
	}
	for _, index := range indexes {
		fmt.Println("index : ", index)
		for _, tag := range index.Tags {
			fmt.Println("tag : ", tag)
		}
	}
	return nil
}

func GetTags() error {

	var tags []*model.Tag
	err := model.DB().Preload("Indexes").Find(&tags).Error
	if err != nil {
		return err
	}
	for _, tag := range tags {
		fmt.Println("fund : ", tag)
		for _, fund := range tag.Indexes {
			fmt.Println("fund : ", fund)
		}
	}
	return nil
}
func GetTagsFunds(code string) error {

	var indexes []*model.Index
	err := model.DB().Preload("Funds").Preload("Tags").Where("code=?", code).Find(&indexes).Error
	if err != nil {
		return err
	}
	for _, index := range indexes {
		fmt.Println("fund : ", index)
		for _, fund := range index.Funds {
			fmt.Println("fund : ", fund)
		}
		for _, tag := range index.Tags {
			fmt.Println("tag : ", tag)
		}
	}
	return nil
}

func AppendIndexTag(indexCode string, newTags []*model.Tag) error {
	return model.DB().Model(&model.Index{Code: indexCode}).Association("Tags").Append(newTags).Error
}

func ReplaceIndexTag(indexCode string, newTags []*model.Tag) error {
	return model.DB().Model(&model.Index{Code: indexCode}).Association("Tags").Replace(newTags).Error
}
