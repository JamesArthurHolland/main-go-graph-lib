package organisation

import (
    "context"
    "github.com/hasura/go-graphql-client"
    "main-go-graph-lib/main_graph_ql"
)

type OrganisationService struct {
    client *graphql.Client
}

type OrganisationAllQuery struct {
    Result []OrganisationEntity `graphql:"organisationAll"`
}
type OrganisationByIdQuery struct {
    Result OrganisationEntity `graphql:"organisation(id: $id)"`
}

func NewOrganisationService(client main_graph_ql.Client) *OrganisationService {
    return &OrganisationService{
        client: client,
    }
}

func(s *OrganisationService) All() ([]*OrganisationEntity, error) {
    var query OrganisationAllQuery
    err := s.client.Query(context.Background(), &query, nil)
    if(err != nil) {
        return nil, err
    }

    var pointerSlice []*OrganisationEntity
    for i := 0; i < len(query.Result); i++ {
    	pointerSlice = append(pointerSlice, &query.Result[i])
    }


    return pointerSlice, nil
}

func(s *OrganisationService) Fetch(id string) (*OrganisationEntity, error) {
    variables := map[string]interface{}{
        "organisationId": graphql.String(id),
    }

    var query OrganisationByIdQuery
    err := s.client.Query(context.Background(), &query, variables)
    if(err != nil) {
        return nil, err
    }

    return &query.Result, nil
}

