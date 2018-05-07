package keycloak

type AdminEventRepresentation struct {
	AuthDetails    *AuthDetailsRepresentation `json:"authDetails,omitempty"`
	Error          *string                    `json:"error,omitempty"`
	OperationType  *string                    `json:"operationType,omitempty"`
	RealmId        *string                    `json:"realmId,omitempty"`
	Representation *string                    `json:"representation,omitempty"`
	ResourcePath   *string                    `json:"resourcePath,omitempty"`
	ResourceType   *string                    `json:"resourceType,omitempty"`
	Time           *int64                     `json:"time,omitempty"`
}

type AuthDetailsRepresentation struct {
	ClientId  *string `json:"clientId,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
	RealmId   *string `json:"realmId,omitempty"`
	UserId    *string `json:"userId,omitempty"`
}

type AuthenticationExecutionExportRepresentation struct {
	Authenticator       *string `json:"authenticator,omitempty"`
	AuthenticatorConfig *string `json:"authenticatorConfig,omitempty"`
	AuthenticatorFlow   *bool   `json:"authenticatorFlow,omitempty"`
	AutheticatorFlow    *bool   `json:"autheticatorFlow,omitempty"`
	FlowAlias           *string `json:"flowAlias,omitempty"`
	Priority            *int32  `json:"priority,omitempty"`
	Requirement         *string `json:"requirement,omitempty"`
	UserSetupAllowed    *bool   `json:"userSetupAllowed,omitempty"`
}

type AuthenticationExecutionInfoRepresentation struct {
	Alias                *string   `json:"alias,omitempty"`
	AuthenticationConfig *string   `json:"authenticationConfig,omitempty"`
	AuthenticationFlow   *bool     `json:"authenticationFlow,omitempty"`
	Configurable         *bool     `json:"configurable,omitempty"`
	DisplayName          *string   `json:"displayName,omitempty"`
	FlowId               *string   `json:"flowId,omitempty"`
	Id                   *string   `json:"id,omitempty"`
	Index                *int32    `json:"index,omitempty"`
	Level                *int32    `json:"level,omitempty"`
	ProviderId           *string   `json:"providerId,omitempty"`
	Requirement          *string   `json:"requirement,omitempty"`
	RequirementChoices   *[]string `json:"requirementChoices,omitempty"`
}

type AuthenticationExecutionRepresentation struct {
	Authenticator       *string `json:"authenticator,omitempty"`
	AuthenticatorConfig *string `json:"authenticatorConfig,omitempty"`
	AuthenticatorFlow   *bool   `json:"authenticatorFlow,omitempty"`
	AutheticatorFlow    *bool   `json:"autheticatorFlow,omitempty"`
	FlowId              *string `json:"flowId,omitempty"`
	Id                  *string `json:"id,omitempty"`
	ParentFlow          *string `json:"parentFlow,omitempty"`
	Priority            *int32  `json:"priority,omitempty"`
	Requirement         *string `json:"requirement,omitempty"`
}

type AuthenticationFlowRepresentation struct {
	Alias                    *string                                        `json:"alias,omitempty"`
	AuthenticationExecutions *[]AuthenticationExecutionExportRepresentation `json:"authenticationExecutions,omitempty"`
	BuiltIn                  *bool                                          `json:"builtIn,omitempty"`
	Description              *string                                        `json:"description,omitempty"`
	Id                       *string                                        `json:"id,omitempty"`
	ProviderId               *string                                        `json:"providerId,omitempty"`
	TopLevel                 *bool                                          `json:"topLevel,omitempty"`
}

type AuthenticatorConfigInfoRepresentation struct {
	HelpText   *string                         `json:"helpText,omitempty"`
	Name       *string                         `json:"name,omitempty"`
	Properties *[]ConfigPropertyRepresentation `json:"properties,omitempty"`
	ProviderId *string                         `json:"providerId,omitempty"`
}

type AuthenticatorConfigRepresentation struct {
	Alias  *string                 `json:"alias,omitempty"`
	Config *map[string]interface{} `json:"config,omitempty"`
	Id     *string                 `json:"id,omitempty"`
}

type CertificateRepresentation struct {
	Certificate *string `json:"certificate,omitempty"`
	Kid         *string `json:"kid,omitempty"`
	PrivateKey  *string `json:"privateKey,omitempty"`
	PublicKey   *string `json:"publicKey,omitempty"`
}

type ClientInitialAccessCreatePresentation struct {
	Count      *int32 `json:"count,omitempty"`
	Expiration *int32 `json:"expiration,omitempty"`
}

type ClientInitialAccessPresentation struct {
	Count          *int32  `json:"count,omitempty"`
	Expiration     *int32  `json:"expiration,omitempty"`
	Id             *string `json:"id,omitempty"`
	RemainingCount *int32  `json:"remainingCount,omitempty"`
	Timestamp      *int32  `json:"timestamp,omitempty"`
	Token          *string `json:"token,omitempty"`
}

type ClientMappingsRepresentation struct {
	Client   *string               `json:"client,omitempty"`
	Id       *string               `json:"id,omitempty"`
	Mappings *[]RoleRepresentation `json:"mappings,omitempty"`
}

type ClientRepresentation struct {
	Access                       *map[string]interface{}         `json:"access,omitempty"`
	AdminUrl                     *string                         `json:"adminUrl,omitempty"`
	Attributes                   *map[string]interface{}         `json:"attributes,omitempty"`
	AuthorizationServicesEnabled *bool                           `json:"authorizationServicesEnabled,omitempty"`
	AuthorizationSettings        *ResourceServerRepresentation   `json:"authorizationSettings,omitempty"`
	BaseUrl                      *string                         `json:"baseUrl,omitempty"`
	BearerOnly                   *bool                           `json:"bearerOnly,omitempty"`
	ClientAuthenticatorType      *string                         `json:"clientAuthenticatorType,omitempty"`
	ClientId                     *string                         `json:"clientId,omitempty"`
	ClientTemplate               *string                         `json:"clientTemplate,omitempty"`
	ConsentRequired              *bool                           `json:"consentRequired,omitempty"`
	DefaultRoles                 *[]string                       `json:"defaultRoles,omitempty"`
	Description                  *string                         `json:"description,omitempty"`
	DirectAccessGrantsEnabled    *bool                           `json:"directAccessGrantsEnabled,omitempty"`
	Enabled                      *bool                           `json:"enabled,omitempty"`
	FrontchannelLogout           *bool                           `json:"frontchannelLogout,omitempty"`
	FullScopeAllowed             *bool                           `json:"fullScopeAllowed,omitempty"`
	Id                           *string                         `json:"id,omitempty"`
	ImplicitFlowEnabled          *bool                           `json:"implicitFlowEnabled,omitempty"`
	Name                         *string                         `json:"name,omitempty"`
	NodeReRegistrationTimeout    *int32                          `json:"nodeReRegistrationTimeout,omitempty"`
	NotBefore                    *int32                          `json:"notBefore,omitempty"`
	Protocol                     *string                         `json:"protocol,omitempty"`
	ProtocolMappers              *[]ProtocolMapperRepresentation `json:"protocolMappers,omitempty"`
	PublicClient                 *bool                           `json:"publicClient,omitempty"`
	RedirectUris                 *[]string                       `json:"redirectUris,omitempty"`
	RegisteredNodes              *map[string]interface{}         `json:"registeredNodes,omitempty"`
	RegistrationAccessToken      *string                         `json:"registrationAccessToken,omitempty"`
	RootUrl                      *string                         `json:"rootUrl,omitempty"`
	Secret                       *string                         `json:"secret,omitempty"`
	ServiceAccountsEnabled       *bool                           `json:"serviceAccountsEnabled,omitempty"`
	StandardFlowEnabled          *bool                           `json:"standardFlowEnabled,omitempty"`
	SurrogateAuthRequired        *bool                           `json:"surrogateAuthRequired,omitempty"`
	UseTemplateConfig            *bool                           `json:"useTemplateConfig,omitempty"`
	UseTemplateMappers           *bool                           `json:"useTemplateMappers,omitempty"`
	UseTemplateScope             *bool                           `json:"useTemplateScope,omitempty"`
	WebOrigins                   *[]string                       `json:"webOrigins,omitempty"`
}

type ClientTemplateRepresentation struct {
	Attributes                *map[string]interface{}         `json:"attributes,omitempty"`
	BearerOnly                *bool                           `json:"bearerOnly,omitempty"`
	ConsentRequired           *bool                           `json:"consentRequired,omitempty"`
	Description               *string                         `json:"description,omitempty"`
	DirectAccessGrantsEnabled *bool                           `json:"directAccessGrantsEnabled,omitempty"`
	FrontchannelLogout        *bool                           `json:"frontchannelLogout,omitempty"`
	FullScopeAllowed          *bool                           `json:"fullScopeAllowed,omitempty"`
	Id                        *string                         `json:"id,omitempty"`
	ImplicitFlowEnabled       *bool                           `json:"implicitFlowEnabled,omitempty"`
	Name                      *string                         `json:"name,omitempty"`
	Protocol                  *string                         `json:"protocol,omitempty"`
	ProtocolMappers           *[]ProtocolMapperRepresentation `json:"protocolMappers,omitempty"`
	PublicClient              *bool                           `json:"publicClient,omitempty"`
	ServiceAccountsEnabled    *bool                           `json:"serviceAccountsEnabled,omitempty"`
	StandardFlowEnabled       *bool                           `json:"standardFlowEnabled,omitempty"`
}

type ComponentExportRepresentation struct {
	Config        *MultivaluedHashMap `json:"config,omitempty"`
	Id            *string             `json:"id,omitempty"`
	Name          *string             `json:"name,omitempty"`
	ProviderId    *string             `json:"providerId,omitempty"`
	SubComponents *MultivaluedHashMap `json:"subComponents,omitempty"`
	SubType       *string             `json:"subType,omitempty"`
}

type ComponentRepresentation struct {
	Config       *MultivaluedHashMap `json:"config,omitempty"`
	Id           *string             `json:"id,omitempty"`
	Name         *string             `json:"name,omitempty"`
	ParentId     *string             `json:"parentId,omitempty"`
	ProviderId   *string             `json:"providerId,omitempty"`
	ProviderType *string             `json:"providerType,omitempty"`
	SubType      *string             `json:"subType,omitempty"`
}

type ComponentTypeRepresentation struct {
	HelpText   *string                         `json:"helpText,omitempty"`
	Id         *string                         `json:"id,omitempty"`
	Metadata   *map[string]interface{}         `json:"metadata,omitempty"`
	Properties *[]ConfigPropertyRepresentation `json:"properties,omitempty"`
}

type ConfigPropertyRepresentation struct {
	DefaultValue *map[string]interface{} `json:"defaultValue,omitempty"`
	HelpText     *string                 `json:"helpText,omitempty"`
	Label        *string                 `json:"label,omitempty"`
	Name         *string                 `json:"name,omitempty"`
	Options      *[]string               `json:"options,omitempty"`
	Secret       *bool                   `json:"secret,omitempty"`
	Type         *string                 `json:"type,omitempty"`
}

type CredentialRepresentation struct {
	Algorithm         *string             `json:"algorithm,omitempty"`
	Config            *MultivaluedHashMap `json:"config,omitempty"`
	Counter           *int32              `json:"counter,omitempty"`
	CreatedDate       *int64              `json:"createdDate,omitempty"`
	Device            *string             `json:"device,omitempty"`
	Digits            *int32              `json:"digits,omitempty"`
	HashIterations    *int32              `json:"hashIterations,omitempty"`
	HashedSaltedValue *string             `json:"hashedSaltedValue,omitempty"`
	Period            *int32              `json:"period,omitempty"`
	Salt              *string             `json:"salt,omitempty"`
	Temporary         *bool               `json:"temporary,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Value             *string             `json:"value,omitempty"`
}

type EventRepresentation struct {
	ClientId  *string                 `json:"clientId,omitempty"`
	Details   *map[string]interface{} `json:"details,omitempty"`
	Error     *string                 `json:"error,omitempty"`
	IpAddress *string                 `json:"ipAddress,omitempty"`
	RealmId   *string                 `json:"realmId,omitempty"`
	SessionId *string                 `json:"sessionId,omitempty"`
	Time      *int64                  `json:"time,omitempty"`
	Type      *string                 `json:"type,omitempty"`
	UserId    *string                 `json:"userId,omitempty"`
}

type FederatedIdentityRepresentation struct {
	IdentityProvider *string `json:"identityProvider,omitempty"`
	UserId           *string `json:"userId,omitempty"`
	UserName         *string `json:"userName,omitempty"`
}

type GlobalRequestResult struct {
	FailedRequests  *[]string `json:"failedRequests,omitempty"`
	SuccessRequests *[]string `json:"successRequests,omitempty"`
}

type GroupRepresentation struct {
	Access      *map[string]interface{} `json:"access,omitempty"`
	Attributes  *map[string]interface{} `json:"attributes,omitempty"`
	ClientRoles *map[string]interface{} `json:"clientRoles,omitempty"`
	Id          *string                 `json:"id,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	Path        *string                 `json:"path,omitempty"`
	RealmRoles  *[]string               `json:"realmRoles,omitempty"`
	SubGroups   *[]GroupRepresentation  `json:"subGroups,omitempty"`
}

type IdentityProviderMapperRepresentation struct {
	Config                 *map[string]interface{} `json:"config,omitempty"`
	Id                     *string                 `json:"id,omitempty"`
	IdentityProviderAlias  *string                 `json:"identityProviderAlias,omitempty"`
	IdentityProviderMapper *string                 `json:"identityProviderMapper,omitempty"`
	Name                   *string                 `json:"name,omitempty"`
}

type IdentityProviderRepresentation struct {
	AddReadTokenRoleOnCreate  *bool                   `json:"addReadTokenRoleOnCreate,omitempty"`
	Alias                     *string                 `json:"alias,omitempty"`
	Config                    *map[string]interface{} `json:"config,omitempty"`
	DisplayName               *string                 `json:"displayName,omitempty"`
	Enabled                   *bool                   `json:"enabled,omitempty"`
	FirstBrokerLoginFlowAlias *string                 `json:"firstBrokerLoginFlowAlias,omitempty"`
	InternalId                *string                 `json:"internalId,omitempty"`
	LinkOnly                  *bool                   `json:"linkOnly,omitempty"`
	PostBrokerLoginFlowAlias  *string                 `json:"postBrokerLoginFlowAlias,omitempty"`
	ProviderId                *string                 `json:"providerId,omitempty"`
	StoreToken                *bool                   `json:"storeToken,omitempty"`
	TrustEmail                *bool                   `json:"trustEmail,omitempty"`
}

type KeysMetadataRepresentation struct {
	Active *map[string]interface{}                                `json:"active,omitempty"`
	Keys   *[]KeysMetadataRepresentationKeyMetadataRepresentation `json:"keys,omitempty"`
}

type KeysMetadataRepresentationKeyMetadataRepresentation struct {
	Certificate      *string `json:"certificate,omitempty"`
	Kid              *string `json:"kid,omitempty"`
	ProviderId       *string `json:"providerId,omitempty"`
	ProviderPriority *int64  `json:"providerPriority,omitempty"`
	PublicKey        *string `json:"publicKey,omitempty"`
	Status           *string `json:"status,omitempty"`
	Type             *string `json:"type,omitempty"`
}

type KeyStoreConfig struct {
	Format           *string `json:"format,omitempty"`
	KeyAlias         *string `json:"keyAlias,omitempty"`
	KeyPassword      *string `json:"keyPassword,omitempty"`
	RealmAlias       *string `json:"realmAlias,omitempty"`
	RealmCertificate *bool   `json:"realmCertificate,omitempty"`
	StorePassword    *string `json:"storePassword,omitempty"`
}

type ManagementPermissionReference struct {
	Enabled          *bool                   `json:"enabled,omitempty"`
	Resource         *string                 `json:"resource,omitempty"`
	ScopePermissions *map[string]interface{} `json:"scopePermissions,omitempty"`
}

type MappingsRepresentation struct {
	ClientMappings *map[string]interface{} `json:"clientMappings,omitempty"`
	RealmMappings  *[]RoleRepresentation   `json:"realmMappings,omitempty"`
}

type MemoryInfoRepresentation struct {
	Free           *int64  `json:"free,omitempty"`
	FreeFormated   *string `json:"freeFormated,omitempty"`
	FreePercentage *int64  `json:"freePercentage,omitempty"`
	Total          *int64  `json:"total,omitempty"`
	TotalFormated  *string `json:"totalFormated,omitempty"`
	Used           *int64  `json:"used,omitempty"`
	UsedFormated   *string `json:"usedFormated,omitempty"`
}

type MultivaluedHashMap struct {
	Empty      *bool  `json:"empty,omitempty"`
	LoadFactor *int32 `json:"loadFactor,omitempty"`
	Threshold  *int32 `json:"threshold,omitempty"`
}

type PartialImportRepresentation struct {
	Clients           *[]ClientRepresentation           `json:"clients,omitempty"`
	Groups            *[]GroupRepresentation            `json:"groups,omitempty"`
	IdentityProviders *[]IdentityProviderRepresentation `json:"identityProviders,omitempty"`
	IfResourceExists  *string                           `json:"ifResourceExists,omitempty"`
	Policy            *string                           `json:"policy,omitempty"`
	Roles             *RolesRepresentation              `json:"roles,omitempty"`
	Users             *[]UserRepresentation             `json:"users,omitempty"`
}

type PasswordPolicyTypeRepresentation struct {
	ConfigType        *string `json:"configType,omitempty"`
	DefaultValue      *string `json:"defaultValue,omitempty"`
	DisplayName       *string `json:"displayName,omitempty"`
	Id                *string `json:"id,omitempty"`
	MultipleSupported *bool   `json:"multipleSupported,omitempty"`
}

type PolicyRepresentation struct {
	Config           *map[string]interface{} `json:"config,omitempty"`
	DecisionStrategy *string                 `json:"decisionStrategy,omitempty"`
	Description      *string                 `json:"description,omitempty"`
	Id               *string                 `json:"id,omitempty"`
	Logic            *string                 `json:"logic,omitempty"`
	Name             *string                 `json:"name,omitempty"`
	Policies         *[]string               `json:"policies,omitempty"`
	Resources        *[]string               `json:"resources,omitempty"`
	Scopes           *[]string               `json:"scopes,omitempty"`
	Type             *string                 `json:"type,omitempty"`
}

type ProfileInfoRepresentation struct {
	DisabledFeatures *[]string `json:"disabledFeatures,omitempty"`
	Name             *string   `json:"name,omitempty"`
}

type ProtocolMapperRepresentation struct {
	Config          *map[string]interface{} `json:"config,omitempty"`
	ConsentRequired *bool                   `json:"consentRequired,omitempty"`
	ConsentText     *string                 `json:"consentText,omitempty"`
	Id              *string                 `json:"id,omitempty"`
	Name            *string                 `json:"name,omitempty"`
	Protocol        *string                 `json:"protocol,omitempty"`
	ProtocolMapper  *string                 `json:"protocolMapper,omitempty"`
}

type ProviderRepresentation struct {
	OperationalInfo *map[string]interface{} `json:"operationalInfo,omitempty"`
	Order           *int32                  `json:"order,omitempty"`
}

type RealmEventsConfigRepresentation struct {
	AdminEventsDetailsEnabled *bool     `json:"adminEventsDetailsEnabled,omitempty"`
	AdminEventsEnabled        *bool     `json:"adminEventsEnabled,omitempty"`
	EnabledEventTypes         *[]string `json:"enabledEventTypes,omitempty"`
	EventsEnabled             *bool     `json:"eventsEnabled,omitempty"`
	EventsExpiration          *int64    `json:"eventsExpiration,omitempty"`
	EventsListeners           *[]string `json:"eventsListeners,omitempty"`
}

type RealmRepresentation struct {
	AccessCodeLifespan                  *int32                                  `json:"accessCodeLifespan,omitempty"`
	AccessCodeLifespanLogin             *int32                                  `json:"accessCodeLifespanLogin,omitempty"`
	AccessCodeLifespanUserAction        *int32                                  `json:"accessCodeLifespanUserAction,omitempty"`
	AccessTokenLifespan                 *int32                                  `json:"accessTokenLifespan,omitempty"`
	AccessTokenLifespanForImplicitFlow  *int32                                  `json:"accessTokenLifespanForImplicitFlow,omitempty"`
	AccountTheme                        *string                                 `json:"accountTheme,omitempty"`
	ActionTokenGeneratedByAdminLifespan *int32                                  `json:"actionTokenGeneratedByAdminLifespan,omitempty"`
	ActionTokenGeneratedByUserLifespan  *int32                                  `json:"actionTokenGeneratedByUserLifespan,omitempty"`
	AdminEventsDetailsEnabled           *bool                                   `json:"adminEventsDetailsEnabled,omitempty"`
	AdminEventsEnabled                  *bool                                   `json:"adminEventsEnabled,omitempty"`
	AdminTheme                          *string                                 `json:"adminTheme,omitempty"`
	Attributes                          *map[string]interface{}                 `json:"attributes,omitempty"`
	AuthenticationFlows                 *[]AuthenticationFlowRepresentation     `json:"authenticationFlows,omitempty"`
	AuthenticatorConfig                 *[]AuthenticatorConfigRepresentation    `json:"authenticatorConfig,omitempty"`
	BrowserFlow                         *string                                 `json:"browserFlow,omitempty"`
	BrowserSecurityHeaders              *map[string]interface{}                 `json:"browserSecurityHeaders,omitempty"`
	BruteForceProtected                 *bool                                   `json:"bruteForceProtected,omitempty"`
	ClientAuthenticationFlow            *string                                 `json:"clientAuthenticationFlow,omitempty"`
	ClientScopeMappings                 *map[string]interface{}                 `json:"clientScopeMappings,omitempty"`
	ClientTemplates                     *[]ClientTemplateRepresentation         `json:"clientTemplates,omitempty"`
	Clients                             *[]ClientRepresentation                 `json:"clients,omitempty"`
	Components                          *MultivaluedHashMap                     `json:"components,omitempty"`
	DefaultGroups                       *[]string                               `json:"defaultGroups,omitempty"`
	DefaultLocale                       *string                                 `json:"defaultLocale,omitempty"`
	DefaultRoles                        *[]string                               `json:"defaultRoles,omitempty"`
	DirectGrantFlow                     *string                                 `json:"directGrantFlow,omitempty"`
	DisplayName                         *string                                 `json:"displayName,omitempty"`
	DisplayNameHtml                     *string                                 `json:"displayNameHtml,omitempty"`
	DockerAuthenticationFlow            *string                                 `json:"dockerAuthenticationFlow,omitempty"`
	DuplicateEmailsAllowed              *bool                                   `json:"duplicateEmailsAllowed,omitempty"`
	EditUsernameAllowed                 *bool                                   `json:"editUsernameAllowed,omitempty"`
	EmailTheme                          *string                                 `json:"emailTheme,omitempty"`
	Enabled                             *bool                                   `json:"enabled,omitempty"`
	EnabledEventTypes                   *[]string                               `json:"enabledEventTypes,omitempty"`
	EventsEnabled                       *bool                                   `json:"eventsEnabled,omitempty"`
	EventsExpiration                    *int64                                  `json:"eventsExpiration,omitempty"`
	EventsListeners                     *[]string                               `json:"eventsListeners,omitempty"`
	FailureFactor                       *int32                                  `json:"failureFactor,omitempty"`
	FederatedUsers                      *[]UserRepresentation                   `json:"federatedUsers,omitempty"`
	Groups                              *[]GroupRepresentation                  `json:"groups,omitempty"`
	Id                                  *string                                 `json:"id,omitempty"`
	IdentityProviderMappers             *[]IdentityProviderMapperRepresentation `json:"identityProviderMappers,omitempty"`
	IdentityProviders                   *[]IdentityProviderRepresentation       `json:"identityProviders,omitempty"`
	InternationalizationEnabled         *bool                                   `json:"internationalizationEnabled,omitempty"`
	KeycloakVersion                     *string                                 `json:"keycloakVersion,omitempty"`
	LoginTheme                          *string                                 `json:"loginTheme,omitempty"`
	LoginWithEmailAllowed               *bool                                   `json:"loginWithEmailAllowed,omitempty"`
	MaxDeltaTimeSeconds                 *int32                                  `json:"maxDeltaTimeSeconds,omitempty"`
	MaxFailureWaitSeconds               *int32                                  `json:"maxFailureWaitSeconds,omitempty"`
	MinimumQuickLoginWaitSeconds        *int32                                  `json:"minimumQuickLoginWaitSeconds,omitempty"`
	NotBefore                           *int32                                  `json:"notBefore,omitempty"`
	OfflineSessionIdleTimeout           *int32                                  `json:"offlineSessionIdleTimeout,omitempty"`
	OtpPolicyAlgorithm                  *string                                 `json:"otpPolicyAlgorithm,omitempty"`
	OtpPolicyDigits                     *int32                                  `json:"otpPolicyDigits,omitempty"`
	OtpPolicyInitialCounter             *int32                                  `json:"otpPolicyInitialCounter,omitempty"`
	OtpPolicyLookAheadWindow            *int32                                  `json:"otpPolicyLookAheadWindow,omitempty"`
	OtpPolicyPeriod                     *int32                                  `json:"otpPolicyPeriod,omitempty"`
	OtpPolicyType                       *string                                 `json:"otpPolicyType,omitempty"`
	OtpSupportedApplications            *[]string                               `json:"otpSupportedApplications,omitempty"`
	PasswordPolicy                      *string                                 `json:"passwordPolicy,omitempty"`
	PermanentLockout                    *bool                                   `json:"permanentLockout,omitempty"`
	ProtocolMappers                     *[]ProtocolMapperRepresentation         `json:"protocolMappers,omitempty"`
	QuickLoginCheckMilliSeconds         *int64                                  `json:"quickLoginCheckMilliSeconds,omitempty"`
	Realm                               *string                                 `json:"realm,omitempty"`
	RefreshTokenMaxReuse                *int32                                  `json:"refreshTokenMaxReuse,omitempty"`
	RegistrationAllowed                 *bool                                   `json:"registrationAllowed,omitempty"`
	RegistrationEmailAsUsername         *bool                                   `json:"registrationEmailAsUsername,omitempty"`
	RegistrationFlow                    *string                                 `json:"registrationFlow,omitempty"`
	RememberMe                          *bool                                   `json:"rememberMe,omitempty"`
	RequiredActions                     *[]RequiredActionProviderRepresentation `json:"requiredActions,omitempty"`
	ResetCredentialsFlow                *string                                 `json:"resetCredentialsFlow,omitempty"`
	ResetPasswordAllowed                *bool                                   `json:"resetPasswordAllowed,omitempty"`
	RevokeRefreshToken                  *bool                                   `json:"revokeRefreshToken,omitempty"`
	Roles                               *RolesRepresentation                    `json:"roles,omitempty"`
	ScopeMappings                       *[]ScopeMappingRepresentation           `json:"scopeMappings,omitempty"`
	SmtpServer                          *map[string]interface{}                 `json:"smtpServer,omitempty"`
	SslRequired                         *string                                 `json:"sslRequired,omitempty"`
	SsoSessionIdleTimeout               *int32                                  `json:"ssoSessionIdleTimeout,omitempty"`
	SsoSessionMaxLifespan               *int32                                  `json:"ssoSessionMaxLifespan,omitempty"`
	SupportedLocales                    *[]string                               `json:"supportedLocales,omitempty"`
	UserFederationMappers               *[]UserFederationMapperRepresentation   `json:"userFederationMappers,omitempty"`
	UserFederationProviders             *[]UserFederationProviderRepresentation `json:"userFederationProviders,omitempty"`
	Users                               *[]UserRepresentation                   `json:"users,omitempty"`
	VerifyEmail                         *bool                                   `json:"verifyEmail,omitempty"`
	WaitIncrementSeconds                *int32                                  `json:"waitIncrementSeconds,omitempty"`
}

type RequiredActionProviderRepresentation struct {
	Alias         *string                 `json:"alias,omitempty"`
	Config        *map[string]interface{} `json:"config,omitempty"`
	DefaultAction *bool                   `json:"defaultAction,omitempty"`
	Enabled       *bool                   `json:"enabled,omitempty"`
	Name          *string                 `json:"name,omitempty"`
	ProviderId    *string                 `json:"providerId,omitempty"`
}

type ResourceOwnerRepresentation struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ResourceRepresentation struct {
	Id          *string                      `json:"id,omitempty"`
	Icon_uri    *string                      `json:"icon_uri,omitempty"`
	Name        *string                      `json:"name,omitempty"`
	Owner       *ResourceOwnerRepresentation `json:"owner,omitempty"`
	Policies    *[]PolicyRepresentation      `json:"policies,omitempty"`
	Scopes      *[]ScopeRepresentation       `json:"scopes,omitempty"`
	Type        *string                      `json:"type,omitempty"`
	TypedScopes *[]ScopeRepresentation       `json:"typedScopes,omitempty"`
	Uri         *string                      `json:"uri,omitempty"`
}

type ResourceServerRepresentation struct {
	AllowRemoteResourceManagement *bool                     `json:"allowRemoteResourceManagement,omitempty"`
	ClientId                      *string                   `json:"clientId,omitempty"`
	Id                            *string                   `json:"id,omitempty"`
	Name                          *string                   `json:"name,omitempty"`
	Policies                      *[]PolicyRepresentation   `json:"policies,omitempty"`
	PolicyEnforcementMode         *string                   `json:"policyEnforcementMode,omitempty"`
	Resources                     *[]ResourceRepresentation `json:"resources,omitempty"`
	Scopes                        *[]ScopeRepresentation    `json:"scopes,omitempty"`
}

type RoleRepresentation struct {
	ClientRole         *bool                         `json:"clientRole,omitempty"`
	Composite          *bool                         `json:"composite,omitempty"`
	Composites         *RoleRepresentationComposites `json:"composites,omitempty"`
	ContainerId        *string                       `json:"containerId,omitempty"`
	Description        *string                       `json:"description,omitempty"`
	Id                 *string                       `json:"id,omitempty"`
	Name               *string                       `json:"name,omitempty"`
	ScopeParamRequired *bool                         `json:"scopeParamRequired,omitempty"`
}

type RoleRepresentationComposites struct {
	Client *map[string]interface{} `json:"client,omitempty"`
	Realm  *[]string               `json:"realm,omitempty"`
}

type RolesRepresentation struct {
	Client *map[string]interface{} `json:"client,omitempty"`
	Realm  *[]RoleRepresentation   `json:"realm,omitempty"`
}

type ScopeMappingRepresentation struct {
	Client         *string   `json:"client,omitempty"`
	ClientTemplate *string   `json:"clientTemplate,omitempty"`
	Roles          *[]string `json:"roles,omitempty"`
	Self           *string   `json:"self,omitempty"`
}

type ScopeRepresentation struct {
	IconUri   *string                   `json:"iconUri,omitempty"`
	Id        *string                   `json:"id,omitempty"`
	Name      *string                   `json:"name,omitempty"`
	Policies  *[]PolicyRepresentation   `json:"policies,omitempty"`
	Resources *[]ResourceRepresentation `json:"resources,omitempty"`
}

type ServerInfoRepresentation struct {
	BuiltinProtocolMappers *map[string]interface{}             `json:"builtinProtocolMappers,omitempty"`
	ClientImporters        *[]map[string]interface{}           `json:"clientImporters,omitempty"`
	ClientInstallations    *map[string]interface{}             `json:"clientInstallations,omitempty"`
	ComponentTypes         *map[string]interface{}             `json:"componentTypes,omitempty"`
	Enums                  *map[string]interface{}             `json:"enums,omitempty"`
	IdentityProviders      *[]map[string]interface{}           `json:"identityProviders,omitempty"`
	MemoryInfo             *MemoryInfoRepresentation           `json:"memoryInfo,omitempty"`
	PasswordPolicies       *[]PasswordPolicyTypeRepresentation `json:"passwordPolicies,omitempty"`
	ProfileInfo            *ProfileInfoRepresentation          `json:"profileInfo,omitempty"`
	ProtocolMapperTypes    *map[string]interface{}             `json:"protocolMapperTypes,omitempty"`
	Providers              *map[string]interface{}             `json:"providers,omitempty"`
	SocialProviders        *[]map[string]interface{}           `json:"socialProviders,omitempty"`
	SystemInfo             *SystemInfoRepresentation           `json:"systemInfo,omitempty"`
	Themes                 *map[string]interface{}             `json:"themes,omitempty"`
}

type SpiInfoRepresentation struct {
	Internal  *bool                   `json:"internal,omitempty"`
	Providers *map[string]interface{} `json:"providers,omitempty"`
}

type SynchronizationResult struct {
	Added   *int32  `json:"added,omitempty"`
	Failed  *int32  `json:"failed,omitempty"`
	Ignored *bool   `json:"ignored,omitempty"`
	Removed *int32  `json:"removed,omitempty"`
	Status  *string `json:"status,omitempty"`
	Updated *int32  `json:"updated,omitempty"`
}

type SystemInfoRepresentation struct {
	FileEncoding   *string `json:"fileEncoding,omitempty"`
	JavaHome       *string `json:"javaHome,omitempty"`
	JavaRuntime    *string `json:"javaRuntime,omitempty"`
	JavaVendor     *string `json:"javaVendor,omitempty"`
	JavaVersion    *string `json:"javaVersion,omitempty"`
	JavaVm         *string `json:"javaVm,omitempty"`
	JavaVmVersion  *string `json:"javaVmVersion,omitempty"`
	OsArchitecture *string `json:"osArchitecture,omitempty"`
	OsName         *string `json:"osName,omitempty"`
	OsVersion      *string `json:"osVersion,omitempty"`
	ServerTime     *string `json:"serverTime,omitempty"`
	Uptime         *string `json:"uptime,omitempty"`
	UptimeMillis   *int64  `json:"uptimeMillis,omitempty"`
	UserDir        *string `json:"userDir,omitempty"`
	UserLocale     *string `json:"userLocale,omitempty"`
	UserName       *string `json:"userName,omitempty"`
	UserTimezone   *string `json:"userTimezone,omitempty"`
	Version        *string `json:"version,omitempty"`
}

type UserConsentRepresentation struct {
	ClientId               *string                 `json:"clientId,omitempty"`
	CreatedDate            *int64                  `json:"createdDate,omitempty"`
	GrantedClientRoles     *map[string]interface{} `json:"grantedClientRoles,omitempty"`
	GrantedProtocolMappers *map[string]interface{} `json:"grantedProtocolMappers,omitempty"`
	GrantedRealmRoles      *[]string               `json:"grantedRealmRoles,omitempty"`
	LastUpdatedDate        *int64                  `json:"lastUpdatedDate,omitempty"`
}

type UserFederationMapperRepresentation struct {
	Config                        *map[string]interface{} `json:"config,omitempty"`
	FederationMapperType          *string                 `json:"federationMapperType,omitempty"`
	FederationProviderDisplayName *string                 `json:"federationProviderDisplayName,omitempty"`
	Id                            *string                 `json:"id,omitempty"`
	Name                          *string                 `json:"name,omitempty"`
}

type UserFederationProviderRepresentation struct {
	ChangedSyncPeriod *int32                  `json:"changedSyncPeriod,omitempty"`
	Config            *map[string]interface{} `json:"config,omitempty"`
	DisplayName       *string                 `json:"displayName,omitempty"`
	FullSyncPeriod    *int32                  `json:"fullSyncPeriod,omitempty"`
	Id                *string                 `json:"id,omitempty"`
	LastSync          *int32                  `json:"lastSync,omitempty"`
	Priority          *int32                  `json:"priority,omitempty"`
	ProviderName      *string                 `json:"providerName,omitempty"`
}

type UserRepresentation struct {
	Access                     *map[string]interface{}            `json:"access,omitempty"`
	Attributes                 *map[string]interface{}            `json:"attributes,omitempty"`
	ClientConsents             *[]UserConsentRepresentation       `json:"clientConsents,omitempty"`
	ClientRoles                *map[string]interface{}            `json:"clientRoles,omitempty"`
	CreatedTimestamp           *int64                             `json:"createdTimestamp,omitempty"`
	Credentials                *[]CredentialRepresentation        `json:"credentials,omitempty"`
	DisableableCredentialTypes *[]string                          `json:"disableableCredentialTypes,omitempty"`
	Email                      *string                            `json:"email,omitempty"`
	EmailVerified              *bool                              `json:"emailVerified,omitempty"`
	Enabled                    *bool                              `json:"enabled,omitempty"`
	FederatedIdentities        *[]FederatedIdentityRepresentation `json:"federatedIdentities,omitempty"`
	FederationLink             *string                            `json:"federationLink,omitempty"`
	FirstName                  *string                            `json:"firstName,omitempty"`
	Groups                     *[]string                          `json:"groups,omitempty"`
	Id                         *string                            `json:"id,omitempty"`
	LastName                   *string                            `json:"lastName,omitempty"`
	NotBefore                  *int32                             `json:"notBefore,omitempty"`
	Origin                     *string                            `json:"origin,omitempty"`
	RealmRoles                 *[]string                          `json:"realmRoles,omitempty"`
	RequiredActions            *[]string                          `json:"requiredActions,omitempty"`
	Self                       *string                            `json:"self,omitempty"`
	ServiceAccountClientId     *string                            `json:"serviceAccountClientId,omitempty"`
	Username                   *string                            `json:"username,omitempty"`
}

type UserSessionRepresentation struct {
	Clients    *map[string]interface{} `json:"clients,omitempty"`
	Id         *string                 `json:"id,omitempty"`
	IpAddress  *string                 `json:"ipAddress,omitempty"`
	LastAccess *int64                  `json:"lastAccess,omitempty"`
	Start      *int64                  `json:"start,omitempty"`
	UserId     *string                 `json:"userId,omitempty"`
	Username   *string                 `json:"username,omitempty"`
}
