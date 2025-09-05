package goods

// 商品销售单价
type Sale struct {
	ID        int32  `gorm:"column:id;primarykey"`
	GoodsID   int32  `gorm:"column:goods_id;not null;default:0;comment:商品ID"`
	UnitID    int32  `gorm:"column:unit_id;not null;default:0;comment:单位ID"`
	Price     int32  `gorm:"column:price;not null;default:0;comment:销售单价 分"`
	DeStocks  int32  `gorm:"column:de_stocks;not null;default:0;comment:库存扣减数"`
	Pics      string `gorm:"column:pics;size:255;not null;default:'';comment:图片，逗号分隔"`
	Barcode   string `gorm:"column:barcode;size:50;not null;default:'';comment:条形码"`
	Spec      string `gorm:"column:spec;size:50;not null;default:'';comment:规格描述"`
	Class     uint8  `gorm:"column:class;type:tinyint unsigned;not null;default:1;comment:类别 1零售 2批发"`
	Status    uint8  `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:状态 1生效 2失效"`
	Mark      string `gorm:"column:mark;size:255;not null;default:'';comment:备注"`
	CreatedAt int64  `gorm:"column:created_at;type:bigint;not null;default:0"`
}

func (s *Sale) TableName() string {
	return "goods_sales"
}
