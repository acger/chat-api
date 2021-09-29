package handler

import (
	"net/http"

	"github.com/acger/chat-api/internal/logic/chat"
	"github.com/acger/chat-api/internal/svc"
	"github.com/acger/chat-api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ChatHistorySaveHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatHistorySaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewChatHistorySaveLogic(r.Context(), ctx)
		resp, err := l.ChatHistorySave(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
