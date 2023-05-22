package httpservice

import (
	"context"
	"time"

	"github.com/docktermj/go-http/senzinghttpapi"
	ogenHttp "github.com/ogen-go/ogen/http"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// HttpServiceImpl is...
type HttpServiceImpl struct {
	senzinghttpapi.UnimplementedHandler
}

// ----------------------------------------------------------------------------
// internal methods
// ----------------------------------------------------------------------------

func (httpService *HttpServiceImpl) getOptSzLinks() senzinghttpapi.OptSzLinks {
	var result senzinghttpapi.OptSzLinks
	szLinks := senzinghttpapi.SzLinks{
		Self:                 senzinghttpapi.NewOptString("SelfBob"),
		OpenApiSpecification: senzinghttpapi.NewOptString("OpenApiSpecificationBob"),
	}
	result = senzinghttpapi.NewOptSzLinks(szLinks)
	return result
}

func (httpService *HttpServiceImpl) getOptSzMeta() senzinghttpapi.OptSzMeta {
	var result senzinghttpapi.OptSzMeta
	szMeta := senzinghttpapi.SzMeta{
		Server:                     senzinghttpapi.NewOptString("ServerBob"),
		HttpMethod:                 senzinghttpapi.NewOptSzHttpMethod(senzinghttpapi.SzHttpMethodGET),
		HttpStatusCode:             senzinghttpapi.NewOptInt(200),
		Timestamp:                  senzinghttpapi.NewOptDateTime(time.Now()),
		Version:                    senzinghttpapi.NewOptString("VersionBob"),
		RestApiVersion:             senzinghttpapi.NewOptString("RestApiVersionBob"),
		NativeApiVersion:           senzinghttpapi.NewOptString("NativeApiVersionBob"),
		NativeApiBuildVersion:      senzinghttpapi.NewOptString("NativeApiBuildVersionBob"),
		NativeApiBuildNumber:       senzinghttpapi.NewOptString("NativeApiBuildNumberBob"),
		NativeApiBuildDate:         senzinghttpapi.NewOptDateTime(time.Now()),
		ConfigCompatibilityVersion: senzinghttpapi.NewOptString("ConfigCompatibilityVersionBob"),
		Timings:                    senzinghttpapi.NewOptNilSzMetaTimings(map[string]int64{}),
	}
	result = senzinghttpapi.NewOptSzMeta(szMeta)
	return result
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The SaySomething method simply prints the 'Something' value in the type-struct.

Input
  - ctx: A context to control lifecycle.

Output
  - Nothing is returned, except for an error.  However, something is printed.
    See the example output.
*/
func (httpService *HttpServiceImpl) SaySomething(ctx context.Context) error {
	return nil
}

// AddDataSources implements addDataSources operation.
//
// Obtains the current default configuration, adds the specified data
// sources and sets the modified configuration as the new default
// configuration -- returning the set of all configured data sources.
// **NOTE:** This operation may not be allowed.  Some conditions that
// might cause this operation to be forbidden are:
// 1. The server does not have administrative functions enabled.
// 2. The server is running in "read-only" mode.
// 3. The server is started with a file-based configuration specified
// by `G2CONFIGFILE` property in the initialziation parameters.
// 4. The server is started with a specific configuration ID and
// therefore cannot modify the configuration and change to a new
// configuration.
//
// POST /data-sources
func (httpService *HttpServiceImpl) AddDataSources(ctx context.Context, req senzinghttpapi.AddDataSourcesReq, params senzinghttpapi.AddDataSourcesParams) (r senzinghttpapi.AddDataSourcesRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// AddRecord implements addRecord operation.
//
// This operation loads a single record using the data source identified by
// the data source code in the request path.  The record will be identified
// uniquely within the data source by the record ID provided in the request
// path.  The provided record in the request body is described in JSON
// using the [Senzing Generic Entity Specification](https://senzing.zendesk.
// com/hc/en-us/articles/231925448-Generic-Entity-Specification).
// The provided JSON record may omit the `RECORD_ID`, but if it contains a
// `RECORD_ID` then it **must** match the record ID provided on the request
// path.  The record ID is returned as part of the response.
// **NOTE:** The `withInfo` parameter will return the entity resolution
// info pertaining to the load.  This can be used to update a search index
// or external data mart.   Additionally, Senzing API Server provides a
// means to have the "raw" entity resolution info (from the underlying
// native Senzing API) automatically sent to a messaging service such as
// those provided by Amazon SQS, Rabbit MQ or Kafka regardless of the
// `withInfo` query parameter value.
//
// PUT /data-sources/{dataSourceCode}/records/{recordId}
func (httpService *HttpServiceImpl) AddRecord(ctx context.Context, req senzinghttpapi.AddRecordReq, params senzinghttpapi.AddRecordParams) (r senzinghttpapi.AddRecordRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// AddRecordWithReturnedRecordId implements addRecordWithReturnedRecordId operation.
//
// This operation loads a single record using the data source identified by
// the data source code in the request path.  The provided record in the
// request body is described in JSON using the
// [Senzing Generic Entity Specification](https://senzing.zendesk.
// com/hc/en-us/articles/231925448-Generic-Entity-Specification).
// The provided record may contain a `RECORD_ID` to identify it uniquely
// among other records in the same data source, but if it does not then a
// record ID will be automatically generated.  The record ID is returned
// as part of the response.
// **NOTE:** The `withInfo` parameter will return the entity resolution
// info pertaining to the load.  This can be used to update a search index
// or external data mart.   Additionally, Senzing API Server provides a
// means to have the "raw" entity resolution info (from the underlying
// native Senzing API) automatically sent to a messaging service such as
// those provided by Amazon SQS, Rabbit MQ or Kafka regardless of the
// `withInfo` query parameter value.
//
// POST /data-sources/{dataSourceCode}/records
func (httpService *HttpServiceImpl) AddRecordWithReturnedRecordId(ctx context.Context, req senzinghttpapi.AddRecordWithReturnedRecordIdReq, params senzinghttpapi.AddRecordWithReturnedRecordIdParams) (r senzinghttpapi.AddRecordWithReturnedRecordIdRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// AnalyzeBulkRecords implements analyzeBulkRecords operation.
//
// Provides a means to analyze a bulk data file of records prior to loading
// it.  The records are encoded as a JSON array of JSON objects, a single
// JSON object per line in JSON-lines file format, or as a CSV with one
// record per row.  The data should be in pre-mapped format using JSON
// property names or CSV column names as described by the
// [Senzing Generic Entity Specification](https://senzing.zendesk.
// com/hc/en-us/articles/231925448-Generic-Entity-Specification).
// **SCALABILITY NOTE:** This operation can be invoked in three ways.  In
// order of increasingly better scalability these are listed below:
// 1. Standard HTTP Request/Response
// 2. HTTP Request with SSE Response (see below)
// 3. HTTP Upgrade Request for Web Sockets (see below)
// Standard HTTP Request/Response (method 1) has the worst scalability
// because a long-running operation will tie up a Web Server thread **and**
// continue until complete even if the client aborts the operation since
// no data is written back to the client until complete and therefore the
// terminated connection will not be detected.  SSE (method 2) mitigates
// the problem of detecting when a client has aborted the operation
// because periodic progress responses are written to the client and
// therefore a terminated connection will be detected.  However, the best
// way to invoke this operation is via Web Sockets (method 3) which not
// only can detect disconnection of the client, but it also upgrades the
// request to use its own thread outside the Web Server thread pool.
// **SSE NOTE:** This end-point supports "Server-sent Events" (SSE) via the
// `text/event-stream` media type.  This support is activated by adding the
// `Accept: text/event-stream` header to a request to override the
// default `application/json` media type.  Further, the end-point will behave
// similarly to its standard operation but will produce `progress` events
// at regular intervals that are equivalent to its `200` response schema.
// Upon success, the final event will be `completed` with the same response
// schema as a `200` response.  Upon failure, the final event will be
// `failed` with same `SzErrorResponse` schema as the `4xx` or `5xx`.
// **WEB SOCKETS NOTE**: If invoking via Web Sockets then the client may
// send text or binary chunks of the JSON, JSON-Lines or CSV bulk data file
// as messages.  In Web Sockets, text messages are *always* sent as UTF-8.
// If the file's character encoding is unknown then the client should send
// binary messages and the server will attempt to auto-detect the character
// encoding.  Each message should adhere to the maximum message size
// specified by the `webSocketsMessageMaxSize` property returned by the
// `GET /server-info` end-point.  The end of file is detected when the
// number of seconds specified by the `eofSendTimeout` query parameter have
// elapsed without a new message being received.
//
// POST /bulk-data/analyze
func (httpService *HttpServiceImpl) AnalyzeBulkRecords(ctx context.Context, req senzinghttpapi.AnalyzeBulkRecordsReq, params senzinghttpapi.AnalyzeBulkRecordsParams) (r senzinghttpapi.AnalyzeBulkRecordsRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// DeleteRecord implements deleteRecord operation.
//
// This operation deletes a single record identified by the data source
// code and record ID in the request path.
// **NOTE:** The `withInfo` parameter will return the entity resolution
// info pertaining to the delete.  This can be used to update a search
// index or external data mart.   Additionally, Senzing API Server provides
// a means to have the "raw" entity resolution info (from the underlying
// native Senzing API) automatically sent to a messaging service such as
// those provided by Amazon SQS, Rabbit MQ or Kafka regardless of the
// `withInfo` query parameter value.
//
// DELETE /data-sources/{dataSourceCode}/records/{recordId}
func (httpService *HttpServiceImpl) DeleteRecord(ctx context.Context, params senzinghttpapi.DeleteRecordParams) (r senzinghttpapi.DeleteRecordRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// FindEntityNetwork implements findEntityNetwork operation.
//
// This operation finds the entity network around one or more entities.
// This attempts to find paths between the specified entities.  If no
// paths exist, then island networks are returned with each island network
// containing up to a specified number of related entities.  The entities
// are identified by their entity IDs or by data source code and record ID
// pairs for constituent records of those entities.
// **NOTE:** If the first entity is identified by entity ID then the
// subsequent entities must also be identified entity ID.  Similarly, if
// the first entity is identified by the data source code and record ID
// of a consistuent record then the subsequent entities must also be
// identified by the data source code and record ID of constituent records.
// **ALSO NOTE:** Bear in mind that entity ID's are transient and may be
// recycled or repurposed as new records are loaded and entities resolve,
// unresolve and re-resolve.
//
// GET /entity-networks
func (httpService *HttpServiceImpl) FindEntityNetwork(ctx context.Context, params senzinghttpapi.FindEntityNetworkParams) (r senzinghttpapi.FindEntityNetworkRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// FindEntityPath implements findEntityPath operation.
//
// This operation finds the path between two entities and returns a
// description of that entity path (if any) or a response indicating that
// there is no path between the entities.  The subject entities are either
// identfieid by entity ID or by data source code and record ID pairs for
// constituent records of those entities.
// **NOTE:** If the first entity is identified by entity ID then the second
// must also be identified an entity ID.  Similarly, if the first entity is
// identified by data source code and record ID then the second must also
// be identified by data source code and record ID.
// **ALSO NOTE:** Bear in mind that entity ID's are transient and may be
// recycled or repurposed as new records are loaded and entities resolve,
// unresolve and re-resolve.
//
// GET /entity-paths
func (httpService *HttpServiceImpl) FindEntityPath(ctx context.Context, params senzinghttpapi.FindEntityPathParams) (r senzinghttpapi.FindEntityPathRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// GetActiveConfig implements getActiveConfig operation.
//
// This operation returns the JSON configuration that is currently being
// used by the native Senzing API initialized by the running server.  No
// processing or interpretation is performed on the JSON.  This may differ
// from the registered "default configuration" which the server would
// use if no other configuration were provided.
//
// GET /configs/active
func (httpService *HttpServiceImpl) GetActiveConfig(ctx context.Context) (r senzinghttpapi.GetActiveConfigRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// GetAttributeType implements getAttributeType operation.
//
// This operation will provide a description of a single attribute type for
// the specified attribute type code.
//
// GET /attribute-types/{attributeCode}
func (httpService *HttpServiceImpl) GetAttributeType(ctx context.Context, params senzinghttpapi.GetAttributeTypeParams) (r senzinghttpapi.GetAttributeTypeRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// GetAttributeTypes implements getAttributeTypes operation.
//
// This operation will provide a list of attribute types that are
// configured.  The client can filter the returned list of attribute types
// using various query parameters.
//
// GET /attribute-types
func (httpService *HttpServiceImpl) GetAttributeTypes(ctx context.Context, params senzinghttpapi.GetAttributeTypesParams) (r senzinghttpapi.GetAttributeTypesRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// GetDataSource implements getDataSource operation.
//
// This operation provides details on a specific data source identified
// by the data source code in the requested path.
//
// GET /data-sources/{dataSourceCode}
func (httpService *HttpServiceImpl) GetDataSource(ctx context.Context, params senzinghttpapi.GetDataSourceParams) (r senzinghttpapi.GetDataSourceRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// GetDataSources implements getDataSources operation.
//
// This operation will provide a list of data source codes as well as a
// list of detailed descriptions of each data source.
//
// GET /data-sources
func (httpService *HttpServiceImpl) GetDataSources(ctx context.Context, params senzinghttpapi.GetDataSourcesParams) (r senzinghttpapi.GetDataSourcesRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// GetEntityByEntityId implements getEntityByEntityId operation.
//
// Gets the details on a resolved entity that is identified by the entity
// ID specified in the request path.
// **NOTE:** Bear in mind that entity ID's are transient and may be
// recycled or repurposed as new records are loaded and entities resolve,
// unresolve and re-resolve.  An alternative way to identify an entity is
// by one of its constituent records using
// `GET /data-sources/{dataSourceCode}/records/{recordId}/entity`.
//
// GET /entities/{entityId}
func (httpService *HttpServiceImpl) GetEntityByEntityId(ctx context.Context, params senzinghttpapi.GetEntityByEntityIdParams) (r senzinghttpapi.GetEntityByEntityIdRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// GetEntityByRecordId implements getEntityByRecordId operation.
//
// Gets the details on a resolved entity that contains the record
// identified by the data source code and record ID in the specified
// request path.
//
// GET /data-sources/{dataSourceCode}/records/{recordId}/entity
func (httpService *HttpServiceImpl) GetEntityByRecordId(ctx context.Context, params senzinghttpapi.GetEntityByRecordIdParams) (r senzinghttpapi.GetEntityByRecordIdRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// GetRecord implements getRecord operation.
//
// Gets details on a specific entity record identified by the data source
// code and record ID specified in the request path.
//
// GET /data-sources/{dataSourceCode}/records/{recordId}
func (httpService *HttpServiceImpl) GetRecord(ctx context.Context, params senzinghttpapi.GetRecordParams) (r senzinghttpapi.GetRecordRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// GetServerInfo implements getServerInfo operation.
//
// This operation will provides server information describing the options
// with which the server was started.  This can be used to determine if
// the admin operations are enabled or if only read operations may be
// invoked.  This also allows the client to know the maximum message size
// for Web Sockets communication.
//
// GET /server-info
func (httpService *HttpServiceImpl) GetServerInfo(ctx context.Context) (r senzinghttpapi.GetServerInfoRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// GetTemplateConfig implements getTemplateConfig operation.
//
// This operation returns a template base JSON configuration that can be
// modified or customized by the caller.  The returned template is
// according to the underlying native Senzing API and may change between
// version upgrades to Senzing.  No processing or interpretation is
// performed on the JSON.  This will likely differ from the registered
// "default configuration" or currently "active configuration" being used
// by the running API server.
//
// GET /configs/template
func (httpService *HttpServiceImpl) GetTemplateConfig(ctx context.Context) (r senzinghttpapi.GetTemplateConfigRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// GetVirtualEntityByRecordIds implements getVirtualEntityByRecordIds operation.
//
// This operation simulates the resolution of the one or more specified
// records into a single entity and returns the simulated "virtual"
// entity.  The subject records are identified by data source code and
// record ID pairs.
//
// GET /virtual-entities
func (httpService *HttpServiceImpl) GetVirtualEntityByRecordIds(ctx context.Context, params senzinghttpapi.GetVirtualEntityByRecordIdsParams) (r senzinghttpapi.GetVirtualEntityByRecordIdsRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// Heartbeat implements heartbeat operation.
//
// The heartbeat operation can be used to ensure that the HTTP server is
// indeed running, but this operation does not call upon the underlying
// native Senzing API and therefore does not ensure the Senzing
// initialization or configuration is valid.
//
// GET /heartbeat
func (httpService *HttpServiceImpl) Heartbeat(ctx context.Context) (r *senzinghttpapi.SzBaseResponse, _ error) {
	r = &senzinghttpapi.SzBaseResponse{
		Links: httpService.getOptSzLinks(),
		Meta:  httpService.getOptSzMeta(),
	}
	return r, nil
}

// HowEntityByEntityID implements howEntityByEntityID operation.
//
// This operation provides an anlysis of how the records in an entity
// resolved.  The subject entity is identified by the entity ID in the
// request path.
// **NOTE:** Bear in mind that entity ID's are transient and may be
// recycled or repurposed as new records are loaded and entities resolve,
// unresolve and re-resolve.  An alternative way to identify an entity is
// by one of its constituent records using
// `GET /data-sources/{dataSourceCode}/records/{recordId}/entity/how`.
//
// GET /entities/{entityId}/how
func (httpService *HttpServiceImpl) HowEntityByEntityID(ctx context.Context, params senzinghttpapi.HowEntityByEntityIDParams) (r senzinghttpapi.HowEntityByEntityIDRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// HowEntityByRecordID implements howEntityByRecordID operation.
//
// This operation provides an anlysis of how the records in an entity
// resolved.  The subject entity is the one containing the record
// identified by the data source code and record ID in the request path.
//
// GET /data-sources/{dataSourceCode}/records/{recordId}/entity/how
func (httpService *HttpServiceImpl) HowEntityByRecordID(ctx context.Context, params senzinghttpapi.HowEntityByRecordIDParams) (r senzinghttpapi.HowEntityByRecordIDRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// License implements license operation.
//
// This operation will obtain license information for the underlying
// native Senzing API.
//
// GET /license
func (httpService *HttpServiceImpl) License(ctx context.Context, params senzinghttpapi.LicenseParams) (r senzinghttpapi.LicenseRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// LoadBulkRecords implements loadBulkRecords operation.
//
// Provides a means to load a bulk data file of records.  The records are
// encoded as a JSON array of JSON objects, a single JSON object per line
// in JSON-lines file format, or as a CSV with one record per row.  The
// data should be in pre-mapped format using JSON property names or CSV
// column names as described by the
// [Senzing Generic Entity Specification](https://senzing.zendesk.
// com/hc/en-us/articles/231925448-Generic-Entity-Specification).
// **SCALABILITY NOTE:** This operation can be invoked in three ways.  In
// order of increasingly better scalability these are listed below:
// 1. Standard HTTP Request/Response
// 2. HTTP Request with SSE Response (see below)
// 3. HTTP Upgrade Request for Web Sockets (see below)
// Standard HTTP Request/Response (method 1) has the worst scalability
// because a long-running operation will tie up a Web Server thread **and**
// continue until complete even if the client aborts the operation since
// no data is written back to the client until complete and therefore the
// terminated connection will not be detected.  SSE (method 2) mitigates
// the problem of detecting when a client has aborted the operation
// because periodic progress responses are written to the client and
// therefore a terminated connection will be detected.  However, the best
// way to invoke this operation is via Web Sockets (method 3) which not
// only can detect disconnection of the client, but it also upgrades the
// request to use its own thread outside the Web Server thread pool.
// **SSE NOTE:** This end-point supports "Server-sent Events" (SSE) via the
// `text/event-stream` media type.  This support is activated by adding the
// `Accept: text/event-stream` header to a request to override the
// default `application/json` media type.  Further, the end-point will behave
// similarly to its standard operation but will produce `progress` events
// at regular intervals that are equivalent to its `200` response schema.
// Upon success, the final event will be `completed` with the same response
// schema as a `200` response.  Upon failure, the final event will be
// `failed` with same `SzErrorResponse` schema as the `4xx` or `5xx`.
// **WEB SOCKETS NOTE**: If invoking via Web Sockets then the client may
// send text or binary chunks of the JSON, JSON-Lines or CSV bulk data file
// as messages.  In Web Sockets, text messages are *always* sent as UTF-8.
// If the file's character encoding is unknown then the client should send
// binary messages and the server will attempt to auto-detect the character
// encoding.  Each message should adhere to the maximum message size
// specified by the `webSocketsMessageMaxSize` property returned by the
// `GET /server-info` end-point.  The end of file is detected when the
// number of seconds specified by the `eofSendTimeout` query parameter have
// elapsed without a new message being received.
//
// POST /bulk-data/load
func (httpService *HttpServiceImpl) LoadBulkRecords(ctx context.Context, req senzinghttpapi.LoadBulkRecordsReq, params senzinghttpapi.LoadBulkRecordsParams) (r senzinghttpapi.LoadBulkRecordsRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// OpenApiSpecification implements openApiSpecification operation.
//
// This operation can be used to obtain the Open API specification in JSON
// format.  The specification can either be the `data` field of a standard
// response (i.e.: a response with a `meta`, `links` and `data` field) or
// as raw format where the root JSON document is the Open API specification
// JSON.
//
// GET /specifications/open-api
func (httpService *HttpServiceImpl) OpenApiSpecification(ctx context.Context) (r senzinghttpapi.OpenApiSpecificationOKDefault, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// ReevaluateEntity implements reevaluateEntity operation.
//
// Reevaluates an entity identified by the entity ID specified via the
// `entityId` query parameter.
// **NOTE:** The `withInfo` parameter will return the entity resolution
// info pertaining to the reevaluation.  This can be used to update a
// search index or external data mart.   Additionally, Senzing API Server
// provides a means to have the "raw" entity resolution info (from the
// underlying native Senzing API) automatically sent to a messaging service
// such as those provided by Amazon SQS, Rabbit MQ or Kafka regardless of
// the `withInfo` query parameter value.
// **ALSO NOTE:** Bear in mind that entity ID's are transient and may be
// recycled or repurposed as new records are loaded and entities resolve,
// unresolve and re-resolve.
//
// POST /reevaluate-entity
func (httpService *HttpServiceImpl) ReevaluateEntity(ctx context.Context, params senzinghttpapi.ReevaluateEntityParams) (r senzinghttpapi.ReevaluateEntityRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// ReevaluateRecord implements reevaluateRecord operation.
//
// This operation reevaluates a single record identified by the data source
// code and record ID in the request path.
// **NOTE:** The `withInfo` parameter will return the entity resolution
// info pertaining to the reevaluation.  This can be used to update a
// search index or external data mart.   Additionally, Senzing API Server
// provides a means to have the "raw" entity resolution info (from the
// underlying native Senzing API) automatically sent to a messaging service
// such as those provided by Amazon SQS, Rabbit MQ or Kafka regardless of
// the `withInfo` query parameter value.
//
// POST /data-sources/{dataSourceCode}/records/{recordId}/reevaluate
func (httpService *HttpServiceImpl) ReevaluateRecord(ctx context.Context, params senzinghttpapi.ReevaluateRecordParams) (r senzinghttpapi.ReevaluateRecordRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// Root implements root operation.
//
// The root operation can be used to ensure that the HTTP server is
// indeed running, but this operation does not call upon the underlying
// native Senzing API and therefore does not ensure the Senzing
// initialization or configuration is valid.
//
// GET /
func (httpService *HttpServiceImpl) Root(ctx context.Context) (r *senzinghttpapi.SzBaseResponse, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// SearchEntitiesByGet implements searchEntitiesByGet operation.
//
// This operation finds all entities that would resolve or relate to the
// search candidate features specified by the `attr` and/or `attrs` query
// parameters.  The search candidate features are treated as if they
// belonged to an inbound record being loaded, thus the attribute names are
// given by the [Senzing Generic Entity Specification](https://senzing.zendesk.
// com/hc/en-us/articles/231925448-Generic-Entity-Specification).
// If including the search candidate features as query parameters presents
// privacy concerns due to sensitivity of the data, they can alternately
// be sent in the request body using the `POST /search-entities` endpoint.
// **NOTE:** This operation differs from a keyword search in that it uses
// deterministic entity resolution rules to determine the result set.  This
// means that features that are considered "generic" (i.e.: overly common)
// will be ignored just as they are during entity resolution and will not
// yield search results.  For example, searching on a gender by itself will
// return no results rather than half of all entities.  Similarly, a phone
// number such as `555-1212` may yield no results.
//
// GET /entities
func (httpService *HttpServiceImpl) SearchEntitiesByGet(ctx context.Context, params senzinghttpapi.SearchEntitiesByGetParams) (r senzinghttpapi.SearchEntitiesByGetRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// SearchEntitiesByPost implements searchEntitiesByPost operation.
//
// This operation finds all entities that would resolve or relate to the
// search candidate features specified in JSON request body.  The search
// candidate features are treated as if they belonged to an inbound record
// being loaded.  The JSON format of the request body is defined by the
// [Senzing Generic Entity Specification](https://senzing.zendesk.
// com/hc/en-us/articles/231925448-Generic-Entity-Specification).
// This operation is similar to the `GET /entities` endpoint in function
// except that it provides a means to avoid specifying potentially
// sensitive data in query parameters, but instead in the request body.
// **NOTE:** This operation differs from a keyword search in that it uses
// deterministic entity resolution rules to determine the result set.  This
// means that features that are considered "generic" (i.e.: overly common)
// will be ignored just as they are during entity resolution and will not
// yield search results.  For example, searching on a gender by itself will
// return no results rather than half of all entities.  Similarly, a phone
// number such as `555-1212` may yield no results.
//
// POST /search-entities
func (httpService *HttpServiceImpl) SearchEntitiesByPost(ctx context.Context, req senzinghttpapi.SearchEntitiesByPostReq, params senzinghttpapi.SearchEntitiesByPostParams) (r senzinghttpapi.SearchEntitiesByPostRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// Version implements version operation.
//
// This operation will obtain the full version information for the server.
// Much of the same information is available in the `meta` segment of
// every JSON response.
//
// GET /version
func (httpService *HttpServiceImpl) Version(ctx context.Context, params senzinghttpapi.VersionParams) (r senzinghttpapi.VersionRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// WhyEntities implements whyEntities operation.
//
// This operation provides an anlysis of why two entities related, did not
// relate or did not resolve.  The entities are identified either by
// entity ID's or by data source code and record ID pairs for constituent
// records of those entities.
// **NOTE:** If the first entity is identified by entity ID then the second
// must also be identified an entity ID.  Similarly, if the first entity is
// identified by data source code and record ID then the second must also
// be identified by data source code and record ID.
// **ALSO NOTE:** Bear in mind that entity ID's are transient and may be
// recycled or repurposed as new records are loaded and entities resolve,
// unresolve and re-resolve.
//
// GET /why/entities
func (httpService *HttpServiceImpl) WhyEntities(ctx context.Context, params senzinghttpapi.WhyEntitiesParams) (r senzinghttpapi.WhyEntitiesRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// WhyEntityByEntityID implements whyEntityByEntityID operation.
//
// This operation provides an anlysis of why the records in an entity
// resolved.  The subject entity is identified by the entity ID in the
// request path.
// **NOTE:** Bear in mind that entity ID's are transient and may be
// recycled or repurposed as new records are loaded and entities resolve,
// unresolve and re-resolve.  An alternative way to identify an entity is
// by one of its constituent records using
// `GET /data-sources/{dataSourceCode}/records/{recordId}/entity/why`.
//
// GET /entities/{entityId}/why
func (httpService *HttpServiceImpl) WhyEntityByEntityID(ctx context.Context, params senzinghttpapi.WhyEntityByEntityIDParams) (r senzinghttpapi.WhyEntityByEntityIDRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// WhyEntityByRecordID implements whyEntityByRecordID operation.
//
// This operation provides an anlysis of why the records in an entity
// resolved.  The subject entity is the one containing the record
// identified by the data source code and record ID in the request path.
//
// GET /data-sources/{dataSourceCode}/records/{recordId}/entity/why
func (httpService *HttpServiceImpl) WhyEntityByRecordID(ctx context.Context, params senzinghttpapi.WhyEntityByRecordIDParams) (r senzinghttpapi.WhyEntityByRecordIDRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}

// WhyRecords implements whyRecords operation.
//
// This operation provides an anlysis of two records identified by data
// source code and record ID in respective qeury parameters resolved or
// did not resolve.
//
// GET /why/records
func (httpService *HttpServiceImpl) WhyRecords(ctx context.Context, params senzinghttpapi.WhyRecordsParams) (r senzinghttpapi.WhyRecordsRes, _ error) {
	return r, ogenHttp.ErrNotImplemented
}
