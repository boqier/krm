package main

import (
	"github.com/boqier/krm/utils/logs"
	kubernetes "k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/etc/rancher/k3s/k3s.yaml")
	if err != nil {
		logs.Error(map[string]interface{}{"module": "main"}, "读取k3s.yaml失败")
		panic(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
		logs.Error(map[string]interface{}{"module": "main"}, "创建clientSet失败")
	}

}
