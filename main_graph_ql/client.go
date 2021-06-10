package main_graph_ql

import (
    "fmt"
    "github.com/hasura/go-graphql-client"
    "log"
    "os"
)

type DataPersistenceHost string
type DataPersistencePort string
type DataPersistenceUrl string

type Client *graphql.Client

func CreateMainGraphQLClient(url DataPersistenceUrl) Client {
    return graphql.NewClient(string(url), nil)
}

func GetDataPersistenceUrl(host DataPersistenceHost, port DataPersistencePort) DataPersistenceUrl {
    return DataPersistenceUrl(fmt.Sprintf("http://%s:%s/api/graphql", host, port))
}

func GetPersistenceUrlParts() (DataPersistenceHost, DataPersistencePort) {
    portString := os.Getenv("DP_MAIN_PORT")
    if portString == "" {
        log.Fatalln("DP_MAIN_PORT env var not set.")
    }

    hostString := os.Getenv("DP_MAIN_HOST")
    if hostString == "" {
        log.Fatalln("DP_MAIN_HOST env var not set.")
    }

    return DataPersistenceHost(hostString), DataPersistencePort(portString)
}