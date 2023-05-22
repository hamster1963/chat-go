package m_chat_content

type AddChatContent struct {
	Content   *string `json:"content"     ` // 聊天内容
	SendId    *int    `json:"send_id"     ` // 发送者ID
	ReceiveId *int    `json:"receive_id"  ` // 接收者ID
}
