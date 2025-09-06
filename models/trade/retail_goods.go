package trade

// 零售商品
type RetailGoods struct {
	ID        int32   // 主键ID
	Name      string  `gorm:"size:50;not null;default:''"`                       // 商品名称
	GoodsID   int32   `gorm:"column:goods_id;not null;default:0"`                // 商品ID
	SaleID    int32   `gorm:"column:sale_id;not null;default:0"`                 // 销售ID
	Price     int32   `gorm:"not null;default 0;comment:单价 分"`                   // 商品价格
	Num       float64 `gorm:"type:decimal(12,3);not null;default:1.0"`           // 商品数量
	Unit      int32   `gorm:"not null;default 0;comment:单位"`                     // 商品单位
	Spec      string  `gorm:"size:50;not null;default:''"`                       // 规格
	DeStocks  int32   `gorm:"not null;default:0;comment:每单位库存扣减数"`               // 库存扣减数
	Mark      string  `gorm:"size:255;not null;default:''"`                      // 备注
	OrderID   string  `gorm:"index;column:order_id;size:40;not null;default:''"` // 零售单ID
	UpdatedAt int64   `gorm:"not null;default:0"`                                // 更新时间
	CreatedAt int64   `gorm:"not null;default:0"`                                // 创建时间
}
