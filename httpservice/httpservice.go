package httpservice

import (
	"context"
	// "fmt"
	"sync"
	"time"

	api "github.com/docktermj/go-http/senzinghttpapi"
	"github.com/senzing/g2-sdk-go/g2api"
	"github.com/senzing/go-logging/logging"
	"github.com/senzing/go-observing/observer"
	"github.com/senzing/go-sdk-abstract-factory/factory"
	"google.golang.org/grpc"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// HttpServiceImpl is...
type HttpServiceImpl struct {
	api.UnimplementedHandler
	abstractFactory                factory.SdkAbstractFactory
	abstractFactorySyncOnce        sync.Once
	g2configSingleton              g2api.G2config
	g2configSyncOnce               sync.Once
	GrpcTarget                     string
	GrpcDialOptions                []grpc.DialOption
	logger                         logging.LoggingInterface
	LogLevelName                   string
	ObserverOrigin                 string
	Observers                      []observer.Observer
	ObserverUrl                    string
	Port                           int
	SenzingEngineConfigurationJson string
	SenzingModuleName              string
	SenzingVerboseLogging          int
}

// ----------------------------------------------------------------------------
// internal methods
// ----------------------------------------------------------------------------

// --- Logging ----------------------------------------------------------------

// Get the Logger singleton.
func (httpService *HttpServiceImpl) getLogger() logging.LoggingInterface {
	var err error = nil
	if httpService.logger == nil {
		loggerOptions := []interface{}{
			&logging.OptionCallerSkip{Value: 3},
		}
		httpService.logger, err = logging.NewSenzingToolsLogger(ComponentId, IdMessages, loggerOptions...)
		if err != nil {
			panic(err)
		}
	}
	return httpService.logger
}

// Trace method entry.
func (httpService *HttpServiceImpl) traceEntry(messageNumber int, details ...interface{}) {
	httpService.getLogger().Log(messageNumber, details...)
}

// Trace method exit.
func (httpService *HttpServiceImpl) traceExit(messageNumber int, details ...interface{}) {
	httpService.getLogger().Log(messageNumber, details...)
}

// --- Errors -----------------------------------------------------------------

// Create error.
func (httpService *HttpServiceImpl) error(messageNumber int, details ...interface{}) error {
	return httpService.getLogger().NewError(messageNumber, details...)
}

// --- Services ---------------------------------------------------------------

func (httpService *HttpServiceImpl) getAbstractFactory() factory.SdkAbstractFactory {
	httpService.abstractFactorySyncOnce.Do(func() {
		if len(httpService.GrpcTarget) == 0 {
			httpService.abstractFactory = &factory.SdkAbstractFactoryImpl{}
		} else {
			// TODO: Make a call to go-grpcing.getGrpcAddress() and go-grpcing.
			httpService.abstractFactory = &factory.SdkAbstractFactoryImpl{
				GrpcTarget:      httpService.GrpcTarget,
				GrpcDialOptions: httpService.GrpcDialOptions,
			}
		}
	})
	return httpService.abstractFactory
}

// Singleton pattern for g2config.
// See https://medium.com/golang-issue/how-singleton-pattern-works-with-golang-2fdd61cd5a7f
func (httpService *HttpServiceImpl) getG2config(ctx context.Context) g2api.G2config {
	var err error = nil
	httpService.g2configSyncOnce.Do(func() {
		httpService.g2configSingleton, err = httpService.getAbstractFactory().GetG2config(ctx)
		if err != nil {
			panic(err)
		}
		if httpService.g2configSingleton.GetSdkId(ctx) == factory.ImplementedByBase {
			err = httpService.g2configSingleton.Init(ctx, httpService.SenzingModuleName, httpService.SenzingEngineConfigurationJson, httpService.SenzingVerboseLogging)
			if err != nil {
				panic(err)
			}
		}
	})
	return httpService.g2configSingleton
}

// --- Misc -------------------------------------------------------------------

func (httpService *HttpServiceImpl) getOptSzLinks() api.OptSzLinks {
	var result api.OptSzLinks
	szLinks := api.SzLinks{
		Self:                 api.NewOptString("SelfBob"),
		OpenApiSpecification: api.NewOptString("OpenApiSpecificationBob"),
	}
	result = api.NewOptSzLinks(szLinks)
	return result
}

func (httpService *HttpServiceImpl) getOptSzMeta() api.OptSzMeta {
	var result api.OptSzMeta
	szMeta := api.SzMeta{
		Server:                     api.NewOptString("ServerBob"),
		HttpMethod:                 api.NewOptSzHttpMethod(api.SzHttpMethodGET),
		HttpStatusCode:             api.NewOptInt(200),
		Timestamp:                  api.NewOptDateTime(time.Now()),
		Version:                    api.NewOptString("VersionBob"),
		RestApiVersion:             api.NewOptString("RestApiVersionBob"),
		NativeApiVersion:           api.NewOptString("NativeApiVersionBob"),
		NativeApiBuildVersion:      api.NewOptString("NativeApiBuildVersionBob"),
		NativeApiBuildNumber:       api.NewOptString("NativeApiBuildNumberBob"),
		NativeApiBuildDate:         api.NewOptDateTime(time.Now()),
		ConfigCompatibilityVersion: api.NewOptString("ConfigCompatibilityVersionBob"),
		Timings:                    api.NewOptNilSzMetaTimings(map[string]int64{}),
	}
	result = api.NewOptSzMeta(szMeta)
	return result
}

// ----------------------------------------------------------------------------
// Interface methods
// See https://github.com/docktermj/go-http/blob/main/senzinghttpapi/oas_unimplemented_gen.go
// ----------------------------------------------------------------------------

func (httpService *HttpServiceImpl) AddDataSources(ctx context.Context, req api.AddDataSourcesReq, params api.AddDataSourcesParams) (r api.AddDataSourcesRes, _ error) {
	var err error = nil

	// URl parameters.

	// dataSource := params.DataSource
	// withRaw := params.WithRaw

	// TODO: Call Senzing

	// g2Config := httpService.getG2config(ctx)
	// configHandle, err := g2Config.Create(ctx)
	// inputJson := fmt.Sprintf(`{"DSRC_CODE": "%s"}`, params.DataSource)
	// response, err := g2Config.AddDataSource(ctx, configHandle, inputJson)
	// err = g2Config.Close(ctx, configHandle)

	// type SzDataSource struct {
	// 	// The data source code.
	// 	DataSourceCode OptString `json:"dataSourceCode"`
	// 	// The data source ID. The value can be null when used for input in creating a data source to
	// 	// indicate that the data source ID should be auto-generated.
	// 	DataSourceId OptNilInt32 `json:"dataSourceId"`
	// }

	szDataSource := &api.SzDataSource{
		DataSourceCode: api.NewOptString("DataSourceCodeBob"),
		DataSourceId:   api.NewOptNilInt32(1),
	}

	// type SzDataSourcesResponseDataDataSourceDetails map[string]SzDataSource

	szDataSourcesResponseDataDataSourceDetails := &api.SzDataSourcesResponseDataDataSourceDetails{
		"xxxBob": *szDataSource,
	}

	// type OptSzDataSourcesResponseDataDataSourceDetails struct {
	// 	Value SzDataSourcesResponseDataDataSourceDetails
	// 	Set   bool
	// }

	optSzDataSourcesResponseDataDataSourceDetails := &api.OptSzDataSourcesResponseDataDataSourceDetails{
		Value: *szDataSourcesResponseDataDataSourceDetails,
		Set:   true,
	}

	// type SzDataSourcesResponseData struct {
	// 	// The list of data source codes for the configured data sources.
	// 	DataSources []string `json:"dataSources"`
	// 	// The list of `SzDataSource` instances describing the data sources that are configured.
	// 	DataSourceDetails OptSzDataSourcesResponseDataDataSourceDetails `json:"dataSourceDetails"`
	// }

	szDataSourcesResponseData := &api.SzDataSourcesResponseData{
		DataSources:       []string{"Bobber"},
		DataSourceDetails: *optSzDataSourcesResponseDataDataSourceDetails,
	}

	// type OptSzDataSourcesResponseData struct {
	// 	Value SzDataSourcesResponseData
	// 	Set   bool
	// }

	optSzDataSourcesResponseData := &api.OptSzDataSourcesResponseData{
		Value: *szDataSourcesResponseData,
		Set:   true,
	}

	// type SzDataSourcesResponse struct {
	// 	Data OptSzDataSourcesResponseData `json:"data"`
	// }

	r = &api.SzDataSourcesResponse{
		Data: *optSzDataSourcesResponseData,
	}

	// Condensed version of "r"

	r = &api.SzDataSourcesResponse{
		Data: api.OptSzDataSourcesResponseData{
			Set: true,
			Value: api.SzDataSourcesResponseData{
				DataSources: []string{"Bobber"},
				DataSourceDetails: api.OptSzDataSourcesResponseDataDataSourceDetails{
					Set: true,
					Value: api.SzDataSourcesResponseDataDataSourceDetails{
						"xxxBob": api.SzDataSource{
							DataSourceCode: api.NewOptString("DataSourceCodeBob"),
							DataSourceId:   api.NewOptNilInt32(1),
						},
					},
				},
			},
		},
	}

	return r, err
}

func (httpService *HttpServiceImpl) Heartbeat(ctx context.Context) (r *api.SzBaseResponse, _ error) {
	var err error = nil
	r = &api.SzBaseResponse{
		Links: httpService.getOptSzLinks(),
		Meta:  httpService.getOptSzMeta(),
	}
	return r, err
}
