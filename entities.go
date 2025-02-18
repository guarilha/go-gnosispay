package gnosispay

import "time"

type ApiError struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Code    int    `json:"code,omitempty"`
}

type SignUpResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type User struct {
	Email             string                 `json:"email,omitempty"`
	FirstName         string                 `json:"firstName,omitempty"`
	LastName          string                 `json:"lastName,omitempty"`
	SignInWallets     []EoaAccount           `json:"signInWallets,omitempty"`
	SafeWallets       []SafeAccount          `json:"safeWallets,omitempty"`
	KycStatus         *KycStatus             `json:"kycStatus,omitempty"`
	AvailableFeatures *UserAvailableFeatures `json:"availableFeatures,omitempty"`
	Cards             []Card                 `json:"cards,omitempty"`
	BankingDetails    *BankingDetails        `json:"bankingDetails,omitempty"`
}

type KycStatus string

// List of KYCStatus
const (
	NOT_STARTED_KycStatus            KycStatus = "notStarted"
	DOCUMENTS_REQUESTED_KycStatus    KycStatus = "documentsRequested"
	PENDING_KycStatus                KycStatus = "pending"
	PROCESSING_KycStatus             KycStatus = "processing"
	APPROVED_KycStatus               KycStatus = "approved"
	RESUBMISSION_REQUESTED_KycStatus KycStatus = "resubmissionRequested"
	REJECTED_KycStatus               KycStatus = "rejected"
	REQUIRES_ACTION_KycStatus        KycStatus = "requiresAction"
)

type Card struct {
	Id             string    `json:"id"`
	LastFourDigits string    `json:"lastFourDigits"`
	ActivatedAt    time.Time `json:"activatedAt,omitempty"`
}

type CardStatus struct {
	ActivatedAt string  `json:"activatedAt,omitempty"`
	StatusCode  float64 `json:"statusCode"`
	IsFrozen    bool    `json:"isFrozen"`
	IsStolen    bool    `json:"isStolen"`
	IsLost      bool    `json:"isLost"`
	IsBlocked   bool    `json:"isBlocked"`
	IsVoid      bool    `json:"isVoid"`
}

type GetTransactionsFilters struct {
	CardTokens          *string `json:"cardTokens"`
	Before              *string `json:"before"`
	After               *string `json:"after"`
	BillingCurrency     *string `json:"billingCurrency"`
	TransactionCurrency *string `json:"transactionCurrency"`
	MCC                 *string `json:"mcc"`
}

type Country struct {
	Name    string `json:"name,omitempty"`
	Numeric string `json:"numeric,omitempty"`
	Alpha2  string `json:"alpha2,omitempty"`
	Alpha3  string `json:"alpha3,omitempty"`
}

type Merchant struct {
	Name    string   `json:"name,omitempty"`
	City    string   `json:"city,omitempty"`
	Country *Country `json:"country,omitempty"`
}

type Currency struct {
	Symbol   string `json:"symbol,omitempty"`
	Code     string `json:"code,omitempty"`
	Decimals int32  `json:"decimals,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Transaction struct {
	Status string `json:"status,omitempty"`
	To     string `json:"to,omitempty"`
	Value  string `json:"value,omitempty"`
	Data   string `json:"data,omitempty"`
	Hash   string `json:"hash,omitempty"`
}

type CardEvent struct {
	Kind                string        `json:"kind,omitempty"`
	Status              string        `json:"status,omitempty"`
	CreatedAt           time.Time     `json:"createdAt,omitempty"`
	ClearedAt           time.Time     `json:"clearedAt,omitempty"`
	Country             *Country      `json:"country,omitempty"`
	IsPending           bool          `json:"isPending,omitempty"`
	Mcc                 string        `json:"mcc,omitempty"`
	Merchant            *Merchant     `json:"merchant,omitempty"`
	BillingAmount       string        `json:"billingAmount,omitempty"`
	BillingCurrency     *Currency     `json:"billingCurrency,omitempty"`
	TransactionAmount   string        `json:"transactionAmount,omitempty"`
	TransactionCurrency *Currency     `json:"transactionCurrency,omitempty"`
	Transactions        []Transaction `json:"transactions,omitempty"`
}

type EoaAccount struct {
	Id        string    `json:"id,omitempty"`
	Address   string    `json:"address,omitempty"`
	UserId    string    `json:"userId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type SafeAccount struct {
	Id          string    `json:"id,omitempty"`
	Address     string    `json:"address,omitempty"`
	Salt        *string   `json:"salt,omitempty"`
	ChainId     string    `json:"chainId,omitempty"`
	UserId      string    `json:"userId,omitempty"`
	TokenSymbol string    `json:"tokenSymbol,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}

type UserAvailableFeatures struct {
	MoneriumIban bool `json:"moneriumIban,omitempty"`
}

type BankingDetails struct {
	Id                 string    `json:"id,omitempty"`
	Address            string    `json:"address,omitempty"`
	MoneriumIban       string    `json:"moneriumIban,omitempty"`
	MoneriumBic        string    `json:"moneriumBic,omitempty"`
	MoneriumIbanStatus string    `json:"moneriumIbanStatus,omitempty"`
	UserId             string    `json:"userId,omitempty"`
	CreatedAt          time.Time `json:"createdAt,omitempty"`
	UpdatedAt          time.Time `json:"updatedAt,omitempty"`
}

type IbanDetails struct {
	Iban       string `json:"iban,omitempty"`
	Bic        string `json:"bic,omitempty"`
	IbanStatus string `json:"ibanStatus,omitempty"`
}

type IbanOrderCounterpart struct {
	Details struct {
		Name string `json:"name,omitempty"`
	} `json:"details"`
	Identifier struct {
		Standard string `json:"standard"`
		Address  string `json:"address,omitempty"`
		Chain    string `json:"chain,omitempty"`
		Iban     string `json:"iban,omitempty"`
	}
}

type IbanOrderMetadata struct {
	PlacedAt string `json:"placedAt"`
}

type IbanOrder struct {
	Id          string                `json:"id"`       // Unique identifier for the order.
	Kind        string                `json:"kind"`     // Type of order.
	Currency    string                `json:"currency"` // Currency of the order.
	Amount      string                `json:"amount"`   // Amount of the transaction in string format.
	Address     string                `json:"address"`  // Ethereum address associated with the order.
	Counterpart *IbanOrderCounterpart `json:"counterpart"`
	Memo        string                `json:"memo,omitempty"` // Optional memo for the order.
	State       string                `json:"state"`          // State of the order.
	Meta        *IbanOrderMetadata    `json:"meta"`
}

type AccountBalances struct {
	Total     string `json:"total"`     // The total balance for this account (spendable and pending).
	Spendable string `json:"spendable"` // The amount that can be spent from this account.
	Pending   string `json:"pending"`   // The amount that is being reviewed for spending.
}

type KycIntegration struct {
	Type string `json:"type"` // The type of the KYC integration.
	Url  string `json:"url"`  // The url that needs to be opened to follow through the SumSub KYC flow.
}

type UserReferrals struct {
	IsOgTokenHolder    bool  `json:"isOgTokenHolder,omitempty"` // Indicates if the user is an OG token holder
	PendingReferrals   int32 `json:"pendingReferrals"`          // Number of pending referrals
	CompletedReferrals int32 `json:"completedReferrals"`        // Number of completed referrals
}

type UserReferralCode struct {
	UserId       string `json:"userId"`       // ID of the authenticated user
	ReferrerCode string `json:"referrerCode"` // User's referral code for sharing
}

type KycQuestion struct {
	Question string   `json:"question"` // The text of the question.
	Answers  []string `json:"answers"`  // The possible answers to the question.
}

type KycAnswer struct {
	Question string `json:"question"` // The text of the question being answered.
	Answer   string `json:"answer"`   // The user's answer to the question.
}

type ApiGenericResponse struct {
	Ok      bool   `json:"ok,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type KycPhoneVerification struct {
	PhoneNumber string `json:"phoneNumber"` // The phone number to verify
}

type KycPhoneVerificationCheck struct {
	Code string `json:"code"` // The phone number to verify
}

type KycImportPartnerApplicant struct {
	ForClientId string  `json:"forClientId"`
	TtlInSecs   float64 `json:"ttlInSecs"`
}

type KycImportPartnerApplicantResponse struct {
	Token       string `json:"token"`
	ForClientId string `json:"forClientId"`
}

type DelayTransaction struct {
	Id              string    `json:"id,omitempty"`              // Unique identifier for the delayed transaction.
	SafeAddress     string    `json:"safeAddress,omitempty"`     // The Safe contract address associated with the transaction.
	TransactionData string    `json:"transactionData,omitempty"` // Data payload of the transaction.
	EnqueueTaskId   string    `json:"enqueueTaskId,omitempty"`   // Identifier of the task that enqueued this transaction.
	DispatchTaskId  string    `json:"dispatchTaskId,omitempty"`  // Identifier of the task responsible for dispatching this transaction.
	ReadyAt         time.Time `json:"readyAt,omitempty"`         // Timestamp indicating when the transaction is ready for processing.
	OperationType   string    `json:"operationType,omitempty"`   // Type of operation being performed.
	UserId          string    `json:"userId,omitempty"`          // Identifier of the user associated with the transaction.
	Status          string    `json:"status,omitempty"`          // Current status of the transaction.
	CreatedAt       time.Time `json:"createdAt,omitempty"`       // Timestamp of when the transaction was created.
}

type SafeConfig struct {
	HasNoApprovals bool   `json:"hasNoApprovals,omitempty"` // Indicates whether the safe has no approvals.
	IsDeployed     bool   `json:"isDeployed,omitempty"`     // Indicates whether the safe is deployed.
	Address        string `json:"address,omitempty"`        // The address of the safe, if available.
	TokenSymbol    string `json:"tokenSymbol,omitempty"`    // The token symbol associated with the safe.
	FiatSymbol     string `json:"fiatSymbol,omitempty"`     // The fiat symbol derived from the token symbol.
}
