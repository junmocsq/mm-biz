package trade

// 批发单
type Wholesale struct {
	OrderID    string `gorm:"column:order_id;size:40;primarykey"`
	TotalPrice int32  `gorm:"column:total_price;not null;default:0;comment:总价"`
	PayPrice   int32  `gorm:"column:pay_price;not null;default:0;comment:支付价"`
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
