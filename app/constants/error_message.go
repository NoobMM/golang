package constants

// ErrorMessages constants
const (
	ErrorMessageUnableProcessRequest         = "unable to process the request"
	ErrorMessageMissingRequiredParameter     = "required parameter(s) is missing"
	ErrorMessageUnableProcessParameter       = "unable to process parameters"
	ErrorMessageUnableProcessEntity          = "unable to process the entity"
	ErrorMessageParameterInvalid             = "parameter(s) is invalid"
	ErrorMessageDatabaseError                = "database error(s)"
	ErrorMessageStorageError                 = "storage error(s)"
	ErrorMessageGCPError                     = "payment server error(s)"
	ErrorMessageInternalError                = "internal server error(s)"
	ErrorMessageInvalidLoginCred             = "invalid login credential"
	ErrorMessageUnableAuthen                 = "unable to authenticate the request"
	ErrorMessageDuplicatedUsername           = "duplicated user name"
	ErrorMessageDuplicatedRequester          = "duplicated requesters"
	ErrorMessageExistedMerchant              = "merchant is existed"
	ErrorMessageNoRelatedRequesterRecords    = "related requester records not found"
	ErrorMessageNoRelatedBankAccountRecords  = "related bank account records not found"
	ErrorMessageInvalidSOF                   = "invalid source of funds"
	ErrorMessageUnauthorized                 = "not authorized"
	ErrorMessageForbidden                    = "forbidden"
	ErrorMessageNotFound                     = "not found"
	ErrorMessageMerchantAlreadyDeleted       = "merchant status is deleted"
	ErrorMessageUserOrPassword               = "invalid username or password"
	ErrorMessageInsufficient                 = "insufficient"
	ErrorMessageTxUnMatchUser                = "transaction is not match to user"
	ErrorMessageFmtInvalidFormat             = "%s has invalid format"
	ErrorMessageFmtDatetimeConflict          = "%s must be equal or later than %s"
	ErrorMessageFmtSuspended                 = "%s is suspended"
	ErrorMessageFmtRequired                  = "%s is required"
	ErrorMessageFmtMax                       = "%s cannot be longer than %s"
	ErrorMessageFmtMin                       = "%s must be longer than %s"
	ErrorMessageFmtEmail                     = "invalid email format for %s"
	ErrorMessageFmtSize                      = "%s cannot be larger than %s"
	ErrorMessageFmtOversize                  = "%s too large"
	ErrorMessageFmtLen                       = "%s must be %s characters long"
	ErrorMessageFmtConvert                   = "unable to convert %s to %s"
	ErrorMessageFmtDuplicated                = "%s is already existed"
	ErrorMessageFmtEntryExpired              = "%s is already expired"
	ErrorMessageFmtPassword                  = "password must be at least 8 characters, 1 uppercase, 1 number, and 1 special character"
	ErrorMessageFmtMustMatch                 = "%s must match with %s"
	ErrorMessageFmtOneOf                     = "%s must be one of [%s]"
	ErrorMessageFmtDatetime                  = "%s is invalid datetime format: the format is %s"
	ErrorMessageFmtConflict                  = "%s is conflicted with %s"
	ErrorMessageFmtNotFound                  = "%s not found"
	ErrorMessageFmtUnsupportedVersion        = "not supported version"
	ErrorMessageFmtUnsupported               = "%s is not supported"
	ErrorMessageFmtNotActive                 = "%s is not active"
	ErrorMessageFmtTxAlreadyVoidedOrRefunded = "transaction is already voided or refunded"
	ErrorMessageFmtTxAlreadyXX               = "transaction is already %s"
	ErrorMessageFmtTxStatus                  = "transaction status is not %s"
	ErrorMessageFmtTxType                    = "transaction type is not %s"
	ErrorMessageFmtTxVoidCutOff              = "cannot void after %s"
	ErrorMessageFmtGCPQrStatusUnmatched      = "QR status is not %s"
)