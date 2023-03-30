package one_to_many

import (
	"CoolGoPkg/apply_gorm/v1/model"
	"fmt"
)

func CreateTags() error {
	tags := []*model.Tag{
		{Name: "初始标签", Level: 0},
		{ID: 2, Name: "次级标签", Level: 1, Father: 1},
		{ID: 6, Name: "初始标签3", Level: 0, Father: 0, Children: []*model.Tag{{
			ID: 7, Name: "次级标签4", Level: 1, Children: []*model.Tag{{
				ID: 9, Name: "次次级标签1", Level: 2, Children: []*model.Tag{
					{
						ID: 12, Name: "次次次级标签1", Level: 3,
					},
				},
			}, {
				ID: 10, Name: "次次级标签2", Level: 2,
			}},
		}, {
			ID: 8, Name: "次级标签5", Level: 1, Children: []*model.Tag{{
				ID: 11, Name: "次次级标签3", Level: 2,
			}},
		},
		}},
	}

	for _, tag := range tags {
		err := model.DB().FirstOrCreate(tag).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func GetIndexTag(id int) error {
	tag := model.Tag{}
	err := model.DB().Preload("Children").Where("id=?", id).Find(&tag).Error
	if err != nil {
		return err
	}

	fmt.Println("first : ", tag)
	err = GetChildren(tag.Children)
	if err != nil {
		return err
	}
	fmt.Println("father tag : ", tag)
	PrintChildTag(tag.Children)

	return nil
}

func GetChildren(children []*model.Tag) error {
	if len(children) <= 0 {
		return nil
	}

	for _, child := range children {
		err := model.DB().Preload("Children").Where("id=?", child.ID).Find(&child).Error
		if err != nil {
			return err
		}

		err = GetChildren(child.Children)
		if err != nil {
			return err
		}
	}

	return nil
}

func PrintChildTag(tags []*model.Tag) {
	if len(tags) <= 0 {
		return
	}

	for _, tag := range tags {
		fmt.Println("child tag :", tag)
		PrintChildTag(tag.Children)
	}
}
