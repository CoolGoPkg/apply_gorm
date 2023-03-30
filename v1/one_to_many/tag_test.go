package one_to_many

import (
	"CoolGoPkg/apply_gorm/v1/model"
	"fmt"
	"testing"
)

func TestCreateTags(t *testing.T) {
	model.InitDB()
	err := CreateTags()
	if err != nil {
		t.Log("err : ", err)
		return
	}
}

func TestGetIndexTag(t *testing.T) {
	model.InitDB()
	err := GetIndexTag(11)
	if err != nil {
		t.Log("err :", err)
		return
	}
}

func TestCreateTags2(t *testing.T) {
	model.InitDB()
	tags := []*model.Tag{
		{Name: "初始标签1", Level: 0},
		{Name: "初始标签2", Level: 0},
		{Name: "初始标签3", Level: 0, Father: 0, Children: []*model.Tag{{
			Name: "次级标签1", Level: 1, Children: []*model.Tag{{
				Name: "次次级标签1", Level: 2, Children: []*model.Tag{
					{
						Name: "次次次级标签1", Level: 3,
					},
					{
						Name: "次次次级标签2", Level: 3,
					},
				},
			}, {Name: "次次级标签2", Level: 2}},
		}, {
			Name: "次级标签2", Level: 1, Children: []*model.Tag{{
				Name: "次次级标签3", Level: 2,
			}},
		},
		}},
		{Name: "初始标签4", Level: 0, Children: []*model.Tag{
			{Name: "次级标签4", Level: 1, Children: []*model.Tag{{
				Name: "次次级标签4", Level: 2, Children: []*model.Tag{{
					Name: "次次次级标签4", Level: 3, Children: []*model.Tag{{
						Name: "次次次次级标签4", Level: 4,
					}},
				}},
			}},
			}},
		},
	}
	for _, tag := range tags {
		fmt.Println("range rag name", tag.Name)
		err := tag.Create()
		if err != nil {
			t.Log("err : ", err)
			return
		}
	}
}
func TestCreateTags3(t *testing.T) {
	model.InitDB()
	tags := []*model.Tag{
		{ID: 11, Name: "初始标签4", Level: 0, Children: []*model.Tag{
			{ID: 12, Name: "次次级标签6", Level: 2},
		}},
	}
	for _, tag := range tags {
		fmt.Println("range rag name", tag.Name)
		err := tag.Create()
		if err != nil {
			t.Log("err : ", err)
			return
		}
	}
}
