// Code generated by goctl. DO NOT EDIT.
package types

type ChatMessage struct {
	Id      int64  `json:"id"`
	Uid     uint64 `json:"uid"`
	ToUid   uint64 `json:"to_uid"`
	Message string `json:"message"`
}

type ChatMessageSaveReq struct {
	Uid     uint64 `json:"uid"`
	ToUid   uint64 `json:"to_uid"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type ChatMessageReq struct {
	Uid      uint64 `json:"uid"`
	ToUid    uint64 `json:"to_uid"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"page_size"`
}

type ChatMessageRsp struct {
	Code    int64          `json:"code"`
	Message string         `json:"message"`
	Chat    []*ChatMessage `json:"chat"`
}

type ChatHistorySaveReq struct {
	ToUid uint64 `json:"to_uid"`
}

type Rsp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type User struct {
	Id      uint64 `json:"id"`
	Account string `json:"account"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Status  bool   `json:"status"`
}

type ChatHistoryRsp struct {
	Code    int64   `json:"code"`
	Message string  `json:"message"`
	User    []*User `json:"user"`
}

type ChatNumberRsp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Number  int64  `json:"number"`
}
