package model

type Stock struct {
	ID          int64   `gorm:"primary_key" json:"id"`
	ProdCode    string  `gorm:"type:varchar(100);" json:"prod_code"`   // 代码
	ProdName    string  `gorm:"type:varchar(100);" json:"prod_name"`   // 名称
	TradeStatus string  `gorm:"type:varchar(50);" json:"trade_status"` // 交易状态
	Last        float64 `gorm:"type:double" json:"last"`               // 价格

	Fiance Fiance // 经济数据
	// External External // 额外数据
}

// Fiance 个股经济数据
type Fiance struct {
	ID                  int64   `gorm:"primary_key" json:"id"`
	Eps                 float64 `json:"eps"`                   // 每股收益
	Bps                 float64 `json:"bps"`                   // 每股净资产
	TotalShares         float64 `json:"total_shares"`          // 总股本
	CirculationShares   float64 `json:"circulation_shares"`    // 流通股本
	UnCirculationShares float64 `json:"un_circulation_shares"` // 非流通股本

	StockID int64
}

// FianceWithStock 个股经济数据
type FianceWithStock struct {
	ID                  int64   `json:"id"`
	Eps                 float64 `json:"eps"`                   // 每股收益
	Bps                 float64 `json:"bps"`                   // 每股净资产
	TotalShares         float64 `json:"total_shares"`          // 总股本
	CirculationShares   float64 `json:"circulation_shares"`    // 流通股本
	UnCirculationShares float64 `json:"un_circulation_shares"` // 非流通股本

	StockID int64
	Stock   Stock
}

// External 个股附加数据
type External struct {
	ID     int64   `gorm:"primary_key" json:"id"`
	Day3Px float64 `gorm:"day_3_px" json:"day_3_px"` // 近3日涨幅
	Day5Px float64 `gorm:"day_5_px" json:"day_5_px"` // 近5日涨幅

	ProdCode string
}

// Company 个股所属公司
type Company struct {
	ID          int64  `gorm:"primary_key" json:"id"`
	CompanyName string `gorm:"company_name" json:"company_name"`
	CompanyCode string `gorm:"company_code" json:"company_code"`
}
