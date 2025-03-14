package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
)

func BodyEquals(t assert.TestingT, obj interface{}, recorder *httptest.ResponseRecorder) {
	bytes, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	actual := string(bytes)

	JSONEquals(t, obj, actual)
}

func JSONEquals(t assert.TestingT, obj interface{}, expected string) {
	bytes, err := json.Marshal(obj)
	assert.Nil(t, err)
	objJSON := string(bytes)

	assert.JSONEq(t, expected, objJSON)
}
