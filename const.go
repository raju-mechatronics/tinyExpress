package te

// Request Types
type RequestMethod = string

const (
	RequestTypeGet     RequestMethod = "GET"
	RequestTypePost    RequestMethod = "POST"
	RequestTypeDelete  RequestMethod = "DELETE"
	RequestTypePut     RequestMethod = "PUT"
	RequestTypePatch   RequestMethod = "PATCH"
	RequestTypeHead    RequestMethod = "HEAD"
	RequestTypeOptions RequestMethod = "OPTIONS"
	RequestTypeConnect RequestMethod = "CONNECT"
	RequestTypeTrace   RequestMethod = "TRACE"
	RequestTypeAny     RequestMethod = "*"
)

const (
	STATUS_CODE_OK                    = 200
	STATUS_CODE_CREATED               = 201
	STATUS_CODE_ACCEPTED              = 202
	STATUS_CODE_NO_CONTENT            = 204
	STATUS_CODE_BAD_REQUEST           = 400
	STATUS_CODE_UNAUTHORIZED          = 401
	STATUS_CODE_FORBIDDEN             = 403
	STATUS_CODE_NOT_FOUND             = 404
	STATUS_CODE_INTERNAL_SERVER_ERROR = 500
	STATUS_CODE_NOT_IMPLEMENTED       = 501
	STATUS_CODE_BAD_GATEWAY           = 502
	STATUS_CODE_SERVICE_UNAVAILABLE   = 503
)

const (
	ERROR_MESSAGE_BAD_REQUEST           = "Bad Request"
	ERROR_MESSAGE_UNAUTHORIZED          = "Unauthorized"
	ERROR_MESSAGE_FORBIDDEN             = "Forbidden"
	ERROR_MESSAGE_NOT_FOUND             = "Not Found"
	ERROR_MESSAGE_METHOD_NOT_ALLOWED    = "Method Not Allowed"
	ERROR_MESSAGE_REQUEST_TIMEOUT       = "Request Timeout"
	ERROR_MESSAGE_CONFLICT              = "Conflict"
	ERROR_MESSAGE_INTERNAL_SERVER_ERROR = "Internal Server Error"
	ERROR_MESSAGE_NOT_IMPLEMENTED       = "Not Implemented"
	ERROR_MESSAGE_BAD_GATEWAY           = "Bad Gateway"
	ERROR_MESSAGE_SERVICE_UNAVAILABLE   = "Service Unavailable"
	ERROR_MESSAGE_GATEWAY_TIMEOUT       = "Gateway Timeout"
)

const (
	ERROR_RESPONSE_ALREADY_SENT = "response already sent"
)
