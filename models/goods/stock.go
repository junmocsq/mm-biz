package goods

// 库存
type Stock struct {
	ID       int32 `gorm:"column:id;primarykey"`
	GoodsID  int32 `gorm:"column:goods_id;not null;default:0;comment:商品ID"`
	Stocks   int32 `gorm:"column:stocks;not null;default:0;comment:库存数量"`
	BranchID int32 `gorm:"column:branch_id;not null;default:0;comment:分店ID"`
}

func (i *Stock) TableName() string {
	return "goods_stocks"
}
