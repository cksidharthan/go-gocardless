package consts

const (
	NordigenBaseURL = "https://ob.nordigen.com/api"
	APIVersion      = "v2"
	TestBaseURL     = "https://localhost:8000/api"
)

const (
	AccountsPath            = "/accounts/:id/"
	AccountBalancesPath     = "/accounts/:id/balances/"
	AccountDetailsPath      = "/accounts/:id/details/"
	AccountTransactionsPath = "/accounts/:id/transactions/"
)

const (
	AccountsTransactionPremiumPath = "/accounts/premium/:id/transactions/"
)

const (
	AgreementsEndusersPath      = "/agreements/enduser/"
	AgreementsEnduserPath       = "/agreements/enduser/:id/"
	AgreementsEnduserAcceptPath = "/agreements/enduser/:id/accept/"
)

const (
	InstitutionsPath = "/institutions/"
)

const (
	PaymentsPath                  = "/payments/"
	PaymentPath                   = "/payments/:id/"
	PaymentSubmitPath             = "/payments/:id/submit/"
	PaymentsAccountPath           = "/payments/account/"
	PaymentsCreditorsPath         = "/payments/creditors/"
	PaymentsCreditorPath          = "/payments/creditors/:id/"
	PaymentsFieldsInstitutionPath = "/payments/fields/:institution_id/"
)

const (
	RequisitionsPath = "/requisitions/"
	RequisitionPath  = "/requisitions/:id/"
)

const (
	TokenNewPath     = "/token/new/"
	TokenRefreshPath = "/token/refresh/"
)

const (
	AustriaInstitutionID    = "AT"
	BelgiumInstitutionID    = "BE"
	BulgariaInstitutionID   = "BG"
	CroatiaInstitutionID    = "HR"
	CyprusInstitutionID     = "CY"
	CzechiaInstitutionID    = "CZ"
	DenmarkInstitutionID    = "DK"
	EstoniaInstitutionID    = "EE"
	FinlandInstitutionID    = "FI"
	FranceInstitutionID     = "FR"
	GermanyInstitutionID    = "DE"
	GreeceInstitutionID     = "GR"
	HungaryInstitutionID    = "HU"
	IrelandInstitutionID    = "IE"
	IcelandInstitutionID    = "IS"
	ItalyInstitutionID      = "IT"
	LatviaInstitutionID     = "LV"
	LiechtensteinID         = "LI"
	LithuaniaInstitutionID  = "LT"
	LuxembourgInstitutionID = "LU"
	MaltaInstitutionID      = "MT"
	NetherlandsInstitution  = "NL"
	NorwayInstitutionID     = "NO"
	PolandInstitutionID     = "PL"
	PortugalInstitutionID   = "PT"
	RomaniaInstitutionID    = "RO"
	SlovakiaInstitutionID   = "SK"
	SloveniaInstitutionID   = "SI"
	SpainInstitutionID      = "ES"
	SwedenInstitutionID     = "SE"
	UKInstitutionID         = "GB"
	USAInstitutionID        = "US"
)

const (
	TestInstitutionID = "REVOLUT_REVOLT21"
)
