package initcontroller

import (
	"context"

	"github.com/boqier/krm/config"
	"github.com/boqier/krm/utils/logs"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var Clusterconfig *rest.Config

func InitK8sConfig() {

	var err error

	// 优先尝试从集群内部加载配置
	Clusterconfig, err = rest.InClusterConfig()
	if err != nil {
		logs.Info(map[string]interface{}{"err": err}, "get in cluster config failed, try get from host")
		// 从外部文件加载配置
		Clusterconfig, err = clientcmd.BuildConfigFromFlags("", "/etc/rancher/k3s/k3s.yaml")
		if err != nil {
			logs.Error(map[string]interface{}{"err": err}, "get config failed")
			panic(err)
		}
	}

	// 创建客户端
	config.ClusterClientSet, err = kubernetes.NewForConfig(Clusterconfig)
	if err != nil {
		logs.Error(map[string]interface{}{"err": err}, "clientSet创建失败")
		panic(err)
	}
	logs.Info(map[string]interface{}{}, "加载客户端成功")

	// 检查并创建元数据命名空间
	_, err = config.ClusterClientSet.CoreV1().Namespaces().Get(context.TODO(), config.MetadataNamespace, metav1.GetOptions{})
	if err != nil {
		logs.Info(map[string]interface{}{"Namespace": config.MetadataNamespace}, "无元数据命名空间，尝试创建")
		var metadataNameSpace corev1.Namespace
		metadataNameSpace.Name = config.MetadataNamespace

		_, err = config.ClusterClientSet.CoreV1().Namespaces().Create(context.TODO(), &metadataNameSpace, metav1.CreateOptions{})
		if err != nil {
			logs.Error(map[string]interface{}{"Namespace": config.MetadataNamespace}, "创建元数据命名空间失败")
			panic(err)
		}
		logs.Info(map[string]interface{}{"Namespace": config.MetadataNamespace}, "创建元数据命名空间成功")
	} else {
		logs.Info(map[string]interface{}{"Namespace": config.MetadataNamespace}, "元数据命名空间已存在无需创建")
	}
	//初始化kubeconfig，存在config的map里。
	ListOptions := metav1.ListOptions{
		LabelSelector: config.ClusterConfigSecretLabelKey + "=" + config.ClusterConfigSecretLabelValue,
	}
	config.ClusterKubeconfig = make(map[string]string)
	secretList, err := config.ClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).List(context.TODO(), ListOptions)
	if err != nil {
		logs.Error(map[string]interface{}{"err": err}, "获取集群配置secret失败")
		panic(err)
	}
	for _, v := range secretList.Items {
		clusterID := v.Name
		kubeconfig := string(v.Data["kubeconfig"])
		config.ClusterKubeconfig[clusterID] = string(kubeconfig)
	}
	logs.Debug(map[string]interface{}{"集群配置": config.ClusterKubeconfig}, "初始化kubeconfig成功")
}
