package handler

import (
	"net/http"

	"github.com/acger/chat-api/internal/logic/chat"
	"github.com/acger/chat-api/internal/svc"
	"github.com/acger/chat-api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ChatMessageSaveHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatMessageSaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewChatMessageSaveLogic(r.Context(), ctx)
		resp, err := l.ChatMessageSave(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
