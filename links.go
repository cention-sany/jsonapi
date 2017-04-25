package jsonapi

type ServerInformation interface {
	GetBaseURL() string
	GetPrefix() string
}

// LinksWithSI provides server information to implementer for constructing the
// self link object. Return nil pointer mean no link object will be contructed.
// This is lower priority than Linkable and only active if marshal with server
// information.
type SILinkable interface {
	LinksWithSI(ServerInformation) *Links
}

// RelationshipLinksWithSI provides server information to implementer for
// constructing the relationships link object. Return nil pointer mean no link
// object will be contructed. This is lower priority than RelationshipLinkable
// and only active if marshal with server information.
type SIRelationshipLinkable interface {
	RelationshipLinksWithSI(string, ServerInformation) *Links
}

func MarshalManyWithSI(models []interface{},
	si ServerInformation) (*ManyPayload, error) {
	return mMany(models, si)
}

func MarshalOneWithSI(model interface{}, si ServerInformation) (*OnePayload,
	error) {
	return mOne(model, si)
}
