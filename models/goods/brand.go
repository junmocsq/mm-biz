package goods

// 品牌
type Brand struct {
	ID        int32  `gorm:"column:id;primarykey"`
	CompanyID int32  `gorm:"column:company_id;not null;default:0;comment:公司ID"`
	Name      string `gorm:"column:name;size:50;comment:品牌名;not null;default:''"`
	Status    uint8  `gorm:"type:tinyint unsigned;not null;default:1;comment:1 生效 2 失效"`
	SortIdx   int32  `gorm:"column:sort_idx;not null;default:0;comment:排序码"`
}

func (b *Brand) TableName() string {
	return "goods_brands"
}
