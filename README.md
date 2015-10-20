## Sensu Golang API library ##

This library is an implementation of Sensu 0.20's REST API for Golang.

## Example ##

```go
// Create a new configration struct
config := sensu.DefaultConfig()
config.Address = "<your_sensu_API_server:port>"

// Create a new API Client
sensuAPI, err := sensu.NewAPIClient(config)
if err != nil {
	// catch errors in client creation
}

// Get All clients that Sensu API server knows about
var clients []sensu.Client
_, e := sensuAPI.getClients(&data)

fmt.Printf(First Client: %+v\n", clients[0])
```

## Documentation ##

https://gowalker.org/github.com/jefflaplante/sensulib

## License ##

GPLv3
