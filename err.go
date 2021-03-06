package jsonapivalidator

import "errors"

var (
	// ErrInvalidDocumentStructure is the error returned when the JSON was not
	// a POJO at the top level
	ErrInvalidDocumentStructure = errors.New(`A JSON object MUST be at the ` +
		`root of every JSON API request and response containing data. This ` +
		`object defines a document's "top level".`)

	// ErrAtLeastOneRoot is the error returned when a document is missing a top
	// level member
	ErrAtLeastOneRoot = errors.New("A document MUST contain at least one of the" +
		"following top-level members: /data, /errors or /meta")
	// ErrRootDataAndErrors is the error returned when the top level document
	// contains both data and errors members
	ErrRootDataAndErrors = errors.New("The members /data and /errors MUST NOT " +
		"coexist in the same document")
	// ErrRootIncludedWithoutData is the error returned when the top level
	// document contains an included member without a data member
	ErrRootIncludedWithoutData = errors.New("The /included member MUST NOT be " +
		"be present if there is no /data member")

	// ErrInvalidDataType is for data not being a hash array or null JSON type
	ErrInvalidDataType = errors.New("/data must contain a value that is a {}, " +
		"[] or null")

	// ErrInvalidIncludedType is for errors not being an array
	ErrInvalidIncludedType = errors.New("/included must contain a []")

	// ErrInvalidErrorsType is for errors not being an array
	ErrInvalidErrorsType = errors.New("/errors must contain a []")
	// ErrNotErrorObject is the error when the error object is not a {} with
	// string keys
	ErrNotErrorObject = errors.New("The value at the error key was " +
		"not a valid error object {}")
	// ErrInvalidErrorMember is for when the errors object has an unexpected
	// member key
	ErrInvalidErrorMember = errors.New("Invalid member to /errors")

	// ErrNotJSONAPIObject is for when the "jsonapi" member in the root doc was
	// not a {}
	ErrNotJSONAPIObject = errors.New("/jsonapi must contain a value that is a {}")

	// ErrResourceObjectMissingID is returned when a resouce object did not have
	// any value for the id key
	ErrResourceObjectMissingID = errors.New("A resource object MUST contain an " +
		"id")
	// ErrResourceObjectMissingType is returned when a resouce object did not have
	// any value for the type key
	ErrResourceObjectMissingType = errors.New("A resource object MUST contain a" +
		" type")
	// ErrResourceObjectIdentifierMissingID is returned when a resouce object did
	// not have any value for the id key
	ErrResourceObjectIdentifierMissingID = errors.New("A resource object MUST " +
		"contain an id")
	// ErrResourceObjectIdentifierMissingType is returned when a resouce object
	// did not have any value for the type key
	ErrResourceObjectIdentifierMissingType = errors.New("A resource object MUST" +
		" contain a type")
	// ErrNotAResource is returend when something should have been a resource
	// object or a resource identifier
	ErrNotAResource = errors.New("Was not a resource object or a single " +
		"resource identifier object")

	// ErrIDNotString is for when an id member was anything but a string
	ErrIDNotString = errors.New("id was not a string")
	// ErrTypeNotString is for when a type member was anything but a string
	ErrTypeNotString = errors.New("type was not a string")
	// ErrInvalidURL is the error returned when a value should have been a URL but
	// was an invalid URL
	ErrInvalidURL = errors.New("URL was not a valid URL")

	// ErrNotAttributesObject is the error when the "attributes" value was not a
	// hash {}
	ErrNotAttributesObject = errors.New("The value at the attributes key was " +
		"not a valid attributes object {}")

	// ErrNotMetaObject is the error when the "attributes" value was not a
	// hash {}
	ErrNotMetaObject = errors.New("The value at the meta key was " +
		"not a valid meta object {}")

	// ErrNotRelationshipsObject is the error when the "relationships" value was
	// not a hash {}
	ErrNotRelationshipsObject = errors.New("The value at the relationships key " +
		"was not a valid relationships object {}")
	// ErrRelationshipsObjectOneRequiredMember is used in the case where a
	// relationships object did not have at least one of the required members
	ErrRelationshipsObjectOneRequiredMember = errors.New("A relationship object" +
		" MUST contain at least one of the following: links, data or meta")
	// ErrNotRelationshipObject is the error when a key in a "relationships" {}
	// did not contain a {} also
	ErrNotRelationshipObject = errors.New("The value for the relationship was " +
		"not a valid relationship object {}")
	// ErrInvalidResourceLinkage is the error that is returned if the resource
	// linkage was not represented
	ErrInvalidResourceLinkage = errors.New("Resource linkage MUST be " +
		"represented as one of the following: null for empty-to-one " +
		"relationships, an empty array [] for empty to-many relationships," +
		"a single resource identifier object for non-empty to-one relationships," +
		"or an array of resource identifier objects for non-empty to-many " +
		"relationships")

	// ErrNotLinksObject is the error returned if the links object could not be
	// type asserted as a {string => *}
	ErrNotLinksObject = errors.New("The value at the links key was " +
		"not a valid links object {}")
	// ErrNotLinkObject is the error returned if the link object could not be
	// type asserted as a {string => *}
	ErrNotLinkObject = errors.New("The value at the link key was " +
		"not a valid link object {}")
	// ErrInvalidLinkType is the error returned if a link object was something
	// other than a {} or string
	ErrInvalidLinkType = errors.New("A link MUST be represented as either: a " +
		"a string containing the link’s URL OR a link object")
	// ErrInvalidLinkMember is the error returned if a link object member was not
	// href or meta
	ErrInvalidLinkMember = errors.New("Invalid member to link object; only " +
		"href and meta are permitted")
)
