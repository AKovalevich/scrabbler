package config

import (
	"fmt"
	"sync"
	"time"
	"reflect"

	"github.com/BurntSushi/toml"
	"github.com/containous/flaeg"
	"github.com/AKovalevich/scrabbler/entrypoint"
	log "github.com/AKovalevich/scrabbler/log/logrus"
)


const (
	// DefaultGraceTimeout controls how long Scrabbler serves pending requests
	// prior to shutting down.
	DefaultGraceTimeout = 10 * time.Second

	// DefaultConfigFileName path to configuration file
	DefaultConfigPath = "configuration.default.toml"
)

// The main Scrabbler configuration
type ScrabblerConfiguration struct {
	sync.RWMutex
	EntryPointList	[]*entrypoint.Entrypoint
	// Main configuration
	Debug			bool						`toml:"debug" short:"d" description:"Enable debug mode" export:"true"`
	LogLevel		string						`toml:"log_level" short:"l" description:"Log level" export:"true"`
	ConfigFilePath	string						`toml:"config_file_path" short:"c" description:"Path to configuration directory, load configuration.toml file in a directory" export:"true"`
	// Scrabbler server configuration
	ServerPort		int							`toml:"server_port" short:"sp" description:"Scrabbler web server port" export:"true"`
	ServerHost		string						`toml:"server_host" short:"sd" description:"Scrabbler web server host" export:"true"`
	// Web UI configuration
	WebUI			bool 						`toml:"web_ui" short:"w" description:"Run service with web UI" export:"true"`
	WebUIPort		int							`toml:"web_ui_port" short:"wp" description:"Web UI port" export:"true"`
	WebUIHost		string						`toml:"web_ui_host" short:"wh" description:"Web UI host" export:"true"`
	// Shutdown configuration
	GraceTimeOut 	flaeg.Duration 				`toml:"grace_time_out" short:"g" description:"Duration to give active requests a chance to finish before Scrabbler stops" export:"true"`
	// Entrypoints configuration
	EntryPoints 	entrypoint.EntrypointList 	`toml:"entry_points" short:"e" description:"Scrabbler server entry points" export:"true"`
}

func NewScrabblerConfiguration() *ScrabblerConfiguration {
	return &ScrabblerConfiguration{}
}

func NewScrabblerDefaultConfiguration() *ScrabblerConfiguration {
	// Will take data from configuration file
	return &ScrabblerConfiguration{}
}

// Reload configuration
func (config *ScrabblerConfiguration) Reload() {
	config.Load()
}

// Load configuration from file
func (config *ScrabblerConfiguration) Load() {
	log.Do.Info("Load configuration...")
	var configFilePath string
	var tmpScreabblerConfiguration ScrabblerConfiguration

	if config.ConfigFilePath != "" {
		configFilePath = config.ConfigFilePath
	} else {
		configFilePath = DefaultConfigPath
	}

	if _, err := toml.DecodeFile(configFilePath, &tmpScreabblerConfiguration); err != nil {
		log.Do.Error(err)
	}

	tmpConfigStructureValues := reflect.ValueOf(&tmpScreabblerConfiguration).Elem()
	realConfigStructureValues := reflect.ValueOf(config).Elem()
	realConfigStructureTypes := reflect.TypeOf(*config)

	for indexRealConfigStructure := 0; indexRealConfigStructure < realConfigStructureValues.NumField(); indexRealConfigStructure++ {
		var fieldName string
		var field reflect.Value

		// Get field from real configuration structure
		fieldName = realConfigStructureTypes.Field(indexRealConfigStructure).Name
		field = realConfigStructureValues.FieldByName(fieldName)

		// Check if field is empty
		if isEmpty(field.Interface()) && field.CanSet() {
			config.Lock()
			// Try to get field from tmp configuration structure (form configuration file)
			tmpField := tmpConfigStructureValues.FieldByName(fieldName)
			f := tmpField.Interface()
			val := reflect.ValueOf(f)

			// Try to set values to global configuration structure
			switch field.Kind() {
			case reflect.Bool:
				field.SetBool(val.Bool())
				break
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				field.SetInt(val.Int())
				break
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				field.SetUint(val.Uint())
				break
			case reflect.Float32, reflect.Float64:
				field.SetFloat(val.Float())
				break
			case reflect.Complex64, reflect.Complex128:
				field.SetComplex(val.Complex())
				break
			case reflect.String:
				field.SetString(val.String())
				break
			case reflect.Array:
				break
			case reflect.Slice:
				break
			}
			config.Unlock()
		}
	}
	fmt.Printf("%+v\n", config.EntryPointList)
}

// Check if field is empty
func isEmpty(object interface{}) bool {
	if object == nil {
		return true
	} else if object == "" {
		return true
	} else if object == false {
		return true
	} else if object == nil {
		return true
	} else if object == 0 {
		return true
	}

	//Then see if it's a struct
	if reflect.ValueOf(object).Kind() == reflect.Struct {
		// and create an empty copy of the struct object to compare against
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		if reflect.DeepEqual(object, empty) {
			return true
		}
	}

	return false
}

func CustomEntryPointParsers() (reflect.Type, *entrypoint.EntrypointList) {
	// Add custom parsers to flaeg
	return reflect.TypeOf(entrypoint.EntrypointList{}), &entrypoint.EntrypointList{}
}