package common

//资源列表结构体，可为用户需求资源和总资源结构体列表
type Resource struct {
	Cpu   int
	Mem   int
	Time  float64
	Level int
}

//某个用户获得的资源结构体
/*type UsedResource struct {
	//用户需求满足状态：已满足、等待
	State string
	//用户获得的CPU数目
	GetCpu int
	//用户获得的内存数目
	GetMem int
	//用户分配的端口数目或其他唯一识别标识, 不包括主节点
	IdNum int
	//用户分配的主节点，即用户click对应的节点
	MaterNodeId string
	//用户获得的所有端口或其他唯一识别标识列表
	Id []string
}*/
