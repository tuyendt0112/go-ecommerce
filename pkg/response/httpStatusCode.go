package response

const (
	ErrCodeSuccess      = 20001 // success
	ErrCodeParamInvalid = 20003 // Email is invalid
	ErrAuthentication   = 20004 // Email is invalid

	ErrCodeUserHasExist = 20005 // user has already registered
)

// mess
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrAuthentication:   "Error Authentication please check again",
	ErrCodeUserHasExist: "User has already registered",
}
