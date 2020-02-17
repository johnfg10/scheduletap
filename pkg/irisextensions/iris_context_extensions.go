package irisextensions

import "github.com/kataras/iris"

var (
	sendPrivateDetails bool
)

func init() {
	sendPrivateDetails = false
}

// SendPrivateDetails sets the private field sendPrivateDetails which enabled sending error information to the end user
// Primarly used for development you almost certainly done want this in production
func SendPrivateDetails(sendDetails bool) {
	sendPrivateDetails = sendDetails
}

// IsErrorPresent takes a error and returns a boolean if there is a error
func IsErrorPresent(err error) bool {
	if err != nil {
		return true
	}
	return false
}

// FinnishOnError finnishs the current request with a json error message
func FinnishOnError(ctx iris.Context, err error, statuscodeOptional int) bool {
	if IsErrorPresent(err) {
		ctx.Application().Logger().Error(err)
		var resp APIResponse

		if sendPrivateDetails {
			resp = NewErrorResponse(err.Error())
		} else {
			resp = NewErrorResponse("Error details have been logged in the application logger")
		}

		ctx.JSON(&resp)

		if statuscodeOptional != 0 {
			ctx.StatusCode(statuscodeOptional)
		} else {
			ctx.StatusCode(500)
		}

		return true
	}
	return false
}

// FinnishOnErrorDebug finnishes the current request with a json error message and debug log
func FinnishOnErrorDebug(ctx iris.Context, err error, statuscodeOptional int) bool {
	if IsErrorPresent(err) {
		ctx.Application().Logger().Debug(err)
		var resp APIResponse

		if sendPrivateDetails {
			resp = NewErrorResponse(err.Error())
		} else {
			resp = NewErrorResponse("Error details have been logged in the application logger")
		}

		ctx.JSON(&resp)

		if statuscodeOptional != 0 {
			ctx.StatusCode(statuscodeOptional)
		} else {
			ctx.StatusCode(500)
		}

		return true
	}
	return false
}
