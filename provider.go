package main_graphql_lib

import (
    "go.uber.org/dig"
    "github.com/JamesArthurHolland/thumbin/main-go-graph-lib/bundle"
    "github.com/JamesArthurHolland/thumbin/main-go-graph-lib/youtube_channel"
    "github.com/JamesArthurHolland/thumbin/main-go-graph-lib/stripe_card"
    "github.com/JamesArthurHolland/thumbin/main-go-graph-lib/organisation"
    "github.com/JamesArthurHolland/thumbin/main-go-graph-lib/user"
)

func BuildContainer(container *dig.Container) *dig.Container {
	container.Provide(bundle.NewService)
	container.Provide(youtube_channel.NewService)
	container.Provide(stripe_card.NewService)
	container.Provide(organisation.NewService)
	container.Provide(user.NewService)

	return container
}
