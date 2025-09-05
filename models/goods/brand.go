package goods

type Brand struct {
	ID        int    `gorm:"column:id;primarykey"`
	CompanyID int    `gorm:"column:company_id"`
	Name      string `gorm:"column:name;size:50;comment:品牌名"`
	Status    uint8  `gorm:"type:tinyint;not null;default:1;comment:1 生效 2 失效"`
	SortIdx   int    `gorm:"column:sort_idx;not null;default:0;comment:排序码"`
}
