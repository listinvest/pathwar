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
	ErrComposeBuild                         ErrCode = 3008
	ErrComposeBundle                        ErrCode = 3009
	ErrComposeReadDab                       ErrCode = 3010
	ErrComposeParseDab                      ErrCode = 3011
	ErrComposeParseConfig                   ErrCode = 3012
	ErrComposeCreateTempDir                 ErrCode = 3013
	ErrComposeForceRecreateDown             ErrCode = 3014
	ErrComposeRunCreate                     ErrCode = 3015
	ErrComposeRunUp                         ErrCode = 3016
	ErrGetPWInitBinary                      ErrCode = 3017
	ErrWritePWInitFileHeader                ErrCode = 3018
	ErrWritePWInitFile                      ErrCode = 3019
	ErrMarshalPWInitConfigFile              ErrCode = 3020
	ErrWritePWInitConfigFileHeader          ErrCode = 3021
	ErrWritePWInitConfigFile                ErrCode = 3022
	ErrWritePWInitCloseTarWriter            ErrCode = 3023
	ErrCopyPWInitToContainer                ErrCode = 3024
	ErrComposeGetContainersInfo             ErrCode = 3025
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
	ErrDockerAPIContainerList               ErrCode = 8001
	ErrDockerAPIContainerRemove             ErrCode = 8002
	ErrDockerAPIImageRemove                 ErrCode = 8003
	ErrDockerAPIContainerCreate             ErrCode = 8004
	ErrDockerAPIContainerExecCreate         ErrCode = 8005
	ErrDockerAPIContainerExecAttach         ErrCode = 8006
	ErrDockerAPIContainerExecStart          ErrCode = 8007
	ErrDockerAPIContainerExecInspect        ErrCode = 8008
	ErrDockerAPIImagePull                   ErrCode = 8009
	ErrDockerAPINetworkList                 ErrCode = 8010
	ErrDockerAPINetworkCreate               ErrCode = 8011
	ErrDockerAPINetworkRemove               ErrCode = 8012
	ErrDockerAPIExitCode                    ErrCode = 8013
	ErrExecuteOnInitHook                    ErrCode = 9001
	ErrRemoveInitConfig                     ErrCode = 9002
)

var ErrCode_name = map[int32]string{
	0:    "Undefined",
	666:  "TODO",
	777:  "ErrNotImplemented",
	888:  "ErrDeprecated",
	999:  "ErrInternal",
	101:  "ErrInvalidInput",
	102:  "ErrMissingInput",
	103:  "ErrUnauthenticated",
	1001: "ErrSSOGetOIDC",
	1002: "ErrSSOInvalidPublicKey",
	1003: "ErrSSOFailedKeycloakRequest",
	1004: "ErrSSOInvalidKeycloakResponse",
	1005: "ErrSSOLogout",
	1006: "ErrSSOInitKeycloakClient",
	1007: "ErrSSOInvalidBearer",
	1008: "ErrSSOKeycloakError",
	2001: "ErrDBNotFound",
	2002: "ErrDBInternal",
	2003: "ErrDBRunMigrations",
	2004: "ErrDBInit",
	2005: "ErrDBConnect",
	2006: "ErrDBAutoMigrate",
	2007: "ErrDBAddForeignKey",
	3001: "ErrComposeInvalidPath",
	3002: "ErrComposeDirectoryNotFound",
	3003: "ErrComposeReadConfig",
	3004: "ErrComposeInvalidConfig",
	3005: "ErrComposeMarshalConfig",
	3006: "ErrComposeCreateTempFile",
	3007: "ErrComposeWriteTempFile",
	3008: "ErrComposeBuild",
	3009: "ErrComposeBundle",
	3010: "ErrComposeReadDab",
	3011: "ErrComposeParseDab",
	3012: "ErrComposeParseConfig",
	3013: "ErrComposeCreateTempDir",
	3014: "ErrComposeForceRecreateDown",
	3015: "ErrComposeRunCreate",
	3016: "ErrComposeRunUp",
	3017: "ErrGetPWInitBinary",
	3018: "ErrWritePWInitFileHeader",
	3019: "ErrWritePWInitFile",
	3020: "ErrMarshalPWInitConfigFile",
	3021: "ErrWritePWInitConfigFileHeader",
	3022: "ErrWritePWInitConfigFile",
	3023: "ErrWritePWInitCloseTarWriter",
	3024: "ErrCopyPWInitToContainer",
	3025: "ErrComposeGetContainersInfo",
	4001: "ErrGetUserIDFromContext",
	4002: "ErrMissingChallengeValidation",
	4003: "ErrInvalidSeason",
	4004: "ErrTeamNotInSeason",
	4005: "ErrGetChallengeSubscription",
	4006: "ErrInitSnowflake",
	4007: "ErrUpdateChallengeSubscription",
	4008: "ErrCreateChallengeValidation",
	4009: "ErrGetChallengeValidation",
	4010: "ErrInvalidTeam",
	4011: "ErrChallengeAlreadySubscribed",
	4012: "ErrCreateChallengeSubscription",
	4013: "ErrFindOrganizations",
	4014: "ErrGetSeasonFromSeasonChallenge",
	4015: "ErrGetUserTeamFromSeason",
	4016: "ErrInvalidSeasonID",
	4017: "ErrUserHasNoTeamForSeason",
	4018: "ErrGetSeasonChallenges",
	4019: "ErrGetSeason",
	4020: "ErrSeasonDenied",
	4021: "ErrAlreadyHasTeamForSeason",
	4022: "ErrReservedName",
	4023: "ErrCheckOrganizationUniqueName",
	4024: "ErrCreateOrganization",
	4025: "ErrOrganizationAlreadyHasTeamForSeason",
	4026: "ErrGetOrganization",
	4027: "ErrGetSeasonChallenge",
	4028: "ErrCannotCreateTeamForSoloOrganization",
	4029: "ErrUserNotInOrganization",
	4030: "ErrCreateTeam",
	4031: "ErrGetTeam",
	4032: "ErrGetTeams",
	4033: "ErrGetUser",
	4034: "ErrUpdateUser",
	4035: "ErrUpdateTeam",
	4036: "ErrUpdateOrganization",
	4037: "ErrNewUserFromClaims",
	4038: "ErrGetOAuthUser",
	4039: "ErrDifferentUserBetweenTokenAndDatabase",
	4040: "ErrLoadUserSeasons",
	4041: "ErrGetUserOrganizations",
	4042: "ErrGetSeasons",
	4043: "ErrGetUserBySubject",
	4044: "ErrEmailAddressNotVerified",
	4045: "ErrGetDefaultSeason",
	4046: "ErrCommitUserTransaction",
	4047: "ErrUpdateActiveSeason",
	4048: "ErrMissingContextMetadata",
	4049: "ErrNoTokenProvided",
	4050: "ErrGetTokenWithClaims",
	4051: "ErrNoTokenInContext",
	4052: "ErrGetSubjectFromToken",
	4053: "ErrGetSubjectFromContext",
	4054: "ErrGetActiveSeasonMembership",
	4055: "ErrGetTokenFromContext",
	4056: "ErrChallengeAlreadyClosed",
	4057: "ErrGetAgent",
	4058: "ErrSaveAgent",
	4059: "ErrListChallengeInstances",
	5001: "ErrServerListen",
	5002: "ErrServerRegisterGateway",
	6001: "ErrInitLogger",
	6002: "ErrStartService",
	6003: "ErrInitServer",
	6004: "ErrGroupTerminated",
	6005: "ErrGetSSOClientFromFlags",
	6006: "ErrDumpDatabase",
	6007: "ErrGetDBInfo",
	6008: "ErrGetSSOWhoami",
	6009: "ErrGetSSOLogout",
	6010: "ErrGetSSOClaims",
	6011: "ErrInitDockerClient",
	6012: "ErrInitDB",
	6013: "ErrConfigureDB",
	6014: "ErrInitSSOClient",
	6015: "ErrInitService",
	7001: "ErrAgentGetContainersInfo",
	7002: "ErrCheckNginxContainer",
	7003: "ErrRemoveNginxContainer",
	7004: "ErrBuildNginxContainer",
	7005: "ErrStartNginxContainer",
	7006: "ErrParsingTemplate",
	7007: "ErrWriteConfigFileHeader",
	7008: "ErrWriteConfigFile",
	7009: "ErrCloseTarWriter",
	7010: "ErrCopyNginxConfigToContainer",
	7011: "ErrNginxNewConfigCheckFailed",
	7012: "ErrNginxSendCommandNewConfigCheck",
	7013: "ErrNginxSendCommandNewConfigRemove",
	7014: "ErrNginxSendCommandConfigReplace",
	7015: "ErrNginxSendCommandReloadConfig",
	7016: "ErrNginxConnectNetwork",
	7017: "ErrContainerConnectNetwork",
	7018: "ErrNatPortOpening",
	7019: "ErrBuildNginxConfig",
	7020: "ErrExecuteTemplate",
	7021: "ErrWriteBytesToHashBuilder",
	7022: "ErrReadBytesFromHashBuilder",
	7023: "ErrGeneratePrefixHash",
	7024: "ErrCleanPathwarInstances",
	7025: "ErrParseInitConfig",
	7026: "ErrUpPathwarInstance",
	7027: "ErrUpdateNginx",
	8001: "ErrDockerAPIContainerList",
	8002: "ErrDockerAPIContainerRemove",
	8003: "ErrDockerAPIImageRemove",
	8004: "ErrDockerAPIContainerCreate",
	8005: "ErrDockerAPIContainerExecCreate",
	8006: "ErrDockerAPIContainerExecAttach",
	8007: "ErrDockerAPIContainerExecStart",
	8008: "ErrDockerAPIContainerExecInspect",
	8009: "ErrDockerAPIImagePull",
	8010: "ErrDockerAPINetworkList",
	8011: "ErrDockerAPINetworkCreate",
	8012: "ErrDockerAPINetworkRemove",
	8013: "ErrDockerAPIExitCode",
	9001: "ErrExecuteOnInitHook",
	9002: "ErrRemoveInitConfig",
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
	"ErrComposeBuild":                         3008,
	"ErrComposeBundle":                        3009,
	"ErrComposeReadDab":                       3010,
	"ErrComposeParseDab":                      3011,
	"ErrComposeParseConfig":                   3012,
	"ErrComposeCreateTempDir":                 3013,
	"ErrComposeForceRecreateDown":             3014,
	"ErrComposeRunCreate":                     3015,
	"ErrComposeRunUp":                         3016,
	"ErrGetPWInitBinary":                      3017,
	"ErrWritePWInitFileHeader":                3018,
	"ErrWritePWInitFile":                      3019,
	"ErrMarshalPWInitConfigFile":              3020,
	"ErrWritePWInitConfigFileHeader":          3021,
	"ErrWritePWInitConfigFile":                3022,
	"ErrWritePWInitCloseTarWriter":            3023,
	"ErrCopyPWInitToContainer":                3024,
	"ErrComposeGetContainersInfo":             3025,
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
	"ErrDockerAPIContainerList":               8001,
	"ErrDockerAPIContainerRemove":             8002,
	"ErrDockerAPIImageRemove":                 8003,
	"ErrDockerAPIContainerCreate":             8004,
	"ErrDockerAPIContainerExecCreate":         8005,
	"ErrDockerAPIContainerExecAttach":         8006,
	"ErrDockerAPIContainerExecStart":          8007,
	"ErrDockerAPIContainerExecInspect":        8008,
	"ErrDockerAPIImagePull":                   8009,
	"ErrDockerAPINetworkList":                 8010,
	"ErrDockerAPINetworkCreate":               8011,
	"ErrDockerAPINetworkRemove":               8012,
	"ErrDockerAPIExitCode":                    8013,
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
	// 2091 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x58, 0xd9, 0x6f, 0x1c, 0xc7,
	0xd1, 0x97, 0x81, 0xef, 0x13, 0xa1, 0x49, 0x64, 0x96, 0x47, 0xb6, 0xd6, 0x96, 0x6c, 0x8e, 0xed,
	0x38, 0x16, 0x90, 0x63, 0xf5, 0x10, 0x60, 0x81, 0xbc, 0x10, 0xe0, 0x72, 0x49, 0x6a, 0x61, 0x69,
	0x49, 0xf0, 0xb0, 0x80, 0xbc, 0x35, 0x77, 0x6a, 0x67, 0x3b, 0x9c, 0xed, 0x5e, 0xf7, 0xf4, 0x90,
	0x62, 0xfe, 0x03, 0xe7, 0x21, 0xc8, 0x73, 0xde, 0x72, 0xc7, 0xce, 0x7d, 0xc7, 0xb7, 0x75, 0xdf,
	0xb6, 0xe4, 0xfb, 0xc8, 0x61, 0x4b, 0x49, 0x1c, 0xdf, 0x57, 0x0e, 0xe7, 0x0e, 0xba, 0xa7, 0x7a,
	0x76, 0x66, 0x49, 0xf9, 0x8d, 0xac, 0x5f, 0x55, 0x75, 0xf5, 0xaf, 0x8e, 0xae, 0x59, 0x6f, 0x3b,
	0x2a, 0xd5, 0x96, 0x21, 0x56, 0xfb, 0x4a, 0x6a, 0xe9, 0x8f, 0xf6, 0x99, 0xee, 0xae, 0x31, 0x55,
	0x25, 0xf1, 0xae, 0x4f, 0x47, 0x5c, 0x77, 0xd3, 0xe5, 0x6a, 0x5b, 0xf6, 0xf6, 0x46, 0x32, 0x92,
	0x7b, 0xad, 0xde, 0x72, 0xda, 0xb1, 0xff, 0xd9, 0x7f, 0xec, 0x5f, 0x99, 0xfd, 0x27, 0xbe, 0x74,
	0x9b, 0x37, 0x32, 0xa5, 0xd4, 0xa4, 0x0c, 0xd1, 0xdf, 0xee, 0x6d, 0x5b, 0x12, 0x21, 0x76, 0xb8,
	0xc0, 0x10, 0xb6, 0xf8, 0xdb, 0xbc, 0xff, 0x5b, 0x9c, 0x6d, 0xcc, 0xc2, 0x57, 0xfe, 0xdf, 0xdf,
	0xe9, 0x5d, 0x33, 0xa5, 0x54, 0x4b, 0xea, 0x66, 0xaf, 0x1f, 0x63, 0x0f, 0x85, 0xc6, 0x10, 0xee,
	0xde, 0xea, 0xfb, 0xde, 0xf6, 0x29, 0xa5, 0x1a, 0xd8, 0x57, 0xd8, 0x66, 0x46, 0xf6, 0xc1, 0x56,
	0x1f, 0xbc, 0x8f, 0x4c, 0x29, 0xd5, 0x14, 0x1a, 0x95, 0x60, 0x31, 0xbc, 0x32, 0xe2, 0xef, 0xf0,
	0x46, 0xad, 0x64, 0x95, 0xc5, 0x3c, 0x6c, 0x8a, 0x7e, 0xaa, 0x01, 0x49, 0x78, 0x80, 0x27, 0x09,
	0x17, 0x51, 0x26, 0xec, 0xf8, 0x3b, 0x3d, 0x7f, 0x4a, 0xa9, 0x25, 0xc1, 0x52, 0xdd, 0x45, 0xa1,
	0x79, 0xe6, 0x34, 0xa2, 0x73, 0x16, 0x16, 0x66, 0x67, 0x50, 0xcf, 0x36, 0x1b, 0x93, 0xf0, 0xea,
	0x88, 0xbf, 0xdb, 0xdb, 0x99, 0xc9, 0xc8, 0xf1, 0x5c, 0xba, 0x1c, 0xf3, 0xf6, 0x1d, 0xb8, 0x0e,
	0xaf, 0x8d, 0xf8, 0x37, 0x7b, 0xbb, 0x33, 0x70, 0x9a, 0xf1, 0x18, 0xc3, 0x3b, 0x70, 0xbd, 0x1d,
	0x4b, 0xb6, 0x32, 0x8f, 0x77, 0xa5, 0x98, 0x68, 0x78, 0x7d, 0xc4, 0xbf, 0xd5, 0xbb, 0xa9, 0x64,
	0x3e, 0x50, 0x49, 0xfa, 0x52, 0x24, 0x08, 0x6f, 0x8c, 0xf8, 0xd7, 0x78, 0x1f, 0xcd, 0x74, 0xf6,
	0xcb, 0x48, 0xa6, 0x1a, 0xde, 0x1c, 0xf1, 0x6f, 0xf2, 0xae, 0x77, 0x66, 0x5c, 0x3b, 0x9b, 0xc9,
	0x98, 0xa3, 0xd0, 0xf0, 0xd6, 0x88, 0x7f, 0xbd, 0xb7, 0xa3, 0xe4, 0xb5, 0x8e, 0x4c, 0xa1, 0x82,
	0xb7, 0x0b, 0x88, 0x33, 0x9a, 0x52, 0x4a, 0x2a, 0x78, 0x67, 0xc4, 0x91, 0x58, 0x6f, 0x49, 0x3d,
	0x2d, 0x53, 0x11, 0xc2, 0xc5, 0xd1, 0x5c, 0x96, 0xd3, 0xf8, 0xc4, 0xa8, 0x5f, 0xb1, 0xe4, 0x34,
	0xea, 0xf3, 0xa9, 0x38, 0xc0, 0x23, 0xc5, 0x34, 0x97, 0x22, 0x81, 0x27, 0x47, 0xfd, 0xab, 0xbd,
	0x6d, 0xa4, 0xcc, 0x35, 0x3c, 0x35, 0x4a, 0x61, 0x37, 0xea, 0x93, 0x52, 0x08, 0x6c, 0x6b, 0x78,
	0x7a, 0xd4, 0xbf, 0xce, 0x03, 0x2b, 0x9a, 0x48, 0xb5, 0xcc, 0x8c, 0x11, 0x9e, 0x19, 0xb8, 0x9c,
	0x08, 0xc3, 0x69, 0xa9, 0x90, 0x47, 0xc2, 0xf0, 0xf7, 0xec, 0xa8, 0xbf, 0xcb, 0xbb, 0xce, 0x56,
	0x45, 0xaf, 0x2f, 0x13, 0x74, 0x04, 0x33, 0xdd, 0x85, 0xfb, 0x2a, 0xc4, 0x2d, 0x61, 0x0d, 0xae,
	0xb0, 0xad, 0xa5, 0x5a, 0xcf, 0xa3, 0xbf, 0xbf, 0xe2, 0xdf, 0xe0, 0x5d, 0x3b, 0xd0, 0x98, 0x47,
	0x16, 0x4e, 0x4a, 0xd1, 0xe1, 0x11, 0x3c, 0x50, 0xf1, 0x6f, 0xf4, 0x2a, 0x1b, 0x1c, 0x13, 0xfa,
	0xe0, 0x10, 0x7a, 0x80, 0xa9, 0xa4, 0xcb, 0x62, 0x42, 0x1f, 0xaa, 0x10, 0xf7, 0x84, 0x4e, 0x2a,
	0x64, 0x1a, 0x17, 0xb1, 0xd7, 0x9f, 0xe6, 0x31, 0xc2, 0xc3, 0x43, 0xc6, 0x07, 0x15, 0x2f, 0xa0,
	0x8f, 0x54, 0xfc, 0x6b, 0x6d, 0xbd, 0x11, 0x5a, 0x4f, 0x79, 0x1c, 0xc2, 0xa3, 0x15, 0xe2, 0x25,
	0x97, 0x8a, 0x30, 0x46, 0x38, 0x5c, 0xa1, 0x7a, 0x2f, 0x5c, 0xa0, 0xc1, 0x96, 0xe1, 0x48, 0x85,
	0xf8, 0x22, 0xf9, 0x1c, 0x53, 0x09, 0x1a, 0xe0, 0x68, 0xa5, 0xcc, 0x97, 0x05, 0x28, 0xec, 0x63,
	0x43, 0x71, 0x0d, 0xc2, 0x6e, 0x70, 0x05, 0xc7, 0x87, 0xd8, 0x9c, 0x96, 0xaa, 0x8d, 0xf3, 0xd8,
	0xb6, 0x4a, 0x0d, 0xb9, 0x26, 0xe0, 0x44, 0x85, 0x2a, 0xc7, 0x05, 0x93, 0x8a, 0xcc, 0x05, 0x9c,
	0x1c, 0xba, 0xd3, 0x7c, 0x2a, 0x96, 0xfa, 0x70, 0xca, 0x05, 0x39, 0x83, 0x7a, 0xee, 0xa0, 0xa9,
	0x88, 0x3a, 0x17, 0x4c, 0xad, 0xc3, 0x69, 0xc7, 0x9f, 0x65, 0x26, 0x83, 0x0c, 0x37, 0xfb, 0x90,
	0x85, 0xa8, 0xe0, 0x8c, 0xb3, 0x1b, 0x82, 0xe1, 0x6c, 0xc5, 0x0f, 0xbc, 0x5d, 0xa6, 0x55, 0xb3,
	0x74, 0x64, 0x50, 0x76, 0x3b, 0xab, 0x70, 0xae, 0xe2, 0x7f, 0xcc, 0x1b, 0x2b, 0x5b, 0x0e, 0x60,
	0x72, 0x7f, 0x7e, 0x93, 0xd3, 0x0b, 0x3e, 0x1e, 0xab, 0xf8, 0xb7, 0x78, 0x37, 0x0e, 0xc1, 0xb1,
	0x4c, 0x70, 0x91, 0x65, 0x22, 0x05, 0x8f, 0x0f, 0xf2, 0xdf, 0x5f, 0xcf, 0x34, 0x16, 0xe5, 0xa4,
	0x14, 0x9a, 0x71, 0x81, 0x0a, 0x2e, 0x0c, 0x31, 0x39, 0x83, 0x3a, 0x07, 0x93, 0xa6, 0xe8, 0x48,
	0xb8, 0xe8, 0x32, 0x31, 0x83, 0x7a, 0x29, 0x41, 0xd5, 0x6c, 0x4c, 0x2b, 0xd9, 0x33, 0x4a, 0x78,
	0x48, 0xc3, 0x57, 0x03, 0x9a, 0x08, 0x34, 0x91, 0x26, 0xbb, 0x2c, 0x8e, 0x51, 0x44, 0x78, 0xa7,
	0xa9, 0x50, 0xdb, 0x6b, 0xf0, 0xb5, 0x80, 0xea, 0x85, 0xea, 0x76, 0x01, 0x59, 0x22, 0x05, 0x7c,
	0x3d, 0x20, 0xea, 0x16, 0x91, 0xf5, 0xcc, 0x8c, 0x14, 0x04, 0x7c, 0x23, 0xa0, 0x98, 0x4c, 0x30,
	0xce, 0xdf, 0x42, 0xba, 0x9c, 0xb4, 0x15, 0xef, 0x5b, 0x8f, 0xdf, 0x1c, 0x78, 0xe4, 0x7a, 0x41,
	0xc8, 0xb5, 0x4e, 0xcc, 0x56, 0x10, 0xbe, 0x15, 0x10, 0xa5, 0x4b, 0xfd, 0x90, 0x69, 0xdc, 0xdc,
	0xf6, 0xdb, 0x01, 0x71, 0x96, 0xd5, 0xc3, 0x66, 0x01, 0x7f, 0x27, 0xf0, 0xc7, 0xbc, 0x1b, 0x86,
	0x02, 0x28, 0xe0, 0xf7, 0x04, 0xfe, 0x0e, 0xef, 0xea, 0xc1, 0x85, 0xcc, 0x05, 0xe0, 0x5e, 0xc7,
	0x44, 0x6e, 0x31, 0x11, 0x2b, 0x64, 0xe1, 0x3a, 0x9d, 0xbe, 0x8c, 0x21, 0x7c, 0xd7, 0x05, 0x38,
	0x74, 0x76, 0x29, 0xc0, 0xef, 0x05, 0x34, 0x08, 0xa6, 0xb9, 0x08, 0x67, 0x55, 0xc4, 0x04, 0xff,
	0x02, 0x0d, 0xad, 0xef, 0x07, 0xfe, 0x6d, 0x5e, 0x90, 0x05, 0x96, 0x91, 0x65, 0x72, 0x91, 0xfd,
	0x95, 0x3b, 0x83, 0x1f, 0x04, 0x94, 0x72, 0xca, 0x98, 0x09, 0x6f, 0xa0, 0x07, 0x3f, 0x74, 0xbc,
	0x97, 0xd2, 0xd1, 0x6c, 0xc0, 0x8f, 0xdc, 0xb5, 0x8d, 0xd1, 0x3e, 0x96, 0xb4, 0xa4, 0xb5, 0x94,
	0x8a, 0x0c, 0x7f, 0x1c, 0xd0, 0xe3, 0x91, 0x9f, 0x9e, 0x9f, 0x99, 0xc0, 0x4f, 0x02, 0x9a, 0x9f,
	0x39, 0x08, 0x3f, 0x0d, 0xa8, 0xd3, 0xb2, 0xff, 0x1b, 0x28, 0x38, 0x86, 0xf0, 0xb3, 0x80, 0x1a,
	0x83, 0xe8, 0xd9, 0xc7, 0x92, 0xf2, 0x31, 0x3f, 0x77, 0x66, 0xf3, 0x98, 0xa0, 0x5a, 0xc5, 0xb0,
	0xc5, 0x7a, 0x08, 0xbf, 0xc8, 0xa9, 0xeb, 0x62, 0x7b, 0xa5, 0x48, 0xcb, 0x92, 0xe0, 0x77, 0xa5,
	0x68, 0x95, 0x7e, 0x19, 0xb8, 0x89, 0x62, 0xf9, 0x2d, 0x6a, 0xc1, 0xaf, 0x02, 0xff, 0x93, 0xde,
	0xed, 0x53, 0x4a, 0x15, 0xa5, 0x57, 0x8a, 0xe1, 0xbe, 0x60, 0x30, 0x0e, 0x4a, 0x5e, 0xee, 0x77,
	0x27, 0x6c, 0xe4, 0x00, 0x1e, 0x70, 0x27, 0x4c, 0x32, 0x21, 0xa4, 0x76, 0x23, 0x2b, 0xf3, 0x2b,
	0x63, 0x59, 0x72, 0xf4, 0xa0, 0x4b, 0x92, 0x21, 0xdb, 0x56, 0x7f, 0x09, 0x7e, 0x28, 0xa0, 0xb7,
	0x6c, 0xe0, 0x05, 0x1e, 0x0e, 0xfc, 0x51, 0xcf, 0xcb, 0xce, 0xb6, 0x82, 0x47, 0x02, 0xda, 0x1a,
	0x48, 0x90, 0xc0, 0xa3, 0x05, 0x15, 0xe3, 0x18, 0x0e, 0x3b, 0x3f, 0x59, 0x4b, 0x58, 0xd9, 0x91,
	0xb2, 0xcc, 0xba, 0x3a, 0xea, 0xee, 0x95, 0xc9, 0x4a, 0xb1, 0x1c, 0x73, 0x05, 0xd9, 0xc2, 0x35,
	0xe3, 0xc0, 0xf6, 0x7f, 0xcc, 0x78, 0x2f, 0x81, 0xe3, 0x2e, 0x57, 0x86, 0xa7, 0x89, 0x54, 0x77,
	0xed, 0x01, 0x27, 0x02, 0xff, 0x53, 0xde, 0x1e, 0xf3, 0x42, 0xf2, 0x4e, 0x07, 0x15, 0x0a, 0x1b,
	0x4b, 0x1d, 0xf5, 0x1a, 0xa2, 0x58, 0x94, 0x2b, 0x28, 0x26, 0x44, 0xd8, 0x60, 0x9a, 0x2d, 0xb3,
	0x04, 0xe1, 0xa4, 0xe3, 0x7a, 0xbf, 0x64, 0xa1, 0x51, 0xcc, 0x78, 0x4d, 0xe0, 0x54, 0x50, 0x9e,
	0x3c, 0xe5, 0x5e, 0x38, 0xed, 0x6e, 0x91, 0x67, 0x22, 0x81, 0x33, 0x01, 0x4d, 0x7d, 0xb2, 0xa8,
	0x9b, 0xe6, 0xfb, 0xbc, 0x79, 0xcb, 0xcf, 0xba, 0xaa, 0x9b, 0xea, 0x31, 0x1e, 0x4f, 0x84, 0xa1,
	0xc2, 0x24, 0x69, 0x49, 0x7d, 0x27, 0x2a, 0xde, 0x31, 0x65, 0x79, 0xae, 0x60, 0xda, 0xc0, 0x0e,
	0x4b, 0x63, 0x57, 0xc6, 0xe7, 0x83, 0xc1, 0x0b, 0xda, 0xe3, 0x59, 0x47, 0x29, 0x26, 0x12, 0xd6,
	0xb6, 0xec, 0x3c, 0x56, 0x66, 0x6e, 0xa2, 0xad, 0xf9, 0x2a, 0x92, 0xe9, 0xe3, 0xae, 0xa3, 0xdc,
	0x74, 0xcc, 0xa6, 0xe6, 0x01, 0xd4, 0x2c, 0x64, 0x9a, 0xc1, 0x05, 0x77, 0xf5, 0x96, 0xb4, 0xb4,
	0xcc, 0x29, 0xb9, 0xca, 0x43, 0x0c, 0xe1, 0x62, 0xa1, 0xcc, 0x2c, 0x72, 0x90, 0xeb, 0x2e, 0x71,
	0xfe, 0x84, 0x8b, 0x94, 0x8c, 0x9a, 0xc2, 0x0d, 0xe3, 0x27, 0x8b, 0x0d, 0x9a, 0x5d, 0xdc, 0xe4,
	0xca, 0x6a, 0xc1, 0x53, 0x85, 0xa9, 0x50, 0x00, 0x9d, 0xed, 0xd3, 0x6e, 0x2c, 0xce, 0xa0, 0x2e,
	0xde, 0xe1, 0x00, 0xf6, 0x96, 0x51, 0x25, 0x5d, 0xde, 0x87, 0x67, 0x0a, 0xee, 0xad, 0xcf, 0xa2,
	0xfd, 0xb3, 0xee, 0xaa, 0xc3, 0xe3, 0xcf, 0xbe, 0x47, 0x21, 0x3c, 0x57, 0xa8, 0xd5, 0x89, 0xc8,
	0xac, 0x7d, 0xcf, 0xbb, 0x89, 0xb1, 0xc0, 0x56, 0x31, 0x13, 0xbd, 0xe0, 0x9c, 0xec, 0xe7, 0xc9,
	0x60, 0xf2, 0x36, 0x45, 0xa2, 0x99, 0x68, 0x63, 0x02, 0x2f, 0x0e, 0x26, 0x8a, 0x5a, 0x45, 0xab,
	0x85, 0x02, 0xee, 0xde, 0xe3, 0xd6, 0x4b, 0x2b, 0x9d, 0xc7, 0xc8, 0xc8, 0xd5, 0x0c, 0xd3, 0xb8,
	0xc6, 0xd6, 0xe1, 0x8b, 0x7b, 0xa8, 0x50, 0xcc, 0x63, 0xb1, 0x5f, 0x46, 0x11, 0x2a, 0x78, 0xb7,
	0xea, 0x1c, 0x69, 0xa6, 0xb4, 0xb1, 0xe3, 0x6d, 0x84, 0xf7, 0xaa, 0x05, 0xcd, 0xcc, 0x19, 0xbc,
	0x5f, 0x75, 0x93, 0x40, 0xc9, 0xb4, 0xbf, 0x88, 0xaa, 0xc7, 0x85, 0xdd, 0xae, 0xff, 0x5c, 0x2d,
	0xf0, 0xb9, 0x30, 0x9b, 0xed, 0xb2, 0x86, 0x91, 0xe9, 0x98, 0x45, 0x09, 0xfc, 0xc5, 0x9d, 0xd0,
	0x48, 0x7b, 0xfd, 0xbc, 0xd6, 0xff, 0x5a, 0x1d, 0x4c, 0x49, 0xb3, 0x78, 0x76, 0x24, 0xfc, 0xad,
	0x3a, 0x68, 0xa1, 0x85, 0x85, 0xd9, 0x83, 0x5d, 0xc9, 0x7a, 0x1c, 0x3e, 0x28, 0x4b, 0x69, 0x91,
	0xfe, 0x7b, 0x59, 0x4a, 0x05, 0xf1, 0x8f, 0x2a, 0x15, 0x84, 0x09, 0xbb, 0x21, 0xdb, 0x2b, 0xa8,
	0x68, 0xb3, 0xfe, 0x67, 0x95, 0x96, 0x5c, 0x8b, 0xd4, 0xe1, 0x5f, 0x55, 0x7a, 0xb8, 0xb2, 0x1d,
	0x22, 0x55, 0xd8, 0xa8, 0xc3, 0xbf, 0xab, 0xc5, 0xc7, 0xd4, 0xdd, 0x04, 0xfe, 0x53, 0xcd, 0x1f,
	0x39, 0x9e, 0x33, 0xf4, 0xdf, 0x2a, 0x25, 0xc8, 0xe6, 0x6b, 0xe3, 0xb2, 0xf0, 0x7c, 0x8d, 0x4a,
	0xc4, 0x4e, 0xe9, 0x56, 0xc4, 0xc5, 0xa1, 0xc1, 0xae, 0xf1, 0x42, 0x8d, 0xfa, 0x79, 0x1e, 0x7b,
	0x72, 0x15, 0x87, 0xd0, 0x17, 0x9d, 0xa9, 0x5d, 0x32, 0x87, 0xc0, 0x5f, 0x3b, 0xd0, 0xe6, 0x6b,
	0x08, 0xfc, 0x4d, 0x8d, 0x52, 0x64, 0x16, 0x48, 0x2e, 0x22, 0xb3, 0x25, 0xc6, 0x66, 0xd5, 0xfb,
	0x6d, 0xad, 0xb8, 0x3d, 0x6d, 0x58, 0xae, 0x7e, 0x57, 0x2b, 0xee, 0x6e, 0x85, 0xb5, 0xea, 0xa5,
	0x9a, 0xdb, 0x64, 0xcb, 0xbb, 0xd4, 0xcb, 0x35, 0xf7, 0xc4, 0xcb, 0xfe, 0xba, 0x0b, 0xa2, 0xc3,
	0xa3, 0xe2, 0x42, 0x75, 0xa9, 0x46, 0x7d, 0x64, 0xf1, 0x16, 0xae, 0x65, 0x2a, 0x96, 0x8f, 0xec,
	0xa3, 0x0a, 0x2e, 0xd7, 0xfc, 0xdb, 0xbd, 0x5b, 0x9c, 0xca, 0x02, 0x8a, 0xd0, 0x4c, 0x16, 0x26,
	0xc2, 0xb2, 0x36, 0xfc, 0xbe, 0xe6, 0xef, 0xf1, 0x6e, 0xfd, 0x30, 0xbd, 0x8c, 0x48, 0xf8, 0x43,
	0xcd, 0xff, 0xb8, 0x77, 0xf3, 0x26, 0x8a, 0x4e, 0xab, 0x1f, 0xb3, 0x36, 0xc2, 0x1f, 0x6b, 0xb4,
	0x3d, 0x0c, 0xab, 0xcd, 0x63, 0x2c, 0xf3, 0x8f, 0x8d, 0x57, 0x1c, 0xd5, 0xee, 0x82, 0xe6, 0x5b,
	0xa8, 0x85, 0x7a, 0x4d, 0xaa, 0x15, 0xf8, 0x53, 0x8d, 0xc6, 0x68, 0x7e, 0xe1, 0x21, 0x85, 0x57,
	0x1d, 0x75, 0x2d, 0xa6, 0xe7, 0xa4, 0xd2, 0xb3, 0x7d, 0x14, 0x5c, 0x44, 0xf0, 0x5a, 0x8d, 0x6a,
	0xb4, 0x94, 0x5d, 0x73, 0xde, 0xeb, 0x2e, 0x0b, 0x53, 0x87, 0xb0, 0x9d, 0x66, 0x3b, 0xbe, 0xcd,
	0xde, 0x1b, 0xee, 0x2c, 0xcb, 0x7e, 0x7d, 0x5d, 0x63, 0xb2, 0x28, 0xf7, 0xb1, 0xa4, 0x6b, 0x5d,
	0xa0, 0x82, 0x37, 0x6b, 0xb4, 0x27, 0x9a, 0x2f, 0x0d, 0x8b, 0x9b, 0xf6, 0x2b, 0x6a, 0xbc, 0x55,
	0xcb, 0xc7, 0xa8, 0x40, 0xf3, 0xf1, 0x36, 0xa7, 0xb0, 0xc3, 0x0f, 0x19, 0x15, 0x78, 0xdb, 0x15,
	0xc7, 0x64, 0x8c, 0x4c, 0xcc, 0x65, 0x3f, 0x07, 0x0c, 0x46, 0xcd, 0x3b, 0xc5, 0xa2, 0xc2, 0xc1,
	0xde, 0x0d, 0xef, 0xd6, 0xe8, 0x35, 0x5c, 0xea, 0x0f, 0x19, 0xc1, 0x7b, 0x35, 0x6a, 0x99, 0xec,
	0x29, 0xb0, 0xb7, 0x84, 0xf7, 0x6b, 0xd4, 0x32, 0x59, 0x67, 0x4e, 0xcc, 0x35, 0x73, 0xee, 0xcc,
	0xfc, 0x82, 0xc3, 0xe3, 0x74, 0x8b, 0x8d, 0x38, 0xa5, 0xf7, 0xc8, 0x38, 0xf5, 0x4d, 0xae, 0xd1,
	0xec, 0xb1, 0x08, 0x09, 0x3d, 0x7a, 0x65, 0x7b, 0xfa, 0xe2, 0x39, 0x36, 0x4e, 0x79, 0xdf, 0xa8,
	0x61, 0x38, 0x27, 0xad, 0xe3, 0x1f, 0xae, 0x35, 0xa1, 0x35, 0x6b, 0x77, 0xe1, 0xc4, 0x38, 0xad,
	0x61, 0x9b, 0x6b, 0xd9, 0xf6, 0x84, 0x93, 0xe3, 0x54, 0x8f, 0x9b, 0x2b, 0x35, 0x45, 0xd2, 0x37,
	0x6f, 0xf2, 0xa9, 0x71, 0xca, 0x4e, 0xf9, 0x5e, 0x73, 0x69, 0x1c, 0xc3, 0xe9, 0x0d, 0x77, 0xa6,
	0x0a, 0xb3, 0x9c, 0x9d, 0x19, 0x1f, 0xe6, 0x94, 0x50, 0xba, 0xcb, 0xd9, 0x2b, 0xe1, 0xc4, 0xd9,
	0xb9, 0x71, 0xca, 0x61, 0x8e, 0x4f, 0x1d, 0x32, 0x09, 0x0e, 0x11, 0xce, 0x3b, 0x88, 0xca, 0x71,
	0x56, 0x98, 0xdc, 0xef, 0x93, 0x72, 0x05, 0xee, 0x99, 0xa6, 0x1a, 0xce, 0xbc, 0x14, 0x6a, 0xe2,
	0xde, 0xe9, 0xfa, 0x67, 0x2f, 0xbc, 0x3c, 0xb6, 0xe5, 0xd4, 0xa5, 0xb1, 0xab, 0x2e, 0x5c, 0x1a,
	0xbb, 0xea, 0xa5, 0x4b, 0x63, 0x57, 0x7d, 0xf9, 0xf2, 0xd8, 0x96, 0x0b, 0x97, 0xc7, 0xb6, 0x3c,
	0x77, 0x79, 0x6c, 0xcb, 0xe7, 0x76, 0xbb, 0x9f, 0x9b, 0x62, 0x26, 0xc2, 0xbd, 0x91, 0xdc, 0xdb,
	0x5f, 0x89, 0xf6, 0xd2, 0x4f, 0x4f, 0xcb, 0x5b, 0xed, 0x4f, 0x4a, 0x9f, 0xf9, 0x5f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd8, 0x8e, 0x77, 0x19, 0xa3, 0x12, 0x00, 0x00,
}
