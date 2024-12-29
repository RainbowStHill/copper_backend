// Identity server generates distributed ID for other services
// like user service, message service, etc.
package identity_server

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/rainbowsthill/copper_backend/config"
)

// IDGenerator generates distributed ID with specified algorithm.
type IDGenerator interface {
	Generate() Unique
}

// Unique means an ID is unique all over the system.
type Unique interface {
	// Compare returns true or false when the given unique variable
	// is greater than or less than this unique variable respectively.
	Compare(Unique) bool
	String() string
}

// GetIDGenerator finds an ID generator using the specified
// algorithm.
func GetIDGenerator(algorithm string) IDGenerator {
	switch algorithm {
	case "SNOWFLAKE":
		return snowflakeIDGenerator
	default:
		return snowflakeIDGenerator
	}
}

var (
	snowflakeIDGenerator IDGenerator
)

func init() {
	cf, err := filepath.Abs("./config_files/config.yaml")
	if err != nil {
		log.Fatalf("Config file not exist: %v", err)
	}

	err = config.AddConfigFile(cf)
	if err != nil {
		log.Fatalf("Cannot load config file %s: %v", cf, err)
	}

	cfgPath := []string{"id_generator"}

	opt := SnowflakeOpt{}
	snowflakeCfgPath := append(cfgPath, "snowflake")

	timestampBits, err := config.GetBuiltInTypeConfig[int](cf, append(snowflakeCfgPath, "timestamp_bits"))
	assert(err, fmt.Sprintf("Failed to load %v", append(snowflakeCfgPath, "timestamp_bits")))
	opt.TimestampBits = int64(*timestampBits)

	dataCenterIDBits, err := config.GetBuiltInTypeConfig[int](cf, append(snowflakeCfgPath, "data_center_id_bits"))
	assert(err, fmt.Sprintf("Failed to load %v", append(snowflakeCfgPath, "data_center_id_bits")))
	opt.DataCenterIDBits = int64(*dataCenterIDBits)

	instanceIDBits, err := config.GetBuiltInTypeConfig[int](cf, append(snowflakeCfgPath, "instance_id_bits"))
	assert(err, fmt.Sprintf("Failed to load %v", append(snowflakeCfgPath, "instance_id_bits")))
	opt.InstanceIDBits = int64(*instanceIDBits)

	incrIDBits, err := config.GetBuiltInTypeConfig[int](cf, append(snowflakeCfgPath, "incr_id_bits"))
	assert(err, fmt.Sprintf("Failed to load %v", append(snowflakeCfgPath, "incr_id_bits")))
	opt.IncrIDBits = int64(*incrIDBits)

	startsAt, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2024-12-25 00:00:00 +0800 CST")
	opt.StartsAt = startsAt.UnixMilli()

	dataCenter, err := config.GetBuiltInTypeConfig[int](cf, append(snowflakeCfgPath, "data_center"))
	assert(err, fmt.Sprintf("Failed to load %v", append(snowflakeCfgPath, "data_center")))
	opt.DataCenter = int64(*dataCenter)

	instance, err := config.GetBuiltInTypeConfig[int](cf, append(snowflakeCfgPath, "instance"))
	assert(err, fmt.Sprintf("Failed to load %v", append(snowflakeCfgPath, "instance")))
	opt.Instance = int64(*instance)

	snowflakeIDGenerator = NewSnowflake(opt)
}

func assert(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
