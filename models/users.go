package models

// TableName 指定用户表表名.
func (User) TableName() string {
	return "users"
}

// User 定义用户表对应的结构.
type User struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Gender    int8   `json:"gender"`
	AvatarURL string `json:"avatar_url"`
	Status    int8   `json:"status"`
	CreateAt  string `json:"create_at"`
	UpdateAt  string `json:"update_at"`
}

// GetUserList 获取用户列表.
func GetUserList() (users []User, err error) {
	if err = dbConn.Find(&users).Error; err != nil {
		return
	}
	return
}
