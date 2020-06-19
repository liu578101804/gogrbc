package cache

import (
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego"
	"time"
	"fmt"
	"encoding/json"
	"errors"
)

var (
	bm cache.Cache
	cacheType string
)

const (
	CacheType_Redis = "redis"
	CacheType_Memory = "memory"
)
const (
	CacheTypeError = "未知缓存类型"
)

func init()  {

	var (
		err error
		cacheConf string
	)

	//缓存类型
	cacheType = beego.AppConfig.DefaultString("cacheType","memory")

	beego.Info("缓存类型：", cacheType)

	switch cacheType {
	case CacheType_Memory:
		cleanTime := beego.AppConfig.DefaultInt("cleanTime", 60)
		cacheConf = fmt.Sprintf(`{"interval":%d}`,cleanTime)
	case CacheType_Redis:
		redisKey := beego.AppConfig.DefaultString("redisKey","dev")
		redisConn := beego.AppConfig.DefaultString("redisConn",":6379")
		redisDbNum := beego.AppConfig.DefaultInt("redisDbNum",0)
		redisPassword := beego.AppConfig.DefaultString("redisPassword","")

		cacheConf = fmt.Sprintf(`{"key":"%s","conn":"%s","dbNum":"%d","password":"%s"}`,redisKey,redisConn,redisDbNum,redisPassword)
	default:
		panic(errors.New(CacheTypeError))
	}

	bm, err = cache.NewCache(cacheType , cacheConf)

	if err != nil {
		beego.Error(err)
	}
}

func Set(key string, value interface{}, timed int) {

	switch cacheType {

	case CacheType_Memory:
		bm.Put(key, value, time.Duration(timed) * time.Second)

	case CacheType_Redis:
		byt,_ := json.Marshal(value)
		err := bm.Put(key, string(byt), time.Duration(timed) * time.Second)
		if err != nil {
			beego.Error(err)
		}

	default:

		panic(errors.New(CacheTypeError))
	}
}

func Get(key string) interface{} {

	var tmp interface{}

	switch cacheType {

	case CacheType_Memory:
		tmp = bm.Get(key)

	case CacheType_Redis:
		res := bm.Get(key).([]uint8)
		json.Unmarshal([]byte(res), &tmp)

	default:
		panic(errors.New(CacheTypeError))

	}

	return tmp
}

func IsExist(key string) bool {
	return bm.IsExist(key)
}

func Delete(key string) error {
	return bm.Delete(key)
}

const (
	C_RolePath = "cache_path_%v"
	C_UserRole = "cache_user_auth_%v"
)