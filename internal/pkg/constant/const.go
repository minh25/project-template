package constant

const (
	EmptyString          = ""
	Authorization        = "Authorization"
	AcceptLanguageHeader = "Accept-Language"
	LanguageKey          = "language"
	LangVi               = "vi"
	LangEn               = "en"
	DefaultErrEn         = "System error"
	DefaultErrVi         = "Lỗi hệ thống"
	DefaultSuccessMsgEn  = "Success"
	DefaultSuccessMsgVi  = "Thành công"
	MessageSuccess       = "Success"
	MaxResponseLength    = 52428800 // 50MB
)

// Response
const (
	SuccessCode       = 200
	InvalidFormatCode = 4000
	OutOfRangeCode    = 4001
	RequiredCode      = 4002
	InvalidCode       = 4003

	UnauthorizedCode = 401

	PermissionDeniedCode = 403

	NotFoundCode = 404

	DuplicatedCode = 4090
	NotMatchCode   = 4091

	InternalErrorCode  = 500
	ServiceUnavailable = 503

	ErrorServiceCode = 5030
)

// Message key
const (
	SuccessMsgKey               = "Success"
	ErrCallIAMMsgKey            = "ErrCallIAM"
	ErrInvalidTokenMsgKey       = "ErrInvalidToken"
	ErrInvalidFormatMsgKey      = "ErrInvalidFormat"
	ErrOutOfRangeMsgKey         = "ErrOutOfRange"
	ErrRequiredMsgKey           = "ErrRequired"
	ErrInvalidMsgKey            = "ErrInvalid"
	ErrUnauthorizedMsgKey       = "ErrUnauthorized"
	ErrNoPermissionMsgKey       = "ErrNoPermission"
	ErrNotFoundMsgKey           = "ErrNotFound"
	ErrDuplicatedMsgKey         = "ErrDuplicated"
	ErrNotMatchMsgKey           = "ErrNotMatch"
	ErrInternalErrorMsgKey      = "ErrInternalError"
	ErrServiceUnavailableMsgKey = "ErrServiceUnavailable"
	ErrExternalServiceMsgKey    = "ErrExternalService"
)

// Field key
const (
	PageField            = "page"
	PageSizeField        = "page_size"
	KeywordField         = "keyword"
	CreatedAtField       = "created_at"
	UpdatedAtField       = "updated_at"
	Brand                = "brand"
	MsgType              = "msgType"
	Name                 = "name"
	Status               = "status"
	Config               = "config"
	ProviderSenderParams = "ProviderParams, SenderParams"
	And                  = " and "
	Provider             = "Provider"
	AuthorizationCode    = "authorizationCode"
	SenderID             = "senderId"
	ProviderID           = "ProviderId"
	ValueType            = "ValueType"
)

// trigger

