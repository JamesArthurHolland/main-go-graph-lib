package bundle

import "github.com/hasura/go-graphql-client"

type BundleEntity struct {
    Id graphql.String `json:"id"`
    ThumbAmount graphql.Int `json:"thumb_amount"`
    CostEur graphql.Int `json:"cost_eur"`
}
