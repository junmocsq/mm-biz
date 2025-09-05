package trade

// 零售
type Retail struct {
	ID        int32  `gorm:"column:id;primarykey"`
	GoodsID   int32  `gorm:"column:goods_id;not null;default:0;comment:商品ID"`
	SaleID    int32  `gorm:"column:sale_id;not null;default:0;comment:销售ID"`
	Name      string `gorm:"column:name;size:255;not null;default:'';comment:商品名称"`
	Num       int32  `gorm:"column:num;not null;default:0;comment:数量"`
	Price     int32  `gorm:"column:price;not null;default:0;comment:单价 分"`
	BranchID  int32  `gorm:"column:branch_id;not null;default:0;comment:分店ID"`
	Uid       string `gorm:"column:uid;size:50;not null;default:'';comment:售货员"`
	Mark      string `gorm:"column:mark;size:255;not null;default:'';comment:备注"`
	PayStatus int8   `gorm:"column:pay_status;type:tinyint;not null;default:1;comment:支付状态 1未支付 2已支付"`
	CreatedAt int64  `gorm:"column:created_at;not null;default:0"`
}

func (r *Retail) TableName() string {
	return "trade_retails"
}
