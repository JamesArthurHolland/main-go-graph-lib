package main_graphql_lib

import (
    "go.uber.org/dig"
    "main-go-graph-lib/bundle"
    "main-go-graph-lib/organisation"
    "main-go-graph-lib/user"
)

func BuildContainer(container *dig.Container) *dig.Container {
	container.Provide(bundle.NewBundleService)
	container.Provide(organisation.NewOrganisationService)
	container.Provide(user.NewUserService)

	return container
}
