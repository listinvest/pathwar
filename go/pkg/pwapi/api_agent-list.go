package pwapi

import (
	"context"

	"pathwar.land/v2/go/pkg/errcode"
)

func (svc *service) AgentList(context.Context, *AgentList_Input) (*AgentList_Output, error) {
	// FIXME: check if client is admin
	return nil, errcode.ErrNotImplemented
}
