package stripe_card

import (
    "context"
    "fmt"
    "github.com/hasura/go-graphql-client"
    "github.com/JamesArthurHolland/thumbin/main-go-graph-lib/main_graph_ql"
)

type Service struct {
    client *graphql.Client
}

type StripeCardAllQuery struct {
    Result []StripeCardEntity `graphql:"stripeCardAll"`
}
type StripeCardByIdQuery struct {
    Result StripeCardEntity `graphql:"stripeCard(id: $id)"`
}
type CreateStripeCardQuery struct {
    Result StripeCardEntity `graphql:"createStripeCard(input: $input)"`
}
type UpdateStripeCardQuery struct {
    Result StripeCardEntity `graphql:"updateStripeCard(id: $id, input: $input)"`
}

type StripeCardByOwnerIdQuery struct {
    Result StripeCardEntity `graphql:"stripeCardByOwnerId(ownerId: $ownerId)"`
}


func NewService(client main_graph_ql.Client) *Service {
    return &Service{
        client: client,
    }
}

func(s *Service) All() ([]*StripeCardEntity, error) {
    var query StripeCardAllQuery
    err := s.client.Query(context.Background(), &query, nil)
    if(err != nil) {
        return nil, err
    }

    var pointerSlice []*StripeCardEntity
    for i := 0; i < len(query.Result); i++ {
    	pointerSlice = append(pointerSlice, &query.Result[i])
    }


    return pointerSlice, nil
}

func(s *Service) Fetch(id string) (*StripeCardEntity, error) {
    variables := map[string]interface{}{
        "stripeCardId": graphql.String(id),
    }

    var query StripeCardByIdQuery
    err := s.client.Query(context.Background(), &query, variables)
    if(err != nil) {
        return nil, err
    }

    return &query.Result, nil
}

func(s *Service) FetchByOwnerId(owner_id string) (*StripeCardEntity, error) {
    variables := map[string]interface{}{
        "ownerId": graphql.String(owner_id),
    }

    var query StripeCardByOwnerIdQuery
    err := s.client.Query(context.Background(), &query, variables)
    if(err != nil) {
        return nil, err
    }

    return &query.Result, err
}


func(s *Service) Create(input *StripeCardInputType) (*StripeCardEntity, error) {
	variables := map[string]interface{}{
		"input": input,
	}

	var query CreateStripeCardQuery
	err := s.client.Mutate(context.TODO(), &query, variables)
	if(err != nil) {
		return nil, err
	}

	return &query.Result, nil
}

func(s *Service) Update(inputEntity *StripeCardEntity) (*StripeCardEntity, error) {
    input := NewInputType(inputEntity)

	variables := map[string]interface{}{
	    "id": graphql.String(input.Id),
		"input": input,
	}

	var query UpdateStripeCardQuery
    fmt.Println("vars:")
    fmt.Println(variables)
	err := s.client.Mutate(context.Background(), &query, variables)
    if(err != nil) {
        fmt.Println("here3-----")
        return nil, err
    }

	return &query.Result, nil
}

