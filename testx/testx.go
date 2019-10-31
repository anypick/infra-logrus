package testx

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra-logrus"
	"github.com/anypick/infra/base/props"
	"github.com/anypick/infra/base/props/container"
	"github.com/sirupsen/logrus"
)

func Init() {
	infra.Register(&container.YamlStarter{})


	baselog.Init()

	infra.Register(&infra.BaseInitializerStarter{})

	source := props.NewYamlSource("./config.yml")

	app := infra.New(*source)
	app.Start()

	logrus.Info("start...")

}

// 测试logrus
func LogrusX() {
	// Nothing
}
