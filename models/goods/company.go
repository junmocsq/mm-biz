package goods

type Company struct {
	ID      int    `gorm:"column:id;primarykey"`
	Name    string `gorm:"column:name;size:50;not null;default:''"`
	Status  uint8  `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:1生效 2无效"`
	SortIdx int    `gorm:"column:sort_idx;not null;default:0;comment:排序码"`
}
