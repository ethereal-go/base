package base

import (
	"github.com/ethereal-go/base/root/database"
	"github.com/ethereal-go/ethereal"
	"github.com/ethereal-go/ethereal/utils"
	"github.com/ethereal-go/ethereal/root/app"
	"github.com/graphql-go/graphql"
	"strconv"
	"golang.org/x/crypto/bcrypt"
	"github.com/ethereal-go/ethereal/root/config"

)

var (
	locale = config.GetCnf("L18N.LOCALE").(string)
	i18n   = ethereal.ConstructorI18N()
)

/**
/ User Type
*/
var usersType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Description:string(i18n.T(locale, "graphQL.UserType.Description")),
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: string(i18n.T(locale, "graphQL.UserType.id")),
		},
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: string(i18n.T(locale, "graphQL.UserType.email")),
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: string(i18n.T(locale, "graphQL.UserType.name")),
		},
		"password": &graphql.Field{
			Type:        graphql.String,
			Description: string(i18n.T(locale, "graphQL.UserType.password")),
		},
		"role": &graphql.Field{
			Type:        roleType,
			Description: string(i18n.T(locale, "graphQL.UserType.role")),
		},
	},
})

/**
/ Create User
*/
var CreateUser = graphql.Field{
	Type:        usersType,
	Description: "Create new user",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"role": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		db := params.Context.Value("*Application").(*app.Application).Db

		email, _ := params.Args["email"].(string)
		name, _ := params.Args["name"].(string)
		password, _ := params.Args["password"].(string)
		role, _ := params.Args["role"].(int)

		hashedPassword, err := utils.HashPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			panic(`Error hash password create User service.`)
		}

		var user = database.User{Email: email, Name: name, Password: string(hashedPassword), RoleID: role}
		db.Create(&user)

		return user, nil
	},
}

var UserField = graphql.Field{
	Type:        graphql.NewList(usersType),
	Description: string(i18n.T(locale, "graphQL.User.Description")),
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// code for local jwt token..

		//jwtAuth := params.Context.Value("middlewareJWTToken").(middlewareJWTToken)
		//
		//if jwtAuth.included == false || jwtAuth.authenticated {

		db := params.Context.Value("*Application").(*app.Application).Db
		var users []*database.User
		db.Find(&users)

		idQuery, isOK := params.Args["id"].(string)

		if isOK {
			for _, user := range users {
				if strconv.Itoa(int(user.ID)) == idQuery {

					var role database.Role
					db.Model(&user).Related(&role)
					user.Role = role
					return []database.User{*user}, nil
				}
			}
		}

		//for _, user := range users {
		//	var role Role
		//	db.Model(&user).Related(&role)
		//	user.Role = role
		//}
		return users, nil
		//}
		//jwtAuth.responseWriter.WriteHeader(jwtAuth.statusError)
		//json.NewEncoder(jwtAuth.responseWriter).Encode(http.StatusText(jwtAuth.statusError))
		//return nil, errors.New(jwtAuth.responseError)
	},
}
