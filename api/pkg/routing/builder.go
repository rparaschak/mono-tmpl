package routing

import "github.com/danielgtaylor/huma/v2"

type Groups struct {
	Public *huma.Group
}

type GroupBuilder struct {
	groups Groups
	tag    string
}

func NewBuilder(parent huma.API, pathPrefix, tag string) *GroupBuilder {
	b := &GroupBuilder{tag: tag}
	tagModifier := func(o *huma.Operation) {
		o.Tags = []string{tag}
	}

	b.groups.Public = huma.NewGroup(parent, pathPrefix)
	b.groups.Public.UseSimpleModifier(tagModifier)

	return b
}

func (b *GroupBuilder) Groups() Groups {
	return b.groups
}

func (b *GroupBuilder) Tag() string {
	return b.tag
}
