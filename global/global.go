package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"github.com/spf13/viper"
	"my-gin-vue-admin/config"
	"github.com/go-redis/redis"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
)
