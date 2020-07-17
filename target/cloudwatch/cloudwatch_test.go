package cloudwatch_test

import (
	"github.com/breeze7086/grabana/target/cloudwatch"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewCloudwatchTargetCanBeCreated(t *testing.T) {
	req := require.New(t)

	namespace := "AWS/Logs"
	metricName := "IncomingBytes"
	statistics := []string{"Average"}
	dimensions := map[string]string{}
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

func TestRefCanBeConfigured(t *testing.T) {
	req := require.New(t)

	target := cloudwatch.New("", "", []string{}, map[string]string{}, "", "",
		cloudwatch.Ref("A"))

	req.Equal("A", target.Ref)
	req.False(target.Hidden)
}

func TestTargetCanBeHidden(t *testing.T) {
	req := require.New(t)

	target := cloudwatch.New("", "", []string{}, map[string]string{}, "", "",
		cloudwatch.Hide())

	req.True(target.Hidden)
}
