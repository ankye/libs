package utils

import (
	"github.com/gonethopper/libs/common"
	"github.com/gonethopper/libs/config"
	log "github.com/gonethopper/libs/logs"
	jsoniter "github.com/json-iterator/go"
)

//LoadLogConfig load log config from log file, if get some error than return nil and error
func LoadLogConfig(configFile string) (*log.LogConfig, error) {
	c := log.NewLogConfig()
	if err := config.ParseConfig(c, configFile); err != nil {
		return nil, err
	}

	//try create log dir
	MakeDirByFile(configFile)

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(&c.Params)
	if err != nil {
		log.Error("json marshal config error:", err)
		return nil, err
	}

	log.SetLogger(c.Adapter, string(b))
	//日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号
	//log.EnableFuncCallDepth(true)

	return c, nil
}

//LoadRouterConfig load router config
func LoadRouterConfig(configFiles []string) (map[int]*common.RouterServer, error) {
	routers := make(map[int]*common.RouterServer)

	for index, file := range configFiles {
		router := new(common.RouterServer)
		err := config.ParseConfig(router, file)
		if err != nil {
			log.Error("read router config failed .", index, file, err)
			return nil, err
		}
		routers[router.DestType] = router
	}
	return routers, nil
}
