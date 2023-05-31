/*
 */
package cmd

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/docktermj/serve-http/httpserver"
	"github.com/senzing/go-common/g2engineconfigurationjson"
	"github.com/senzing/go-grpcing/grpcurl"
	"github.com/senzing/go-observing/observer"
	"github.com/senzing/senzing-tools/constant"
	"github.com/senzing/senzing-tools/envar"
	"github.com/senzing/senzing-tools/helper"
	"github.com/senzing/senzing-tools/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	defaultCommand              string = "/bin/bash"
	defaultConnectionErrorLimit int    = 10
	defaultKeepalivePingTimeout int    = 20
	defaultMaxBufferSizeBytes   int    = 512
	defaultServerAddress        string = "0.0.0.0"
	envarAllowedHostnames       string = "SENZING_TOOLS_XTERM_ALLOWED_HOSTNAMES"
	envarArguments              string = "SENZING_TOOLS_XTERM_ARGUMENTS"
	envarCommand                string = "SENZING_TOOLS_XTERM_COMMAND"
	envarConnectionErrorLimit   string = "SENZING_TOOLS_XTERM_CONNECTION_ERROR_LIMIT"
	envarKeepalivePingTimeout   string = "SENZING_TOOLS_XTERM_KEEPALIVE_PING_TIMEOUT"
	envarMaxBufferSizeBytes     string = "SENZING_TOOLS_XTERM_MAX_BUFFER_SIZE_BYTES"
	envarServerAddress          string = "SENZING_TOOLS_SERVER_ADDRESS"
	optionAllowedHostnames      string = "xterm-allowed-hostnames"
	optionArguments             string = "xterm-arguments"
	optionCommand               string = "xterm-command"
	optionConnectionErrorLimit  string = "xterm-connection-error-limit"
	optionKeepalivePingTimeout  string = "xterm-keepalive-ping-timeout"
	optionMaxBufferSizeBytes    string = "xterm-max-buffer-size-bytes"
	optionServerAddress         string = "server-address"

	defaultConfiguration           string = ""
	defaultDatabaseUrl             string = ""
	defaultEnableAll               bool   = false
	defaultEnableSenzingRestApi    bool   = false
	defaultEnableSwaggerUI         bool   = false
	defaultEnableXterm             bool   = false
	defaultEngineConfigurationJson string = ""
	defaultEngineLogLevel          int    = 0
	defaultGrpcUrl                        = ""
	defaultHttpPort                int    = 8261
	defaultLogLevel                string = "INFO"
	defaultObserverOrigin          string = "serve-http"
	defaultObserverUrl             string = ""
	Short                          string = "serve-http short description"
	Use                            string = "serve-http"
	Long                           string = `
serve-http long description.
	`
)

var (
	defaultEngineModuleName string   = fmt.Sprintf("serve-http-%d", time.Now().Unix())
	defaultAllowedHostnames []string = []string{"localhost", "192.168.1.12"}
	defaultArguments        []string
)

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

// Since init() is always invoked, define command line parameters.
func init() {

	RootCmd.Flags().Int(optionConnectionErrorLimit, defaultConnectionErrorLimit, fmt.Sprintf("Connection re-attempts before terminating [%s]", envarConnectionErrorLimit))
	RootCmd.Flags().Int(optionKeepalivePingTimeout, defaultKeepalivePingTimeout, fmt.Sprintf("Maximum allowable seconds between a ping message and its response [%s]", envarKeepalivePingTimeout))
	RootCmd.Flags().Int(optionMaxBufferSizeBytes, defaultMaxBufferSizeBytes, fmt.Sprintf("Maximum length of terminal input [%s]", envarMaxBufferSizeBytes))
	RootCmd.Flags().String(optionCommand, defaultCommand, fmt.Sprintf("Path of shell command [%s]", envarCommand))
	RootCmd.Flags().String(optionServerAddress, defaultServerAddress, fmt.Sprintf("IP interface server listens on [%s]", envarServerAddress))
	RootCmd.Flags().StringSlice(optionAllowedHostnames, defaultAllowedHostnames, fmt.Sprintf("Comma-delimited list of hostnames permitted to connect to the websocket [%s]", envarAllowedHostnames))
	RootCmd.Flags().StringSlice(optionArguments, defaultArguments, fmt.Sprintf("Comma-delimited list of arguments passed to the terminal command prompt [%s]", envarArguments))

	RootCmd.Flags().Bool("enable-all", defaultEnableSwaggerUI, fmt.Sprintf("Enable all services [%s]", "SENZING_TOOLS_ENABLE_ALL"))
	RootCmd.Flags().Bool("enable-senzing-rest-api", defaultEnableSwaggerUI, fmt.Sprintf("Enable the Senzing REST API service [%s]", "SENZING_TOOLS_ENABLE_SENZING_REST_API"))
	RootCmd.Flags().Bool(option.EnableSwaggerUi, defaultEnableSwaggerUI, fmt.Sprintf("Enable the Swagger UI service [%s]", envar.EnableSwaggerUi))
	RootCmd.Flags().Bool("enable-xterm", defaultEnableXterm, fmt.Sprintf("Enable the XTerm service [%s]", "SENZING_TOOLS_ENABLE_XTERM"))
	RootCmd.Flags().Int(option.HttpPort, defaultHttpPort, fmt.Sprintf("Port to serve HTTP [%s]", envar.HttpPort))
	RootCmd.Flags().Int(option.EngineLogLevel, defaultEngineLogLevel, fmt.Sprintf("Log level for Senzing Engine [%s]", envar.EngineLogLevel))
	RootCmd.Flags().String(option.GrpcUrl, defaultGrpcUrl, fmt.Sprintf("URL of Senzing gRPC service [%s]", envar.GrpcUrl))
	RootCmd.Flags().String(option.Configuration, defaultConfiguration, fmt.Sprintf("Path to configuration file [%s]", envar.Configuration))
	RootCmd.Flags().String(option.DatabaseUrl, defaultDatabaseUrl, fmt.Sprintf("URL of database to initialize [%s]", envar.DatabaseUrl))
	RootCmd.Flags().String(option.EngineConfigurationJson, defaultEngineConfigurationJson, fmt.Sprintf("JSON string sent to Senzing's init() function [%s]", envar.EngineConfigurationJson))
	RootCmd.Flags().String(option.EngineModuleName, defaultEngineModuleName, fmt.Sprintf("Identifier given to the Senzing engine [%s]", envar.EngineModuleName))
	RootCmd.Flags().String(option.LogLevel, defaultLogLevel, fmt.Sprintf("Log level [%s]", envar.LogLevel))
	RootCmd.Flags().String(option.ObserverOrigin, defaultObserverOrigin, fmt.Sprintf("Identify this instance to the Observer [%s]", envar.ObserverOrigin))
	RootCmd.Flags().String(option.ObserverUrl, defaultObserverUrl, fmt.Sprintf("URL of Observer [%s]", envar.ObserverUrl))
}

// If a configuration file is present, load it.
func loadConfigurationFile(cobraCommand *cobra.Command) {
	configuration := ""
	configFlag := cobraCommand.Flags().Lookup(option.Configuration)
	if configFlag != nil {
		configuration = configFlag.Value.String()
	}
	if configuration != "" { // Use configuration file specified as a command line option.
		viper.SetConfigFile(configuration)
	} else { // Search for a configuration file.

		// Determine home directory.

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Specify configuration file name.

		viper.SetConfigName("serve-http")
		viper.SetConfigType("yaml")

		// Define search path order.

		viper.AddConfigPath(home + "/.senzing-tools")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/senzing-tools")
	}

	// If a config file is found, read it in.

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Applying configuration file:", viper.ConfigFileUsed())
	}
}

// Configure Viper with user-specified options.
func loadOptions(cobraCommand *cobra.Command) {
	var err error = nil
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix(constant.SetEnvPrefix)

	// Bools

	boolOptions := map[string]bool{
		"enable-all":              defaultEnableAll,
		"enable-senzing-rest-api": defaultEnableSenzingRestApi,
		option.EnableSwaggerUi:    defaultEnableSwaggerUI,
		"enable-xterm":            defaultEnableXterm,
	}
	for optionKey, optionValue := range boolOptions {
		viper.SetDefault(optionKey, optionValue)
		err = viper.BindPFlag(optionKey, cobraCommand.Flags().Lookup(optionKey))
		if err != nil {
			panic(err)
		}
	}

	// Ints

	intOptions := map[string]int{
		option.EngineLogLevel:      defaultEngineLogLevel,
		option.HttpPort:            defaultHttpPort,
		optionConnectionErrorLimit: defaultConnectionErrorLimit,
		optionKeepalivePingTimeout: defaultKeepalivePingTimeout,
		optionMaxBufferSizeBytes:   defaultMaxBufferSizeBytes,
	}
	for optionKey, optionValue := range intOptions {
		viper.SetDefault(optionKey, optionValue)
		err = viper.BindPFlag(optionKey, cobraCommand.Flags().Lookup(optionKey))
		if err != nil {
			panic(err)
		}
	}

	// Strings

	stringOptions := map[string]string{
		option.Configuration:           defaultConfiguration,
		option.DatabaseUrl:             defaultDatabaseUrl,
		option.EngineConfigurationJson: defaultEngineConfigurationJson,
		option.EngineModuleName:        defaultEngineModuleName,
		option.LogLevel:                defaultLogLevel,
		option.ObserverOrigin:          defaultObserverOrigin,
		option.ObserverUrl:             defaultObserverUrl,
		option.GrpcUrl:                 defaultGrpcUrl,
		optionCommand:                  defaultCommand,
		optionServerAddress:            defaultServerAddress,
	}
	for optionKey, optionValue := range stringOptions {
		viper.SetDefault(optionKey, optionValue)
		err = viper.BindPFlag(optionKey, cobraCommand.Flags().Lookup(optionKey))
		if err != nil {
			panic(err)
		}
	}

	// StringSlice

	stringSliceOptions := map[string][]string{
		optionAllowedHostnames: defaultAllowedHostnames,
		optionArguments:        defaultArguments,
	}
	for optionKey, optionValue := range stringSliceOptions {
		viper.SetDefault(optionKey, optionValue)
		err = viper.BindPFlag(optionKey, cobraCommand.Flags().Lookup(optionKey))
		if err != nil {
			panic(err)
		}
	}

}

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Used in construction of cobra.Command
func PreRun(cobraCommand *cobra.Command, args []string) {
	loadConfigurationFile(cobraCommand)
	loadOptions(cobraCommand)
	cobraCommand.SetVersionTemplate(constant.VersionTemplate)
}

// Used in construction of cobra.Command
func RunE(_ *cobra.Command, _ []string) error {
	var err error = nil
	ctx := context.TODO()

	// Build senzingEngineConfigurationJson.

	senzingEngineConfigurationJson := viper.GetString(option.EngineConfigurationJson)
	if len(senzingEngineConfigurationJson) == 0 {
		senzingEngineConfigurationJson, err = g2engineconfigurationjson.BuildSimpleSystemConfigurationJson(viper.GetString(option.DatabaseUrl))
		if err != nil {
			return err
		}
	}

	// Determine if gRPC is being used.

	grpcUrl := viper.GetString(option.GrpcUrl)
	grpcTarget := ""
	grpcDialOptions := []grpc.DialOption{}
	if len(grpcUrl) > 0 {
		grpcTarget, grpcDialOptions, err = grpcurl.Parse(ctx, grpcUrl)
		if err != nil {
			return err
		}
	}

	// Build observers.
	//  viper.GetString(option.ObserverUrl),

	observers := []observer.Observer{}

	// Create object and Serve.

	httpServer := &httpserver.HttpServerImpl{
		ApiUrlRoutePrefix:              "api",
		EnableAll:                      viper.GetBool("enable-all"),
		EnableSenzingRestAPI:           viper.GetBool("enable-senzing-rest-api"),
		EnableSwaggerUI:                viper.GetBool(option.EnableSwaggerUi),
		EnableXterm:                    viper.GetBool("enable-xterm"),
		GrpcDialOptions:                grpcDialOptions,
		GrpcTarget:                     grpcTarget,
		LogLevelName:                   viper.GetString(option.LogLevel),
		ObserverOrigin:                 viper.GetString(option.ObserverOrigin),
		Observers:                      observers,
		SenzingEngineConfigurationJson: senzingEngineConfigurationJson,
		SenzingModuleName:              viper.GetString(option.EngineModuleName),
		SenzingVerboseLogging:          viper.GetInt(option.EngineLogLevel),
		ServerAddress:                  viper.GetString(optionServerAddress),
		ServerPort:                     viper.GetInt(option.HttpPort),
		SwaggerUrlRoutePrefix:          "swagger",
		XtermAllowedHostnames:          viper.GetStringSlice(optionAllowedHostnames),
		XtermArguments:                 viper.GetStringSlice(optionArguments),
		XtermCommand:                   viper.GetString(optionCommand),
		XtermConnectionErrorLimit:      viper.GetInt(optionConnectionErrorLimit),
		XtermKeepalivePingTimeout:      viper.GetInt(optionKeepalivePingTimeout),
		XtermMaxBufferSizeBytes:        viper.GetInt(optionMaxBufferSizeBytes),
		XtermUrlRoutePrefix:            "xterm",
	}
	err = httpServer.Serve(ctx)
	return err
}

// Used in construction of cobra.Command
func Version() string {
	return helper.MakeVersion(githubVersion, githubIteration)
}

// ----------------------------------------------------------------------------
// Command
// ----------------------------------------------------------------------------

// RootCmd represents the command.
var RootCmd = &cobra.Command{
	Use:     Use,
	Short:   Short,
	Long:    Long,
	PreRun:  PreRun,
	RunE:    RunE,
	Version: Version(),
}
