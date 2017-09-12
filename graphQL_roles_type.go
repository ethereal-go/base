package ethereal

import (
	"github.com/ethereal-go/ethereal/root/app"
	"github.com/graphql-go/graphql"
	"strconv"
)

var roleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.ID,
			Description: string(i18n.T(locale, "graphQL.RoleType.id")),
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: string(i18n.T(locale, "graphQL.RoleType.name")),
		},
		"display_name": &graphql.Field{
			Type:        graphql.String,
			Description: string(i18n.T(locale, "graphQL.RoleType.display_name")),
		},
		"description": &graphql.Field{
			Type:        graphql.String,
			Description: string(i18n.T(locale, "graphQL.RoleType.description")),
		},
	},
})

var RoleField = graphql.Field{
	Type:        graphql.NewList(roleType),
	Description: string(i18n.T(locale, "graphQL.Role.Description")),
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.ID,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := params.Context.Value("*Application").(*app.Application).Db
		var roles []Role
		db.Find(&roles)

		idQuery, isOK := params.Args["id"].(string)

		if isOK {
			for _, role := range roles {
				if strconv.FormatInt(int64(role.ID), 10) == idQuery {
					return []Role{role}, nil
				}
			}
		}
		return roles, nil
	},
}
