package config

// 这里是定义表的  操作数据库表的结构体 是什么表就写什么文件名
type User struct { // 默认表名是 `users`
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

// 表示配置表的名称   这里就指定操作的表了
func (User) TableName() string {
	return "user"
}
