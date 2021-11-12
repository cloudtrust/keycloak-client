package toolbox

import _ "github.com/golang/mock/mockgen/model"

//go:generate mockgen -destination=./mock/logger.go -package=mock -mock_names=Logger=Logger github.com/cloudtrust/keycloak-client/toolbox Logger
