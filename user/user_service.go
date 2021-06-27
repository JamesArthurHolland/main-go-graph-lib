package user

import (
    "context"
    "fmt"
    "github.com/hasura/go-graphql-client"
    "github.com/JamesArthurHolland/thumbin/main-go-graph-lib/main_graph_ql"
)

type Service struct {
    client *graphql.Client
}

type UserAllQuery struct {
    Result []UserEntity `graphql:"userAll"`
}
type UserByIdQuery struct {
    Result UserEntity `graphql:"user(id: $id)"`
}
type CreateUserQuery struct {
    Result UserEntity `graphql:"createUser(input: $input)"`
}
type UpdateUserQuery struct {
    Result UserEntity `graphql:"updateUser(id: $id, input: $input)"`
}

type UserByEmailQuery struct {
    Result UserEntity `graphql:"userByEmail(email: $email)"`
}


func NewService(client main_graph_ql.Client) *Service {
    return &Service{
        client: client,
    }
}

func(s *Service) All() ([]*UserEntity, error) {
    var query UserAllQuery
    err := s.client.Query(context.Background(), &query, nil)
    if(err != nil) {
        return nil, err
    }

    var pointerSlice []*UserEntity
    for i := 0; i < len(query.Result); i++ {
    	pointerSlice = append(pointerSlice, &query.Result[i])
    }


    return pointerSlice, nil
}

func(s *Service) Fetch(id string) (*UserEntity, error) {
    variables := map[string]interface{}{
        "userId": graphql.String(id),
    }

    var query UserByIdQuery
    err := s.client.Query(context.Background(), &query, variables)
    if(err != nil) {
        return nil, err
    }

    return &query.Result, nil
}

func(s *Service) FetchByEmail(email string) (*UserEntity, error) {
    variables := map[string]interface{}{
        "email": graphql.String(email),
    }

    var query UserByEmailQuery
    err := s.client.Query(context.Background(), &query, variables)
    if(err != nil) {
        return nil, err
    }

    return &query.Result, err
}


func(s *Service) Create(input *UserInputType) (*UserEntity, error) {
	variables := map[string]interface{}{
		"input": input,
	}

	var query CreateUserQuery
	err := s.client.Mutate(context.Background(), &query, variables)
	if(err != nil) {
		return nil, err
	}

	return &query.Result, nil
}

func(s *Service) Update(inputEntity *UserEntity) (*UserEntity, error) {
    input := NewInputType(inputEntity)

	variables := map[string]interface{}{
	    "id": graphql.String(input.Id),
		"input": input,
	}

	var query UpdateUserQuery
    fmt.Println("vars:")
    fmt.Println(variables)
	json, err := s.client.MutateRaw(context.Background(), &query, variables)
    if(err != nil) {
        fmt.Println("here3-----")
        return nil, err
    }
    fmt.Println("here4-----")
    fmt.Println(string(*json))

	return &query.Result, nil
}

