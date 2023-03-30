package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Index struct {
	Code  string  `gorm:"column:code;primary_key" json:"code"`
	Name  string  `gorm:"column:name" json:"name"`
	Funds []*Fund `gorm:"save_associations:false"`
	Tags  []*Tag  `gorm:"many2many:fund_index_tag;foreignkey:name;jointable_foreignkey:name;association_foreignkey:name;association_jointable_foreignkey:name;save_associations:true"`
}

func (in *Index) Create() error {
	if in.Code == "" {
		return errors.New("index code 不可为空")
	}

	err := DB().Debug().First(&Index{Code: in.Code}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = DB().Create(in).Error
			if err != nil {
				return err
			}
		} else {
			return err
		}

	} else {
		err = DB().Debug().Save(in).Error
		if err != nil {
			return err
		}
	}

	for _, fund := range in.Funds {
		fund.IndexCode = in.Code
		err = fund.Create()
		if err != nil {
			return err
		}
	}

	err = DB().Model(&Index{Code: in.Code}).Association("Funds").Replace(in.Funds).Error
	if err != nil {
		return err
	}

	for _, tag := range in.Tags {
		if tag.ID == 0 {
			return errors.New("未知的标签")
		}
	}
	return DB().Model(&Index{Code: in.Code}).Association("Tags").Replace(in.Tags).Error

}

type indexes []*Index

type Fund struct {
	Code      string `gorm:"column:code;primary_key" json:"code"`
	Name      string `gorm:"column:name" json:"name"`
	Desc      string `gorm:"column:desc" json:"desc"`
	IndexCode string `gorm:"index_code"`
	Index     *Index
}

func (f *Fund) Create() error {
	if f.Code == "" {
		return errors.New("fund code 不可为空")
	}

	err := DB().Debug().First(&Fund{Code: f.Code}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = DB().Create(f).Error
			if err != nil {
				return err
			}
		}
		return err
	}

	return DB().Debug().Save(f).Error
}

type Tag struct {
	ID       int    `gorm:"column:id;primary_key" json:"id"`
	Name     string `gorm:"column:name"`
	Level    int    `gorm:"column:level"`
	Father   int    `gorm:"column:father"`
	Children []*Tag `gorm:"foreignkey:Father"`
	Indexes  []*Tag `gorm:"many2many:fund_index_tag;foreignkey:name;jointable_foreignkey:tag_name;save_associations:true"`
	//Indexes  []*Index `gorm:"many2many:fund_index_tag;save_associations:false"`
}

func CreateOneTag(tag *Tag) error {
	fmt.Println("create tag name", tag.Name, tag.Level)
	if tag.ID == 0 {
		return DB().Debug().Create(tag).Error
	}
	return DB().Debug().FirstOrCreate(tag).Error
}

func (t *Tag) Create() error {
	err := CreateOneTag(t)
	if err != nil {
		return err
	}

	if len(t.Children) > 0 {
		for _, child := range t.Children {
			child.Father = t.ID
			err = CreateOneTag(child)
			if err != nil {
				return err
			}
		}

		for _, cc := range t.Children {
			fmt.Println("children : ", cc)
		}
		err = DB().Model(&Tag{ID: t.ID}).Debug().Association("Children").Replace(t.Children).Error
		if err != nil {
			return err
		}

		// 递归处理child.Children
		for _, child := range t.Children {
			if len(child.Children) > 0 {
				for _, c := range child.Children {
					err = c.Create()
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}
