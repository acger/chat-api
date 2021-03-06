// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	chatGroup "github.com/acger/chat-api/internal/handler/chatGroup"
	"github.com/acger/chat-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/chat/message",
				Handler: chatGroup.ChatMessageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/chat/message/save",
				Handler: chatGroup.ChatMessageSaveHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/chat/history",
				Handler: chatGroup.ChatHistoryHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/chat/history/save",
				Handler: chatGroup.ChatHistorySaveHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/chat/history/number",
				Handler: chatGroup.ChatHistoryNumberHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
