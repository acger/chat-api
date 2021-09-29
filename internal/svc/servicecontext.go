package svc

import (
	"github.com/acger/chat-api/internal/config"
	"github.com/acger/chat-svc/chat"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	ChatSvc chat.Chat
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		ChatSvc: chat.NewChat(zrpc.MustNewClient(c.ChatSvc)),
	}
}
