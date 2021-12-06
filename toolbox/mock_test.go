package toolbox

import _ "github.com/golang/mock/mockgen/model"

//go:generate mockgen --build_flags=--mod=mod -destination=./mock/logger.go -package=mock -mock_names=Logger=Logger github.com/cloudtrust/keycloak-client/v2/toolbox Logger
