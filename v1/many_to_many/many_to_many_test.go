package many_to_many

import (
	"CoolGoPkg/apply_gorm/v1/model"
	"testing"
)

func TestCreateFundAndTag(t *testing.T) {
	model.InitDB()
	err := CreateFundAndTag()
	if err != nil {
		t.Log("err :", err)
	}
	return
}

func TestGetFunds(t *testing.T) {
	model.InitDB()
	err := GetFunds()
	if err != nil {
		t.Log("err :", err)
	}
	return
}

func TestGetTags(t *testing.T) {
	model.InitDB()
	err := GetTags()
	if err != nil {
		t.Log("err :", err)
	}
	return
}

func TestGetTagsFunds(t *testing.T) {
	model.InitDB()
	err := GetTagsFunds("mmmmmmmm.SS")
	if err != nil {
		t.Log("err :", err)
	}
	return
}

func TestAppendIndexTag(t *testing.T) {
	model.InitDB()
	err := AppendIndexTag("mmmmmmmm.SS", []*model.Tag{{
		Name:  "testTag12",
		Level: 0,
	}, {
		ID:    24,
		Name:  "testTag9",
		Level: 0,
	}})
	if err != nil {
		t.Log("err :", err)
		return
	}
}

func TestReplaceIndexTag(t *testing.T) {
	model.InitDB()
	err := ReplaceIndexTag("mmmmmmmm.SS", []*model.Tag{{
		Name:  "testTag12",
		Level: 0,
	}, {
		ID:    24,
		Name:  "testTag9",
		Level: 0,
	}})
	if err != nil {
		t.Log("err :", err)
		return
	}
}
