package agent

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	mytool "github.com/SGNYYYY/gomall/app/ai/infra/tool"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

// ptrOf returns a pointer to the given value.
func ptrOf[T any](v T) *T {
	return &v
}

func Run(ctx context.Context, content string, userId uint32) {
	// 初始化 tools
	listTool, _ := mytool.NewListOrderTool(ctx)
	placeOrderTool, _ := mytool.NewPlaceOrderTool(ctx)
	// searchTool, _ := mytool.NewSearchTool(ctx)
	myTools := []tool.BaseTool{
		listTool,
		placeOrderTool,
		// searchTool,
	}

	// 创建并配置 ChatModel
	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Region:  "cn-beijing",
		Model:   "doubao-1-5-pro-32k-250115",
		Timeout: ptrOf(300 * time.Second),
	})
	if err != nil {
		log.Fatal(err)
	}
	// 获取工具信息并绑定到 ChatModel
	toolInfos := make([]*schema.ToolInfo, 0, len(myTools))
	for _, tool := range myTools {
		info, err := tool.Info(ctx)
		if err != nil {
			log.Fatal(err)
		}
		toolInfos = append(toolInfos, info)
	}
	err = chatModel.BindTools(toolInfos)
	if err != nil {
		log.Fatal(err)
	}

	// 创建 tools 节点
	myToolsNode, err := compose.NewToolNode(ctx, &compose.ToolsNodeConfig{
		Tools: myTools,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 构建完整的处理链
	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()
	chain.
		AppendChatModel(chatModel, compose.WithNodeName("chat_model")).
		AppendToolsNode(myToolsNode, compose.WithNodeName("tools"))

	// 编译并运行 chain
	agent, err := chain.Compile(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 运行示例
	resp, err := agent.Invoke(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: fmt.Sprintf("%s,ID为%d", content, userId),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	// 输出结果
	for _, msg := range resp {
		fmt.Println(msg.Content)
	}
}
