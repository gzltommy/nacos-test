package main

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func main() {
	// 创建 clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "e525eafa-f7d7-4029-83d9-008937f9d468", // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建 clientConfig 的另一种方式
	//clientConfig := *constant.NewClientConfig(
	//	constant.WithNamespaceId("e525eafa-f7d7-4029-83d9-008937f9d468"), //当namespace是public时，此处填空字符串。
	//	constant.WithTimeoutMs(5000),
	//	constant.WithNotLoadCacheAtStart(true),
	//	constant.WithLogDir("/tmp/nacos/log"),
	//	constant.WithCacheDir("/tmp/nacos/cache"),
	//	constant.WithLogLevel("debug"),
	//)

	// 至少一个 ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "console1.nacos.io",
			ContextPath: "/nacos",
			Port:        80,
			Scheme:      "http",
		},
		{
			IpAddr:      "console2.nacos.io",
			ContextPath: "/nacos",
			Port:        80,
			Scheme:      "http",
		},
	}

	// 创建 serverConfig 的另一种方式
	//serverConfigs := []constant.ServerConfig{
	//	*constant.NewServerConfig(
	//		"console1.nacos.io",
	//		80,
	//		constant.WithScheme("http"),
	//		constant.WithContextPath("/nacos"),
	//	),
	//	*constant.NewServerConfig(
	//		"console2.nacos.io",
	//		80,
	//		constant.WithScheme("http"),
	//		constant.WithContextPath("/nacos"),
	//	),
	//}

	//// 创建服务发现客户端
	//_, _ := clients.CreateNamingClient(map[string]interface{}{
	//	"serverConfigs": serverConfigs,
	//	"clientConfig":  clientConfig,
	//})
	//
	//// 创建动态配置客户端
	//_, _ := clients.CreateConfigClient(map[string]interface{}{
	//	"serverConfigs": serverConfigs,
	//	"clientConfig":  clientConfig,
	//})

	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, _ := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	_ = namingClient
	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, _ := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	_ = configClient
}
