package user

// 批发商
type Wholesaler struct {
	ID        int32  `gorm:"column:id;primarykey"`
	Name      string `gorm:"column:name;size:50;not null;default:'';comment:名称"`
	Mobile    string `gorm:"column:mobile;size:20;not null;default:'';comment:手机号"`
	Address   string `gorm:"column:address;size:255;not null;default:'';comment:地址"`
	Status    uint8  `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:1生效 2失效"`
	UpdatedAt int64  `gorm:"column:updated_at;type:bigint;not null;default:0"`
	CreatedAt int64  `gorm:"column:created_at;type:bigint;not null;default:0"`
}

func (w *Wholesaler) TableName() string {
	return "wholesalers"
}
