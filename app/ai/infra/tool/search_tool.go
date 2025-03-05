package tool

import (
	"context"
	"time"

	"github.com/cloudwego/eino-ext/components/tool/duckduckgo"
	"github.com/cloudwego/eino-ext/components/tool/duckduckgo/ddgsearch"
	"github.com/cloudwego/eino/components/tool"
)

func NewSearchTool(ctx context.Context) (tool.BaseTool, error) {
	// 创建 duckduckgo Search 工具
	searchTool, err := duckduckgo.NewTool(ctx, &duckduckgo.Config{
		ToolName:   "duckduckgo_search",                        // 工具名称
		ToolDesc:   "search web for information by duckduckgo", // 工具描述
		Region:     ddgsearch.RegionWT,                         // 搜索地区
		MaxResults: 10,                                         // 每页结果数量
		SafeSearch: ddgsearch.SafeSearchOff,                    // 安全搜索级别
		TimeRange:  ddgsearch.TimeRangeAll,                     // 时间范围
		DDGConfig: &ddgsearch.Config{
			Timeout:    300 * time.Second,
			Cache:      true,
			MaxRetries: 5,
		}, // DuckDuckGo 配置
	})
	return searchTool, err
}
