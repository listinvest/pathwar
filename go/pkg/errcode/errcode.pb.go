// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errcode.proto

package errcode

import (
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ErrCode int32

const (
	Undefined                               ErrCode = 0
	TODO                                    ErrCode = 666
	ErrNotImplemented                       ErrCode = 777
	ErrDeprecated                           ErrCode = 888
	ErrInternal                             ErrCode = 999
	ErrInvalidInput                         ErrCode = 101
	ErrMissingInput                         ErrCode = 102
	ErrUnauthenticated                      ErrCode = 103
	ErrRestrictedArea                       ErrCode = 104
	ErrSSOGetOIDC                           ErrCode = 1001
	ErrSSOInvalidPublicKey                  ErrCode = 1002
	ErrSSOFailedKeycloakRequest             ErrCode = 1003
	ErrSSOInvalidKeycloakResponse           ErrCode = 1004
	ErrSSOLogout                            ErrCode = 1005
	ErrSSOInitKeycloakClient                ErrCode = 1006
	ErrSSOInvalidBearer                     ErrCode = 1007
	ErrSSOKeycloakError                     ErrCode = 1008
	ErrDBNotFound                           ErrCode = 2001
	ErrDBInternal                           ErrCode = 2002
	ErrDBRunMigrations                      ErrCode = 2003
	ErrDBInit                               ErrCode = 2004
	ErrDBConnect                            ErrCode = 2005
	ErrDBAutoMigrate                        ErrCode = 2006
	ErrDBAddForeignKey                      ErrCode = 2007
	ErrComposeInvalidPath                   ErrCode = 3001
	ErrComposeDirectoryNotFound             ErrCode = 3002
	ErrComposeReadConfig                    ErrCode = 3003
	ErrComposeInvalidConfig                 ErrCode = 3004
	ErrComposeMarshalConfig                 ErrCode = 3005
	ErrComposeCreateTempFile                ErrCode = 3006
	ErrComposeWriteTempFile                 ErrCode = 3007
	ErrComposeCloseTempFile                 ErrCode = 3008
	ErrComposeBuild                         ErrCode = 3009
	ErrComposeBundle                        ErrCode = 30010
	ErrComposeReadDab                       ErrCode = 3011
	ErrComposeParseDab                      ErrCode = 3012
	ErrComposeParseConfig                   ErrCode = 3013
	ErrComposeCreateTempDir                 ErrCode = 3014
	ErrComposeUpdateTempFile                ErrCode = 3015
	ErrComposeForceRecreateDown             ErrCode = 3016
	ErrComposeRunCreate                     ErrCode = 3017
	ErrComposeRunUp                         ErrCode = 3018
	ErrGetPWInitBinary                      ErrCode = 3019
	ErrWritePWInitFileHeader                ErrCode = 3020
	ErrWritePWInitFile                      ErrCode = 3021
	ErrMarshalPWInitConfigFile              ErrCode = 3022
	ErrWritePWInitConfigFileHeader          ErrCode = 3023
	ErrWritePWInitConfigFile                ErrCode = 3024
	ErrWritePWInitCloseTarWriter            ErrCode = 3025
	ErrCopyPWInitToContainer                ErrCode = 3026
	ErrComposeGetContainersInfo             ErrCode = 3027
	ErrMissingPwinitConfig                  ErrCode = 3028
	ErrGetUserIDFromContext                 ErrCode = 4001
	ErrMissingChallengeValidation           ErrCode = 4002
	ErrInvalidSeason                        ErrCode = 4003
	ErrTeamNotInSeason                      ErrCode = 4004
	ErrGetChallengeSubscription             ErrCode = 4005
	ErrInitSnowflake                        ErrCode = 4006
	ErrUpdateChallengeSubscription          ErrCode = 4007
	ErrCreateChallengeValidation            ErrCode = 4008
	ErrGetChallengeValidation               ErrCode = 4009
	ErrInvalidTeam                          ErrCode = 4010
	ErrChallengeAlreadySubscribed           ErrCode = 4011
	ErrCreateChallengeSubscription          ErrCode = 4012
	ErrFindOrganizations                    ErrCode = 4013
	ErrGetSeasonFromSeasonChallenge         ErrCode = 4014
	ErrGetUserTeamFromSeason                ErrCode = 4015
	ErrInvalidSeasonID                      ErrCode = 4016
	ErrUserHasNoTeamForSeason               ErrCode = 4017
	ErrGetSeasonChallenges                  ErrCode = 4018
	ErrGetSeason                            ErrCode = 4019
	ErrSeasonDenied                         ErrCode = 4020
	ErrAlreadyHasTeamForSeason              ErrCode = 4021
	ErrReservedName                         ErrCode = 4022
	ErrCheckOrganizationUniqueName          ErrCode = 4023
	ErrCreateOrganization                   ErrCode = 4024
	ErrOrganizationAlreadyHasTeamForSeason  ErrCode = 4025
	ErrGetOrganization                      ErrCode = 4026
	ErrGetSeasonChallenge                   ErrCode = 4027
	ErrCannotCreateTeamForSoloOrganization  ErrCode = 4028
	ErrUserNotInOrganization                ErrCode = 4029
	ErrCreateTeam                           ErrCode = 4030
	ErrGetTeam                              ErrCode = 4031
	ErrGetTeams                             ErrCode = 4032
	ErrGetUser                              ErrCode = 4033
	ErrUpdateUser                           ErrCode = 4034
	ErrUpdateTeam                           ErrCode = 4035
	ErrUpdateOrganization                   ErrCode = 4036
	ErrNewUserFromClaims                    ErrCode = 4037
	ErrGetOAuthUser                         ErrCode = 4038
	ErrDifferentUserBetweenTokenAndDatabase ErrCode = 4039
	ErrLoadUserSeasons                      ErrCode = 4040
	ErrGetUserOrganizations                 ErrCode = 4041
	ErrGetSeasons                           ErrCode = 4042
	ErrGetUserBySubject                     ErrCode = 4043
	ErrEmailAddressNotVerified              ErrCode = 4044
	ErrGetDefaultSeason                     ErrCode = 4045
	ErrCommitUserTransaction                ErrCode = 4046
	ErrUpdateActiveSeason                   ErrCode = 4047
	ErrMissingContextMetadata               ErrCode = 4048
	ErrNoTokenProvided                      ErrCode = 4049
	ErrGetTokenWithClaims                   ErrCode = 4050
	ErrNoTokenInContext                     ErrCode = 4051
	ErrGetSubjectFromToken                  ErrCode = 4052
	ErrGetSubjectFromContext                ErrCode = 4053
	ErrGetActiveSeasonMembership            ErrCode = 4054
	ErrGetTokenFromContext                  ErrCode = 4055
	ErrChallengeAlreadyClosed               ErrCode = 4056
	ErrGetAgent                             ErrCode = 4057
	ErrSaveAgent                            ErrCode = 4058
	ErrListChallengeInstances               ErrCode = 4059
	ErrChallengeAdd                         ErrCode = 4060
	ErrCouponAlreadyValidatedBySameTeam     ErrCode = 4061
	ErrCouponExpired                        ErrCode = 4062
	ErrCouponNotFound                       ErrCode = 4063
	ErrUserDoesNotBelongToTeam              ErrCode = 4064
	ErrServerListen                         ErrCode = 5001
	ErrServerRegisterGateway                ErrCode = 5002
	ErrInitLogger                           ErrCode = 6001
	ErrStartService                         ErrCode = 6002
	ErrInitServer                           ErrCode = 6003
	ErrGroupTerminated                      ErrCode = 6004
	ErrGetSSOClientFromFlags                ErrCode = 6005
	ErrDumpDatabase                         ErrCode = 6006
	ErrGetDBInfo                            ErrCode = 6007
	ErrGetSSOWhoami                         ErrCode = 6008
	ErrGetSSOLogout                         ErrCode = 6009
	ErrGetSSOClaims                         ErrCode = 6010
	ErrInitDockerClient                     ErrCode = 6011
	ErrInitDB                               ErrCode = 6012
	ErrConfigureDB                          ErrCode = 6013
	ErrInitSSOClient                        ErrCode = 6014
	ErrInitService                          ErrCode = 6015
	ErrAgentGetContainersInfo               ErrCode = 7001
	ErrCheckNginxContainer                  ErrCode = 7002
	ErrRemoveNginxContainer                 ErrCode = 7003
	ErrBuildNginxContainer                  ErrCode = 7004
	ErrStartNginxContainer                  ErrCode = 7005
	ErrParsingTemplate                      ErrCode = 7006
	ErrWriteConfigFileHeader                ErrCode = 7007
	ErrWriteConfigFile                      ErrCode = 7008
	ErrCloseTarWriter                       ErrCode = 7009
	ErrCopyNginxConfigToContainer           ErrCode = 7010
	ErrNginxNewConfigCheckFailed            ErrCode = 7011
	ErrNginxSendCommandNewConfigCheck       ErrCode = 7012
	ErrNginxSendCommandNewConfigRemove      ErrCode = 7013
	ErrNginxSendCommandConfigReplace        ErrCode = 7014
	ErrNginxSendCommandReloadConfig         ErrCode = 7015
	ErrNginxConnectNetwork                  ErrCode = 7016
	ErrContainerConnectNetwork              ErrCode = 7017
	ErrNatPortOpening                       ErrCode = 7018
	ErrBuildNginxConfig                     ErrCode = 7019
	ErrExecuteTemplate                      ErrCode = 7020
	ErrWriteBytesToHashBuilder              ErrCode = 7021
	ErrReadBytesFromHashBuilder             ErrCode = 7022
	ErrGeneratePrefixHash                   ErrCode = 7023
	ErrCleanPathwarInstances                ErrCode = 7024
	ErrParseInitConfig                      ErrCode = 7025
	ErrUpPathwarInstance                    ErrCode = 7026
	ErrUpdateNginx                          ErrCode = 7027
	ErrAgentUpdateState                     ErrCode = 7028
	ErrDockerAPIContainerList               ErrCode = 8001
	ErrDockerAPIContainerRemove             ErrCode = 8002
	ErrDockerAPIImageRemove                 ErrCode = 8003
	ErrDockerAPIContainerCreate             ErrCode = 8004
	ErrDockerAPIContainerExecCreate         ErrCode = 8005
	ErrDockerAPIContainerExecAttach         ErrCode = 8006
	ErrDockerAPIContainerExecStart          ErrCode = 8007
	ErrDockerAPIContainerExecInspect        ErrCode = 8008
	ErrDockerAPIImagePull                   ErrCode = 8009
	ErrDockerAPIImageInspect                ErrCode = 8010
	ErrDockerAPINetworkList                 ErrCode = 8011
	ErrDockerAPINetworkCreate               ErrCode = 8012
	ErrDockerAPINetworkRemove               ErrCode = 8013
	ErrDockerAPIExitCode                    ErrCode = 8014
	ErrExecuteOnInitHook                    ErrCode = 9001
	ErrRemoveInitConfig                     ErrCode = 9002
)

var ErrCode_name = map[int32]string{
	0:     "Undefined",
	666:   "TODO",
	777:   "ErrNotImplemented",
	888:   "ErrDeprecated",
	999:   "ErrInternal",
	101:   "ErrInvalidInput",
	102:   "ErrMissingInput",
	103:   "ErrUnauthenticated",
	104:   "ErrRestrictedArea",
	1001:  "ErrSSOGetOIDC",
	1002:  "ErrSSOInvalidPublicKey",
	1003:  "ErrSSOFailedKeycloakRequest",
	1004:  "ErrSSOInvalidKeycloakResponse",
	1005:  "ErrSSOLogout",
	1006:  "ErrSSOInitKeycloakClient",
	1007:  "ErrSSOInvalidBearer",
	1008:  "ErrSSOKeycloakError",
	2001:  "ErrDBNotFound",
	2002:  "ErrDBInternal",
	2003:  "ErrDBRunMigrations",
	2004:  "ErrDBInit",
	2005:  "ErrDBConnect",
	2006:  "ErrDBAutoMigrate",
	2007:  "ErrDBAddForeignKey",
	3001:  "ErrComposeInvalidPath",
	3002:  "ErrComposeDirectoryNotFound",
	3003:  "ErrComposeReadConfig",
	3004:  "ErrComposeInvalidConfig",
	3005:  "ErrComposeMarshalConfig",
	3006:  "ErrComposeCreateTempFile",
	3007:  "ErrComposeWriteTempFile",
	3008:  "ErrComposeCloseTempFile",
	3009:  "ErrComposeBuild",
	30010: "ErrComposeBundle",
	3011:  "ErrComposeReadDab",
	3012:  "ErrComposeParseDab",
	3013:  "ErrComposeParseConfig",
	3014:  "ErrComposeCreateTempDir",
	3015:  "ErrComposeUpdateTempFile",
	3016:  "ErrComposeForceRecreateDown",
	3017:  "ErrComposeRunCreate",
	3018:  "ErrComposeRunUp",
	3019:  "ErrGetPWInitBinary",
	3020:  "ErrWritePWInitFileHeader",
	3021:  "ErrWritePWInitFile",
	3022:  "ErrMarshalPWInitConfigFile",
	3023:  "ErrWritePWInitConfigFileHeader",
	3024:  "ErrWritePWInitConfigFile",
	3025:  "ErrWritePWInitCloseTarWriter",
	3026:  "ErrCopyPWInitToContainer",
	3027:  "ErrComposeGetContainersInfo",
	3028:  "ErrMissingPwinitConfig",
	4001:  "ErrGetUserIDFromContext",
	4002:  "ErrMissingChallengeValidation",
	4003:  "ErrInvalidSeason",
	4004:  "ErrTeamNotInSeason",
	4005:  "ErrGetChallengeSubscription",
	4006:  "ErrInitSnowflake",
	4007:  "ErrUpdateChallengeSubscription",
	4008:  "ErrCreateChallengeValidation",
	4009:  "ErrGetChallengeValidation",
	4010:  "ErrInvalidTeam",
	4011:  "ErrChallengeAlreadySubscribed",
	4012:  "ErrCreateChallengeSubscription",
	4013:  "ErrFindOrganizations",
	4014:  "ErrGetSeasonFromSeasonChallenge",
	4015:  "ErrGetUserTeamFromSeason",
	4016:  "ErrInvalidSeasonID",
	4017:  "ErrUserHasNoTeamForSeason",
	4018:  "ErrGetSeasonChallenges",
	4019:  "ErrGetSeason",
	4020:  "ErrSeasonDenied",
	4021:  "ErrAlreadyHasTeamForSeason",
	4022:  "ErrReservedName",
	4023:  "ErrCheckOrganizationUniqueName",
	4024:  "ErrCreateOrganization",
	4025:  "ErrOrganizationAlreadyHasTeamForSeason",
	4026:  "ErrGetOrganization",
	4027:  "ErrGetSeasonChallenge",
	4028:  "ErrCannotCreateTeamForSoloOrganization",
	4029:  "ErrUserNotInOrganization",
	4030:  "ErrCreateTeam",
	4031:  "ErrGetTeam",
	4032:  "ErrGetTeams",
	4033:  "ErrGetUser",
	4034:  "ErrUpdateUser",
	4035:  "ErrUpdateTeam",
	4036:  "ErrUpdateOrganization",
	4037:  "ErrNewUserFromClaims",
	4038:  "ErrGetOAuthUser",
	4039:  "ErrDifferentUserBetweenTokenAndDatabase",
	4040:  "ErrLoadUserSeasons",
	4041:  "ErrGetUserOrganizations",
	4042:  "ErrGetSeasons",
	4043:  "ErrGetUserBySubject",
	4044:  "ErrEmailAddressNotVerified",
	4045:  "ErrGetDefaultSeason",
	4046:  "ErrCommitUserTransaction",
	4047:  "ErrUpdateActiveSeason",
	4048:  "ErrMissingContextMetadata",
	4049:  "ErrNoTokenProvided",
	4050:  "ErrGetTokenWithClaims",
	4051:  "ErrNoTokenInContext",
	4052:  "ErrGetSubjectFromToken",
	4053:  "ErrGetSubjectFromContext",
	4054:  "ErrGetActiveSeasonMembership",
	4055:  "ErrGetTokenFromContext",
	4056:  "ErrChallengeAlreadyClosed",
	4057:  "ErrGetAgent",
	4058:  "ErrSaveAgent",
	4059:  "ErrListChallengeInstances",
	4060:  "ErrChallengeAdd",
	4061:  "ErrCouponAlreadyValidatedBySameTeam",
	4062:  "ErrCouponExpired",
	4063:  "ErrCouponNotFound",
	4064:  "ErrUserDoesNotBelongToTeam",
	5001:  "ErrServerListen",
	5002:  "ErrServerRegisterGateway",
	6001:  "ErrInitLogger",
	6002:  "ErrStartService",
	6003:  "ErrInitServer",
	6004:  "ErrGroupTerminated",
	6005:  "ErrGetSSOClientFromFlags",
	6006:  "ErrDumpDatabase",
	6007:  "ErrGetDBInfo",
	6008:  "ErrGetSSOWhoami",
	6009:  "ErrGetSSOLogout",
	6010:  "ErrGetSSOClaims",
	6011:  "ErrInitDockerClient",
	6012:  "ErrInitDB",
	6013:  "ErrConfigureDB",
	6014:  "ErrInitSSOClient",
	6015:  "ErrInitService",
	7001:  "ErrAgentGetContainersInfo",
	7002:  "ErrCheckNginxContainer",
	7003:  "ErrRemoveNginxContainer",
	7004:  "ErrBuildNginxContainer",
	7005:  "ErrStartNginxContainer",
	7006:  "ErrParsingTemplate",
	7007:  "ErrWriteConfigFileHeader",
	7008:  "ErrWriteConfigFile",
	7009:  "ErrCloseTarWriter",
	7010:  "ErrCopyNginxConfigToContainer",
	7011:  "ErrNginxNewConfigCheckFailed",
	7012:  "ErrNginxSendCommandNewConfigCheck",
	7013:  "ErrNginxSendCommandNewConfigRemove",
	7014:  "ErrNginxSendCommandConfigReplace",
	7015:  "ErrNginxSendCommandReloadConfig",
	7016:  "ErrNginxConnectNetwork",
	7017:  "ErrContainerConnectNetwork",
	7018:  "ErrNatPortOpening",
	7019:  "ErrBuildNginxConfig",
	7020:  "ErrExecuteTemplate",
	7021:  "ErrWriteBytesToHashBuilder",
	7022:  "ErrReadBytesFromHashBuilder",
	7023:  "ErrGeneratePrefixHash",
	7024:  "ErrCleanPathwarInstances",
	7025:  "ErrParseInitConfig",
	7026:  "ErrUpPathwarInstance",
	7027:  "ErrUpdateNginx",
	7028:  "ErrAgentUpdateState",
	8001:  "ErrDockerAPIContainerList",
	8002:  "ErrDockerAPIContainerRemove",
	8003:  "ErrDockerAPIImageRemove",
	8004:  "ErrDockerAPIContainerCreate",
	8005:  "ErrDockerAPIContainerExecCreate",
	8006:  "ErrDockerAPIContainerExecAttach",
	8007:  "ErrDockerAPIContainerExecStart",
	8008:  "ErrDockerAPIContainerExecInspect",
	8009:  "ErrDockerAPIImagePull",
	8010:  "ErrDockerAPIImageInspect",
	8011:  "ErrDockerAPINetworkList",
	8012:  "ErrDockerAPINetworkCreate",
	8013:  "ErrDockerAPINetworkRemove",
	8014:  "ErrDockerAPIExitCode",
	9001:  "ErrExecuteOnInitHook",
	9002:  "ErrRemoveInitConfig",
}

var ErrCode_value = map[string]int32{
	"Undefined":                               0,
	"TODO":                                    666,
	"ErrNotImplemented":                       777,
	"ErrDeprecated":                           888,
	"ErrInternal":                             999,
	"ErrInvalidInput":                         101,
	"ErrMissingInput":                         102,
	"ErrUnauthenticated":                      103,
	"ErrRestrictedArea":                       104,
	"ErrSSOGetOIDC":                           1001,
	"ErrSSOInvalidPublicKey":                  1002,
	"ErrSSOFailedKeycloakRequest":             1003,
	"ErrSSOInvalidKeycloakResponse":           1004,
	"ErrSSOLogout":                            1005,
	"ErrSSOInitKeycloakClient":                1006,
	"ErrSSOInvalidBearer":                     1007,
	"ErrSSOKeycloakError":                     1008,
	"ErrDBNotFound":                           2001,
	"ErrDBInternal":                           2002,
	"ErrDBRunMigrations":                      2003,
	"ErrDBInit":                               2004,
	"ErrDBConnect":                            2005,
	"ErrDBAutoMigrate":                        2006,
	"ErrDBAddForeignKey":                      2007,
	"ErrComposeInvalidPath":                   3001,
	"ErrComposeDirectoryNotFound":             3002,
	"ErrComposeReadConfig":                    3003,
	"ErrComposeInvalidConfig":                 3004,
	"ErrComposeMarshalConfig":                 3005,
	"ErrComposeCreateTempFile":                3006,
	"ErrComposeWriteTempFile":                 3007,
	"ErrComposeCloseTempFile":                 3008,
	"ErrComposeBuild":                         3009,
	"ErrComposeBundle":                        30010,
	"ErrComposeReadDab":                       3011,
	"ErrComposeParseDab":                      3012,
	"ErrComposeParseConfig":                   3013,
	"ErrComposeCreateTempDir":                 3014,
	"ErrComposeUpdateTempFile":                3015,
	"ErrComposeForceRecreateDown":             3016,
	"ErrComposeRunCreate":                     3017,
	"ErrComposeRunUp":                         3018,
	"ErrGetPWInitBinary":                      3019,
	"ErrWritePWInitFileHeader":                3020,
	"ErrWritePWInitFile":                      3021,
	"ErrMarshalPWInitConfigFile":              3022,
	"ErrWritePWInitConfigFileHeader":          3023,
	"ErrWritePWInitConfigFile":                3024,
	"ErrWritePWInitCloseTarWriter":            3025,
	"ErrCopyPWInitToContainer":                3026,
	"ErrComposeGetContainersInfo":             3027,
	"ErrMissingPwinitConfig":                  3028,
	"ErrGetUserIDFromContext":                 4001,
	"ErrMissingChallengeValidation":           4002,
	"ErrInvalidSeason":                        4003,
	"ErrTeamNotInSeason":                      4004,
	"ErrGetChallengeSubscription":             4005,
	"ErrInitSnowflake":                        4006,
	"ErrUpdateChallengeSubscription":          4007,
	"ErrCreateChallengeValidation":            4008,
	"ErrGetChallengeValidation":               4009,
	"ErrInvalidTeam":                          4010,
	"ErrChallengeAlreadySubscribed":           4011,
	"ErrCreateChallengeSubscription":          4012,
	"ErrFindOrganizations":                    4013,
	"ErrGetSeasonFromSeasonChallenge":         4014,
	"ErrGetUserTeamFromSeason":                4015,
	"ErrInvalidSeasonID":                      4016,
	"ErrUserHasNoTeamForSeason":               4017,
	"ErrGetSeasonChallenges":                  4018,
	"ErrGetSeason":                            4019,
	"ErrSeasonDenied":                         4020,
	"ErrAlreadyHasTeamForSeason":              4021,
	"ErrReservedName":                         4022,
	"ErrCheckOrganizationUniqueName":          4023,
	"ErrCreateOrganization":                   4024,
	"ErrOrganizationAlreadyHasTeamForSeason":  4025,
	"ErrGetOrganization":                      4026,
	"ErrGetSeasonChallenge":                   4027,
	"ErrCannotCreateTeamForSoloOrganization":  4028,
	"ErrUserNotInOrganization":                4029,
	"ErrCreateTeam":                           4030,
	"ErrGetTeam":                              4031,
	"ErrGetTeams":                             4032,
	"ErrGetUser":                              4033,
	"ErrUpdateUser":                           4034,
	"ErrUpdateTeam":                           4035,
	"ErrUpdateOrganization":                   4036,
	"ErrNewUserFromClaims":                    4037,
	"ErrGetOAuthUser":                         4038,
	"ErrDifferentUserBetweenTokenAndDatabase": 4039,
	"ErrLoadUserSeasons":                      4040,
	"ErrGetUserOrganizations":                 4041,
	"ErrGetSeasons":                           4042,
	"ErrGetUserBySubject":                     4043,
	"ErrEmailAddressNotVerified":              4044,
	"ErrGetDefaultSeason":                     4045,
	"ErrCommitUserTransaction":                4046,
	"ErrUpdateActiveSeason":                   4047,
	"ErrMissingContextMetadata":               4048,
	"ErrNoTokenProvided":                      4049,
	"ErrGetTokenWithClaims":                   4050,
	"ErrNoTokenInContext":                     4051,
	"ErrGetSubjectFromToken":                  4052,
	"ErrGetSubjectFromContext":                4053,
	"ErrGetActiveSeasonMembership":            4054,
	"ErrGetTokenFromContext":                  4055,
	"ErrChallengeAlreadyClosed":               4056,
	"ErrGetAgent":                             4057,
	"ErrSaveAgent":                            4058,
	"ErrListChallengeInstances":               4059,
	"ErrChallengeAdd":                         4060,
	"ErrCouponAlreadyValidatedBySameTeam":     4061,
	"ErrCouponExpired":                        4062,
	"ErrCouponNotFound":                       4063,
	"ErrUserDoesNotBelongToTeam":              4064,
	"ErrServerListen":                         5001,
	"ErrServerRegisterGateway":                5002,
	"ErrInitLogger":                           6001,
	"ErrStartService":                         6002,
	"ErrInitServer":                           6003,
	"ErrGroupTerminated":                      6004,
	"ErrGetSSOClientFromFlags":                6005,
	"ErrDumpDatabase":                         6006,
	"ErrGetDBInfo":                            6007,
	"ErrGetSSOWhoami":                         6008,
	"ErrGetSSOLogout":                         6009,
	"ErrGetSSOClaims":                         6010,
	"ErrInitDockerClient":                     6011,
	"ErrInitDB":                               6012,
	"ErrConfigureDB":                          6013,
	"ErrInitSSOClient":                        6014,
	"ErrInitService":                          6015,
	"ErrAgentGetContainersInfo":               7001,
	"ErrCheckNginxContainer":                  7002,
	"ErrRemoveNginxContainer":                 7003,
	"ErrBuildNginxContainer":                  7004,
	"ErrStartNginxContainer":                  7005,
	"ErrParsingTemplate":                      7006,
	"ErrWriteConfigFileHeader":                7007,
	"ErrWriteConfigFile":                      7008,
	"ErrCloseTarWriter":                       7009,
	"ErrCopyNginxConfigToContainer":           7010,
	"ErrNginxNewConfigCheckFailed":            7011,
	"ErrNginxSendCommandNewConfigCheck":       7012,
	"ErrNginxSendCommandNewConfigRemove":      7013,
	"ErrNginxSendCommandConfigReplace":        7014,
	"ErrNginxSendCommandReloadConfig":         7015,
	"ErrNginxConnectNetwork":                  7016,
	"ErrContainerConnectNetwork":              7017,
	"ErrNatPortOpening":                       7018,
	"ErrBuildNginxConfig":                     7019,
	"ErrExecuteTemplate":                      7020,
	"ErrWriteBytesToHashBuilder":              7021,
	"ErrReadBytesFromHashBuilder":             7022,
	"ErrGeneratePrefixHash":                   7023,
	"ErrCleanPathwarInstances":                7024,
	"ErrParseInitConfig":                      7025,
	"ErrUpPathwarInstance":                    7026,
	"ErrUpdateNginx":                          7027,
	"ErrAgentUpdateState":                     7028,
	"ErrDockerAPIContainerList":               8001,
	"ErrDockerAPIContainerRemove":             8002,
	"ErrDockerAPIImageRemove":                 8003,
	"ErrDockerAPIContainerCreate":             8004,
	"ErrDockerAPIContainerExecCreate":         8005,
	"ErrDockerAPIContainerExecAttach":         8006,
	"ErrDockerAPIContainerExecStart":          8007,
	"ErrDockerAPIContainerExecInspect":        8008,
	"ErrDockerAPIImagePull":                   8009,
	"ErrDockerAPIImageInspect":                8010,
	"ErrDockerAPINetworkList":                 8011,
	"ErrDockerAPINetworkCreate":               8012,
	"ErrDockerAPINetworkRemove":               8013,
	"ErrDockerAPIExitCode":                    8014,
	"ErrExecuteOnInitHook":                    9001,
	"ErrRemoveInitConfig":                     9002,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4240057316120df7, []int{0}
}

func init() {
	proto.RegisterEnum("pathwar.errcode.ErrCode", ErrCode_name, ErrCode_value)
}

func init() { proto.RegisterFile("errcode.proto", fileDescriptor_4240057316120df7) }

var fileDescriptor_4240057316120df7 = []byte{
	// 2215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x58, 0x47, 0x73, 0x1c, 0xc7,
	0x15, 0x26, 0xab, 0x6c, 0xa1, 0x34, 0x36, 0x85, 0xa7, 0xa1, 0xc8, 0x95, 0x48, 0x09, 0xa3, 0x60,
	0x93, 0x2e, 0x87, 0xe5, 0xc1, 0x55, 0x5b, 0xe5, 0x0b, 0xaa, 0xb0, 0x01, 0xe0, 0x96, 0xc8, 0x05,
	0x0a, 0x0b, 0x88, 0x55, 0xbe, 0x35, 0x76, 0xde, 0xce, 0xb6, 0x31, 0xdb, 0x3d, 0xea, 0xe9, 0x41,
	0xf0, 0x3f, 0x90, 0x4f, 0x3e, 0xfb, 0xe6, 0x6c, 0xc9, 0x39, 0x5b, 0x54, 0xa4, 0x98, 0xa3, 0x98,
	0x94, 0x13, 0x45, 0xda, 0x96, 0x45, 0xe5, 0x60, 0x5b, 0xce, 0xae, 0xee, 0xe9, 0x9e, 0x9d, 0x59,
	0x80, 0xba, 0xed, 0xbe, 0xd4, 0xef, 0x7d, 0x2f, 0xf4, 0xeb, 0x71, 0x36, 0xa1, 0x10, 0x1d, 0xee,
	0x63, 0x39, 0x12, 0x5c, 0x72, 0x77, 0x34, 0x22, 0xb2, 0xb7, 0x4c, 0x44, 0xd9, 0x90, 0xb7, 0x7d,
	0x29, 0xa0, 0xb2, 0x97, 0x2c, 0x94, 0x3b, 0xbc, 0xbf, 0x2b, 0xe0, 0x01, 0xdf, 0xa5, 0xe5, 0x16,
	0x92, 0xae, 0xfe, 0xa7, 0xff, 0xe8, 0x5f, 0xa9, 0xfe, 0xe7, 0x8f, 0xef, 0x70, 0x46, 0x1a, 0x42,
	0xd4, 0xb8, 0x8f, 0xee, 0x26, 0xe7, 0xfa, 0x79, 0xe6, 0x63, 0x97, 0x32, 0xf4, 0x61, 0x83, 0x7b,
	0xbd, 0xf3, 0x89, 0xb9, 0xe9, 0xfa, 0x34, 0x7c, 0xeb, 0x93, 0xee, 0x56, 0xe7, 0xc6, 0x86, 0x10,
	0x2d, 0x2e, 0x9b, 0xfd, 0x28, 0xc4, 0x3e, 0x32, 0x89, 0x3e, 0xdc, 0x77, 0x9d, 0xeb, 0x3a, 0x9b,
	0x1a, 0x42, 0xd4, 0x31, 0x12, 0xd8, 0x21, 0x8a, 0xf6, 0xd1, 0x75, 0x2e, 0x38, 0x9f, 0x6a, 0x08,
	0xd1, 0x64, 0x12, 0x05, 0x23, 0x21, 0xbc, 0x36, 0xe2, 0x6e, 0x76, 0x46, 0x35, 0x65, 0x89, 0x84,
	0xd4, 0x6f, 0xb2, 0x28, 0x91, 0x80, 0x86, 0xb8, 0x97, 0xc6, 0x31, 0x65, 0x41, 0x4a, 0xec, 0xba,
	0x5b, 0x1d, 0xb7, 0x21, 0xc4, 0x3c, 0x23, 0x89, 0xec, 0x21, 0x93, 0x34, 0x35, 0x1a, 0xb8, 0x5b,
	0xf4, 0xf9, 0xb3, 0x18, 0x4b, 0x41, 0x3b, 0x12, 0xfd, 0x09, 0x81, 0x04, 0x7a, 0xe6, 0xf8, 0x76,
	0x7b, 0x7a, 0x0a, 0xe5, 0x74, 0xb3, 0x5e, 0x83, 0xd7, 0x47, 0xdc, 0xed, 0xce, 0xd6, 0x94, 0x66,
	0xce, 0x9b, 0x49, 0x16, 0x42, 0xda, 0xb9, 0x1b, 0x57, 0xe1, 0xea, 0x88, 0x7b, 0xbb, 0xb3, 0x3d,
	0x65, 0x4e, 0x12, 0x1a, 0xa2, 0x7f, 0x37, 0xae, 0x76, 0x42, 0x4e, 0x16, 0x67, 0xf1, 0xde, 0x04,
	0x63, 0x09, 0x6f, 0x8c, 0xb8, 0x77, 0x3a, 0xb7, 0x15, 0xd4, 0x07, 0x22, 0x71, 0xc4, 0x59, 0x8c,
	0xf0, 0xe6, 0x88, 0x7b, 0xa3, 0xf3, 0xe9, 0x54, 0x66, 0x0f, 0x0f, 0x78, 0x22, 0xe1, 0xad, 0x11,
	0xf7, 0x36, 0xe7, 0x66, 0xab, 0x46, 0xa5, 0xd5, 0xa9, 0x85, 0x14, 0x99, 0x84, 0xb7, 0x47, 0xdc,
	0x9b, 0x9d, 0xcd, 0x05, 0xab, 0x55, 0x24, 0x02, 0x05, 0xbc, 0x93, 0xe3, 0x58, 0xa5, 0x86, 0x10,
	0x5c, 0xc0, 0xbb, 0x23, 0x16, 0xdb, 0x6a, 0x8b, 0xcb, 0x49, 0x9e, 0x30, 0x1f, 0xce, 0x8d, 0x66,
	0xb4, 0x0c, 0xdd, 0xf3, 0xa3, 0x6e, 0x49, 0x63, 0x56, 0xaf, 0xce, 0x26, 0x6c, 0x2f, 0x0d, 0x04,
	0x91, 0x94, 0xb3, 0x18, 0x2e, 0x8c, 0xba, 0x37, 0x38, 0xd7, 0x1b, 0x61, 0x2a, 0xe1, 0xe2, 0xa8,
	0x71, 0xbb, 0x5e, 0xad, 0x71, 0xc6, 0xb0, 0x23, 0xe1, 0xa9, 0x51, 0x77, 0x8b, 0x03, 0x9a, 0x34,
	0x91, 0x48, 0x9e, 0x2a, 0x23, 0x3c, 0x3d, 0x30, 0x39, 0xe1, 0xfb, 0x93, 0x5c, 0x20, 0x0d, 0x98,
	0xc2, 0xef, 0x99, 0x51, 0x77, 0x9b, 0xb3, 0x45, 0x17, 0x4b, 0x3f, 0xe2, 0x31, 0x5a, 0x80, 0x89,
	0xec, 0xc1, 0x83, 0x25, 0x83, 0xad, 0xe1, 0xd5, 0xa9, 0xc0, 0x8e, 0xe4, 0x62, 0x35, 0xf3, 0x7e,
	0x7f, 0xc9, 0xbd, 0xc5, 0xb9, 0x69, 0x20, 0x31, 0x8b, 0xc4, 0xaf, 0x71, 0xd6, 0xa5, 0x01, 0x3c,
	0x54, 0x72, 0x6f, 0x75, 0x4a, 0x6b, 0x0c, 0x1b, 0xee, 0xc3, 0x43, 0xdc, 0xbd, 0x44, 0xc4, 0x3d,
	0x12, 0x1a, 0xee, 0x23, 0x25, 0x83, 0xbd, 0xe1, 0xd6, 0x04, 0x12, 0x89, 0x73, 0xd8, 0x8f, 0x26,
	0x69, 0x88, 0xf0, 0xe8, 0x90, 0xf2, 0x3e, 0x41, 0x73, 0xdc, 0xc7, 0x86, 0xb8, 0xb5, 0x90, 0xc7,
	0x03, 0xee, 0xe3, 0x25, 0xf7, 0x26, 0x5d, 0xa4, 0x86, 0x5b, 0x4d, 0x68, 0xe8, 0xc3, 0x81, 0x92,
	0xbb, 0x55, 0xa3, 0x96, 0x51, 0x99, 0x1f, 0x22, 0xec, 0xbf, 0xba, 0xd1, 0x74, 0x49, 0x2e, 0xbe,
	0x3a, 0x59, 0x80, 0x83, 0x25, 0x03, 0xa7, 0xa1, 0xcf, 0x10, 0x11, 0xa3, 0x62, 0x1c, 0x2a, 0x15,
	0xe1, 0xd4, 0x0c, 0x13, 0xd5, 0xe1, 0x61, 0xc7, 0xb2, 0xa8, 0xea, 0x54, 0xc0, 0x91, 0xa1, 0x98,
	0xe7, 0x23, 0x3f, 0x1f, 0xf3, 0xd1, 0xa1, 0x5c, 0x4c, 0x72, 0xd1, 0xc1, 0x59, 0xec, 0x68, 0x1b,
	0x75, 0xbe, 0xcc, 0xe0, 0x58, 0xc9, 0xd4, 0x9d, 0xf5, 0x35, 0x61, 0xe9, 0x09, 0x70, 0x7c, 0x28,
	0xe6, 0xd9, 0x84, 0xcd, 0x47, 0x70, 0xc2, 0xc6, 0x30, 0x85, 0x72, 0x66, 0x9f, 0xaa, 0xa7, 0x2a,
	0x65, 0x44, 0xac, 0xc2, 0x49, 0xeb, 0x89, 0xc6, 0x35, 0x65, 0x29, 0x1f, 0x76, 0x23, 0xf1, 0x51,
	0xc0, 0x29, 0xab, 0x37, 0xc4, 0x86, 0xd3, 0x25, 0xd7, 0x73, 0xb6, 0xa9, 0xfe, 0x4f, 0x93, 0x99,
	0xb2, 0xd2, 0xe0, 0xb5, 0xc0, 0x99, 0x92, 0x7b, 0x97, 0x33, 0x56, 0xd4, 0x1c, 0xb0, 0x8d, 0xf9,
	0x27, 0xd7, 0x39, 0x3d, 0x67, 0xe3, 0x6c, 0xc9, 0xbd, 0xc3, 0xb9, 0x75, 0x88, 0xad, 0x33, 0x4c,
	0x52, 0x92, 0x80, 0x73, 0x03, 0x24, 0xa3, 0xd5, 0x54, 0x62, 0x8e, 0xd7, 0x38, 0x93, 0x84, 0x32,
	0x14, 0x70, 0x7e, 0x08, 0xc9, 0x29, 0x94, 0x19, 0x33, 0x6e, 0xb2, 0x2e, 0x87, 0x0b, 0x25, 0x33,
	0x70, 0xcc, 0x20, 0x9b, 0x59, 0xa6, 0x99, 0x13, 0x70, 0xd1, 0x66, 0x71, 0x0a, 0xe5, 0x7c, 0x8c,
	0xa2, 0x59, 0x9f, 0x14, 0xbc, 0xaf, 0x2c, 0xe0, 0x8a, 0x84, 0x6f, 0x7b, 0x66, 0xd8, 0x18, 0xd5,
	0x5a, 0x8f, 0x84, 0x21, 0xb2, 0x00, 0xef, 0x51, 0xc5, 0xaf, 0xdb, 0x18, 0xbe, 0xe3, 0x99, 0x16,
	0x35, 0x2d, 0xd1, 0x46, 0x12, 0x73, 0x06, 0xdf, 0xf5, 0x0c, 0xae, 0x73, 0x48, 0xfa, 0x6a, 0x2a,
	0x33, 0xc3, 0xf8, 0x9e, 0x67, 0x1c, 0x56, 0x9e, 0x5a, 0x7b, 0xed, 0x64, 0x21, 0xee, 0x08, 0x1a,
	0x69, 0x8b, 0xdf, 0x1f, 0x58, 0xa4, 0xb2, 0xcd, 0xf8, 0x72, 0x37, 0x24, 0x8b, 0x08, 0x3f, 0xf0,
	0x0c, 0xde, 0x69, 0x2d, 0xad, 0xaf, 0xfb, 0x43, 0xcf, 0x00, 0x9a, 0x16, 0xcb, 0x7a, 0x0e, 0xff,
	0xc8, 0x73, 0xc7, 0x9c, 0x5b, 0x86, 0x1c, 0xc8, 0xf1, 0xef, 0xf7, 0xdc, 0xcd, 0xce, 0x0d, 0x83,
	0x80, 0x54, 0x00, 0xf0, 0x80, 0x45, 0x22, 0xd3, 0x98, 0x08, 0x05, 0x12, 0x7f, 0xd5, 0x9c, 0xbe,
	0x80, 0x3e, 0xfc, 0xd8, 0x3a, 0x38, 0x74, 0x76, 0xc1, 0xc1, 0x9f, 0x78, 0x66, 0xc6, 0x4c, 0x52,
	0xe6, 0x4f, 0x8b, 0x80, 0x30, 0xfa, 0x75, 0x33, 0x0f, 0x7f, 0xea, 0xb9, 0x9f, 0x71, 0xbc, 0xd4,
	0xb1, 0x14, 0x2c, 0x95, 0x8b, 0xf4, 0x57, 0x66, 0x0c, 0x7e, 0xe6, 0x99, 0x7a, 0x30, 0x19, 0x53,
	0xee, 0x0d, 0xe4, 0xe0, 0xe7, 0x16, 0xf7, 0x42, 0x3a, 0x9a, 0x75, 0xf8, 0x85, 0x0d, 0x5b, 0x29,
	0xed, 0x26, 0x71, 0x8b, 0x6b, 0x4d, 0x2e, 0x8c, 0xe2, 0x2f, 0x3d, 0x53, 0x26, 0xd9, 0xe9, 0xd9,
	0x99, 0x31, 0xfc, 0xca, 0x33, 0xa3, 0x39, 0x63, 0xc2, 0xaf, 0x3d, 0xd3, 0x86, 0xe9, 0xff, 0x3a,
	0x32, 0x8a, 0x3e, 0xfc, 0xc6, 0x33, 0x5d, 0x63, 0xe0, 0xd9, 0x4d, 0xe2, 0xe2, 0x31, 0xbf, 0xb5,
	0x6a, 0xb3, 0x18, 0xa3, 0x58, 0x42, 0xbf, 0x45, 0xfa, 0x08, 0xbf, 0xcb, 0xa0, 0xeb, 0x61, 0x67,
	0x31, 0x0f, 0xcb, 0x3c, 0xa3, 0xf7, 0x26, 0xa8, 0x85, 0x7e, 0xef, 0xd9, 0x69, 0xa4, 0xf1, 0xcd,
	0x4b, 0xc1, 0x1f, 0x3c, 0xf7, 0x0b, 0xce, 0x8e, 0x86, 0x10, 0x79, 0xea, 0xb5, 0x7c, 0x78, 0xd0,
	0x1b, 0xcc, 0x8a, 0x82, 0x95, 0xfd, 0xf6, 0x84, 0xb5, 0x18, 0xc0, 0x43, 0xf6, 0x84, 0x1a, 0x61,
	0x8c, 0x4b, 0x3b, 0xee, 0x52, 0xbb, 0x3c, 0xe4, 0x05, 0x43, 0x0f, 0xdb, 0x24, 0x29, 0xb0, 0x75,
	0xf5, 0x17, 0xd8, 0x8f, 0x78, 0xe6, 0x9a, 0x1c, 0x58, 0x81, 0x47, 0x3d, 0x77, 0xd4, 0x71, 0xd2,
	0xb3, 0x35, 0xe1, 0x31, 0xcf, 0xec, 0x29, 0x86, 0x10, 0xc3, 0xe3, 0x39, 0x11, 0x65, 0x18, 0x0e,
	0x58, 0x3b, 0x69, 0x4b, 0x68, 0xda, 0x13, 0x45, 0x9a, 0x36, 0x75, 0xd0, 0xc6, 0x95, 0xd2, 0x0a,
	0xbe, 0x1c, 0xb2, 0x05, 0xd9, 0xc2, 0x65, 0x65, 0x40, 0xf7, 0x7f, 0x48, 0x68, 0x3f, 0x86, 0xc3,
	0x36, 0x57, 0x0a, 0xa7, 0x89, 0x44, 0xf6, 0xf4, 0x01, 0x47, 0x3c, 0xf7, 0x8b, 0xce, 0x4e, 0x75,
	0xf9, 0xd2, 0x6e, 0x17, 0x05, 0x32, 0xed, 0x4b, 0x15, 0xe5, 0x32, 0x22, 0x9b, 0xe3, 0x8b, 0xc8,
	0x26, 0x98, 0x5f, 0x27, 0x92, 0x2c, 0x90, 0x18, 0xe1, 0xa8, 0xc5, 0x7a, 0x0f, 0x27, 0xbe, 0x12,
	0x4c, 0x71, 0x8d, 0xe1, 0x98, 0x57, 0x9c, 0x3c, 0xc5, 0x5e, 0x38, 0x6e, 0xa3, 0xc8, 0x32, 0x11,
	0xc3, 0x09, 0xcf, 0x5c, 0x09, 0x46, 0xa3, 0xaa, 0x9a, 0xef, 0x6b, 0x6a, 0x4d, 0x38, 0x69, 0xab,
	0xae, 0xd1, 0x27, 0x34, 0x9c, 0xf0, 0x7d, 0x81, 0x71, 0xdc, 0xe2, 0xf2, 0x1e, 0x14, 0xb4, 0xab,
	0xca, 0xf2, 0x54, 0x4e, 0xb5, 0x8e, 0x5d, 0x92, 0x84, 0xb6, 0x8c, 0x4f, 0x7b, 0x83, 0x8b, 0xaa,
	0x4f, 0xd3, 0x8e, 0x12, 0x84, 0xc5, 0xa4, 0xa3, 0xd1, 0x39, 0x53, 0x44, 0x6e, 0xa2, 0x23, 0xe9,
	0x12, 0x1a, 0xd5, 0x27, 0x6d, 0x47, 0xd9, 0xe9, 0x98, 0x4e, 0xcd, 0xbd, 0x28, 0x89, 0x4f, 0x24,
	0x81, 0xb3, 0x36, 0xf4, 0x16, 0xd7, 0xb0, 0xcc, 0x08, 0xbe, 0x44, 0x7d, 0xf4, 0xe1, 0x5c, 0xae,
	0xcc, 0x34, 0x67, 0x1f, 0x95, 0x3d, 0x83, 0xf9, 0x79, 0xeb, 0xa9, 0x51, 0x6a, 0x32, 0x3b, 0x8c,
	0x2f, 0xe4, 0x1b, 0x34, 0x0d, 0x5c, 0xe5, 0x4a, 0x4b, 0xc1, 0xc5, 0xdc, 0x54, 0xc8, 0x31, 0xad,
	0xee, 0x53, 0x76, 0x2c, 0x4e, 0xa1, 0xcc, 0xc7, 0xb0, 0x17, 0xfb, 0x0b, 0x28, 0xe2, 0x1e, 0x8d,
	0xe0, 0xe9, 0x9c, 0x79, 0x6d, 0x33, 0xaf, 0xff, 0x8c, 0x0d, 0x75, 0x78, 0xfc, 0xe9, 0xcb, 0xca,
	0x87, 0x67, 0x73, 0xb5, 0x3a, 0x11, 0xa8, 0x8d, 0xf2, 0x39, 0x3b, 0x31, 0xda, 0x64, 0x09, 0x53,
	0xd2, 0xf3, 0xd6, 0xc8, 0x1e, 0x1a, 0x0f, 0x26, 0x6f, 0x93, 0xc5, 0x92, 0xb0, 0x0e, 0xc6, 0xf0,
	0x82, 0x2d, 0xb7, 0xc1, 0x21, 0xbe, 0x0f, 0x2f, 0x7a, 0xee, 0xe7, 0x9c, 0xbb, 0x74, 0x82, 0x92,
	0x28, 0xeb, 0x69, 0x33, 0xaf, 0xd1, 0xaf, 0xae, 0xb6, 0x49, 0x3f, 0xad, 0xf2, 0x97, 0xec, 0xbd,
	0x91, 0x4a, 0x36, 0x56, 0x22, 0x2a, 0xd0, 0x87, 0x97, 0xbd, 0x6c, 0xeb, 0x51, 0xe4, 0x6c, 0xdb,
	0x7b, 0xc5, 0x16, 0x8d, 0xca, 0x79, 0x9d, 0xa3, 0x2a, 0x98, 0x2a, 0x86, 0x9c, 0x05, 0x73, 0x7a,
	0x34, 0xc2, 0xa5, 0xc1, 0x84, 0x13, 0x4b, 0xa8, 0xbd, 0x46, 0x06, 0xf7, 0xed, 0xb4, 0x9b, 0xb4,
	0xa6, 0xce, 0x62, 0xa0, 0xe8, 0x62, 0x8a, 0x48, 0x5c, 0x26, 0xab, 0xf0, 0x8d, 0x9d, 0xa6, 0x70,
	0xd5, 0xe5, 0xb5, 0x87, 0x07, 0x01, 0x0a, 0x78, 0xaf, 0x6c, 0x0d, 0x49, 0x22, 0xa4, 0xd2, 0xa3,
	0x1d, 0x84, 0xf7, 0xcb, 0x39, 0xc9, 0xd4, 0x18, 0x7c, 0x50, 0xb6, 0x93, 0x49, 0xf0, 0x24, 0x9a,
	0x43, 0xd1, 0xa7, 0x4c, 0xbf, 0x2f, 0x3e, 0x2c, 0xe7, 0xf2, 0xdb, 0x9e, 0x4e, 0xd7, 0x76, 0x95,
	0xa1, 0xc9, 0x90, 0x04, 0x31, 0xfc, 0xd5, 0x9e, 0x50, 0x4f, 0xfa, 0x51, 0xd6, 0x7b, 0x7f, 0x2b,
	0x0f, 0xa6, 0xb6, 0xda, 0xb1, 0xbb, 0x1c, 0xfe, 0x5e, 0x1e, 0xb4, 0x74, 0xbb, 0x3d, 0xbd, 0xaf,
	0xc7, 0x49, 0x9f, 0xc2, 0x47, 0x45, 0xaa, 0x79, 0x33, 0xfc, 0xa3, 0x48, 0x35, 0x05, 0xfa, 0xcf,
	0xb2, 0x29, 0x50, 0xe5, 0x76, 0x9d, 0x77, 0x16, 0x51, 0x98, 0x47, 0xc4, 0xbf, 0xca, 0x66, 0x9f,
	0xd7, 0x9c, 0x2a, 0xfc, 0xbb, 0x6c, 0x2e, 0xd2, 0x74, 0xd7, 0x48, 0x04, 0xd6, 0xab, 0xf0, 0x9f,
	0x72, 0xfe, 0x72, 0xb7, 0x91, 0xc0, 0x7f, 0xcb, 0xd9, 0xa5, 0x4b, 0x33, 0x84, 0xfe, 0x57, 0x36,
	0x05, 0xa3, 0xeb, 0x67, 0xed, 0x66, 0xf3, 0x5c, 0xc5, 0x94, 0xac, 0xbe, 0x35, 0x5a, 0x01, 0x65,
	0x2b, 0x83, 0xc5, 0xe8, 0xf9, 0x8a, 0x99, 0x2f, 0xb3, 0xd8, 0xe7, 0x4b, 0x38, 0xc4, 0x7d, 0xc1,
	0xaa, 0xea, 0x8d, 0x79, 0x88, 0xf9, 0xa2, 0x65, 0xea, 0x7c, 0x0d, 0x31, 0x5f, 0xaa, 0x98, 0x14,
	0xa9, 0x65, 0x98, 0xb2, 0x40, 0xed, 0xb4, 0xa1, 0xda, 0x4b, 0x5f, 0xae, 0xe4, 0x57, 0xbd, 0x35,
	0x9b, 0xe0, 0x2b, 0x95, 0xfc, 0xa2, 0x99, 0xdb, 0x01, 0x2f, 0x55, 0x6c, 0x7d, 0x16, 0x17, 0xbf,
	0x57, 0x2b, 0x76, 0xe5, 0xe0, 0xd1, 0xaa, 0x75, 0xa2, 0x4b, 0x83, 0xfc, 0xf6, 0x77, 0xb9, 0x62,
	0xfa, 0x5a, 0xf3, 0x5b, 0xb8, 0x9c, 0x8a, 0x68, 0x3c, 0xd2, 0xf7, 0x23, 0x5c, 0xa9, 0xb8, 0x3b,
	0x9c, 0x3b, 0xac, 0x48, 0x1b, 0x99, 0xaf, 0x26, 0x1d, 0x61, 0x7e, 0x51, 0x1a, 0xfe, 0x58, 0x71,
	0x77, 0x3a, 0x77, 0x7e, 0x9c, 0x5c, 0x0a, 0x24, 0xfc, 0xa9, 0xe2, 0x7e, 0xd6, 0xb9, 0x7d, 0x1d,
	0x41, 0x2b, 0x15, 0x85, 0xa4, 0x83, 0xf0, 0xe7, 0x8a, 0xd9, 0x66, 0x86, 0xc5, 0x66, 0x31, 0xe4,
	0xd9, 0xbb, 0xea, 0x35, 0x0b, 0xb5, 0x0d, 0x50, 0x3d, 0xfb, 0x5a, 0x28, 0x97, 0xb9, 0x58, 0x84,
	0xbf, 0x54, 0x4c, 0x87, 0x66, 0x01, 0x0f, 0x09, 0xbc, 0x6e, 0xa1, 0x6b, 0x11, 0x39, 0xc3, 0x85,
	0x9c, 0x8e, 0x90, 0x51, 0x16, 0xc0, 0xd5, 0x8a, 0xa9, 0xd1, 0x42, 0x76, 0xd5, 0x79, 0x6f, 0xd8,
	0x2c, 0x34, 0x56, 0xb0, 0x93, 0xa4, 0x2f, 0x12, 0x9d, 0xbd, 0x37, 0xed, 0x59, 0x1a, 0xfd, 0xea,
	0xaa, 0xc4, 0x78, 0x8e, 0xef, 0x26, 0x71, 0x4f, 0x9b, 0x40, 0x01, 0x6f, 0x55, 0xcc, 0xde, 0xaa,
	0x5e, 0x4d, 0x9a, 0xaf, 0xda, 0x2f, 0x2f, 0xf1, 0x76, 0x25, 0x1b, 0xeb, 0x0c, 0xd5, 0x3b, 0x75,
	0x46, 0x60, 0x97, 0xae, 0x28, 0x11, 0x78, 0xc7, 0x16, 0x47, 0x2d, 0x44, 0xc2, 0x66, 0xd2, 0x0f,
	0x22, 0x83, 0xd1, 0xf7, 0x6e, 0xbe, 0xa8, 0x70, 0xf0, 0x48, 0x80, 0xf7, 0x2a, 0xe6, 0x76, 0x9e,
	0x8f, 0x86, 0x94, 0xe0, 0xfd, 0x8a, 0x69, 0x99, 0xf4, 0x6a, 0xd2, 0x51, 0xc2, 0x07, 0x36, 0x72,
	0xdd, 0x32, 0x29, 0xa7, 0x2d, 0x55, 0x80, 0x1f, 0x56, 0x4c, 0x33, 0xa5, 0x3d, 0x3b, 0x31, 0xd3,
	0xcc, 0x50, 0x55, 0x93, 0x0d, 0x0e, 0x8c, 0x9b, 0xf8, 0xd6, 0xf2, 0x4d, 0xe2, 0x9f, 0x18, 0x37,
	0x1d, 0x95, 0x49, 0x34, 0xfb, 0x24, 0x40, 0xc3, 0x3d, 0x78, 0x6d, 0x7d, 0xf3, 0x70, 0x3b, 0x34,
	0x6e, 0x2a, 0x62, 0xad, 0x84, 0xca, 0x86, 0x91, 0x3a, 0xfc, 0xf1, 0x52, 0x13, 0x52, 0x92, 0x4e,
	0x0f, 0x8e, 0x8c, 0x9b, 0x85, 0x71, 0x7d, 0x29, 0xdd, 0xb8, 0x70, 0x74, 0xdc, 0x54, 0xea, 0xfa,
	0x42, 0x4d, 0x16, 0x47, 0x6a, 0x7b, 0x38, 0x36, 0x6e, 0xf2, 0x56, 0x8c, 0x6b, 0x26, 0x09, 0x43,
	0x38, 0x3e, 0x6e, 0xf2, 0x56, 0xe4, 0x59, 0xd5, 0x13, 0x6b, 0x20, 0x31, 0xa5, 0xa9, 0x21, 0x3d,
	0x39, 0x3e, 0x0c, 0xb9, 0xe1, 0x9a, 0x50, 0x4f, 0x5d, 0x8b, 0x6f, 0x20, 0x3d, 0x3d, 0x6e, 0x92,
	0x9f, 0xf1, 0x1b, 0x2b, 0xaa, 0x32, 0x7c, 0x84, 0x33, 0x96, 0x65, 0xea, 0x78, 0x9a, 0xa9, 0xa2,
	0xd9, 0xcd, 0xf9, 0x22, 0xdc, 0x3f, 0x69, 0x4a, 0x20, 0xb5, 0x92, 0x2b, 0xa6, 0x07, 0x26, 0xab,
	0x5f, 0x39, 0xfb, 0xea, 0xd8, 0x86, 0x63, 0x97, 0xc7, 0x36, 0x9e, 0xbd, 0x3c, 0xb6, 0xf1, 0xd2,
	0xe5, 0xb1, 0x8d, 0xdf, 0xbc, 0x32, 0xb6, 0xe1, 0xec, 0x95, 0xb1, 0x0d, 0xcf, 0x5e, 0x19, 0xdb,
	0xf0, 0xd5, 0xed, 0xf6, 0x4b, 0x5d, 0x48, 0x98, 0xbf, 0x2b, 0xe0, 0xbb, 0xa2, 0xc5, 0x60, 0x97,
	0xf9, 0x6a, 0xb7, 0x70, 0x9d, 0xfe, 0x1a, 0xf7, 0xe5, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x30,
	0x9d, 0x6b, 0x37, 0xde, 0x13, 0x00, 0x00,
}
