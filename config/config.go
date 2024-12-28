package config

// 网关配置=服务+网关
type AppGatewayConfig struct {
	*AppServiceConfig
	*GatewayConfig
}

// 服务配置=数据+服务
type AppServiceConfig struct {
	*RepoConfig
	*ServiceConfig
}


// 网关配置
type GatewayConfig struct {
	GatewayHttpPort int `json:"http_port"` // 网关对外暴露的http端口
}
