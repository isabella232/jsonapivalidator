package test

import (
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

func TestValidate_validLinks(t *testing.T) {
	data := []byte(`{
		"meta": {},
		"links": {
			"self": "http://example.com/articles?page[number]=3&page[size]=1",
    	"first": "http://example.com/articles?page[number]=1&page[size]=1",
    	"prev": "http://example.com/articles?page[number]=2&page[size]=1",
    	"next": "http://example.com/articles?page[number]=4&page[size]=1",
    	"last": "http://example.com/articles?page[number]=13&page[size]=1",
  		"related": {
    		"href": "http://example.com/articles/1/comments",
    		"meta": {
      		"count": 10
    		}
  		}
		}
	}`)

	expectedResult(t, data, noError, noWarning)
}

func TestValidate_invalidLinks(t *testing.T) {
	data := []byte(`{
		"meta": {},
	  "links": 5
	}`)

	expectedResult(t, data, jsonapivalidator.ErrNotLinksObject, noWarning)
}

func TestValidate_invalidLinkValue(t *testing.T) {
	data := []byte(`{
		"meta": {},
	  "links": {"aren": []}
	}`)

	expectedResult(t, data, jsonapivalidator.ErrInvalidLinkType, noWarning)
}

// TODO: dependent on a validateURL implementation
// func TestValidate_invalidLinkURL(t *testing.T) {
// 	data := []byte(`{
// 		"meta": {},
// 	  "links": {"aren": "25"}
// 	}`)
//
// 	expectedResult(t, data, jsonapivalidator.ErrInvalidURL)
// }

func TestValidate_invalidLinkObjectMember(t *testing.T) {
	data := []byte(`{
		"meta": {},
	  "links": {
			"aren": {"foo": "bar"}
		}
	}`)

	expectedResult(t, data, jsonapivalidator.ErrInvalidLinkMember, noWarning)
}

func TestValidate_invalidTopLevelLinksMember(t *testing.T) {
	data := []byte(`{
		"meta": {},
	  "links": {
			"aren": {
    		"href": "http://example.com/articles/1/comments",
    		"meta": {
      		"count": 10
    		}
  		}
		}
	}`)

	expectedResultHasErrors(t, data, 1)
}
