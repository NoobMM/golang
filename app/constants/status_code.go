package constants

// Code constant
const (
	// Success code
	StatusCodeSuccess = uint(1000)

	// Error that is caused by input
	StatusCodeMissingRequiredParameters = uint(1101)
	StatusCodeInvalidParameters         = uint(1102)
	StatusCodeEntryExpired              = uint(1103)
	StatusCodeDuplicatedEntry           = uint(1104)
	StatusCodeEntryNotFound             = uint(1105)
	StatusCodeUnprocessableEntity       = uint(1106)
	StatusCodeNotActive                 = uint(1201)
	StatusCodeInputUnsupported          = uint(1202)
	StatusCodeDataUnsupported           = uint(1203)
	StatusCodeGCPQrStatus               = uint(1204)
	StatusCodeTxStatus                  = uint(1205)
	StatusCodeVoidCutOff                = uint(1206)
	StatusCodeTxType                    = uint(1207)
	StatusCodeFileInvalid               = uint(1208)
	StatusCodeTxUnMatchUser             = uint(1209)
	StatusCodeInsufficient              = uint(1210)

	// Error that is caused by GCP
	StatusCodeGenericGCPError = uint(1300)

	// Error that is caused by our own code
	StatusCodeGenericInternalError = uint(8900)
	StatusCodeDatabaseError        = uint(8901)
	StatusCodeStorageError         = uint(8902)

	// Error that is related to security
	StatusCodeAuthError     = uint(9500)
	StatusCodeSecurityError = uint(9900)
	StatusCodeUnauthorized  = uint(9901)
	StatusCodeForbidden     = uint(9902)
)
