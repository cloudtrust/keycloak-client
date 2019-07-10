module github.com/cloudtrust/keycloak-client

go 1.12

replace github.com/gbrlsnchs/jwt => github.com/gbrlsnchs/jwt/v2 v2.0.0

exclude github.com/gbrlsnchs/jwt/v2 v2.0.0-00010101000000-000000000000

require (
	github.com/coreos/go-oidc v2.0.0+incompatible
	github.com/gbrlsnchs/jwt v0.0.0-00010101000000-000000000000
	github.com/gbrlsnchs/jwt/v2 v2.0.0-alpha.0 // indirect
	github.com/nbio/st v0.0.0-20140626010706-e9e8d9816f32 // indirect
	github.com/pkg/errors v0.8.1
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
	google.golang.org/appengine v1.6.1 // indirect
	gopkg.in/h2non/gentleman.v2 v2.0.3
	gopkg.in/square/go-jose.v2 v2.3.1 // indirect
)
