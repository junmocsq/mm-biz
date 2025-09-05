package goods

type Goods struct {
	ID         int    `json:"id" gorm:"type:int;primarykey"`
	Name       string `json:"name" gorm:"type:string;size:255;not null;default:''"`
	Company    string `json:"company" gorm:"type:string;size:50;not null;default:'';comment:公司名"`
	Brand      string `gorm:"type:string;size:50"`
	Img        string
	ShelfLife  int
	CategoryID int
	Status     int8
	Mark       string
	CreatedAt  int64
}
