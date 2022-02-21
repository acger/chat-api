package chatGroup

import (
	"context"
	"encoding/json"
	"github.com/acger/chat-svc/chat"

	"github.com/acger/chat-api/internal/svc"
	"github.com/acger/chat-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatHistorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatHistorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChatHistorySaveLogic {
	return ChatHistorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatHistorySaveLogic) ChatHistorySave(req types.ChatHistorySaveReq) (*types.Rsp, error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()

	l.svcCtx.ChatSvc.ChatHistorySave(l.ctx, &chat.CHSaveReq{
		Uid:   uint64(uid),
		ToUid: req.ToUid,
	})

	return &types.Rsp{}, nil
}
