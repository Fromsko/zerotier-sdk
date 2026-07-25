// 云端 Central API V2 使用示例
package main

import (
	"context"
	"fmt"
	"log"

	zerotier "github.com/fromsko/zerotier-sdk"
	centralv2 "github.com/fromsko/zerotier-sdk/central/v2"
)

func main() {
	fmt.Println("========== ZeroTier Central V2 API ==========")

	// Service Account Token 获取方式：
	// https://docs.zerotier.com/central/v2/ 创建 Service Account 后生成 API Key
	token := "your_service_account_token_here" // 替换为你的 Token

	client, err := zerotier.NewCentralV2(token)
	if err != nil {
		log.Fatalf("创建 V2 客户端失败: %v", err)
	}

	ctx := context.Background()

	// 列出组织
	fmt.Println("\n=== 组织列表 ===")
	orgs, err := client.ListOrgsWithResponse(ctx, nil)
	if err != nil {
		log.Printf("列出组织失败: %v", err)
	} else if orgs.JSON200 != nil {
		if len(orgs.JSON200.Items) == 0 {
			fmt.Println("暂无组织")
		}
		for _, o := range orgs.JSON200.Items {
			fmt.Printf("组织: %s (%s)\n", o.Name, o.Id)
		}
	}

	// 列出网络组
	fmt.Println("\n=== 网络组列表 ===")
	groups, err := client.ListNetworkGroupsWithResponse(ctx, nil)
	if err != nil {
		log.Printf("列出网络组失败: %v", err)
	} else if groups.JSON200 != nil {
		for _, g := range groups.JSON200.Items {
			fmt.Printf("网络组: %s (%s)\n", g.Name, g.Id)
		}
	}

	// 列出首个网络组下的网络
	if groups.JSON200 != nil && len(groups.JSON200.Items) > 0 {
		groupID := groups.JSON200.Items[0].Id
		fmt.Printf("\n=== 网络组 %s 的网络 ===\n", groupID)
		networks, err := client.ListNetworkGroupNetworksWithResponse(ctx, groupID, &centralv2.ListNetworkGroupNetworksParams{})
		if err != nil {
			log.Printf("列出网络失败: %v", err)
		} else if networks.JSON200 != nil {
			for _, n := range networks.JSON200.Items {
				fmt.Printf("网络: %s (%s)\n", n.Name, n.Id)
			}
		}
	}

	fmt.Println("\n========== 完成 ==========")
}
