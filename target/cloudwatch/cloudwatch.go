package cloudwatch

type Option func(target *Cloudwatch)

type Cloudwatch struct {
	Namespace  string
	MetricName string
	Statistics []string
	Dimensions map[string]string
	Region     string
	Period     string

	Ref        string
	Hidden     bool
	Alias      string
	Expression string
}

func New(namespace string, metricName string, statistics []string, dimensions map[string]string, region string, period string, options ...Option) *Cloudwatch {
	cloudwatch := &Cloudwatch{
		Namespace:  namespace,
		MetricName: metricName,
		Statistics: statistics,
		Dimensions: dimensions,
		Region:     region,
		Period:     period,
	}

	for _, opt := range options {
		opt(cloudwatch)
	}

	return cloudwatch
}

func Ref(ref string) Option {
	return func(cloudwatch *Cloudwatch) {
		cloudwatch.Ref = ref
	}
}

func Hide() Option {
	return func(cloudwatch *Cloudwatch) {
		cloudwatch.Hidden = true
	}
}

func Alias(alias string) Option {
	return func(cloudwatch *Cloudwatch) {
		cloudwatch.Alias = alias
	}
}

func Expression(expr string) Option {
	return func(cloudwatch *Cloudwatch) {
		cloudwatch.Expression = expr
	}
}
