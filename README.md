Kami GoRelic
=======

NewRelic middleware for kami framework.

## Usage

```go
import(
        "github.com/david4shure/kamigorelic"
        "github.com/guregu/kami"
)

func main(){

     kamigorelic.InitNewrelicAgent("YOUR_NEWRELIC_LICENSE_KEY", "YOUR_APPLICATION_NAME", true)
     kami.Use("/", kamigorelic.Handler)

     kami.Serve()
}
```

## Authors

* [David Shure](http://github.com/david4shure)
