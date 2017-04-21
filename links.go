package jsonapi

type ServerInformation interface {
	GetBaseURL() string
	GetPrefix() string
}

type SILinkable interface {
	LinksWithSI(ServerInformation) *Links
}

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
