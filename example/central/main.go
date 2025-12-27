// 云端 Central API 使用示例
package main

import (
	"fmt"
	"log"

	zerotier "github.com/fromsko/zerotier-sdk"
)

func main() {
	fmt.Println("========== ZeroTier Central API（云端管理）==========")

	// ========================================
	// 1. 创建客户端
	// ========================================

	// API Token 获取方式：
	// 1. 登录 https://my.zerotier.com
	// 2. 进入 Account 页面
	// 3. 创建 API Token

	apiToken := "your_api_token_here" // 替换为你的 Token

	// 方式一：默认配置
	client := zerotier.NewCentral(apiToken)

	// 方式二：自定义配置
	// client := zerotier.NewCentral(apiToken,
	// 	zerotier.WithCentralBaseURL("https://api.zerotier.com/api/v1"),
	// )

	// ========================================
	// 2. 获取状态
	// ========================================
	fmt.Println("\n=== Central 状态 ===")
	status, err := client.Status()
	if err != nil {
		log.Printf("获取状态失败: %v", err)
	} else {
		fmt.Printf("API 版本: %s\n", status.APIVersion)
		fmt.Printf("版本: %s\n", status.Version)
		fmt.Printf("只读模式: %v\n", status.ReadOnlyMode)
		if status.User != nil {
			fmt.Printf("当前用户: %s\n", status.User.DisplayName)
			fmt.Printf("用户ID: %s\n", status.User.ID)
		}
	}

	// ========================================
	// 3. 网络管理
	// ========================================
	fmt.Println("\n=== 网络列表 ===")
	networks, err := client.Networks().List()
	if err != nil {
		log.Printf("获取网络列表失败: %v", err)
	} else {
		if len(networks) == 0 {
			fmt.Println("暂无网络")
		}
		for _, n := range networks {
			fmt.Printf("网络: %s\n", n.ID)
			fmt.Printf("  名称: %s\n", n.Config.Name)
			fmt.Printf("  私有: %v\n", n.Config.Private)
			fmt.Printf("  MTU: %d\n", n.Config.MTU)
			fmt.Printf("  在线成员: %d\n", n.OnlineMemberCount)
			fmt.Printf("  授权成员: %d\n", n.AuthorizedMemberCount)
			fmt.Printf("  总成员数: %d\n", n.TotalMemberCount)
			if len(n.Config.IPAssignmentPools) > 0 {
				pool := n.Config.IPAssignmentPools[0]
				fmt.Printf("  IP池: %s - %s\n", pool.IPRangeStart, pool.IPRangeEnd)
			}
		}
	}

	// 获取网络详情
	// network, err := client.Networks().Get("network_id")

	// 创建网络
	// config := zerotier.NewCentralNetworkConfig().
	// 	Name("My New Network").
	// 	Private(true).
	// 	EnableBroadcast(true).
	// 	MTU(2800).
	// 	MulticastLimit(32).
	// 	AddRoute("10.147.20.0/24", nil).
	// 	AddIPPool("10.147.20.1", "10.147.20.254").
	// 	V4AssignMode(true).
	// 	DNS("zt.example.com", "8.8.8.8", "8.8.4.4").
	// 	Build()
	// network, err := client.Networks().Create(config)
	// if err != nil {
	// 	log.Printf("创建网络失败: %v", err)
	// } else {
	// 	fmt.Printf("创建成功: %s\n", network.ID)
	// }

	// 更新网络
	// updateConfig := zerotier.NewCentralNetworkConfig().
	// 	Name("Updated Network Name").
	// 	Build()
	// network, err := client.Networks().Update("network_id", updateConfig)

	// 删除网络
	// err := client.Networks().Delete("network_id")

	// ========================================
	// 4. 成员管理
	// ========================================
	if len(networks) > 0 {
		networkID := networks[0].ID
		fmt.Printf("\n=== 网络 %s 的成员 ===\n", networkID)

		members, err := client.Networks().Members(networkID).List()
		if err != nil {
			log.Printf("获取成员列表失败: %v", err)
		} else {
			if len(members) == 0 {
				fmt.Println("暂无成员")
			}
			for _, m := range members {
				fmt.Printf("成员: %s\n", m.NodeID)
				fmt.Printf("  名称: %s\n", m.Name)
				fmt.Printf("  描述: %s\n", m.Description)
				fmt.Printf("  授权: %v\n", m.Config.Authorized)
				fmt.Printf("  IP: %v\n", m.Config.IPAssignments)
				fmt.Printf("  物理地址: %s\n", m.PhysicalAddress)
				fmt.Printf("  客户端版本: %s\n", m.ClientVersion)
			}
		}
	}

	// 获取成员详情
	// member, err := client.Networks().Members("network_id").Get("member_id")

	// 授权成员（快捷方式）
	// member, err := client.Networks().Members("network_id").Authorize("member_id")

	// 取消授权（快捷方式）
	// member, err := client.Networks().Members("network_id").Deauthorize("member_id")

	// 更新成员配置
	// memberConfig := zerotier.NewCentralMemberConfig().
	// 	Name("my-device").
	// 	Description("My laptop").
	// 	Authorized(true).
	// 	ActiveBridge(false).
	// 	NoAutoAssignIPs(false).
	// 	IPAssignments("10.147.20.100", "10.147.20.101").
	// 	Build()
	// member, err := client.Networks().Members("network_id").Update("member_id", memberConfig)

	// 删除成员
	// err := client.Networks().Members("network_id").Delete("member_id")

	fmt.Println("\n========== 完成 ==========")
}
