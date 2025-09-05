package other

// 总部
type HeadQuarters struct {
	ID   int32  `gorm:"column:id;primarykey"`
	Name string `gorm:"column:name;size:50;not null;default:'';comment:名称"`
	Desc string `gorm:"column:desc;size:255;not null;default:'';comment:描述"`
	Mark string `gorm:"column:mark;size:255;not null;default:'';comment:说明"`
}

func (h *HeadQuarters) TableName() string {
	return "other_head_quarters"
}
