package logic

import (
	"context"
	"encoding/json"
	"github.com/acger/chat-svc/chat"

	"github.com/acger/chat-api/internal/svc"
	"github.com/acger/chat-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChatHistoryNumberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatHistoryNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChatHistoryNumberLogic {
	return ChatHistoryNumberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatHistoryNumberLogic) ChatHistoryNumber() (*types.ChatNumberRsp, error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()

	r, _ := l.svcCtx.ChatSvc.ChatNumber(l.ctx, &chat.ChatNumberReq{
		Id: uint64(uid),
	})

	return &types.ChatNumberRsp{Number: r.Number}, nil
}
