package test

import (
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

func TestValidate_invalidJSONAPI(t *testing.T) {
	data := []byte(`{
  	"data": {},
  	"jsonapi": [1,2,3]
  }`)

	expectedResult(t, data, jsonapivalidator.ErrNotJSONAPIObject, noWarning)
}

func TestValidate_invalidTopLevelJSONObject(t *testing.T) {
	data := []byte(`[]`)

	expectedResult(t, data, jsonapivalidator.ErrInvalidDocumentStructure, noWarning)
}
