package middlewares

//// AuthCheck
//func AuthCheck() echo.MiddlewareFunc {
//	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
//		return func(c echo.Context) error {
//			authHeader := c.Request().Header.Get("Authorization")
//			if authHeader == "" {
//				uber.EchoScaLog.Error(fmt.Sprintf("未授权 authHeader: %v", authHeader))
//				return utils.ToJson(c, utils.Res.Response(false, "未授权", code.FAILED))
//			}
//
//			parts := strings.SplitN(authHeader, " ", 2)
//			if !(len(parts) == 2 && parts[0] == "Bearer") {
//				uber.EchoScaLog.Error(fmt.Sprintf("非法Token %v", authHeader))
//				return utils.ToJson(c, utils.Res.Response(false, "非法Token", code.FAILED))
//			}
//
//			_, err := jwt.ParseToken(parts[1])
//			if err != nil {
//				uber.EchoScaLog.Error(fmt.Sprintf("Token认证失败 %v", authHeader))
//				return utils.ToJson(c, utils.Res.Response(false, "Token认证失败", code.FAILED))
//			}
//			return handlerFunc(c)
//		}
//	}
//}
