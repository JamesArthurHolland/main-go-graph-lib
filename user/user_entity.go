package user

import "github.com/hasura/go-graphql-client"

type UserEntity struct {
    Id graphql.String `json:"id"`
    Email graphql.String `json:"email"`
    FirstName graphql.String `json:"first_name"`
    LastName graphql.String `json:"last_name"`
    Gravatar graphql.String `json:"gravatar"`
    PreferredCurrency graphql.String `json:"preferred_currency"`
    StripeAccountId graphql.String `json:"stripe_account_id"`
    StripeCustomerId graphql.String `json:"stripe_customer_id"`
    StripeDefaultCard graphql.String `json:"stripe_default_card"`
    Confirmed graphql.Boolean `json:"confirmed"`
    LockVersion int64   `json:"lock_version"`
}

type UserInputType struct {
    Id         graphql.String    `json:"id"`
    Email         graphql.String    `json:"email"`
    FirstName         graphql.String    `json:"first_name"`
    LastName         graphql.String    `json:"last_name"`
    Gravatar         graphql.String    `json:"gravatar"`
    PreferredCurrency         graphql.String    `json:"preferred_currency"`
    StripeAccountId         graphql.String    `json:"stripe_account_id"`
    StripeCustomerId         graphql.String    `json:"stripe_customer_id"`
    StripeDefaultCard         graphql.String    `json:"stripe_default_card"`
    Confirmed         graphql.Boolean    `json:"confirmed"`
    LockVersion int64   `json:"lock_version"`
}

func NewInputType(entity *UserEntity) *UserInputType {
    return &UserInputType{
        Id: entity.Id,
        Email: entity.Email,
        FirstName: entity.FirstName,
        LastName: entity.LastName,
        Gravatar: entity.Gravatar,
        PreferredCurrency: entity.PreferredCurrency,
        StripeAccountId: entity.StripeAccountId,
        StripeCustomerId: entity.StripeCustomerId,
        StripeDefaultCard: entity.StripeDefaultCard,
        Confirmed: entity.Confirmed,
        LockVersion: entity.LockVersion,
    }
}
