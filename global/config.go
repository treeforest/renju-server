package global

import (
	"github.com/treeforest/gos/utils/config"
	"github.com/treeforest/gos/utils/config/encoder/yaml"
	"github.com/treeforest/gos/utils/config/source"
	"github.com/treeforest/gos/utils/config/source/file"
	"github.com/treeforest/logger"
	"os"
	"path/filepath"
)

func initConfig() error {
	// 创建配置文件对象
	var err error = nil

	globalConfig, err = config.NewConfig()
	if err != nil {
		log.Errorf("new config error: %v", err)
		return err
	}

	// 加载配置文件
	path := filepath.Join(os.Getenv("GoPath"), "src", "github.com", "treeforest", "renju-server", "config", "renju.yaml")
	if err = globalConfig.Load(file.NewSource(
		file.WithPath(path),
		source.WithEncoder(yaml.NewEncoder()),
	)); err != nil {
		log.Errorf("config load error: %v", err)
		return err
	}

	return nil
}
