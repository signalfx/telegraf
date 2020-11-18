# SignalFx Output Plugin

```toml
    ## SignalFx API Token
    APIToken = "my-secret-key" # required.

    ## Ingest URL
    DatapointIngestURL = "https://ingest.signalfx.com/v2/datapoint"
    EventIngestURL = "https://ingest.signalfx.com/v2/event"

    ## Exclude metrics by metric name
    Exclude = ["plugin.metric_name", "plugin2.metric_name"]

    ## Events or String typed metrics are omitted by default,
    ## with the exception of host property events which are emitted by 
    ## the SignalFx Metadata Plugin.  If you require a string typed metric
    ## you must specify the metric name in the following list
    Include = ["plugin.metric_name", "plugin2.metric_name"]

    ## When a Telegraf metric type is present, the plugin will map it to the
    ## closest SignalFx metric type (gauge, count, or counter).  However, many
    ## Telegraf inputs leave the metrics as untyped, so this plugin then
    ## defaults to gauge.  If a given metric is a gauge, then all is well.
    ## However, if the metric is either a counter or a cumulative counter,
    ## then you can set it with the options below.

    ## Metrics that should be treated as a SignalFx cumulative counter.
    Counter = ["plugin.metric_name, "plugin2.metric_name"]

    ## Metrics that should be treated as a SignalFx counter.
    Count = ["plugin.metric_name, "plugin2.metric_name"]
```