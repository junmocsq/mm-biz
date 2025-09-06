package trade

// 批发单
type Wholesale struct {
	OrderID    string `gorm:"column:order_id;size:40;primarykey"`
	TotalPrice int32  `gorm:"column:total_price;not null;default:0;comment:总价"`
	PayPrice   int32  `gorm:"column:pay_price;not null;default:0;comment:实际支付价"`
	Status     uint8  `gorm:"not null;default:2;comment:1支付 2未支付"`
	Class      uint8  `gorm:"not null;default:1;comment:分类 1销售 2退货 3陈列 4废弃 5报损"`
	TotalNum   int32  `gorm:"column:total_num;not null;default:0;comment:总数量"`
	BranchID   int32  `gorm:"column:branch_id;not null;default:0;comment:分店ID"`
	UID        int32  `gorm:"index;column:uid;not null;default:0;comment:操作员ID"`
	WID        int32  `gorm:"index;column:wid;not null;default:0;comment:批发商ID"`
	Mark       string `gorm:"column:mark;size:255;not null;default:'';comment:备注"`
	UpdatedAt  int64  `gorm:"column:updated_at;not null;default:0"`
	CreatedAt  int64  `gorm:"column:created_at;not null;default:0"`
}

func (w *Wholesale) TableName() string {
	return "trade_wholesales"
}
