package gocardless

const (
	BaseURL     = "https://bankaccountdata.gocardless.com/api"
	APIVersion  = "v2"
	TestBaseURL = "https://localhost:8000/api"
)

const (
	AccountsPath = "/accounts/"
)

const (
	AgreementsEndusersPath = "/agreements/enduser/"
)

const (
	InstitutionsPath = "/institutions/"
)

const (
	RequisitionsPath = "/requisitions/"
)

const (
	TokenNewPath     = "/token/new/"
	TokenRefreshPath = "/token/refresh/"
)

const (
	AustriaCountryID    = "AT"
	BelgiumCountryID    = "BE"
	BulgariaCountryID   = "BG"
	CroatiaCountryID    = "HR"
	CyprusCountryID     = "CY"
	CzechiaCountryID    = "CZ"
	DenmarkCountryID    = "DK"
	EstoniaCountryID    = "EE"
	FinlandCountryID    = "FI"
	FranceCountryID     = "FR"
	GermanyCountryID    = "DE"
	GreeceCountryID     = "GR"
	HungaryCountryID    = "HU"
	IrelandCountryID    = "IE"
	IcelandCountryID    = "IS"
	ItalyCountryID      = "IT"
	LatviaCountryID     = "LV"
	LiechtensteinID     = "LI"
	LithuaniaCountryID  = "LT"
	LuxembourgCountryID = "LU"
	MaltaCountryID      = "MT"
	NetherlandsCountry  = "NL"
	NorwayCountryID     = "NO"
	PolandCountryID     = "PL"
	PortugalCountryID   = "PT"
	RomaniaCountryID    = "RO"
	SlovakiaCountryID   = "SK"
	SloveniaCountryID   = "SI"
	SpainCountryID      = "ES"
	SwedenCountryID     = "SE"
	UKCountryID         = "GB"
	USACountryID        = "US"
	SandboxCountryID    = "XX"
)

const (
	LangEN = "EN"
)

const (
	TestInstitutionID                = "SANDBOXFINANCE_SFIN0000"
	TestInstitutionMaxHistoricalDays = 90
)

type RequisitionStatus struct {
	Short       string
	Long        string
	Description string
	Stage       int
}

const RequisitionStatusCreated = "CR"
const RequisitionStatusConsentGiven = "GC"
const RequisitionStatusUndergoingAuthentication = "UA"
const RequisitionStatusRejected = "RJ"
const RequisitionStatusSelectAccounts = "SA"
const RequisitionStatusGrantingAccess = "GA"
const RequisitionStatusLinked = "LN"
const RequisitionStatusExpired = "EX"

var RequisitionStatuses = map[string]RequisitionStatus{
	RequisitionStatusCreated: {
		Short:       RequisitionStatusCreated,
		Long:        "CREATED",
		Description: "Requisition has been successfully created",
		Stage:       1,
	},
	RequisitionStatusConsentGiven: {
		Short:       RequisitionStatusConsentGiven,
		Long:        "GIVING_CONSENT",
		Description: "End-user is giving consent at GoCardless's consent screen",
		Stage:       2,
	},
	RequisitionStatusUndergoingAuthentication: {
		Short:       RequisitionStatusUndergoingAuthentication,
		Long:        "UNDERGOING_AUTHENTICATION",
		Description: "End-user is redirected to the financial institution for authentication",
		Stage:       3,
	},
	RequisitionStatusRejected: {
		Short:       RequisitionStatusRejected,
		Long:        "REJECTED",
		Description: "Either SSN verification has failed or end-user has entered incorrect credentials",
		Stage:       4,
	},
	RequisitionStatusSelectAccounts: {
		Short:       RequisitionStatusSelectAccounts,
		Long:        "SELECTING_ACCOUNTS",
		Description: "End-user is selecting accounts",
		Stage:       5,
	},
	RequisitionStatusGrantingAccess: {
		Short:       RequisitionStatusGrantingAccess,
		Long:        "GRANTING_ACCESS",
		Description: "End-user is granting access to their account information",
		Stage:       6,
	},
	RequisitionStatusLinked: {
		Short:       RequisitionStatusLinked,
		Long:        "LINKED",
		Description: "Account has been successfully linked to requisition",
		Stage:       7,
	},
	RequisitionStatusExpired: {
		Short:       RequisitionStatusExpired,
		Long:        "EXPIRED",
		Description: "Access to accounts has expired as set in End User Agreement",
		Stage:       8,
	},
}
