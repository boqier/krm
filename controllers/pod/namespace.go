package pod

// 1.对那个集群的哪个namespace进行操作
// 2.kubeconfig，secret:   clusterid--》secret——>data-->kubeconfig-->clientSet
// 将kubeconfig存储在变量中map[string]string  clusterID->kubeconfig
// 3.创建clientset
// 4.操作namespace
