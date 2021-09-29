package handler

import (
	"net/http"

	"github.com/acger/chat-api/internal/logic/chat"
	"github.com/acger/chat-api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func ChatHistoryHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewChatHistoryLogic(r.Context(), ctx)
		resp, err := l.ChatHistory()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
