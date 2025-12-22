package controllers

import (
	"errors"

	"github.com/boqier/krm/config"
	"github.com/boqier/krm/utils/logs"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// BaseInit 初始化客户端
func GetClientSet(context *gin.Context) (*kubernetes.Clientset, BasicInfo, error) {
	var basicInfo BasicInfo
	var err error
	requestMethod := context.Request.Method
	if requestMethod == "GET" {
		err = context.ShouldBindQuery(&basicInfo)
	} else if requestMethod == "POST" {
		err = context.ShouldBindJSON(&basicInfo)
	} else {
		err = errors.New("请求方法错误")
	}
	if err != nil {
		logs.Error(nil, "请求出错"+err.Error())
		msg := "请求出错" + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	if basicInfo.ClusterID == "" {
		return nil, basicInfo, errors.New("请求出错123")
	}
	kubeconfig := config.ClusterKubeconfig[basicInfo.ClusterID]
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	if err != nil {

		return nil, basicInfo, errors.New("解析kubeconfig失败")
	}
	//创建客户端工具
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {

		return nil, basicInfo, errors.New("创建客户端失败")
	}
	return clientSet, basicInfo, nil
}
