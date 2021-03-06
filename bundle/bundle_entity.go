package bundle

import "github.com/hasura/go-graphql-client"

type BundleEntity struct {
    Id graphql.String `json:"id"`
    ThumbAmount graphql.Int `json:"thumb_amount"`
    CostEur graphql.Int `json:"cost_eur"`
    LockVersion int64   `json:"lock_version"`
}

type BundleInputType struct {
    Id         graphql.String    `json:"id"`
    ThumbAmount         graphql.Int    `json:"thumb_amount"`
    CostEur         graphql.Int    `json:"cost_eur"`
    LockVersion int64   `json:"lock_version"`
}

func NewInputType(entity *BundleEntity) *BundleInputType {
    return &BundleInputType{
        Id: entity.Id,
        ThumbAmount: entity.ThumbAmount,
        CostEur: entity.CostEur,
        LockVersion: entity.LockVersion,
    }
}
