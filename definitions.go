package keycloak

// AdminEventRepresentation struct
type AdminEventRepresentation struct {
	AuthDetails    *AuthDetailsRepresentation `json:"authDetails,omitempty"`
	Error          *string                    `json:"error,omitempty"`
	OperationType  *string                    `json:"operationType,omitempty"`
	RealmID        *string                    `json:"realmId,omitempty"`
	Representation *string                    `json:"representation,omitempty"`
	ResourcePath   *string                    `json:"resourcePath,omitempty"`
	ResourceType   *string                    `json:"resourceType,omitempty"`
	Time           *int64                     `json:"time,omitempty"`
}

// AuthDetailsRepresentation struct
type AuthDetailsRepresentation struct {
	ClientID  *string `json:"clientId,omitempty"`
	IPAddress *string `json:"ipAddress,omitempty"`
	RealmID   *string `json:"realmId,omitempty"`
	UserID    *string `json:"userId,omitempty"`
}

// AuthenticationExecutionExportRepresentation struct
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

// AuthenticationExecutionInfoRepresentation struct
type AuthenticationExecutionInfoRepresentation struct {
	Alias                *string   `json:"alias,omitempty"`
	AuthenticationConfig *string   `json:"authenticationConfig,omitempty"`
	AuthenticationFlow   *bool     `json:"authenticationFlow,omitempty"`
	Configurable         *bool     `json:"configurable,omitempty"`
	DisplayName          *string   `json:"displayName,omitempty"`
	FlowID               *string   `json:"flowId,omitempty"`
	ID                   *string   `json:"id,omitempty"`
	Index                *int32    `json:"index,omitempty"`
	Level                *int32    `json:"level,omitempty"`
	ProviderID           *string   `json:"providerId,omitempty"`
	Requirement          *string   `json:"requirement,omitempty"`
	RequirementChoices   *[]string `json:"requirementChoices,omitempty"`
}

// AuthenticationExecutionRepresentation struct
type AuthenticationExecutionRepresentation struct {
	Authenticator       *string `json:"authenticator,omitempty"`
	AuthenticatorConfig *string `json:"authenticatorConfig,omitempty"`
	AuthenticatorFlow   *bool   `json:"authenticatorFlow,omitempty"`
	AutheticatorFlow    *bool   `json:"autheticatorFlow,omitempty"`
	FlowID              *string `json:"flowId,omitempty"`
	ID                  *string `json:"id,omitempty"`
	ParentFlow          *string `json:"parentFlow,omitempty"`
	Priority            *int32  `json:"priority,omitempty"`
	Requirement         *string `json:"requirement,omitempty"`
}

// AuthenticationFlowRepresentation struct
type AuthenticationFlowRepresentation struct {
	Alias                    *string                                        `json:"alias,omitempty"`
	AuthenticationExecutions *[]AuthenticationExecutionExportRepresentation `json:"authenticationExecutions,omitempty"`
	BuiltIn                  *bool                                          `json:"builtIn,omitempty"`
	Description              *string                                        `json:"description,omitempty"`
	ID                       *string                                        `json:"id,omitempty"`
	ProviderID               *string                                        `json:"providerId,omitempty"`
	TopLevel                 *bool                                          `json:"topLevel,omitempty"`
}

// AuthenticatorConfigInfoRepresentation struct
type AuthenticatorConfigInfoRepresentation struct {
	HelpText   *string                         `json:"helpText,omitempty"`
	Name       *string                         `json:"name,omitempty"`
	Properties *[]ConfigPropertyRepresentation `json:"properties,omitempty"`
	ProviderID *string                         `json:"providerId,omitempty"`
}

// AuthenticatorConfigRepresentation struct
type AuthenticatorConfigRepresentation struct {
	Alias  *string         `json:"alias,omitempty"`
	Config *map[string]any `json:"config,omitempty"`
	ID     *string         `json:"id,omitempty"`
}

// CertificateRepresentation struct
type CertificateRepresentation struct {
	Certificate *string `json:"certificate,omitempty"`
	Kid         *string `json:"kid,omitempty"`
	PrivateKey  *string `json:"privateKey,omitempty"`
	PublicKey   *string `json:"publicKey,omitempty"`
}

// ClientInitialAccessCreatePresentation struct
type ClientInitialAccessCreatePresentation struct {
	Count      *int32 `json:"count,omitempty"`
	Expiration *int32 `json:"expiration,omitempty"`
}

// ClientInitialAccessPresentation struct
type ClientInitialAccessPresentation struct {
	Count          *int32  `json:"count,omitempty"`
	Expiration     *int32  `json:"expiration,omitempty"`
	ID             *string `json:"id,omitempty"`
	RemainingCount *int32  `json:"remainingCount,omitempty"`
	Timestamp      *int32  `json:"timestamp,omitempty"`
	Token          *string `json:"token,omitempty"`
}

// ClientMapperRepresentation struct
// https://www.keycloak.org/docs-api/9.0/rest-api/index.html#_clientscopeevaluateresource-protocolmapperevaluationrepresentation
type ClientMapperRepresentation struct {
	ContainerID    *string `json:"containerId,omitempty"`
	ContainerName  *string `json:"containerName,omitempty"`
	ContainerType  *string `json:"containerType,omitempty"`
	MapperID       *string `json:"mapperId,omitempty"`
	MapperName     *string `json:"mapperName,omitempty"`
	ProtocolMapper *string `json:"protocolMapper,omitempty"`
}

// ClientMappingsRepresentation struct
type ClientMappingsRepresentation struct {
	Client   *string               `json:"client,omitempty"`
	ID       *string               `json:"id,omitempty"`
	Mappings *[]RoleRepresentation `json:"mappings,omitempty"`
}

// ClientRepresentation struct
type ClientRepresentation struct {
	Access                       *map[string]any                 `json:"access,omitempty"`
	AdminURL                     *string                         `json:"adminUrl,omitempty"`
	Attributes                   *map[string]any                 `json:"attributes,omitempty"`
	AuthorizationServicesEnabled *bool                           `json:"authorizationServicesEnabled,omitempty"`
	AuthorizationSettings        *ResourceServerRepresentation   `json:"authorizationSettings,omitempty"`
	BaseURL                      *string                         `json:"baseUrl,omitempty"`
	BearerOnly                   *bool                           `json:"bearerOnly,omitempty"`
	ClientAuthenticatorType      *string                         `json:"clientAuthenticatorType,omitempty"`
	ClientID                     *string                         `json:"clientId,omitempty"`
	ClientTemplate               *string                         `json:"clientTemplate,omitempty"`
	ConsentRequired              *bool                           `json:"consentRequired,omitempty"`
	DefaultRoles                 *[]string                       `json:"defaultRoles,omitempty"`
	Description                  *string                         `json:"description,omitempty"`
	DirectAccessGrantsEnabled    *bool                           `json:"directAccessGrantsEnabled,omitempty"`
	Enabled                      *bool                           `json:"enabled,omitempty"`
	FrontchannelLogout           *bool                           `json:"frontchannelLogout,omitempty"`
	FullScopeAllowed             *bool                           `json:"fullScopeAllowed,omitempty"`
	ID                           *string                         `json:"id,omitempty"`
	ImplicitFlowEnabled          *bool                           `json:"implicitFlowEnabled,omitempty"`
	Name                         *string                         `json:"name,omitempty"`
	NodeReRegistrationTimeout    *int32                          `json:"nodeReRegistrationTimeout,omitempty"`
	NotBefore                    *int32                          `json:"notBefore,omitempty"`
	Protocol                     *string                         `json:"protocol,omitempty"`
	ProtocolMappers              *[]ProtocolMapperRepresentation `json:"protocolMappers,omitempty"`
	PublicClient                 *bool                           `json:"publicClient,omitempty"`
	RedirectUris                 *[]string                       `json:"redirectUris,omitempty"`
	RegisteredNodes              *map[string]any                 `json:"registeredNodes,omitempty"`
	RegistrationAccessToken      *string                         `json:"registrationAccessToken,omitempty"`
	RootURL                      *string                         `json:"rootUrl,omitempty"`
	Secret                       *string                         `json:"secret,omitempty"`
	ServiceAccountsEnabled       *bool                           `json:"serviceAccountsEnabled,omitempty"`
	StandardFlowEnabled          *bool                           `json:"standardFlowEnabled,omitempty"`
	SurrogateAuthRequired        *bool                           `json:"surrogateAuthRequired,omitempty"`
	UseTemplateConfig            *bool                           `json:"useTemplateConfig,omitempty"`
	UseTemplateMappers           *bool                           `json:"useTemplateMappers,omitempty"`
	UseTemplateScope             *bool                           `json:"useTemplateScope,omitempty"`
	WebOrigins                   *[]string                       `json:"webOrigins,omitempty"`
}

// ClientTemplateRepresentation struct
type ClientTemplateRepresentation struct {
	Attributes                *map[string]any                 `json:"attributes,omitempty"`
	BearerOnly                *bool                           `json:"bearerOnly,omitempty"`
	ConsentRequired           *bool                           `json:"consentRequired,omitempty"`
	Description               *string                         `json:"description,omitempty"`
	DirectAccessGrantsEnabled *bool                           `json:"directAccessGrantsEnabled,omitempty"`
	FrontchannelLogout        *bool                           `json:"frontchannelLogout,omitempty"`
	FullScopeAllowed          *bool                           `json:"fullScopeAllowed,omitempty"`
	ID                        *string                         `json:"id,omitempty"`
	ImplicitFlowEnabled       *bool                           `json:"implicitFlowEnabled,omitempty"`
	Name                      *string                         `json:"name,omitempty"`
	Protocol                  *string                         `json:"protocol,omitempty"`
	ProtocolMappers           *[]ProtocolMapperRepresentation `json:"protocolMappers,omitempty"`
	PublicClient              *bool                           `json:"publicClient,omitempty"`
	ServiceAccountsEnabled    *bool                           `json:"serviceAccountsEnabled,omitempty"`
	StandardFlowEnabled       *bool                           `json:"standardFlowEnabled,omitempty"`
}

// ComponentExportRepresentation struct
type ComponentExportRepresentation struct {
	Config        *MultivaluedHashMap `json:"config,omitempty"`
	ID            *string             `json:"id,omitempty"`
	Name          *string             `json:"name,omitempty"`
	ProviderID    *string             `json:"providerId,omitempty"`
	SubComponents *MultivaluedHashMap `json:"subComponents,omitempty"`
	SubType       *string             `json:"subType,omitempty"`
}

// ComponentRepresentation struct
type ComponentRepresentation struct {
	Config       map[string][]string `json:"config,omitempty"`
	ID           *string             `json:"id,omitempty"`
	Name         *string             `json:"name,omitempty"`
	ParentID     *string             `json:"parentId,omitempty"`
	ProviderID   *string             `json:"providerId,omitempty"`
	ProviderType *string             `json:"providerType,omitempty"`
	SubType      *string             `json:"subType,omitempty"`
}

// ComponentTypeRepresentation struct
type ComponentTypeRepresentation struct {
	HelpText   *string                         `json:"helpText,omitempty"`
	ID         *string                         `json:"id,omitempty"`
	Metadata   *map[string]any                 `json:"metadata,omitempty"`
	Properties *[]ConfigPropertyRepresentation `json:"properties,omitempty"`
}

// ConfigPropertyRepresentation struct
type ConfigPropertyRepresentation struct {
	DefaultValue *map[string]any `json:"defaultValue,omitempty"`
	HelpText     *string         `json:"helpText,omitempty"`
	Label        *string         `json:"label,omitempty"`
	Name         *string         `json:"name,omitempty"`
	Options      *[]string       `json:"options,omitempty"`
	Secret       *bool           `json:"secret,omitempty"`
	Type         *string         `json:"type,omitempty"`
}

// CredentialRepresentation struct
type CredentialRepresentation struct {
	ID             *string `json:"id,omitempty"`
	Type           *string `json:"type,omitempty"`
	UserLabel      *string `json:"userLabel,omitempty"`
	CreatedDate    *int64  `json:"createdDate,omitempty"`
	CredentialData *string `json:"credentialData,omitempty"`
	Value          *string `json:"value,omitempty"`
	Temporary      *bool   `json:"temporary,omitempty"`
}

// EventRepresentation struct
type EventRepresentation struct {
	ClientID  *string         `json:"clientId,omitempty"`
	Details   *map[string]any `json:"details,omitempty"`
	Error     *string         `json:"error,omitempty"`
	IPAddress *string         `json:"ipAddress,omitempty"`
	RealmID   *string         `json:"realmId,omitempty"`
	SessionID *string         `json:"sessionId,omitempty"`
	Time      *int64          `json:"time,omitempty"`
	Type      *string         `json:"type,omitempty"`
	UserID    *string         `json:"userId,omitempty"`
}

// FederatedIdentityRepresentation struct
type FederatedIdentityRepresentation struct {
	IdentityProvider *string `json:"identityProvider,omitempty"`
	UserID           *string `json:"userId,omitempty"`
	UserName         *string `json:"userName,omitempty"`
}

// GlobalRequestResult struct
type GlobalRequestResult struct {
	FailedRequests  *[]string `json:"failedRequests,omitempty"`
	SuccessRequests *[]string `json:"successRequests,omitempty"`
}

// GroupRepresentation struct
type GroupRepresentation struct {
	Access      *map[string]any        `json:"access,omitempty"`
	Attributes  *map[string]any        `json:"attributes,omitempty"`
	ClientRoles *map[string]any        `json:"clientRoles,omitempty"`
	ID          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Path        *string                `json:"path,omitempty"`
	RealmRoles  *[]string              `json:"realmRoles,omitempty"`
	SubGroups   *[]GroupRepresentation `json:"subGroups,omitempty"`
}

// IdentityProviderMapperRepresentation struct
type IdentityProviderMapperRepresentation struct {
	Config                 *map[string]any `json:"config,omitempty"`
	ID                     *string         `json:"id,omitempty"`
	IdentityProviderAlias  *string         `json:"identityProviderAlias,omitempty"`
	IdentityProviderMapper *string         `json:"identityProviderMapper,omitempty"`
	Name                   *string         `json:"name,omitempty"`
}

// IdentityProviderRepresentation struct
type IdentityProviderRepresentation struct {
	AddReadTokenRoleOnCreate  *bool          `json:"addReadTokenRoleOnCreate,omitempty"`
	Alias                     *string        `json:"alias,omitempty"`
	AuthenticateByDefault     *bool          `json:"authenticateByDefault,omitempty"`
	Config                    map[string]any `json:"config,omitempty"`
	DisplayName               *string        `json:"displayName,omitempty"`
	Enabled                   *bool          `json:"enabled,omitempty"`
	FirstBrokerLoginFlowAlias *string        `json:"firstBrokerLoginFlowAlias,omitempty"`
	HideOnLogin               *bool          `json:"hideOnLogin,omitempty"`
	InternalID                *string        `json:"internalId,omitempty"`
	LinkOnly                  *bool          `json:"linkOnly,omitempty"`
	PostBrokerLoginFlowAlias  *string        `json:"postBrokerLoginFlowAlias,omitempty"`
	ProviderID                *string        `json:"providerId,omitempty"`
	StoreToken                *bool          `json:"storeToken,omitempty"`
	TrustEmail                *bool          `json:"trustEmail,omitempty"`
}

// KeysMetadataRepresentation struct
type KeysMetadataRepresentation struct {
	Active *map[string]any                                        `json:"active,omitempty"`
	Keys   *[]KeysMetadataRepresentationKeyMetadataRepresentation `json:"keys,omitempty"`
}

// KeysMetadataRepresentationKeyMetadataRepresentation struct
type KeysMetadataRepresentationKeyMetadataRepresentation struct {
	Certificate      *string `json:"certificate,omitempty"`
	Kid              *string `json:"kid,omitempty"`
	ProviderID       *string `json:"providerId,omitempty"`
	ProviderPriority *int64  `json:"providerPriority,omitempty"`
	PublicKey        *string `json:"publicKey,omitempty"`
	Status           *string `json:"status,omitempty"`
	Type             *string `json:"type,omitempty"`
}

// KeyStoreConfig struct
type KeyStoreConfig struct {
	Format           *string `json:"format,omitempty"`
	KeyAlias         *string `json:"keyAlias,omitempty"`
	KeyPassword      *string `json:"keyPassword,omitempty"`
	RealmAlias       *string `json:"realmAlias,omitempty"`
	RealmCertificate *bool   `json:"realmCertificate,omitempty"`
	StorePassword    *string `json:"storePassword,omitempty"`
}

// ManagementPermissionReference struct
type ManagementPermissionReference struct {
	Enabled          *bool           `json:"enabled,omitempty"`
	Resource         *string         `json:"resource,omitempty"`
	ScopePermissions *map[string]any `json:"scopePermissions,omitempty"`
}

// MappingsRepresentation struct
type MappingsRepresentation struct {
	ClientMappings *map[string]any       `json:"clientMappings,omitempty"`
	RealmMappings  *[]RoleRepresentation `json:"realmMappings,omitempty"`
}

// MemoryInfoRepresentation struct
type MemoryInfoRepresentation struct {
	Free           *int64  `json:"free,omitempty"`
	FreeFormated   *string `json:"freeFormated,omitempty"`
	FreePercentage *int64  `json:"freePercentage,omitempty"`
	Total          *int64  `json:"total,omitempty"`
	TotalFormated  *string `json:"totalFormated,omitempty"`
	Used           *int64  `json:"used,omitempty"`
	UsedFormated   *string `json:"usedFormated,omitempty"`
}

// MultivaluedHashMap struct
type MultivaluedHashMap struct {
	Empty      *bool  `json:"empty,omitempty"`
	LoadFactor *int32 `json:"loadFactor,omitempty"`
	Threshold  *int32 `json:"threshold,omitempty"`
}

// PartialImportRepresentation struct
type PartialImportRepresentation struct {
	Clients           *[]ClientRepresentation           `json:"clients,omitempty"`
	Groups            *[]GroupRepresentation            `json:"groups,omitempty"`
	IdentityProviders *[]IdentityProviderRepresentation `json:"identityProviders,omitempty"`
	IfResourceExists  *string                           `json:"ifResourceExists,omitempty"`
	Policy            *string                           `json:"policy,omitempty"`
	Roles             *RolesRepresentation              `json:"roles,omitempty"`
	Users             *[]UserRepresentation             `json:"users,omitempty"`
}

// PasswordPolicyTypeRepresentation struct
type PasswordPolicyTypeRepresentation struct {
	ConfigType        *string `json:"configType,omitempty"`
	DefaultValue      *string `json:"defaultValue,omitempty"`
	DisplayName       *string `json:"displayName,omitempty"`
	ID                *string `json:"id,omitempty"`
	MultipleSupported *bool   `json:"multipleSupported,omitempty"`
}

// PolicyRepresentation struct
type PolicyRepresentation struct {
	Config           *map[string]any `json:"config,omitempty"`
	DecisionStrategy *string         `json:"decisionStrategy,omitempty"`
	Description      *string         `json:"description,omitempty"`
	ID               *string         `json:"id,omitempty"`
	Logic            *string         `json:"logic,omitempty"`
	Name             *string         `json:"name,omitempty"`
	Policies         *[]string       `json:"policies,omitempty"`
	Resources        *[]string       `json:"resources,omitempty"`
	Scopes           *[]string       `json:"scopes,omitempty"`
	Type             *string         `json:"type,omitempty"`
}

// ProfileInfoRepresentation struct
type ProfileInfoRepresentation struct {
	DisabledFeatures *[]string `json:"disabledFeatures,omitempty"`
	Name             *string   `json:"name,omitempty"`
}

// ProtocolMapperRepresentation struct
type ProtocolMapperRepresentation struct {
	Config          *map[string]any `json:"config,omitempty"`
	ConsentRequired *bool           `json:"consentRequired,omitempty"`
	ConsentText     *string         `json:"consentText,omitempty"`
	ID              *string         `json:"id,omitempty"`
	Name            *string         `json:"name,omitempty"`
	Protocol        *string         `json:"protocol,omitempty"`
	ProtocolMapper  *string         `json:"protocolMapper,omitempty"`
}

// ProviderRepresentation struct
type ProviderRepresentation struct {
	OperationalInfo *map[string]any `json:"operationalInfo,omitempty"`
	Order           *int32          `json:"order,omitempty"`
}

// RealmEventsConfigRepresentation struct
type RealmEventsConfigRepresentation struct {
	AdminEventsDetailsEnabled *bool     `json:"adminEventsDetailsEnabled,omitempty"`
	AdminEventsEnabled        *bool     `json:"adminEventsEnabled,omitempty"`
	EnabledEventTypes         *[]string `json:"enabledEventTypes,omitempty"`
	EventsEnabled             *bool     `json:"eventsEnabled,omitempty"`
	EventsExpiration          *int64    `json:"eventsExpiration,omitempty"`
	EventsListeners           *[]string `json:"eventsListeners,omitempty"`
}

// RealmRepresentation struct
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
	Attributes                          *map[string]*string                     `json:"attributes,omitempty"`
	AuthenticationFlows                 *[]AuthenticationFlowRepresentation     `json:"authenticationFlows,omitempty"`
	AuthenticatorConfig                 *[]AuthenticatorConfigRepresentation    `json:"authenticatorConfig,omitempty"`
	BrowserFlow                         *string                                 `json:"browserFlow,omitempty"`
	BrowserSecurityHeaders              *map[string]any                         `json:"browserSecurityHeaders,omitempty"`
	BruteForceProtected                 *bool                                   `json:"bruteForceProtected,omitempty"`
	ClientAuthenticationFlow            *string                                 `json:"clientAuthenticationFlow,omitempty"`
	ClientScopeMappings                 *map[string]any                         `json:"clientScopeMappings,omitempty"`
	ClientTemplates                     *[]ClientTemplateRepresentation         `json:"clientTemplates,omitempty"`
	Clients                             *[]ClientRepresentation                 `json:"clients,omitempty"`
	Components                          *MultivaluedHashMap                     `json:"components,omitempty"`
	DefaultGroups                       *[]string                               `json:"defaultGroups,omitempty"`
	DefaultLocale                       *string                                 `json:"defaultLocale,omitempty"`
	DefaultRoles                        *[]string                               `json:"defaultRoles,omitempty"`
	DirectGrantFlow                     *string                                 `json:"directGrantFlow,omitempty"`
	DisplayName                         *string                                 `json:"displayName,omitempty"`
	DisplayNameHTML                     *string                                 `json:"displayNameHtml,omitempty"`
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
	ID                                  *string                                 `json:"id,omitempty"`
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
	SMTPServer                          *map[string]any                         `json:"smtpServer,omitempty"`
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

// RequiredActionProviderRepresentation struct
type RequiredActionProviderRepresentation struct {
	Alias         *string         `json:"alias,omitempty"`
	Config        *map[string]any `json:"config,omitempty"`
	DefaultAction *bool           `json:"defaultAction,omitempty"`
	Enabled       *bool           `json:"enabled,omitempty"`
	Name          *string         `json:"name,omitempty"`
	ProviderID    *string         `json:"providerId,omitempty"`
}

// ResourceOwnerRepresentation struct
type ResourceOwnerRepresentation struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ResourceRepresentation struct
type ResourceRepresentation struct {
	ID          *string                      `json:"id,omitempty"`
	IconURI     *string                      `json:"icon_uri,omitempty"`
	Name        *string                      `json:"name,omitempty"`
	Owner       *ResourceOwnerRepresentation `json:"owner,omitempty"`
	Policies    *[]PolicyRepresentation      `json:"policies,omitempty"`
	Scopes      *[]ScopeRepresentation       `json:"scopes,omitempty"`
	Type        *string                      `json:"type,omitempty"`
	TypedScopes *[]ScopeRepresentation       `json:"typedScopes,omitempty"`
	URI         *string                      `json:"uri,omitempty"`
}

// ResourceServerRepresentation struct
type ResourceServerRepresentation struct {
	AllowRemoteResourceManagement *bool                     `json:"allowRemoteResourceManagement,omitempty"`
	ClientID                      *string                   `json:"clientId,omitempty"`
	ID                            *string                   `json:"id,omitempty"`
	Name                          *string                   `json:"name,omitempty"`
	Policies                      *[]PolicyRepresentation   `json:"policies,omitempty"`
	PolicyEnforcementMode         *string                   `json:"policyEnforcementMode,omitempty"`
	Resources                     *[]ResourceRepresentation `json:"resources,omitempty"`
	Scopes                        *[]ScopeRepresentation    `json:"scopes,omitempty"`
}

// RoleRepresentation struct
type RoleRepresentation struct {
	ClientRole         *bool                         `json:"clientRole,omitempty"`
	Composite          *bool                         `json:"composite,omitempty"`
	Composites         *RoleRepresentationComposites `json:"composites,omitempty"`
	ContainerID        *string                       `json:"containerId,omitempty"`
	Description        *string                       `json:"description,omitempty"`
	ID                 *string                       `json:"id,omitempty"`
	Name               *string                       `json:"name,omitempty"`
	ScopeParamRequired *bool                         `json:"scopeParamRequired,omitempty"`
	Attributes         *map[string][]string          `json:"attributes,omitempty"`
}

// RoleRepresentationComposites struct
type RoleRepresentationComposites struct {
	Client *map[string]any `json:"client,omitempty"`
	Realm  *[]string       `json:"realm,omitempty"`
}

// RolesRepresentation struct
type RolesRepresentation struct {
	Client *map[string]any       `json:"client,omitempty"`
	Realm  *[]RoleRepresentation `json:"realm,omitempty"`
}

// ScopeMappingRepresentation struct
type ScopeMappingRepresentation struct {
	Client         *string   `json:"client,omitempty"`
	ClientTemplate *string   `json:"clientTemplate,omitempty"`
	Roles          *[]string `json:"roles,omitempty"`
	Self           *string   `json:"self,omitempty"`
}

// ScopeRepresentation struct
type ScopeRepresentation struct {
	IconURI   *string                   `json:"iconUri,omitempty"`
	ID        *string                   `json:"id,omitempty"`
	Name      *string                   `json:"name,omitempty"`
	Policies  *[]PolicyRepresentation   `json:"policies,omitempty"`
	Resources *[]ResourceRepresentation `json:"resources,omitempty"`
}

// ServerInfoRepresentation struct
type ServerInfoRepresentation struct {
	BuiltinProtocolMappers *map[string]any                     `json:"builtinProtocolMappers,omitempty"`
	ClientImporters        *[]map[string]any                   `json:"clientImporters,omitempty"`
	ClientInstallations    *map[string]any                     `json:"clientInstallations,omitempty"`
	ComponentTypes         *map[string]any                     `json:"componentTypes,omitempty"`
	Enums                  *map[string]any                     `json:"enums,omitempty"`
	IdentityProviders      *[]map[string]any                   `json:"identityProviders,omitempty"`
	MemoryInfo             *MemoryInfoRepresentation           `json:"memoryInfo,omitempty"`
	PasswordPolicies       *[]PasswordPolicyTypeRepresentation `json:"passwordPolicies,omitempty"`
	ProfileInfo            *ProfileInfoRepresentation          `json:"profileInfo,omitempty"`
	ProtocolMapperTypes    *map[string]any                     `json:"protocolMapperTypes,omitempty"`
	Providers              *map[string]any                     `json:"providers,omitempty"`
	SocialProviders        *[]map[string]any                   `json:"socialProviders,omitempty"`
	SystemInfo             *SystemInfoRepresentation           `json:"systemInfo,omitempty"`
	Themes                 *map[string]any                     `json:"themes,omitempty"`
}

// SpiInfoRepresentation struct
type SpiInfoRepresentation struct {
	Internal  *bool           `json:"internal,omitempty"`
	Providers *map[string]any `json:"providers,omitempty"`
}

// SynchronizationResult struct
type SynchronizationResult struct {
	Added   *int32  `json:"added,omitempty"`
	Failed  *int32  `json:"failed,omitempty"`
	Ignored *bool   `json:"ignored,omitempty"`
	Removed *int32  `json:"removed,omitempty"`
	Status  *string `json:"status,omitempty"`
	Updated *int32  `json:"updated,omitempty"`
}

// SystemInfoRepresentation struct
type SystemInfoRepresentation struct {
	FileEncoding   *string `json:"fileEncoding,omitempty"`
	JavaHome       *string `json:"javaHome,omitempty"`
	JavaRuntime    *string `json:"javaRuntime,omitempty"`
	JavaVendor     *string `json:"javaVendor,omitempty"`
	JavaVersion    *string `json:"javaVersion,omitempty"`
	JavaVM         *string `json:"javaVm,omitempty"`
	JavaVMVersion  *string `json:"javaVmVersion,omitempty"`
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

// UserConsentRepresentation struct
type UserConsentRepresentation struct {
	ClientID               *string         `json:"clientId,omitempty"`
	CreatedDate            *int64          `json:"createdDate,omitempty"`
	GrantedClientRoles     *map[string]any `json:"grantedClientRoles,omitempty"`
	GrantedProtocolMappers *map[string]any `json:"grantedProtocolMappers,omitempty"`
	GrantedRealmRoles      *[]string       `json:"grantedRealmRoles,omitempty"`
	LastUpdatedDate        *int64          `json:"lastUpdatedDate,omitempty"`
}

// UserFederationMapperRepresentation struct
type UserFederationMapperRepresentation struct {
	Config                        *map[string]any `json:"config,omitempty"`
	FederationMapperType          *string         `json:"federationMapperType,omitempty"`
	FederationProviderDisplayName *string         `json:"federationProviderDisplayName,omitempty"`
	ID                            *string         `json:"id,omitempty"`
	Name                          *string         `json:"name,omitempty"`
}

// UserFederationProviderRepresentation struct
type UserFederationProviderRepresentation struct {
	ChangedSyncPeriod *int32          `json:"changedSyncPeriod,omitempty"`
	Config            *map[string]any `json:"config,omitempty"`
	DisplayName       *string         `json:"displayName,omitempty"`
	FullSyncPeriod    *int32          `json:"fullSyncPeriod,omitempty"`
	ID                *string         `json:"id,omitempty"`
	LastSync          *int32          `json:"lastSync,omitempty"`
	Priority          *int32          `json:"priority,omitempty"`
	ProviderName      *string         `json:"providerName,omitempty"`
}

// AttributeKey type
type AttributeKey string

// Attributes type
type Attributes map[AttributeKey][]string

// UserRepresentation struct
type UserRepresentation struct {
	Access                     *map[string]bool                   `json:"access,omitempty"`
	Attributes                 *Attributes                        `json:"attributes,omitempty"`
	ClientConsents             *[]UserConsentRepresentation       `json:"clientConsents,omitempty"`
	ClientRoles                *map[string][]string               `json:"clientRoles,omitempty"`
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
	ID                         *string                            `json:"id,omitempty"`
	LastName                   *string                            `json:"lastName,omitempty"`
	NotBefore                  *int32                             `json:"notBefore,omitempty"`
	Origin                     *string                            `json:"origin,omitempty"`
	RealmRoles                 *[]string                          `json:"realmRoles,omitempty"`
	RequiredActions            *[]string                          `json:"requiredActions,omitempty"`
	Self                       *string                            `json:"self,omitempty"`
	ServiceAccountClientID     *string                            `json:"serviceAccountClientId,omitempty"`
	Username                   *string                            `json:"username,omitempty"`
}

// UsersPageRepresentation is used to manage users paging
type UsersPageRepresentation struct {
	Count *int                 `json:"count,omitempty"`
	Users []UserRepresentation `json:"users,omitempty"`
}

// UserSessionRepresentation struct
type UserSessionRepresentation struct {
	Clients    *map[string]any `json:"clients,omitempty"`
	ID         *string         `json:"id,omitempty"`
	IPAddress  *string         `json:"ipAddress,omitempty"`
	LastAccess *int64          `json:"lastAccess,omitempty"`
	Start      *int64          `json:"start,omitempty"`
	UserID     *string         `json:"userId,omitempty"`
	Username   *string         `json:"username,omitempty"`
}

// SmsCodeRepresentation struct
type SmsCodeRepresentation struct {
	Code *string `json:"code,omitempty"`
}

// StatisticsUsersRepresentation elements returned by GetStatisticsUsers
type StatisticsUsersRepresentation struct {
	Total    int64 `json:"total,omitempty"`
	Disabled int64 `json:"disabled,omitempty"`
	Inactive int64 `json:"inactive,omitempty"`
}

// RecoveryCodeRepresentation struct
type RecoveryCodeRepresentation struct {
	Code *string `json:"code,omitempty"`
}

// ActivationCodeRepresentation struct
type ActivationCodeRepresentation struct {
	Code *string `json:"code,omitempty"`
}

// EmailRepresentation struct
type EmailRepresentation struct {
	Recipient   *string                      `json:"recipient,omitempty"`
	Theming     *EmailThemingRepresentation  `json:"theming,omitempty"`
	Attachments *[]AttachementRepresentation `json:"attachments,omitempty"`
}

// EmailThemingRepresentation struct
type EmailThemingRepresentation struct {
	SubjectKey         *string            `json:"subjectKey,omitempty"`
	SubjectParameters  *[]string          `json:"subjectParameters,omitempty"`
	Template           *string            `json:"template,omitempty"`
	TemplateParameters *map[string]string `json:"templateParameters,omitempty"`
	Locale             *string            `json:"locale,omitempty"`
	ThemeRealmName     *string            `json:"themeRealmName,omitempty"`
}

// AttachementRepresentation struct
type AttachementRepresentation struct {
	Filename    *string `json:"filename,omitempty"`
	ContentType *string `json:"contentType,omitempty"`
	Content     *string `json:"content,omitempty"`
}

// SMSRepresentation struct
type SMSRepresentation struct {
	MSISDN  *string                   `json:"msisdn,omitempty"`
	Theming *SMSThemingRepresentation `json:"theming,omitempty"`
}

// SMSThemingRepresentation struct
type SMSThemingRepresentation struct {
	MessageKey        *string   `json:"messageKey,omitempty"`
	MessageParameters *[]string `json:"messageParameters,omitempty"`
	Locale            *string   `json:"locale,omitempty"`
}

// DeletableUserRepresentation struct
type DeletableUserRepresentation struct {
	RealmID   string `json:"realmId,omitempty"`
	RealmName string `json:"realmName,omitempty"`
	UserID    string `json:"userId,omitempty"`
	Username  string `json:"username,omitempty"`
}

// EmailInfoRepresentation struct
type EmailInfoRepresentation struct {
	RealmName    *string `json:"realm,omitempty"`
	CreationDate *int64  `json:"creationDate,omitempty"`
}

// TrustIDAuthTokenRepresentation struct
type TrustIDAuthTokenRepresentation struct {
	Token *string `json:"token"`
}

// LinkedAccountRepresentation struct
type LinkedAccountRepresentation struct {
	Connected      *bool   `json:"connected,omitempty"`
	Social         *bool   `json:"social,omitempty"`
	ProviderAlias  *string `json:"providerAlias,omitempty"`
	ProviderName   *string `json:"providerName,omitempty"`
	DisplayName    *string `json:"displayName,omitempty"`
	LinkedUsername *string `json:"linkedUsername,omitempty"`
}
