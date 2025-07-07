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
		resObj := responseobjects.AuthorizeErrorResObj500{
			ActionStatus:     "ERROR",
			ErrorMessage:     "ServerError",
			ErrorDescription: "System encounter an error. bio sdk servie reports: " + err.Error(),
		}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	recieved := gjson.ParseBytes(bodyBytes)
	additionalParams := recieved.Get("event.request.additionalParams").Array()
	biometricTemplate := additionalParams[0].Get("value").Array()[0].String()
	decodedKey, err := base64.StdEncoding.DecodeString(biometricTemplate)
	if err != nil {
		err = fmt.Errorf("error while decoding recieved biometric key, make sure that key is base64 encoded : %w", err)
		resObj := responseobjects.AuthorizeErrorResObj500{
			ActionStatus:     "ERROR",
			ErrorMessage:     "ServerError",
			ErrorDescription: "System encounter an error. bio sdk servie reports: " + err.Error(),
		}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	probeTemplate, err := controller.sdk.ParseByteArrayToTemplate(&decodedKey)
	if err != nil {
		err = fmt.Errorf("error occured when parsing probe byte data: %w", err)
		resObj := responseobjects.AuthorizeErrorResObj500{
			ActionStatus:     "ERROR",
			ErrorMessage:     "ServerError",
			ErrorDescription: "System encounter an error. bio sdk servie reports: " + err.Error(),
		}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	isMatched, discoveredId, err := controller.sdk.Identify(probeTemplate)
	if err != nil {
		err = fmt.Errorf("error occured when running sdk identify method for probe : %w", err)
		resObj := responseobjects.AuthorizeErrorResObj500{
			ActionStatus:     "ERROR",
			ErrorMessage:     "ServerError",
			ErrorDescription: "System encounter an error. bio sdk servie reports: " + err.Error(),
		}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	log.Println("discover Id: ", discoveredId)
	if isMatched {
		resObj := responseobjects.AuthorizeResObj{
			ActionStatus: "SUCCESS",
			Data: responseobjects.AuthorizeResObj_Data{
				User: responseobjects.AthorizeResObj_User{
					Id: discoveredId,
					Claims: []responseobjects.AuthorizeResObj_Claims{
						{
							Uri:   "http://wso2.org/claims/IsMatched",
							Value: fmt.Sprintf("%t", isMatched),
						},
						{
							Uri:   "http://wso2.org/claims/DiscoveredID",
							Value: discoveredId,
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
	} else {
		resObj := responseobjects.AuthorizeErrorResObj400{
			ActionStatus:     "ERROR",
			ErrorMessage:     "Unauthorized",
			ErrorDescription: "Failed to authorize the request.",
		}
		c.IndentedJSON(http.StatusOK, resObj)
	}

}
