package user

// 用户
type User struct {
	ID        int32  `gorm:"column:id;primarykey"`
	Username  string `gorm:"column:username;size:50;not null;default:'';comment:用户名"`
	Password  string `gorm:"column:password;size:255;not null;default:'';comment:密码"`
	Nickname  string `gorm:"column:nickname;size:50;not null;default:'';comment:昵称"`
	Mobile    string `gorm:"column:mobile;size:20;not null;default:'';comment:手机号"`
	OpenID    string `gorm:"column:open_id;size:100;not null;default:'';comment:微信OpenID"`
	Status    uint8  `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:1生效 2失效"`
	BranchID  int32  `gorm:"column:branch_id;not null;default:0;comment:分店ID"`
	RoleID    int32  `gorm:"column:role_id;not null;default:0;comment:角色ID"`
	Avatar    string `gorm:"column:avatar;size:255;not null;default:'';comment:头像"`
	UpdatedAt int64  `gorm:"column:updated_at;type:bigint;not null;default:0"`
	CreatedAt int64  `gorm:"column:created_at;type:bigint;not null;default:0"`
}

func (u *User) TableName() string {
	return "users"
}
