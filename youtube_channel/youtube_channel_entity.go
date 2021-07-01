package youtube_channel

import "github.com/hasura/go-graphql-client"

type YoutubeChannelEntity struct {
    Id graphql.String `json:"id"`
    OwnerId graphql.String `json:"owner_id"`
    ChannelName graphql.String `json:"channel_name"`
    Verified graphql.Boolean `json:"verified"`
    LockVersion int64   `json:"lock_version"`
}

type YoutubeChannelInputType struct {
    Id         graphql.String    `json:"id"`
    OwnerId         graphql.String    `json:"owner_id"`
    ChannelName         graphql.String    `json:"channel_name"`
    Verified         graphql.Boolean    `json:"verified"`
    LockVersion int64   `json:"lock_version"`
}

func NewInputType(entity *YoutubeChannelEntity) *YoutubeChannelInputType {
    return &YoutubeChannelInputType{
        Id: entity.Id,
        OwnerId: entity.OwnerId,
        ChannelName: entity.ChannelName,
        Verified: entity.Verified,
        LockVersion: entity.LockVersion,
    }
}
