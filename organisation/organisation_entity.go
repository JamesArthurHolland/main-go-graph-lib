package organisation

import "github.com/hasura/go-graphql-client"

type OrganisationEntity struct {
    Id graphql.String `json:"id"`
    Email graphql.String `json:"email"`
    Name graphql.String `json:"name"`
    Gravatar graphql.String `json:"gravatar"`
    PreferredCurrency graphql.String `json:"preferred_currency"`
    LockVersion int64   `json:"lock_version"`
}

type OrganisationInputType struct {
    Id         graphql.String    `json:"id"`
    Email         graphql.String    `json:"email"`
    Name         graphql.String    `json:"name"`
    Gravatar         graphql.String    `json:"gravatar"`
    PreferredCurrency         graphql.String    `json:"preferred_currency"`
    LockVersion int64   `json:"lock_version"`
}

func NewInputType(entity *OrganisationEntity) *OrganisationInputType {
    return &OrganisationInputType{
        Id: entity.Id,
        Email: entity.Email,
        Name: entity.Name,
        Gravatar: entity.Gravatar,
        PreferredCurrency: entity.PreferredCurrency,
        LockVersion: entity.LockVersion,
    }
}
