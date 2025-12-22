package cluster

import (
	"context"

	"github.com/boqier/krm/config"
	"github.com/boqier/krm/utils"
	"github.com/boqier/krm/utils/logs"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func addOrUpdate(ctx *gin.Context, method string) {
	var arg string
	if method == "create" {
		arg = "添加"
	} else if method == "update" {
		arg = "更新"
	}
	clusterConfig := ClusterConfig{}
	returnData := config.NewReturnDate()
	if err := ctx.ShouldBindJSON(&clusterConfig); err != nil {
		returnData.Status = 400
		returnData.Message = arg + "集群配置信息不完整" + err.Error()
		ctx.JSON(200, returnData)
		return
	}
	//判断集群状态是否正常
	clusterStatus, err := clusterConfig.GetClusterStatus()
	if err != nil {
		returnData.Status = 400
		returnData.Message = arg + "集群配置失败,无法获取集群信息" + err.Error()
		logs.Error(map[string]interface{}{"集群名称": clusterConfig.ClusterInfo.DisplayName, "集群ID": clusterConfig.ClusterInfo.ID, "err": err}, arg+"集群配置失败,无法获取集群信息")
		ctx.JSON(200, returnData)
		return
	}
	logs.Info(map[string]interface{}{"集群名称": clusterConfig.ClusterInfo.DisplayName, "集群ID": clusterConfig.ClusterInfo.ID}, arg+"集群配置成功")
	var clusterConfigSecret corev1.Secret
	clusterConfigSecret.Name = clusterConfig.ClusterInfo.ID
	clusterConfigSecret.Namespace = config.MetadataNamespace
	//加标签
	clusterConfigSecret.Labels = make(map[string]string)
	clusterConfigSecret.Labels[config.ClusterConfigSecretLabelKey] = config.ClusterConfigSecretLabelValue
	//添加注释,将结构体转map
	m := utils.Struct2Map(clusterStatus)
	clusterConfigSecret.Annotations = m
	clusterConfigSecret.StringData = map[string]string{
		"kubeconfig": clusterConfig.Kubeconfig,
	}
	if method == "create" {
		_, err = config.ClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Create(context.TODO(), &clusterConfigSecret, metav1.CreateOptions{})
	} else if method == "update" {
		_, err = config.ClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Update(context.TODO(), &clusterConfigSecret, metav1.UpdateOptions{})
	}
	if err != nil {
		logs.Error(map[string]interface{}{"集群Name": clusterConfig.ClusterInfo.DisplayName, "err": err}, arg+"集群配置失败")
		returnData.Status = 400
		returnData.Message = arg + "集群配置失败" + err.Error()
		ctx.JSON(200, returnData)
		return
	}
	config.ClusterKubeconfig[clusterConfig.ClusterInfo.ID] = clusterConfig.Kubeconfig
	logs.Info(map[string]interface{}{"集群Name": clusterConfig.ClusterInfo.DisplayName, "集群ID": clusterConfig.ClusterInfo.ID}, arg+"集群配置成功")
	returnData.Message = arg + "成功"
	returnData.Status = 200
	ctx.JSON(200, returnData)
}
