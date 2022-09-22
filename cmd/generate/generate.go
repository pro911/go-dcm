package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const dsnTpl = "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local"

func main() {
	dsn := fmt.Sprintf(dsnTpl, "test_read_only", "db#01^st$Post", "rm-2zem14s80lyu5c4z7.mysql.rds.aliyuncs.com", 3306, "apipost_dbv3_test_720", "utf8mb4")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("connect mysql fail: %w", err))
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/pkg/dal/query",
	})

	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("user"),              //用户主表
		g.GenerateModel("team"),              //团队表
		g.GenerateModel("project"),           //项目表
		g.GenerateModel("target"),            //APIs
		g.GenerateModel("project_env"),       //项目调试环境表
		g.GenerateModel("project_user"),      //项目和用户关联表
		g.GenerateModel("target_mark"),       //APIS接口状态
		g.GenerateModel("team_invite_email"), //邮箱邀请加入项目或团队 /待注册
		g.GenerateModel("team_place"),        //团队工位表
		g.GenerateModel("user_logout"),       //用户注销表，记录用户注销前的状态
	)
	g.Execute()
}
