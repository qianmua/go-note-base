package middleware

import "github.com/kataras/iris/middleware/basicauth"

/**
 * @author
 * @date 2020/11/13 17:56
 * create by Goland
 * @version 1.0
 */




// BasicAuth middleware sample.
var BasicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin": "password",
	},
})