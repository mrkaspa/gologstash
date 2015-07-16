# gologstash
Go HTTP client for the Logstash server

Just get

```
  go get "github.com/mrkaspa/gologstash"
```

Just import

```go
  import "github.com/mrkaspa/gologstash"
```

Just call log

```go
  var client = golog.Client{URL: "http://localhost:9090", Host: "localhost"}
  client.PostLog("this is a log")
```

Or if you prefer Just send a map

```go
  var client = golog.Client{URL: "http://localhost:9090", Host: "localhost"}
  // The message key is important for Logstash
  data := map[string]string{"message": "demo message", "add": "I'm additional info"}
  client.PostMap(data)
```

And if you need Basic Auth Just set the parameters on the client

```go
  var client = golog.Client{URL: "http://localhost:9090", Host: "localhost", User:"root", Passwor:"top_secret"}
```
