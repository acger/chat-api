package handler

import (
	"net/http"

	"github.com/acger/chat-api/internal/logic/chat"
	"github.com/acger/chat-api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func ChatHistoryNumberHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewChatHistoryNumberLogic(r.Context(), ctx)
		resp, err := l.ChatHistoryNumber()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
