package v1

// ResponseWithMessage 含有message消息的响应
// swagger:response ResponseWithMessage
type ResponseWithMessage struct {
	Message string `json:"message"`
}
