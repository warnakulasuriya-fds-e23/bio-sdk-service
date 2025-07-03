package fingerprintcontroller

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/warnakulasuriya-fds-e23/bio-sdk-service/internal/responseobjects"
)

func (controller *fingerprintController) authorize(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		err = fmt.Errorf("error occured when reading bytes of Request Body: %w", err)
		resObj := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	recieved := gjson.ParseBytes(bodyBytes)
	additionalParams := recieved.Get("event.request.additionalParams").Array()
	biometricKey := additionalParams[0].Get("value").Array()[0].String()
	decodedKey, err := base64.StdEncoding.DecodeString(biometricKey)
	if err != nil {
		err = fmt.Errorf("error while decoding recieved biometric key, make sure that key is base64 encoded : %w", err)
		resObj := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	log.Println("biometric data:", string(decodedKey))

	resObj := responseobjects.AuthorizeResObj{
		ActionStatus: "SUCCESS",
		Data: responseobjects.AuthorizeResObj_Data{
			User: responseobjects.AthorizeResObj_User{
				Id: "9f1ab106-ce85-46b1-8f41-6a071b54eb56",
				Claims: []responseobjects.AuthorizeResObj_Claims{
					{
						Uri:   "http://wso2.org/claims/username",
						Value: "emily",
					},
					{
						Uri:   "http://wso2.org/claims/emailaddress",
						Value: "emily@aol.com",
					},
				},
				UserStore: responseobjects.AuthorizeResObj_UserStore{
					Id:   "UFJJTUFSWQ==",
					Name: "PRIMARY",
				},
			},
		},
	}

	c.IndentedJSON(http.StatusOK, resObj)

}
