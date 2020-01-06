package baselog

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra/base/props"
	"github.com/anypick/infra/base/props/container"
	"github.com/anypick/infra/utils/common"
	"github.com/sirupsen/logrus"
	"testing"
)

func InitT() {
	infra.Register(&container.YamlStarter{})
	Init()

	infra.Register(&infra.BaseInitializerStarter{})

	source := props.NewYamlSource(common.GetAbsolutePath("./testx/config.yml", 1))

	app := infra.New(*source)
	app.Start()

}

//  测试
func Test_Function(t *testing.T) {
	InitT()
	logrus.WithField("type", "get").WithField("instance", "logrus").Info("hello world")
}

func Test_Self(t *testing.T) {

}
