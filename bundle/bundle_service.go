package bundle

import (
    "context"
    "fmt"
    "github.com/hasura/go-graphql-client"
    "github.com/JamesArthurHolland/thumbin/main-go-graph-lib/main_graph_ql"
)

type Service struct {
    client *graphql.Client
}

type BundleAllQuery struct {
    Result []BundleEntity `graphql:"bundleAll"`
}
type BundleByIdQuery struct {
    Result BundleEntity `graphql:"bundle(id: $id)"`
}
type CreateBundleQuery struct {
    Result BundleEntity `graphql:"createBundle(input: $input)"`
}
type UpdateBundleQuery struct {
    Result BundleEntity `graphql:"updateBundle(id: $id, input: $input)"`
}


func NewService(client main_graph_ql.Client) *Service {
    return &Service{
        client: client,
    }
}

func(s *Service) All() ([]*BundleEntity, error) {
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

func(s *Service) Fetch(id string) (*BundleEntity, error) {
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


func(s *Service) Create(input *BundleInputType) (*BundleEntity, error) {
	variables := map[string]interface{}{
		"input": input,
	}

	var query CreateBundleQuery
	err := s.client.Mutate(context.Background(), &query, variables)
	if(err != nil) {
		return nil, err
	}

	return &query.Result, nil
}

func(s *Service) Update(inputEntity *BundleEntity) (*BundleEntity, error) {
    input := NewInputType(inputEntity)

	variables := map[string]interface{}{
	    "id": graphql.String(input.Id),
		"input": input,
	}

	var query UpdateBundleQuery
    fmt.Println("vars:")
    fmt.Println(variables)
	err := s.client.Mutate(context.Background(), &query, variables)
    if(err != nil) {
        fmt.Println("here3-----")
        return nil, err
    }

	return &query.Result, nil
}

