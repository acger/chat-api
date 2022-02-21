package chatGroup

import (
	"context"
	"encoding/json"
	"github.com/acger/chat-svc/chat"
	"github.com/jinzhu/copier"

	"github.com/acger/chat-api/internal/svc"
	"github.com/acger/chat-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChatMessageLogic {
	return ChatMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatMessageLogic) ChatMessage(req types.ChatMessageReq) (*types.ChatMessageRsp, error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()

	r, _ := l.svcCtx.ChatSvc.MessageList(l.ctx, &chat.MsgListReq{
		Uid:      uint64(uid),
		ToUid:    req.ToUid,
		Page:     req.Page,
		PageSize: req.PageSize,
	})

	msg := make([]*types.ChatMessage, len(r.Msg))

	for i, m := range r.Msg {
		item := types.ChatMessage{}
		copier.Copy(&item, &m)
		msg[i] = &item
	}

	return &types.ChatMessageRsp{Chat: msg}, nil
}
