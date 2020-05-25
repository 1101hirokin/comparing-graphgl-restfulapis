package gql

import (
	"errors"
	"log"

	"go-server/model"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetSchemaConfig(db *gorm.DB) graphql.SchemaConfig {
	userObject := graphql.NewObject(graphql.ObjectConfig{
		Name: "Params",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"bio": &graphql.Field{
				Type: graphql.String,
			},
			"url_avatar": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	q := graphql.ObjectConfig{
		Name:        "query",
		Description: "Root query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:        userObject,
				Description: "get user data",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					log.Println("resolve func was called!")
					var err error = nil
					user := model.User{}

					user_id, success := p.Args["id"].(int)
					if !success {
						err = errors.New("Incorrect Param was set")
						return &user, err
					}
					if db.First(&user, user_id).RecordNotFound() {
						err = errors.New("NotFound")
						user = model.User{}
						return &user, err
					}
					return &user, err
				},
			},
			"users": &graphql.Field{
				Type:        graphql.NewList(userObject),
				Description: "get users data",
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					log.Println("resolve func was called!")
					var err error = nil
					users := []model.User{}

					offset, success := p.Args["offset"].(int)
					if !success {
						offset = 0
					}
					limit, success := p.Args["limit"].(int)
					if !success {
						limit = 10
					}

					db.Limit(limit).Offset(offset).Find(&users)
					return &users, err
				},
			},
		},
	}

	m := graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type:        userObject,
				Description: "Create New User",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"bio": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"url_avatar": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var err error = nil
					name, _ := p.Args["name"].(string)
					email, _ := p.Args["email"].(string)
					bio, _ := p.Args["bio"].(string)
					urlAvatar, _ := p.Args["url_avatar"].(string)

					newUser := model.User{
						Name:      name,
						Email:     email,
						Bio:       bio,
						UrlAvatar: urlAvatar,
					}

					db.Create(&newUser)

					return &newUser, err
				},
			},
			"updateUser": &graphql.Field{
				Type:        userObject,
				Description: "Update New User",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"bio": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"url_avatar": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var err error = nil
					user := model.User{}
					user_id, success := p.Args["id"].(int)
					if !success {
						err = errors.New("Incorrect Param was set")
						return &user, err
					}

					if db.First(&user, user_id).RecordNotFound() {
						err = errors.New("record was not found")
					} else {
						name, _ := p.Args["name"].(string)
						email, _ := p.Args["email"].(string)
						bio, _ := p.Args["bio"].(string)
						urlAvatar, _ := p.Args["url_avatar"].(string)

						user.Name = name
						user.Email = email
						user.Bio = bio
						user.UrlAvatar = urlAvatar

						db.Save(&user)
					}
					return &user, err
				},
			},
			"deleteUser": &graphql.Field{
				Type:        userObject,
				Description: "Delete New User",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var err error = nil
					user := model.User{}
					user_id, success := p.Args["id"].(int)
					if !success {
						err = errors.New("Incorrect Param was set")
						return &user, err
					}

					if db.First(&user, user_id).RecordNotFound() {
						err = errors.New("record was not found")
					} else {
						db.Delete(&user)
					}
					return &user, err
				},
			},
		},
	}

	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(q),
		Mutation: graphql.NewObject(m),
	}

	return schemaConfig
}
func GetSchema(sc graphql.SchemaConfig) graphql.Schema {
	schema, err := graphql.NewSchema(sc)
	if err != nil {
		log.Fatalln(err)
	}

	return schema
}
func resolveID(p graphql.ResolveParams) (interface{}, error) {
	return p.Args["id"], nil
}
