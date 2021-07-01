package stripe_card

import "github.com/hasura/go-graphql-client"

type StripeCardEntity struct {
    Id graphql.String `json:"id"`
    OwnerId graphql.String `json:"owner_id"`
    LastFourDigits graphql.String `json:"last_four_digits"`
    LockVersion int64   `json:"lock_version"`
}

type StripeCardInputType struct {
    Id         graphql.String    `json:"id"`
    OwnerId         graphql.String    `json:"owner_id"`
    LastFourDigits         graphql.String    `json:"last_four_digits"`
    LockVersion int64   `json:"lock_version"`
}

func NewInputType(entity *StripeCardEntity) *StripeCardInputType {
    return &StripeCardInputType{
        Id: entity.Id,
        OwnerId: entity.OwnerId,
        LastFourDigits: entity.LastFourDigits,
        LockVersion: entity.LockVersion,
    }
}
