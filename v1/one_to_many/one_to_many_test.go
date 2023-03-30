package one_to_many

import (
	"CoolGoPkg/apply_gorm/v1/model"
	"testing"
)

func TestCreateIndex(t *testing.T) {
	model.InitDB()
	err := CreateIndex()
	if err != nil {
		t.Log("err is : ", err)
		return
	}
	return
}

func TestGetIndex(t *testing.T) {
	model.InitDB()
	err := GetIndex()
	if err != nil {
		t.Log("err is : ", err)
		return
	}
}

func TestGetIndexPreload(t *testing.T) {
	model.InitDB()
	err := GetIndexPreload()
	if err != nil {
		t.Log("err is : ", err)
		return
	}
}

func TestGetIndexAssociation(t *testing.T) {
	model.InitDB()
	code := "000004.SS"
	err := GetIndexAssociation(code)
	if err != nil {
		t.Log("err is : ", err)
		return
	}
}

func TestGetFundPreload(t *testing.T) {
	model.InitDB()
	err := GetFundPreload()
	if err != nil {
		t.Log("err is : ", err)
		return
	}
}

// UPDATE `funds` SET `name` = '5沙', `desc` = '', `index_code` = '000004.SS'  WHERE `funds`.`code` = '400005.SZ'
// INSERT INTO `funds` (`code`,`name`,`desc`,`index_code`) VALUES ('400005.SZ','5沙','','000004.SS')/
func TestAppendFundOfIndex(t *testing.T) {
	model.InitDB()
	indexCode := "xxxxxxx.SS"
	err := AppendFundOfIndex(indexCode, []model.Fund{{
		Code: "xxxxxxx1",
		Name: "5555c",
	}, {
		Code: "xxxxxxx2",
		Name: "4e",
		Desc: "x三沙",
	}})
	if err != nil {
		t.Log("err :", err)
		return
	}
}

// UPDATE `funds` SET `name` = '5沙', `desc` = '', `index_code` = '000004.SS'  WHERE `funds`.`code` = '400005.SZ'
// INSERT INTO `funds` (`code`,`name`,`desc`,`index_code`) VALUES ('400005.SZ','5沙','','000004.SS')
//  UPDATE `funds` SET `index_code` = '<nil>'  WHERE (`code` NOT IN ('400005.SZ','400003.SZ')) AND (`index_code` = '000004.SS')
func TestReplaceFundOfIndex(t *testing.T) {
	model.InitDB()
	indexCode := "000004.SS"
	err := ReplaceFundOfIndex(indexCode, []*model.Fund{{
		Code: "400006.SZ",
		Name: "5沙",
	}, {
		Code: "400003.SZ",
		Name: "4e",
		Desc: "f三沙",
	}})
	if err != nil {
		t.Log("err :", err)
		return
	}
}

func TestIndex(t *testing.T) {
	model.InitDB()
	indexs := []*model.Index{
		{
			Code: "000004.SS",
			Name: "创业指数",
			Funds: []*model.Fund{
				{Code: "400001.SZ", Name: "4e", Desc: "最大基"},
				{Code: "400002.SZ", Name: "4d", Desc: "第二大基金"},
			},
			Tags: []*model.Tag{
				{ID: 14},
			},
		},
		{
			Code: "333333.SZ",
			Name: "测试指数",
			Funds: []*model.Fund{
				{Code: "300008.SZ", Name: "e", Desc: "eee"},
				{Code: "300002.SZ", Name: "d", Desc: "ddd"},
			},
			Tags: []*model.Tag{
				{ID: 6},
				{ID: 7},
			},
		},
	}
	for _, idx := range indexs {
		err := idx.Create()
		if err != nil {
			return
		}
	}
}
