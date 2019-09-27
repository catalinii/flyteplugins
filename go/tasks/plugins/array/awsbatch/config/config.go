package config

import (
	"github.com/lyft/flyteplugins/go/tasks/aws"
	"github.com/lyft/flytestdlib/config"
)

//go:generate pflags Config --default-var defaultConfig

type Config struct {
	JobStoreConfig     JobStoreConfig        `json:"jobStoreConfig" pflag:",Config for job store"`
	JobDefCacheSize    int                   `json:"defCacheSize" pflag:",Maximum job definition cache size as number of items. Caches are used as an optimization to lessen the load on AWS Services."`
	GetRateLimiter     aws.RateLimiterConfig `json:"getRateLimiter" pflag:",Rate limiter config for batch get API."`
	DefaultRateLimiter aws.RateLimiterConfig `json:"defaultRateLimiter" pflag:",Rate limiter config for all batch APIs except get."`
	MaxArrayJobSize    int64                 `json:"maxArrayJobSize" pflag:",Maximum size of array job."`
	MinRetries         int32                 `json:"minRetries" pflag:",Minimum number of retries"`
	MaxRetries         int32                 `json:"maxRetries" pflag:",Maximum number of retries"`
	// Provide additional environment variable pairs that plugin authors will provide to containers
	DefaultEnvVars       map[string]string `json:"defaultEnvVars" pflag:"-,Additional environment variable that should be injected into every resource"`
	MaxErrorStringLength int               `json:"maxErrLength" pflag:",Determines the maximum length of the error string returned for the array."`
	RoleAnnotationKey    string            `json:"roleAnnotationKey" pflag:",Map key to use to lookup role from task annotations."`
	ResyncPeriod         config.Duration   `json:"resyncPeriod" pflag:",Defines the duration for syncing job details from AWS Batch."`
}

type JobStoreConfig struct {
	CacheSize      int `json:"jacheSize" pflag:",Maximum informer cache size as number of items. Caches are used as an optimization to lessen the load on AWS Services."`
	Parallelizm    int `json:"parallelizm"`
	BatchChunkSize int `json:"batchChunkSize" pflag:",Determines the size of each batch sent to GetJobDetails api."`
}

var (
	defaultConfig = &Config{
		JobStoreConfig: JobStoreConfig{
			CacheSize:      10000,
			Parallelizm:    10,
			BatchChunkSize: 1000,
		},
		JobDefCacheSize: 10000,
		MaxArrayJobSize: 5000,
		GetRateLimiter: aws.RateLimiterConfig{
			Rate:  15,
			Burst: 20,
		},
		DefaultRateLimiter: aws.RateLimiterConfig{
			Rate:  15,
			Burst: 20,
		},
		MinRetries:           1,
		MaxRetries:           10,
		MaxErrorStringLength: 500,
	}

	configSection = aws.MustRegisterSubSection("batch", defaultConfig)
)

func GetConfig() *Config {
	return configSection.GetConfig().(*Config)
}