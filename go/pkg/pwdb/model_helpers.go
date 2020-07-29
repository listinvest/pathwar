package pwdb

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/martinlindhe/base36"
	"golang.org/x/crypto/sha3"
	"pathwar.land/pathwar/v2/go/pkg/errcode"
	"pathwar.land/pathwar/v2/go/pkg/pwinit"
)

func (cf *ChallengeFlavor) addSeasonChallengeByID(seasonID int64) {
	if cf.SeasonChallenges == nil {
		cf.SeasonChallenges = []*SeasonChallenge{}
	}
	cf.SeasonChallenges = append(cf.SeasonChallenges, &SeasonChallenge{
		SeasonID: seasonID,
	})
}

func (a *Agent) TagSlice() []string {
	if a.Tags == "" {
		return nil
	}
	return strings.Split(a.Tags, ", ")
}

func (cf ChallengeFlavor) NameAndVersion() string {
	return fmt.Sprintf("%s@%s", cf.Challenge.Name, cf.Version)
}

func (instance *ChallengeInstance) ParseInstanceConfig() (*pwinit.InitConfig, error) {
	var configData pwinit.InitConfig
	err := json.Unmarshal(instance.GetInstanceConfig(), &configData)
	if err != nil {
		return nil, errcode.ErrParseInitConfig.Wrap(err)
	}
	return &configData, nil
}

func ChallengeInstancePrefixHash(instanceID string, userID int64, salt string) (string, error) {
	stringToHash := fmt.Sprintf("%s%d%s", instanceID, userID, salt)
	hashBytes := make([]byte, 8)
	hasher := sha3.NewShake256()
	_, err := hasher.Write([]byte(stringToHash))
	if err != nil {
		return "", errcode.ErrWriteBytesToHashBuilder.Wrap(err)
	}
	_, err = hasher.Read(hashBytes)
	if err != nil {
		return "", errcode.ErrReadBytesFromHashBuilder.Wrap(err)
	}
	userHash := strings.ToLower(base36.EncodeBytes(hashBytes))[:8] // we voluntarily expect short hashes here
	return userHash, nil
}

func (m *SeasonChallenge) GetActiveSubscriptions() []*ChallengeSubscription {
	cs := make([]*ChallengeSubscription, 0)

	for _, subscription := range m.GetSubscriptions() {
		if subscription.GetStatus() == ChallengeSubscription_Active {
			cs = append(cs, subscription)
		}
	}

	return cs
}
