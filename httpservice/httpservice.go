package httpservice

import (
	"context"
	"fmt"
	"time"

	api "github.com/docktermj/go-http/senzinghttpapi"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// HttpServiceImpl is...
type HttpServiceImpl struct {
	api.UnimplementedHandler
}

// ----------------------------------------------------------------------------
// internal methods
// ----------------------------------------------------------------------------

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

	dataSource := params.DataSource

	fmt.Printf(">>>>>> dataSource: %s\n", dataSource)
	fmt.Printf(">>>>>> r: %v\n", r)

	szDataSource := &api.SzDataSource{
		DataSourceCode: api.NewOptString("DataSourceCodeBob"),
		DataSourceId:   api.NewOptNilInt32(1),
	}

	r = &api.SzDataSourcesResponse{
		Data: api.NewOptSzDataSourcesResponseData(api.NewSzDataSourcesResponseData("bob was here")),
	}
	// params.

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
