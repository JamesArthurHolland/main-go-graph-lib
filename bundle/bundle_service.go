package bundle

import (
    "context"
    "github.com/hasura/go-graphql-client"
    "main-go-graph-lib/main_graph_ql"
)

type BundleService struct {
    client *graphql.Client
}

type BundleAllQuery struct {
    Result []BundleEntity `graphql:"bundleAll"`
}
type BundleByIdQuery struct {
    Result BundleEntity `graphql:"bundle(id: $id)"`
}

func NewBundleService(client main_graph_ql.Client) *BundleService {
    return &BundleService{
        client: client,
    }
}

func(s *BundleService) All() ([]*BundleEntity, error) {
    var query BundleAllQuery
    err := s.client.Query(context.Background(), &query, nil)
    if(err != nil) {
        return nil, err
    }

    var pointerSlice []*BundleEntity
    for i := 0; i < len(query.Result); i++ {
    	pointerSlice = append(pointerSlice, &query.Result[i])
    }


    return pointerSlice, nil
}

func(s *BundleService) Fetch(id string) (*BundleEntity, error) {
    variables := map[string]interface{}{
        "bundleId": graphql.String(id),
    }

    var query BundleByIdQuery
    err := s.client.Query(context.Background(), &query, variables)
    if(err != nil) {
        return nil, err
    }

    return &query.Result, nil
}

