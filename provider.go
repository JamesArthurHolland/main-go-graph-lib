package main_graphql_lib

import (
    "go.uber.org/dig"
    "main-go-graph-lib/bundle"
    "main-go-graph-lib/youtube_channel"
    "main-go-graph-lib/user"
    "main-go-graph-lib/organisation"
)

func BuildContainer(container *dig.Container) *dig.Container {
	container.Provide(bundle.NewService)
	container.Provide(youtube_channel.NewService)
	container.Provide(user.NewService)
	container.Provide(organisation.NewService)

	return container
}
