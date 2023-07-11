<div align="center">
    <h3>edgeconfig</h3>
    <br/>
    Go library to interact with <a href="https://vercel.com/docs/storage/edge-config" target="_blank">Vercel Edge Config</a>.
    <br/>
    <br/>
    <a href="https://pkg.go.dev/github.com/thatmattlove/go-vercel-edge-config">
        <img alt="Docs" src="https://img.shields.io/badge/godoc-reference-007D9C.svg?style=for-the-badge">
    </a>
    <a href="https://github.com/thatmattlove/go-vercel-edge-config/actions/workflows/test.yml">
        <img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/thatmattlove/go-vercel-edge-config/test?style=for-the-badge">
    </a>
</div>

# Installation

```
go get -d github.com/thatmattlove/go-vercel-edge-config
```

# Usage

## Example Edge Config Store
```json
{
  "key": "value"
}
```

## Initialize the Client
```go
options := &edgeconfig.ClientOptions{
    EdgeConfigToken: "your-edge-config-token",
    EdgeConfigID: "your-edge-config-id",
}
ec, err := edgeconfig.New(options)

// You can also initialize a client from a connection string:
ec, err := edgeconfig.NewFromConnectionString("https://edge-config.vercel.com/your_edge_config_id_here?token=your_edge_config_read_access_token_here")
```

## Retrieve a Single Value by Key

```go
ec.Item("key")
// value
```

## Retrieve All Items as a Map

```go
ec.Items()
// map[key:value]
```

## Retrieve Edge Config Digest

```go
ec.Digest()
// d53d0989fcf40382f6979a1cb069e70b477606a156cf8dc6c96a7154b6e79a29
```

## Retrieve All Edge Configs

```go
err := ec.API.Authenticate("your_vercel_api_token")
configs, err := ec.API.ListAllEdgeConfigs()
// returns a slice of edge configs
```

![GitHub](https://img.shields.io/github/license/thatmattlove/go-vercel-edge-config?color=000&style=for-the-badge)