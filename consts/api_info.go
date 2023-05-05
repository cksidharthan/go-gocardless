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
	InstitutionPath  = "/institutions/:id/"
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
	TokenNewPath = "/token/new/"
	TokenRefresh = "/token/refresh/"
)
