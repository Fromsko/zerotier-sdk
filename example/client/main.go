// 本地 Service API 使用示例
package main

import (
	"fmt"
	"log"

	zerotier "github.com/fromsko/zerotier-sdk"
)

func main() {
	fmt.Println("========== ZeroTier Service API（本地节点管理）==========")

	// ========================================
	// 1. 创建客户端
	// ========================================

	// 方式一：默认配置（自动读取系统 token）
	client := zerotier.NewClient()

	// 方式二：自定义配置
	// client := zerotier.NewClient(
	// 	zerotier.WithClientBaseURL("http://localhost:9993"),
	// 	zerotier.WithClientToken("your-token"),
	// )

	// 方式三：从文件读取 token
	// client := zerotier.NewClient(
	// 	zerotier.WithClientTokenFile("/path/to/authtoken.secret"),
	// )

	// ========================================
	// 2. 节点状态
	// ========================================
	fmt.Println("\n=== 节点状态 ===")
	status, err := client.Status()
	if err != nil {
		log.Printf("获取状态失败: %v", err)
	} else {
		fmt.Printf("节点地址: %s\n", status.Address)
		fmt.Printf("版本: %s\n", status.Version)
		fmt.Printf("在线: %v\n", status.Online)
		fmt.Printf("公钥: %s...\n", status.PublicIdentity[:20])
	}

	// ========================================
	// 3. 网络管理
	// ========================================
	fmt.Println("\n=== 已加入的网络 ===")
	networks, err := client.Networks().List()
	if err != nil {
		log.Printf("获取网络列表失败: %v", err)
	} else {
		if len(networks) == 0 {
			fmt.Println("暂未加入任何网络")
		}
		for _, n := range networks {
			fmt.Printf("网络: %s\n", n.ID)
			fmt.Printf("  名称: %s\n", n.Name)
			fmt.Printf("  状态: %s\n", n.Status)
			fmt.Printf("  类型: %s\n", n.Type)
			fmt.Printf("  MTU: %d\n", n.MTU)
			fmt.Printf("  分配的IP: %v\n", n.AssignedAddresses)
			fmt.Printf("  允许DNS: %v\n", n.AllowDNS)
			fmt.Printf("  允许托管路由: %v\n", n.AllowManaged)
		}
	}

	// 加入网络
	// network, err := client.Networks().Join("your_network_id")
	// if err != nil {
	// 	log.Printf("加入网络失败: %v", err)
	// } else {
	// 	fmt.Printf("已加入网络: %s\n", network.ID)
	// }

	// 离开网络
	// err = client.Networks().Leave("your_network_id")

	// 获取网络详情
	// network, err := client.Networks().Get("your_network_id")

	// 更新网络设置
	// settings := zerotier.NewNetworkSettings().
	// 	AllowDNS(true).
	// 	AllowManaged(true).
	// 	AllowDefault(false).
	// 	AllowGlobal(false).
	// 	Build()
	// network, err := client.Networks().Update("your_network_id", settings)

	// ========================================
	// 4. Peer 管理
	// ========================================
	fmt.Println("\n=== Peers ===")
	peers, err := client.Peers().List()
	if err != nil {
		log.Printf("获取 Peers 失败: %v", err)
	} else {
		for _, p := range peers {
			fmt.Printf("Peer: %s\n", p.Address)
			fmt.Printf("  角色: %s\n", p.Role)
			fmt.Printf("  版本: %s\n", p.Version)
			fmt.Printf("  延迟: %dms\n", p.Latency)
			if len(p.Paths) > 0 {
				fmt.Printf("  路径数: %d\n", len(p.Paths))
			}
		}
	}

	// 获取指定 Peer
	// peer, err := client.Peers().Get("peer_address")

	// ========================================
	// 5. 控制器（自托管时可用）
	// ========================================
	fmt.Println("\n=== 控制器状态 ===")
	ctrlStatus, err := client.Controller().Status()
	if err != nil {
		log.Printf("获取控制器状态失败（可能未启用）: %v", err)
	} else {
		fmt.Printf("控制器启用: %v\n", ctrlStatus.Controller)
		fmt.Printf("API 版本: %d\n", ctrlStatus.APIVersion)

		if ctrlStatus.Controller {
			// 列出控制器网络
			ctrlNetworks, err := client.Controller().ListNetworks()
			if err != nil {
				log.Printf("获取控制器网络失败: %v", err)
			} else {
				fmt.Printf("管理的网络数: %d\n", len(ctrlNetworks))
				for _, nwid := range ctrlNetworks {
					fmt.Printf("  - %s\n", nwid)
				}
			}
		}
	}

	// 创建网络（自托管控制器）
	// config := zerotier.NewControllerNetworkConfig().
	// 	Name("my-network").
	// 	Private(true).
	// 	EnableBroadcast(true).
	// 	MulticastLimit(32).
	// 	AddRoute("10.147.20.0/24", nil).
	// 	AddIPPool("10.147.20.1", "10.147.20.254").
	// 	V4AssignMode(true).
	// 	Build()
	// network, err := client.Controller().CreateNetwork(status.Address, config)

	// 获取网络配置
	// network, err := client.Controller().GetNetwork("network_id")

	// 更新网络配置
	// client.Controller().UpdateNetwork("network_id", config)

	// 删除网络
	// client.Controller().DeleteNetwork("network_id")

	// 列出网络成员
	// members, err := client.Controller().ListMembers("network_id")

	// 授权成员
	// memberConfig := zerotier.NewControllerMemberConfig().
	// 	Authorized(true).
	// 	IPAssignments("10.147.20.100").
	// 	Build()
	// member, err := client.Controller().UpdateMember("network_id", "member_id", memberConfig)

	// 删除成员
	// client.Controller().DeleteMember("network_id", "member_id")

	fmt.Println("\n========== 完成 ==========")
}
