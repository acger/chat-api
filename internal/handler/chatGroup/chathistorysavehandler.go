package chatGroup

import (
	"net/http"

	"github.com/acger/chat-api/internal/logic/chatGroup"
	"github.com/acger/chat-api/internal/svc"
	"github.com/acger/chat-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatHistorySaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatHistorySaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := chatGroup.NewChatHistorySaveLogic(r.Context(), svcCtx)
		resp, err := l.ChatHistorySave(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
