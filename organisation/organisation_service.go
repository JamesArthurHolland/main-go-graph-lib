package organisation

import (
    "context"
    "fmt"
    "github.com/hasura/go-graphql-client"
    "main-go-graph-lib/main_graph_ql"
)

type Service struct {
    client *graphql.Client
}

type OrganisationAllQuery struct {
    Result []OrganisationEntity `graphql:"organisationAll"`
}
type OrganisationByIdQuery struct {
    Result OrganisationEntity `graphql:"organisation(id: $id)"`
}
type CreateOrganisationQuery struct {
    Result OrganisationEntity `graphql:"createOrganisation(input: $input)"`
}
type UpdateOrganisationQuery struct {
    Result OrganisationEntity `graphql:"updateOrganisation(id: $id, input: $input)"`
}


func NewService(client main_graph_ql.Client) *Service {
    return &Service{
        client: client,
    }
}

func(s *Service) All() ([]*OrganisationEntity, error) {
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

func(s *Service) Fetch(id string) (*OrganisationEntity, error) {
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


func(s *Service) Create(input *OrganisationInputType) (*OrganisationEntity, error) {
	variables := map[string]interface{}{
		"input": input,
	}

	var query CreateOrganisationQuery
	err := s.client.Mutate(context.Background(), &query, variables)
	if(err != nil) {
		return nil, err
	}

	return &query.Result, nil
}

func(s *Service) Update(inputEntity *OrganisationEntity) (*OrganisationEntity, error) {
    input := NewInputType(inputEntity)

	variables := map[string]interface{}{
	    "id": graphql.String(input.Id),
		"input": input,
	}

	var query UpdateOrganisationQuery
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

