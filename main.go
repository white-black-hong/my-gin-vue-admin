package main

import (
	"my-gin-vue-admin/core"
	"my-gin-vue-admin/global"
	"my-gin-vue-admin/initialize"
)

func main()  {
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = initialize.Gorm()
	initialize.MysqlTables(global.GVA_DB)

}
