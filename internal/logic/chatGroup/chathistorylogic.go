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

type ChatHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChatHistoryLogic {
	return ChatHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatHistoryLogic) ChatHistory() (*types.ChatHistoryRsp, error) {

	uid, _ := l.ctx.Value("userId").(json.Number).Int64()

	r, _ := l.svcCtx.ChatSvc.ChatHistoryList(l.ctx, &chat.ChatHistoryReq{
		Id: uint64(uid),
	})

	user := make([]*types.User, len(r.User))

	for i, u := range r.User {
		item := types.User{}
		copier.Copy(&item, &u)
		user[i] = &item
	}

	return &types.ChatHistoryRsp{User: user}, nil
}
