package response

const (
	ErrCodeSuccess      = 20001 // success
	ErrCodeParamInvalid = 20003 // Email is invalid
	ErrAuthentication   = 20004 // Email is invalid
)

// mess
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrAuthentication:"Error Authentication please check again",
}
