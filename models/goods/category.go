package goods

// 商品分类
type Category struct {
	ID       int32  `json:"id" gorm:"column:id;primarykey"`
	ParentID int32  `json:"parent_id" gorm:"column:parent_id;not null;default:0;comment:父分类ID"`
	Name     string `gorm:"type:string;size:50;not null;default:'';comment:分类名"`
	Status   uint8  `gorm:"type:tinyint unsigned;not null;default:1;comment:状态 1有效 2失效"`
	SortIdx  int32  `gorm:"type:int;not null;default:0;comment:排序码"`
}

func (c *Category) TableName() string {
	return "goods_categories"
}
