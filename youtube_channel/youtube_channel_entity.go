package youtube_channel

import "github.com/hasura/go-graphql-client"

type YoutubeChannelEntity struct {
    Id graphql.String `json:"id"`
    OwnerId graphql.String `json:"owner_id"`
    ChannelName graphql.String `json:"channel_name"`
    Verified graphql.Boolean `json:"verified"`
}

type YoutubeChannelInputType struct {
    Id         graphql.String    `json:"id"`
    OwnerId         graphql.String    `json:"owner_id"`
    ChannelName         graphql.String    `json:"channel_name"`
    Verified         graphql.Boolean    `json:"verified"`
}

func NewInputType(entity *YoutubeChannelEntity) *YoutubeChannelInputType {
    return &YoutubeChannelInputType{
        Id: entity.Id,
        OwnerId: entity.OwnerId,
        ChannelName: entity.ChannelName,
        Verified: entity.Verified,
    }
}
