# 8.4.0-alpha.1

This prerelease introduces a new typed API generated from the [elasticsearch-specification](https://github.com/elastic/elasticsearch-specification). This generation from the common specification allows us to provide a complete API which uses an exhaustive hierarchy of types  reflecting the possibilities given by Elasticsearch. 

This new API is the next iteration of the Go client for Elasticsearch, it now lives alongside the existing API, it is in `alpha` state and will gain features over time and releases.

## What's new

The `TypedClient` is built around a fluent builder for easier request creation and a collection of structures and helpers that mimics as closely as possible the Elasticsearch JSON API.

As a first example, here is a search request:
```go
cfg := elasticsearch.Config{
	// Define your configuration
}
es, _ := elasticsearch.NewTypedClient(cfg)
res, err := es.Search().
    Index("index_name").
    Request(&search.Request{
        Query: &types.QueryContainer{
            Match: map[types.Field]types.MatchQuery{
                "name": {Query: "Foo"},
            },
        },
    },
    ).Do(context.Background())
```

The `Request` uses the structures found in the `typedapi/types` package which will lead you along the possibilities. A builder for each structure that allows easier access and declaration is also provided.

More on the specifics and a few examples of standard use-cases can be found in the [TypedAPI section of the documentation](https://www.elastic.co/guide/en/elasticsearch/client/go-api/master/typedapi.html).

## Limitations

While most of the endpoints are covered, a few points are still being worked on and will be part of future releases:

* NDJSON endpoints: `bulk`, `msearch`, `msearch_template`, `ML.post_data`, `find_structure`, to name a few.
* Response and Errors structures with deserialization.

## Transport & config

While being different, the new API uses all the existing layers that were built so far, `elastic-transport-go` remains the preferred transport and all your configuration and credentials applies, same as before.

## Feedback

Feedback is very welcome, play with it, use it, let us know what you think!

# 8.3.0

## API

* `ML.InferTrainedModelDeployment` renamed to `InferTrainedModel`
* `ML.PreviewDatafeed` has two new parameters, `start` and `end`. [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/8.3/ml-preview-datafeed.html)
* `ML.StartTrainedModelDeployment` has three new parameters, `number_of_allocations`, `threads_per_allocation` and `queue_capacity`. [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/start-trained-model-deployment.html)
* `Cluster.DeleteVotingConfigExclusions` has a new `master_timeout` parameter.
* `Cluster.PostVotingConfigExclusions` has a new `master_timeout` parameter.
* `Snapshot.Get` has a new `index_names` parameters (boolean). Whether to include the name of each index in the snapshot. Defaults to true.

**New APIs**

* `Security.HasPrivilegesUserProfile` (Experimental API) [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-has-privileges-user-profile.html)

# 8.2.0
## Client

* Fixed a serialisation error for `retry_on_conflict` in the BulkIndexer. Thanks to @lpflpf for the help!
* Fixed a concurrent map error in the BulkIndexer when custom headers are applied. Thanks to @chzhuo for the contribution! 

## API

**New APIs**

* `Cat.ComponentTemplates`
* `ML.GetMemoryStats` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/get-ml-memory.html)

* `Security.activateUserProfile` (Experimental API) [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-activate-user-profile.html)
* `Security.disableUserProfile` (Experimental API) [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/security-api-disable-user-profile.html)
* `Security.enableUserProfile` (Experimental API) [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/security-api-enable-user-profile.html)
* `Security.getUserProfile` (Experimental API) [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user-profile.html)
* `Security.suggestUserProfiles` (Experimental API) [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/security-api-suggest-user-profile.html)
* `Security.updateUserProfileData` (Experimental API) [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-user-profile-data.html)

# 8.1.0
## API

 * API is generated from the Elasticsearch 8.1.0 specification.

**New parameters**

* `WithWaitForCompletion` for `Indices.Forcemerge` 
* `WithFeatures` for `Indices.Get`
* `WithForce` for `ML.DeleteTrainedModel`

**New APIs**

* `OidcAuthenticate`, `OidcLogout` and `OidcPrepareAuthentication` [see documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api.html#security-openid-apis)
* `TransformResetTransform`

# 8.0.0
## Client

 * The client now uses `elastic-transport-go` dependency which lives in its [own repository](https://github.com/elastic/elastic-transport-go/).
 * With the knewly extracted transport, the `retryOnTimeout` has been replaced with a `retryOnError` callback. This allows to select more finely which error should be retried by the client.
 * `BulkIndexerItem` `Body` field is now an `io.ReadSeeker` allowing reread without increasing memory consumption.
 * `BulkIndexerItem` know correctly uses the `routing` property instead of the deprecated `_routing`.

## API

 * API is generated from the Elasticsearch 8.0.0 specification.




