// 统一接口使用示例
// 更多示例请查看：
//   - example/client/main.go  - 本地 Service API 详细示例
//   - example/central/main.go - 云端 Central API 详细示例
package main

import (
	"fmt"
	"log"

	zerotier "github.com/fromsko/zerotier-sdk"
)

func main() {
	fmt.Println("========== ZeroTier SDK 统一接口示例 ==========")

	// ========================================
	// 本地 Service API
	// ========================================
	fmt.Println("\n--- 本地节点 ---")
	local := zerotier.NewClient()

	status, err := local.Status()
	if err != nil {
		log.Printf("本地状态获取失败: %v", err)
	} else {
		fmt.Printf("节点: %s (v%s)\n", status.Address, status.Version)
	}

	networks, _ := local.Networks().List()
	fmt.Printf("已加入网络: %d 个\n", len(networks))

	peers, _ := local.Peers().List()
	fmt.Printf("Peers: %d 个\n", len(peers))

	// ========================================
	// 云端 Central API
	// ========================================
	fmt.Println("\n--- 云端管理 ---")
	cloud := zerotier.NewCentral("your_api_token")

	cloudStatus, err := cloud.Status()
	if err != nil {
		log.Printf("云端状态获取失败: %v", err)
	} else {
		fmt.Printf("API: v%s\n", cloudStatus.APIVersion)
		if cloudStatus.User != nil {
			fmt.Printf("用户: %s\n", cloudStatus.User.DisplayName)
		}
	}

	cloudNetworks, _ := cloud.Networks().List()
	fmt.Printf("云端网络: %d 个\n", len(cloudNetworks))

	// ========================================
	// Builder 模式
	// ========================================
	fmt.Println("\n--- Builder 示例 ---")

	// 本地网络设置
	localSettings := zerotier.NewNetworkSettings().
		AllowDNS(true).
		AllowManaged(true).
		Build()
	fmt.Printf("本地设置: AllowDNS=%v\n", *localSettings.AllowDNS)

	// 云端网络配置
	cloudConfig := zerotier.NewCentralNetworkConfig().
		Name("Demo Network").
		Private(true).
		AddRoute("10.0.0.0/24", nil).
		AddIPPool("10.0.0.1", "10.0.0.254").
		V4AssignMode(true).
		Build()
	fmt.Printf("云端配置: Name=%s\n", cloudConfig.Name)

	// 云端成员配置
	memberConfig := zerotier.NewCentralMemberConfig().
		Name("my-device").
		Authorized(true).
		Build()
	fmt.Printf("成员配置: Name=%s\n", memberConfig.Name)

	fmt.Println("\n========== 完成 ==========")
	fmt.Println("更多示例请查看 example/client 和 example/central 目录")
}
