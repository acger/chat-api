package chatGroup

import (
	"context"
	"encoding/json"
	"github.com/acger/chat-svc/chat"

	"github.com/acger/chat-api/internal/svc"
	"github.com/acger/chat-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatMessageSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatMessageSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChatMessageSaveLogic {
	return ChatMessageSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatMessageSaveLogic) ChatMessageSave(req types.ChatMessageSaveReq) (*types.Rsp, error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()

	l.svcCtx.ChatSvc.MessageSave(l.ctx, &chat.MsgSaveReq{
		Uid:     uint64(uid),
		ToUid:   req.ToUid,
		Message: req.Message,
		Status:  req.Status,
	})

	return &types.Rsp{}, nil
}
