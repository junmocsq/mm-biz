package goods

// 商品
type Goods struct {
	ID         int32  `json:"id" gorm:"type:int;primarykey"`
	Name       string `json:"name" gorm:"type:string;size:50;not null;default:''"`
	UnitID     int32  `gorm:"type:int;not null;default:0;comment:单位ID"`
	CostPrice  int32  `gorm:"type:int;not null;default:0;comment:成本价 分"`
	PackageNum int32  `gorm:"type:int;not null;default:0;comment:每件包装数量"`
	Img        string `gorm:"type:string;size:255;not null;default:'';comment:商品图片"`
	ShelfLife  int32  `gorm:"type:int;not null;default:0;comment:保质期 天"`
	CategoryID int32  `gorm:"type:int;not null;default:0;comment:分类ID"`
	CompanyID  int32  `gorm:"type:int;not null;default:0;comment:公司ID"`
	BrandID    int32  `gorm:"type:int;not null;default:0;comment:品牌ID"`
	Status     int8   `gorm:"type:tinyint unsigned;not null;default:1;comment:状态 1上架 2下架"`
	Mark       string `gorm:"type:string;size:255;not null;default:'';comment:商品备注"`
	CreatedAt  int64  `gorm:"type:bigint;not null;default:0"`
}

func (g *Goods) TableName() string {
	return "goods"
}
