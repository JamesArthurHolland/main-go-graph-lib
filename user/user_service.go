package user

import (
    "context"
    "github.com/hasura/go-graphql-client"
    "main-go-graph-lib/main_graph_ql"
)

type UserService struct {
    client *graphql.Client
}

type UserAllQuery struct {
    Result []UserEntity `graphql:"userAll"`
}
type UserByIdQuery struct {
    Result UserEntity `graphql:"user(id: $id)"`
}
type UserByEmailQuery struct {
    Result UserEntity `graphql:"userByEmail(email: $email)"`
}


func NewUserService(client main_graph_ql.Client) *UserService {
    return &UserService{
        client: client,
    }
}

func(s *UserService) All() ([]*UserEntity, error) {
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

func(s *UserService) Fetch(id string) (*UserEntity, error) {
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

func(s *UserService) FetchByEmail(email string) (*UserEntity, error) {
    variables := map[string]interface{}{
        "userById": graphql.String(email),
    }

    var query UserByEmailQuery
    err := s.client.Query(context.Background(), &query, variables)
    if(err != nil) {
        return nil, err
    }

    return &query.Result, err
}

