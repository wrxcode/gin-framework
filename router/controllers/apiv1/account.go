package apiv1

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shiyicode/gin-framework/managers"
	"github.com/shiyicode/gin-framework/models"
	"github.com/shiyicode/gin-framework/router/controllers/base-controller"
)

func RegisterAccount(router *gin.RouterGroup) {
	router.POST("account/login", httpHandlerLogin)
	router.POST("account/register", httpHandlerRegister)
}

func httpHandlerLogin(c *gin.Context) {
	account := models.Account{}
	err := c.Bind(&account)
	if err != nil {
		panic(err)
	}
	if flag, token, mess := managers.AccountLogin(account.Email, account.Password); flag == false {
		c.JSON(http.StatusOK, (&basecontroller.Base{}).Fail(mess))
	} else {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    base64.StdEncoding.EncodeToString([]byte(token)),
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.JSON(http.StatusOK, (&basecontroller.Base{}).Success())
	}
}

func httpHandlerRegister(c *gin.Context) {
	account := models.Account{}
	err := c.Bind(&account)
	if err != nil {
		panic(err)
	}
	if flag, userId, mess := managers.AccountRegister(account.Email, account.Password); flag == false {
		c.JSON(http.StatusOK, (&basecontroller.Base{}).Fail(mess))
	} else {
		c.JSON(http.StatusOK, (&basecontroller.Base{}).Success(userId))
	}
}
