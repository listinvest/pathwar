package pwapi

import (
	"context"

	"pathwar.land/pathwar/v2/go/pkg/errcode"
	"pathwar.land/pathwar/v2/go/pkg/pwdb"
)

func (svc *service) AdminListChallengeSubscriptions(ctx context.Context, in *AdminListChallengeSubscriptions_Input) (*AdminListChallengeSubscriptions_Output, error) {
	if !isAdminContext(ctx) {
		return nil, errcode.ErrRestrictedArea
	}
	if in == nil {
		return nil, errcode.ErrMissingInput
	}

	var challengeSubscriptions []*pwdb.ChallengeSubscription
	err := svc.db.
		Preload("Team").
		Preload("Team.Organization").
		Preload("Closer").
		Preload("Validations").
		Preload("SeasonChallenge").
		Preload("SeasonChallenge.Season").
		Preload("SeasonChallenge.Flavor").
		Preload("SeasonChallenge.Flavor.Challenge").
		Preload("Buyer").
		Find(&challengeSubscriptions).Error
	if err != nil {
		return nil, errcode.ErrListChallengeSubscriptions.Wrap(err)
	}

	out := AdminListChallengeSubscriptions_Output{Subscriptions: challengeSubscriptions}
	return &out, nil
}
