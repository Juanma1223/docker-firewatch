package response

import "net/http"

type StatusBody struct {
	Index int    `json:"index"`
	Name  string `json:"name"`
	Http  int    `json:"http"`
}

type Status int

const (
	// GeneralExceptions
	InternalServerError Status = iota
	NotFound
	BadRequest
	Conflict
	Unknown

	// UserExceptions
	InvalidEmail
	InvalidPassword
	InvalidUsername
	EmailAlreadyExists
	UsernameAlreadyExists
	Unauthorized
	Forbidden

	// DBExepctions
	DBQueryError
	DBExecutionError
	DBRowsError
	DBLastRowIdError
	DBScanError
	DBTransactionError
	DBTransactionClosed
	DBCommitError
	DBItemAlreadyExists

	// JsonExceptions
	JsonDecodingError
	JsonEncodingError

	// SuccessfulCodes
	SuccessfulCreation
	SuccessfulDeletion
	SuccessfulUpdate
	SuccessfulSearch

	// FailureCodes
	FailedCreation
	FailedDeletion
	FailedUpdation
	FailedSearch

	// EncryptionExceptions
	EncryptionError
	DecryptionError

	RequestFound
	RequestNotFound
	RequestAlreadyCreated

	PickupInfoFound
	PickupInfoNotFound

	DeliveryInfoFound
	DeliveryInfoNotFound

	SuperdispatchInfoFound
	SuperdispatchInfoNotFound

	TimeIntervalFound

	PlaceFound

	TimerFound
	TimerNotFound

	ExtraInfoFound
	ExtraInfoNotFound

	TimerTypeFound
	TimerTypeNotFound

	BillingInfoFound
	BillingInfoNotFound

	UsageInfoFound
	UsageInfoNotFound

	FileFound
	FileNotFound

	RequestQueueFound
	RequestQueueNotFound

	TypicalRouteFound
	TypicalRouteNotFound
)

var (
	statusMap = map[Status]StatusBody{
		InternalServerError: {Index: 0, Name: "InternalServerError", Http: http.StatusInternalServerError},
		NotFound:            {Index: 1, Name: "NotFound", Http: http.StatusNotFound},
		BadRequest:          {Index: 2, Name: "BadRequest", Http: http.StatusBadRequest},
		Conflict:            {Index: 3, Name: "Conflict", Http: http.StatusConflict},
		Unknown:             {Index: 4, Name: "Unknown", Http: http.StatusNotImplemented},

		InvalidEmail:          {Index: 5, Name: "InvalidEmail", Http: http.StatusBadRequest},
		InvalidPassword:       {Index: 6, Name: "InvalidPassword", Http: http.StatusBadRequest},
		InvalidUsername:       {Index: 7, Name: "InvalidUsername", Http: http.StatusBadRequest},
		EmailAlreadyExists:    {Index: 8, Name: "EmailAlreadyExists", Http: http.StatusConflict},
		UsernameAlreadyExists: {Index: 9, Name: "UsernameAlreadyExists", Http: http.StatusConflict},
		Unauthorized:          {Index: 10, Name: "Unauthorized", Http: http.StatusUnauthorized},
		Forbidden:             {Index: 11, Name: "Forbidden", Http: http.StatusForbidden},

		DBQueryError:        {Index: 12, Name: "DBQueryError", Http: http.StatusInternalServerError},
		DBExecutionError:    {Index: 13, Name: "DBExecutionError", Http: http.StatusInternalServerError},
		DBRowsError:         {Index: 14, Name: "DBRowsError", Http: http.StatusInternalServerError},
		DBLastRowIdError:    {Index: 15, Name: "DBLastRowIdError", Http: http.StatusInternalServerError},
		DBScanError:         {Index: 16, Name: "DBScanError", Http: http.StatusInternalServerError},
		DBTransactionError:  {Index: 17, Name: "DBTransactionError", Http: http.StatusInternalServerError},
		DBTransactionClosed: {Index: 18, Name: "DBTransactionClosed", Http: http.StatusInternalServerError},
		DBCommitError:       {Index: 19, Name: "DBCommitError", Http: http.StatusInternalServerError},
		DBItemAlreadyExists: {Index: 20, Name: "DBItemAlreadyExists", Http: http.StatusConflict},

		JsonDecodingError: {Index: 21, Name: "JsonDecodingError", Http: http.StatusInternalServerError},
		JsonEncodingError: {Index: 22, Name: "JsonEncodingError", Http: http.StatusInternalServerError},

		SuccessfulCreation: {Index: 23, Name: "SuccessfulCreation", Http: http.StatusCreated},
		SuccessfulDeletion: {Index: 24, Name: "SuccessfulDeletion", Http: http.StatusOK},
		SuccessfulUpdate:   {Index: 25, Name: "SuccessfulUpdate", Http: http.StatusOK},
		SuccessfulSearch:   {Index: 26, Name: "SuccessfulSearch", Http: http.StatusOK},

		FailedCreation: {Index: 27, Name: "FailedCreation", Http: http.StatusConflict},
		FailedDeletion: {Index: 28, Name: "FailedDeletion", Http: http.StatusConflict},
		FailedUpdation: {Index: 29, Name: "FailedUpdation", Http: http.StatusConflict},
		FailedSearch:   {Index: 30, Name: "FailedSearch", Http: http.StatusConflict},

		EncryptionError: {Index: 31, Name: "EncryptionError", Http: http.StatusInternalServerError},
		DecryptionError: {Index: 32, Name: "DecryptionError", Http: http.StatusInternalServerError},

		RequestFound:          {Index: 33, Name: "RequestFound", Http: http.StatusOK},
		RequestNotFound:       {Index: 34, Name: "RequestNotFound", Http: http.StatusNotFound},
		RequestAlreadyCreated: {Index: 35, Name: "RequestAlreadyCreated", Http: http.StatusNotFound},

		PickupInfoFound:    {Index: 36, Name: "PickupInfoFound", Http: http.StatusOK},
		PickupInfoNotFound: {Index: 37, Name: "PickupInfoNotFound", Http: http.StatusNotFound},

		DeliveryInfoFound:    {Index: 38, Name: "DeliveryInfoFound", Http: http.StatusOK},
		DeliveryInfoNotFound: {Index: 39, Name: "DeliveryInfoNotFound", Http: http.StatusNotFound},

		SuperdispatchInfoFound:    {Index: 40, Name: "SuperdispatchInfoFound", Http: http.StatusOK},
		SuperdispatchInfoNotFound: {Index: 41, Name: "SuperdispatchInfoNotFound", Http: http.StatusNotFound},

		TimeIntervalFound: {Index: 42, Name: "TimeIntervalFound", Http: http.StatusOK},

		PlaceFound: {Index: 43, Name: "PlaceFound", Http: http.StatusOK},

		TimerFound:    {Index: 44, Name: "TimerFound", Http: http.StatusOK},
		TimerNotFound: {Index: 45, Name: "TimerNotFound", Http: http.StatusNotFound},

		ExtraInfoFound:    {Index: 46, Name: "ExtraInfoFound", Http: http.StatusOK},
		ExtraInfoNotFound: {Index: 47, Name: "ExtraInfoNotFound", Http: http.StatusNotFound},

		TimerTypeFound:    {Index: 48, Name: "TimerTypeFound", Http: http.StatusOK},
		TimerTypeNotFound: {Index: 49, Name: "TimerTypeNotFound", Http: http.StatusNotFound},

		BillingInfoFound:    {Index: 50, Name: "BillingInfoFound", Http: http.StatusOK},
		BillingInfoNotFound: {Index: 51, Name: "BillingInfoNotFound", Http: http.StatusNotFound},

		UsageInfoFound:    {Index: 52, Name: "UsageInfoFound", Http: http.StatusOK},
		UsageInfoNotFound: {Index: 53, Name: "UsageInfoNotFound", Http: http.StatusNotFound},

		FileFound:    {Index: 54, Name: "FileFound", Http: http.StatusOK},
		FileNotFound: {Index: 55, Name: "FileNotFound", Http: http.StatusNotFound},

		RequestQueueFound:    {Index: 56, Name: "RequestQueueFound", Http: http.StatusOK},
		RequestQueueNotFound: {Index: 57, Name: "RequestQueueNotFound", Http: http.StatusNotFound},

		TypicalRouteFound:    {Index: 58, Name: "TypicalRouteFound", Http: http.StatusOK},
		TypicalRouteNotFound: {Index: 59, Name: "TypicalRouteNotFound", Http: http.StatusNotFound},
	}
)

func (s Status) String() string {
	return statusMap[s].Name
}

func (s Status) Index() int {
	return statusMap[s].Index
}

func (s Status) StatusCode() int {
	return statusMap[s].Http
}
