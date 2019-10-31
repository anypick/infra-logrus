/**
用户可能会自定义一些hook,这里将用户定义的hook装载到变量@hooks中，
 */
package helper

import "github.com/sirupsen/logrus"

var hooks = make([]logrus.Hook, 0)

func AddHook(hook logrus.Hook) {
	hooks = append(hooks, hook)
}

func GetHook() []logrus.Hook {
	return hooks
}