package one_to_one

import (
	"CoolGoPkg/apply_gorm/v2/model"
	"testing"
	"time"
)

func TestGetStock(t *testing.T) {
	model.InitDB()
	stock, err := GetStock(1)
	if err != nil {
		t.Log("err : ", err)
		return
	}

	t.Log("info : ", stock)

}

func TestGetFiance(t *testing.T) {
	model.InitDB()
	fiance, err := GetFiance(1)
	if err != nil {
		t.Log("err : ", err)
		return
	}
	t.Log("info : ", fiance)
}

func TestInsertStockInfo(t *testing.T) {
	model.InitDB()
	err := InsertStockInfo(&model.Stock{
		ProdCode:    "000001.SS",
		ProdName:    "上证指数",
		TradeStatus: "Trade",
		Last:        3000.00,
		Fiance: model.Fiance{
			Eps:         1.02,
			Bps:         2.01,
			TotalShares: 1000.00,
		},
	})
	if err != nil {
		t.Log("err : ", err)
		return
	}
}

func TestTruncate(t *testing.T) {
	now := time.Now()
	ttt := now.Add(-15400 * 1e9)
	t.Log(ttt)
	a := ttt.Truncate(14400 * time.Second)
	t.Log(a)
}
