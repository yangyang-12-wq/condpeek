package models

import (
	"time"

	"gorm.io/gorm"
)

// Follow 关注模型 - 用户关注事件
type Follow struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null;index"` // 关注者ID
	EventID   uint           `json:"event_id" gorm:"not null;index"` // 关注的事件ID
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	User  User  `json:"user" gorm:"foreignKey:UserID"`
	Event Event `json:"event" gorm:"foreignKey:EventID"`
}

// FollowResponse 关注响应结构
type FollowResponse struct {
	ID        uint      `json:"id"`
	EventID   uint      `json:"event_id"`
	EventTitle string   `json:"event_title"` // 事件标题
	CreatedAt time.Time `json:"created_at"`
}

// FollowStats 关注统计
type FollowStats struct {
	FollowedEvents []EventResponse `json:"followed_events"`
	TotalCount     int             `json:"total_count"`
}

// AddFollowRequest 添加关注请求
type AddFollowRequest struct {
	EventID uint `json:"event_id" binding:"required"` // 要关注的事件ID
}

// RemoveFollowRequest 取消关注请求
type RemoveFollowRequest struct {
	EventID uint `json:"event_id" binding:"required"` // 要取消关注的事件ID
}

// CheckFollowRequest 检查关注请求
type CheckFollowRequest struct {
	EventID uint `json:"event_id" binding:"required"` // 要检查的事件ID
}

// ToResponse 转换为响应格式
func (f *Follow) ToResponse() FollowResponse {
	return FollowResponse{
		ID:         f.ID,
		EventID:    f.EventID,
		EventTitle: f.Event.Title,
		CreatedAt:  f.CreatedAt,
	}
}



// TableName 设置表名
func (Follow) TableName() string {
	return "follows"
}
