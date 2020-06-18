package cloudwatch_test

import (
	"github.com/breeze7086/grabana/target/cloudwatch"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewCloudwatchTargetCanBeCreated(t *testing.T) {
	req := require.New(t)

	namespace := "LBProvisionedLCU"
	metricName := "ProvisionedLCU"
	statistics := []string{"Average"}
	dimensions := map[string]string{
		"Loadbalancer Name": "9f2e1c28-default-prometheu-0928",
	}
	region := "us-east-1"
	period := ""

	target := cloudwatch.New(namespace, metricName, statistics, dimensions, region, period)

	req.Equal(namespace, target.Namespace)
	req.Equal(metricName, target.MetricName)
	req.Equal(statistics, target.Statistics)
	req.Equal(dimensions, target.Dimensions)
	req.Equal(region, target.Region)
	req.Equal(period, target.Period)
}
