package goods

type Category struct {
	ID       int    `json:"id" gorm:"column:id;primarykey"`
	ParentID int    `json:"parent_id" gorm:"column:parent_id;type:int;not null;default:10"`
	Name     string `json:"name" gorm:"type:string;size:50;not null;default:'';comment:分类名"`
	Status   uint8  `gorm:"type:tinyint unsigned;not null;default:1;comment:状态 1有效 2失效"`
	SortIdx  int    `gorm:"type:int;not null;default:0;comment:排序码"`
}
