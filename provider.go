package main_graphql_lib

import (
    "go.uber.org/dig"
    "main-go-graph-lib/bundle"
    "main-go-graph-lib/youtube_channel"
    "main-go-graph-lib/organisation"
    "main-go-graph-lib/user"
)

func BuildContainer(container *dig.Container) *dig.Container {
	container.Provide(bundle.NewService)
	container.Provide(youtube_channel.NewService)
	container.Provide(organisation.NewService)
	container.Provide(user.NewService)

	return container
}
