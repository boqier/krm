package temp

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 描述创建集群所用的配置信息,通过serect来管理
type ClusterConfig struct {
	ClusterInfo
	Kubeconfig string `json:"kubeconfig"`
}
type ClusterInfo struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"` //集群别名
	District    string `json:"district"`    //集群所在区县
	City        string `json:"city"`        //集群所在城市
}

// 描述集群状态
type ClusterStatus struct {
	ClusterInfo
	Version string `json:"version"` //集群版本
	Status  string `json:"status"`  //集群状态

}

// 判断集群状态
func (c *ClusterConfig) GetClusterStatus() (ClusterStatus, error) {
	ClusterStatus := ClusterStatus{
		ClusterInfo: c.ClusterInfo,
	}
	//通过字符串创建客户端工具
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(c.Kubeconfig))
	if err != nil {
		ClusterStatus.Status = "error"
		return ClusterStatus, err
	}
	//创建客户端工具
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		ClusterStatus.Status = "error"
		return ClusterStatus, err
	}
	ServerVersion, err := clientSet.Discovery().ServerVersion()
	if err != nil {
		ClusterStatus.Status = "error"
		return ClusterStatus, err
	}
	ClusterStatus.Status = "active"
	ClusterStatus.Version = ServerVersion.String()
	return ClusterStatus, nil
}
