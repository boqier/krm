package controllers

// 定义全局数据结构
type BasicInfo struct {
	ClusterID  string      `json:"clusterid" form:"clusterid"`
	Namespace  string      `json:"namespace" form:"namespace"`
	Name       string      `json:"name" form:"name"`
	Item       interface{} `json:"item"`
	DeleteList []string    `json:"deleteList"`
}
