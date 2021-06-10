package user

import "github.com/hasura/go-graphql-client"

type UserEntity struct {
    Id graphql.String `json:"id"`
    Email graphql.String `json:"email"`
    FirstName graphql.String `json:"first_name"`
    LastName graphql.String `json:"last_name"`
    Gravatar graphql.String `json:"gravatar"`
    PreferredCurrency graphql.String `json:"preferred_currency"`
    Confirmed graphql.Boolean `json:"confirmed"`
}

type UserInputType struct {
    Id         graphql.String    `json:"id"`
    Email         graphql.String    `json:"email"`
    FirstName         graphql.String    `json:"first_name"`
    LastName         graphql.String    `json:"last_name"`
    Gravatar         graphql.String    `json:"gravatar"`
    PreferredCurrency         graphql.String    `json:"preferred_currency"`
    Confirmed         graphql.Boolean    `json:"confirmed"`
}

func NewInputType(entity *UserEntity) *UserInputType {
    return &UserInputType{
        Id: entity.Id,
        Email: entity.Email,
        FirstName: entity.FirstName,
        LastName: entity.LastName,
        Gravatar: entity.Gravatar,
        PreferredCurrency: entity.PreferredCurrency,
        Confirmed: entity.Confirmed,
    }
}
