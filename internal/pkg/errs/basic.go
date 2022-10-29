package errs

type Error struct {
	HttpCode int    `json:"http_code,omitempty"`
	Code     string `json:"code"`
	Message  string `json:"message"`
}

var (
	BadRequest = Error{
		Code:    "40001",
		Message: "400 wrong form",
	}
	Speeding = Error{
		Code:    "50001",
		Message: "speeding",
	}
	CaptchaCode = Error{
		Code:    "4008",
		Message: "Incorrect verification code or expired",
	}
	CaptchaCode2 = Error{
		Code:    "4009",
		Message: "Verification code error Or Expired",
	}
	PleaseSignIn = Error{
		HttpCode: 401,
		Code:     "401",
		Message:  "401 please sign in",
	}
	NotData = Error{
		HttpCode: 200,
		Code:     "404",
		Message:  "Not Data",
	}
	LoginFailed = Error{
		Code:    "41001",
		Message: "Login failed Incorrect password or User does not exist",
	}
	WithdrawalPasswordWrong = Error{
		Code:    "41003",
		Message: "Withdrawal password wrong",
	}
	ExUser = Error{
		Code:    "41002",
		Message: "User already exists",
	}
	SqlSystemError = Error{
		Code:    "5002",
		Message: "System upgrade, please try again later",
	}
	ReptileError = Error{
		Code:    "5002",
		Message: "In order to record your crimes, our company will prosecute you for illegal intrusion into the computer information system",
	}
	SystemError = Error{
		Code:    "5001",
		Message: "System upgrade, please try again later",
	}
)

func NewError(code string, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}
