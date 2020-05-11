package itest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func (suite *BaseSuite) TestGetAllEmpty() {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/documents", nil)

	suite.router.ServeHTTP(recorder, request)

	suite.Equal(http.StatusOK, recorder.Code)

	body := make(map[string]interface{})
	err := json.NewDecoder(recorder.Body).Decode(&body)
	suite.Nil(err)

	// does the documents key exists?
	documents, exists := body["documents"]
	suite.True(exists)

	// validate that no documents are returned
	docs := documents.([]interface{})
	suite.Equal(0, len(docs))
}

func (suite *BaseSuite) TestGetAllOne() {
	suite.uploadDocument(usrFile)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/documents", nil)

	suite.router.ServeHTTP(recorder, request)

	suite.Equal(http.StatusOK, recorder.Code)

	body := make(map[string]interface{})
	err := json.NewDecoder(recorder.Body).Decode(&body)
	suite.Nil(err)

	// does the documents key exists?
	documents, exists := body["documents"]
	suite.True(exists)

	// validate that no documents are returned
	docs := documents.([]interface{})
	suite.Equal(1, len(docs))

	docsFirstElement := docs[0].(map[string]interface{})

	suite.Equal("usr", docsFirstElement["id"])
	suite.Equal("User Service", docsFirstElement["name"])
	suite.Equal("Service", docsFirstElement["type"])
	suite.Equal("Team Integration", docsFirstElement["owner"])
	suite.Equal("The central user access.", docsFirstElement["description"])
	suite.Equal("USR", docsFirstElement["shortName"])
	suite.Equal("Team Integration", docsFirstElement["contact"])

	tags := docsFirstElement["tags"].([]interface{})
	suite.Equal(0, len(tags))

	links := docsFirstElement["links"].(map[string]interface{})
	suite.Equal(2, len(links))
	suite.Equal("http://ci.local/user", links["buildchain"])
	suite.Equal("http://wiki.local/user", links["homepage"])

	service := docsFirstElement["service"].(map[string]interface{})

	provides := service["provides"].([]interface{})
	suite.Equal(1, len(provides))

	providesFirstElement := provides[0].(map[string]interface{})
	suite.Equal("Access to all user information", providesFirstElement["description"])
	suite.Equal("user-service", providesFirstElement["serviceName"])
	suite.Equal("https", providesFirstElement["protocol"])
	suite.Equal(float64(9443), providesFirstElement["port"])
	suite.Equal("tcp", providesFirstElement["transportProtocol"])

	dependsOn := service["dependsOn"].(map[string]interface{})

	internal := dependsOn["internal"].([]interface{})
	suite.Equal(1, len(internal))

	internalFirstElement := internal[0].(map[string]interface{})
	suite.Equal("user-db", internalFirstElement["serviceName"])
	suite.Equal("Need to talk to my database.", internalFirstElement["why"])

	external := dependsOn["external"].([]interface{})
	suite.Equal(1, len(external))

	externalFirstElement := external[0].(map[string]interface{})
	suite.Equal("heroku", externalFirstElement["serviceName"])
	suite.Equal("My db is there", externalFirstElement["why"])
}
