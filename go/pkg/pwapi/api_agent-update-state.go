package pwapi

import (
	"context"

	"pathwar.land/v2/go/pkg/errcode"
)

func (svc *service) AgentUpdateState(ctx context.Context, in *AgentUpdateState_Input) (*AgentUpdateState_Output, error) {
	if !isAgentContext(ctx) {
		return nil, errcode.ErrRestrictedArea
	}
	if in == nil {
		return nil, errcode.ErrMissingInput
	}

	return nil, errcode.ErrNotImplemented
}
