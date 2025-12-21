package initcontroller

import (
	"github.com/boqier/krm/utils/logs"
)

func init() {
	//初始化kubeconfig
	logs.Info(map[string]interface{}{"kubeconfig": "/root/.kube/config"}, "init kubeconfig")
	InitK8sConfig()
}
