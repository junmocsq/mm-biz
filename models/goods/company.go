package goods

// 公司
type Company struct {
	ID      int32  `gorm:"column:id;primarykey"`
	Name    string `gorm:"column:name;size:50;not null;default:''"`
	Status  uint8  `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:1生效 2无效"`
	SortIdx int32  `gorm:"column:sort_idx;not null;default:0;comment:排序码"`
}

func (c *Company) TableName() string {
	return "goods_companies"
}
