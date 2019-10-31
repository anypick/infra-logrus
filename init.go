package baselog

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra-logrus/config"
	"github.com/anypick/infra/base/props/container"
)

// 初始化资源，由用户调用
func Init() {
	container.RegisterYamContainer(&config.LogConfig{Prefix: config.YamlPrefix})
	infra.Register(&LogrusStarter{})
}