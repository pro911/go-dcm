package internal

import (
	"apipost-dcm/internal/pkg/conf"
	"apipost-dcm/internal/pkg/dal"
)

func InitProjects() {
	conf.MustInitConf()
	dal.MustInitMySQL()
	dal.MustInitMongo()
	dal.MustInitRedis()
}
