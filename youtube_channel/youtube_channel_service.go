package youtube_channel

import (
    "context"
    "fmt"
    "github.com/hasura/go-graphql-client"
    "main-go-graph-lib/main_graph_ql"
)

type Service struct {
    client *graphql.Client
}

type YoutubeChannelAllQuery struct {
    Result []YoutubeChannelEntity `graphql:"youtubeChannelAll"`
}
type YoutubeChannelByIdQuery struct {
    Result YoutubeChannelEntity `graphql:"youtubeChannel(id: $id)"`
}
type CreateYoutubeChannelQuery struct {
    Result YoutubeChannelEntity `graphql:"createYoutubeChannel(input: $input)"`
}
type UpdateYoutubeChannelQuery struct {
    Result YoutubeChannelEntity `graphql:"updateYoutubeChannel(id: $id, input: $input)"`
}

type YoutubeChannelByChannelNameQuery struct {
    Result YoutubeChannelEntity `graphql:"youtubeChannelByChannelName(channelName: $channelName)"`
}


func NewService(client main_graph_ql.Client) *Service {
    return &Service{
        client: client,
    }
}

func(s *Service) All() ([]*YoutubeChannelEntity, error) {
    var query YoutubeChannelAllQuery
    err := s.client.Query(context.Background(), &query, nil)
    if(err != nil) {
        return nil, err
    }

    var pointerSlice []*YoutubeChannelEntity
    for i := 0; i < len(query.Result); i++ {
    	pointerSlice = append(pointerSlice, &query.Result[i])
    }


    return pointerSlice, nil
}

func(s *Service) Fetch(id string) (*YoutubeChannelEntity, error) {
    variables := map[string]interface{}{
        "youtubeChannelId": graphql.String(id),
    }

    var query YoutubeChannelByIdQuery
    err := s.client.Query(context.Background(), &query, variables)
    if(err != nil) {
        return nil, err
    }

    return &query.Result, nil
}

func(s *Service) FetchByChannelName(channel_name string) (*YoutubeChannelEntity, error) {
    variables := map[string]interface{}{
        "channelName": graphql.String(channel_name),
    }

    var query YoutubeChannelByChannelNameQuery
    err := s.client.Query(context.Background(), &query, variables)
    if(err != nil) {
        return nil, err
    }

    return &query.Result, err
}


func(s *Service) Create(input *YoutubeChannelInputType) (*YoutubeChannelEntity, error) {
	variables := map[string]interface{}{
		"input": input,
	}

	var query CreateYoutubeChannelQuery
	err := s.client.Mutate(context.Background(), &query, variables)
	if(err != nil) {
		return nil, err
	}

	return &query.Result, nil
}

func(s *Service) Update(inputEntity *YoutubeChannelEntity) (*YoutubeChannelEntity, error) {
    input := NewInputType(inputEntity)

	variables := map[string]interface{}{
	    "id": graphql.String(input.Id),
		"input": input,
	}

	var query UpdateYoutubeChannelQuery
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

