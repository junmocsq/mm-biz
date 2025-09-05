package other

// 分店
type Branch struct {
	ID      int32  `gorm:"column:id;primarykey"`
	Name    string `gorm:"column:name;size:50;not null;default:'';comment:分店名称"`
	Status  uint8  `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:1生效 2失效"`
	SortIdx int32  `gorm:"column:sort_idx;not null;default:0;comment:排序码"`
}

func (b *Branch) TableName() string {
	return "other_branches"
}
