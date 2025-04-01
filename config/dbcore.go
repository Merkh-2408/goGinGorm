package config

/*
连接数据库的配置，模版固定下来，只改IP 用户名等一些东西就好了
**/

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 保存数据库的链接的 一个全局变量
// var DB *gorm.DB //后续用于操作数据库
// var err error

func Init() {

	//root:root@tcp(127.0.0.1:3306)/test？
	//和数据库建立连接，返回db对象 在27行
	dsn := "root:123456@tcp(127.0.0.1:3306)/gotest?"
	//连接mysql，获得DB类型实例，用于后面数据库的读写操作
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败，error=" + err.Error())
	}
	// 自动迁移
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate")
	}

	// u1 := User{1, "HWF", 18, "Abc@12.com", 321}

	// // 创建记录
	// db.Create(&u1)
	// db.Create(&User{2, "OK", 16, "足球@12.com", 123})
	// fmt.Println("创建成功")
	// 查询
	// var u = new(User)
	// db.First(&u, 1) // 根据整型主键查找
	// fmt.Printf("%#v\n", u)

	// var uu User
	// db.First(&uu, "age=?", "16")
	// fmt.Printf("%#v\n", uu)

	// // 更新
	// db.Model(&u).Update("age", "17")
	// //Update - 更新多个字段
	// db.Model(&u).Updates(User{Username: "HWF", Email: "跳舞"}) // 仅更新非零值字段
	// // db.Model(&uu).Updates(map[string]interface{}{"Hobby": "演戏", "Name": "赵丽颖"})

	// // 删除
	// db.Delete(&u, 1)

}
