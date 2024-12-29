// Load configurations from config files in read time.
// Supply general and private configurations to specified module.
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/spf13/viper"
)

// config file's absolute path -> configurations
var configurations map[string]*atomic.Value
var configLock *sync.RWMutex

// AddConfigFile registers one more config file to the registry, returns os.NotExist
// when the given config file is not exist.
func AddConfigFile(cfgFileAbsPath string) error {
	// If config file already exists in registry, do nothing.
	configLock.RLock()
	if _, ok := configurations[cfgFileAbsPath]; ok {
		return nil
	}
	configLock.RUnlock()

	// Check if the config file exists.
	if _, err := os.Stat(cfgFileAbsPath); err != nil {
		return err
	}

	path, _ := filepath.Split(cfgFileAbsPath)
	viper.AddConfigPath(path)

	configLock.Lock()
	configurations[cfgFileAbsPath] = new(atomic.Value)
	configLock.Unlock()

	refreshConfigurations()
	return nil
}

// GetConfiguration returns the specified configuration record or nil when
// the record is not found.
func GetConfiguration(path string, recordPath []string) any {
	configLock.RLock()
	defer configLock.RUnlock()

	var records map[string]any
	if config, ok := configurations[path]; !ok {
		return nil
	} else if records, ok = config.Load().(map[string]any); !ok {
		return nil
	}

	for i, recordName := range recordPath {
		if record, ok := records[recordName]; !ok {
			return nil
		} else if i < len(recordPath)-1 {
			if records, ok = record.(map[string]any); !ok {
				return nil
			}
		}
	}

	return records[recordPath[len(recordPath)-1]]
}

// GetBuiltInTypeConfig gets configuration and cast it to specified type.
// If configuration not found or the given type is not correct, return an error.
func GetBuiltInTypeConfig[T int | string](path string, recordPath []string) (*T, error) {
	record := GetConfiguration(path, recordPath)
	if record == nil {
		return nil, fmt.Errorf("record %v not found", recordPath)
	}
	if cfg, ok := record.(T); !ok {
		return nil, fmt.Errorf("illegal type")
	} else {
		return &cfg, nil
	}
}

func init() {
	configurations = make(map[string]*atomic.Value)
	configLock = new(sync.RWMutex)
	go func() { // Load configurations from config files.
		// Refresh configurations for every 5 seconds.
		ticker := time.NewTicker(5 * time.Second)
		// All config files are in yaml format
		viper.SetConfigType("yaml")
		for {
			select {
			case <-ticker.C:
				refreshConfigurations()
			}
		}
	}()
}

func refreshConfigurations() {
	configLock.Lock()
	defer configLock.Unlock()
	for path, config := range configurations {
		// config file must exist.
		if _, err := os.Stat(path); err != nil {
			continue
		}
		_, fname := filepath.Split(path)
		viper.SetConfigName(strings.TrimSuffix(fname, filepath.Ext(fname)))
		if err := viper.ReadInConfig(); err != nil {
			continue
		}
		config.Store(viper.GetViper().Get(strings.TrimSuffix(fname, filepath.Ext(fname))))
	}
}
