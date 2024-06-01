// 模型基类
package models

import "goblog/pkg/types"

// 模型基类
type BaseModel struct {
	ID uint64
}

// 获取 ID 的字符串格式
func (bm BaseModel) GetStringID() string {
	return types.Uint64ToString(bm.ID)
}
