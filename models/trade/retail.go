package trade

// 零售
type Retail struct {
	OrderID    string  `gorm:"column:order_id;size:40;primarykey"`
	TotalNum   float64 `gorm:"column:total_num;not null;default:0;comment:总数量"`
	TotalPrice int32   `gorm:"column:total_price;not null;default:0;comment:总价 分"`
	TotalPay   int32   `gorm:"column:total_pay;not null;default:0;comment:现金 分"`
	Status     int8    `gorm:"column:status;type:tinyint;not null;default:1;comment:单据状态 1未结算 2已结算 3已取消"`
	Class      int8    `gorm:"column:class;type:tinyint;not null;default:1;comment:单据类别 1销售 2退货 3内部 4废弃 5报损"`
	BranchID   int32   `gorm:"column:branch_id;not null;default:0;comment:分店ID"`
	Uid        string  `gorm:"column:uid;size:50;not null;default:'';comment:售货员"`
	Mark       string  `gorm:"column:mark;size:255;not null;default:'';comment:备注"`
	CreatedAt  int64   `gorm:"column:created_at;not null;default:0"`
}

func (r *Retail) TableName() string {
	return "trade_retails"
}
