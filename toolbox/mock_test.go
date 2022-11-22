package toolbox

import _ "github.com/golang/mock/mockgen/model"

//go:generate mockgen --build_flags=--mod=mod -destination=./mock/logger.go -package=mock -mock_names=Logger=Logger github.com/cloudtrust/keycloak-client/v2/toolbox Logger
//go:generate mockgen --build_flags=--mod=mod -destination=./mock/profile.go -package=mock -mock_names=ProfileRetriever=ProfileRetriever,OidcTokenProvider=OidcTokenProvider github.com/cloudtrust/keycloak-client/v2/toolbox ProfileRetriever,OidcTokenProvider

func ptr(value string) *string {
	return &value
}

func ptrMapStringStringPtr(value map[string]*string) *map[string]*string {
	return &value
}
