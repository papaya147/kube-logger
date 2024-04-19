# kube-logger

A Go log client for Kubernetes clusters. Supported config file names by preference in the same path as execution:

- `kube-logger.yaml`
- `kube-logger.json`

### Configuration File Format

| Property                 | Description                                                        | Type       | Mandatory                           | Default Value | Accepted Values |
| ------------------------ | ------------------------------------------------------------------ | ---------- | ----------------------------------- | ------------- | --------------- |
| `namespaces`             | Array of namespaces to watch                                       | `[]string` | false                               | `default`     | any             |
| `pod_prefixes`           | Array of pods within namespace to watch                            | `[]string` | false                               | any           | any             |
| `cluster_provider`       | Provider of the EKS cluster to generate a Kubernetes clientset API | `string`   | true                                | none          | `eks`           |
| `eks.cluster_name`       | Name of the EKS cluster                                            | `string`   | true (if cluster provider is `eks`) | none          | any             |
| `eks.region`             | Region of the EKS cluster                                          | `string`   | true (if cluster provider is `eks`) | none          | any             |
| `eks.access_key`         | Access key with credentials to the EKS cluster                     | `string`   | true (if cluster provider is `eks`) | none          | any             |
| `eks.secret_key`         | Secret key with credentials to the EKS cluster                     | `string`   | true (if cluster provider is `eks`) | none          | any             |
| `console.active`         | Use console sink                                                   | `boolean`  | false                               | `false`       | `true`, `false` |
| `mongo.active`           | Use MongoDB sink                                                   | `boolean`  | false                               | `false`       | `true`, `false` |
| `mongo.connection_uri`   | MongoDB host endpoint with port and credentials                    | `string`   | true (if `mongo.active`)            | none          | any             |
| `mongo.database`         | MongoDB database name (need not exist already)                     | `string`   | true (if `mongo.active`)            | none          | any             |
| `mongo.collection`       | MongoDB collection name (need not exist already)                   | `string`   | true (if `mongo.active`)            | none          | any             |
| `elasticsearch.active`   | Use ElasticSearch sink                                             | `boolean`  | fasle                               | `false`       | `true`, `false` |
| `elasticsearch.host`     | ElasticSearch host endpoint with port                              | `string`   | true (if `elasticsearch.active`)    | none          | any             |
| `elasticsearch.username` | ElasticSearch username                                             | `string`   | true (if `elasticsearch.active`)    | none          | any             |
| `elasticsearch.password` | ElasticSearch password                                             | `string`   | true (if `elasticsearch.active`)    | none          | any             |
| `elasticsearch.index`    | ElasticSearch index name (need not exist already)                  | `string`   | true (if `elasticsearch.active`)    | none          | any             |

#### Example

```yaml
namespaces:
	- default
pod_prefixes:
	- auth
cluster_provider: eks
eks:
	cluster_name: backend-prod
	region: ap-south-1
	access_key: <AWS-ACCESS-KEY>
	secret_key: <AWS-SECRET-KEY>
console:
	active: true
mongo:
	active: true
	connection_uri: mongodb://localhost:27017
	database: logs
	collection: logs
elasticsearch:
	active: true
	host: http://localhost:9200
	username: elastic
	password: elastic
	index: logs
```

### Supported Providers

- AWS EKS

### Supported Log Sinks

- Console (Colored terminal output)
- MongoDB
- ElasticSearch
