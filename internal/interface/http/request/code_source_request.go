package request

// @Author CY Yan
// @Date 2025/5/28 21:24

type CodeSourceCreateRequest struct {
	Name        string `json:"name" `
	Desc        string `json:"desc"`
	RemoteURL   string `json:"remote_url" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	RemoteToken string `json:"remote_token" binding:"required"`
}
