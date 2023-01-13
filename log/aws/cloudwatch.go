package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/go-bamboo/pkg/log/core"
	"github.com/go-kratos/kratos/v2/errors"
	"go.uber.org/zap/zapcore"
)

const queueCap = 1000

type cloudWatchCore struct {
	opts *options
	c    *cloudwatchlogs.Client
	//logGroupStream    *string
	nextSequenceToken *string
	queue             chan *types.InputLogEvent
	wg                sync.WaitGroup
}

type options struct {
	l            zapcore.Level
	accessKey    string
	accessSecret string
	sessionToken string
	endpoint     string
	profile      string
	region       string
	logGroupName string
}

func defaultOptions() *options {
	return &options{
		profile: "default",
		region:  "app",
	}
}

func Level(l zapcore.Level) Option {
	return func(c *options) {
		c.l = l
	}
}

func WithRegion(region string) Option {
	return func(alc *options) {
		alc.region = region
	}
}

func WithAccessKey(ak string) Option {
	return func(alc *options) {
		alc.accessKey = ak
	}
}

func WithAccessSecret(as string) Option {
	return func(alc *options) {
		alc.accessSecret = as
	}
}

func WithSessionToken(sk string) Option {
	return func(alc *options) {
		alc.sessionToken = sk
	}
}

func WithEndpoint(endpoint string) Option {
	return func(alc *options) {
		alc.endpoint = endpoint
	}
}

func WithLogGroupName(logGroupName string) Option {
	return func(alc *options) {
		alc.logGroupName = logGroupName
	}
}

type Option func(alc *options)

// NewCloudWatchCore new a cloud watch logger with options.
func NewCloudWatchCore(options ...Option) (c core.Logger, err error) {
	opts := defaultOptions()
	for _, o := range options {
		o(opts)
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		//config.WithSharedConfigProfile(opts.profile),
		config.WithRegion(opts.region),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     opts.accessKey,
				SecretAccessKey: opts.accessSecret,
				SessionToken:    opts.sessionToken,
			},
		}),
	)
	if err != nil {
		return
	}

	logger := &cloudWatchCore{
		opts:  opts,
		c:     cloudwatchlogs.NewFromConfig(cfg),
		queue: make(chan *types.InputLogEvent, queueCap),
	}
	if err = logger.checkLogGroup(context.TODO()); err != nil {
		return nil, err
	}
	logger.wg.Add(1)
	go func() {
		defer func() {
			logger.wg.Done()
		}()
		tk := time.Tick(time.Second / 2)
		for {
			<-tk
			var inputLogEvents []types.InputLogEvent
		handleGet:
			ev, isOpen := <-logger.queue
			if isOpen {
				inputLogEvents = append(inputLogEvents, *ev)
			}
			if len(logger.queue) > 0 && len(inputLogEvents) < queueCap {
				goto handleGet
			}
			if len(inputLogEvents) > 0 {
				if err := logger.putLogs(context.TODO(), inputLogEvents); err != nil {
					fmt.Printf("failed to send logs to CloudWatch: %v", err)
				}
			}
			if !isOpen {
				fmt.Printf("cloudwatch exit\n")
				break
			}
		}
	}()
	return logger, nil
}

func (c *cloudWatchCore) Close() {
	close(c.queue)
	c.wg.Wait()
}

func (c *cloudWatchCore) With(fields []zapcore.Field) zapcore.Core {
	//clone := c.clone()
	//addFields(clone.enc, fields)
	//return clone
	return nil
}

func (c *cloudWatchCore) Enabled(lvl zapcore.Level) bool {
	return lvl >= c.opts.l
}

func (c *cloudWatchCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if ent.Level >= c.opts.l {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *cloudWatchCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	data := make(map[string]string, len(fields)+4)
	data["name"] = ent.LoggerName
	data["msg"] = ent.Message
	data["caller"] = ent.Caller.String()
	data["stack"] = ent.Stack
	data["level"] = ent.Level.String()
	for _, field := range fields {
		data[field.Key] = toString(field)
	}
	msg, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ev := new(types.InputLogEvent)
	ev.Message = aws.String(string(msg))
	ev.Timestamp = aws.Int64(time.Now().UnixNano() / int64(time.Millisecond))
	//c.putLogs(context.TODO(), ev)
	c.queue <- ev
	return nil
}

func (c *cloudWatchCore) Sync() error {
	return nil
}

// toString 任意类型转string
func toString(f zapcore.Field) string {
	var key string
	switch f.Type {
	case zapcore.BinaryType:
		key = string(f.Interface.([]byte))
	case zapcore.BoolType:
		key = fmt.Sprint(f.Integer == 1)
	case zapcore.ByteStringType:
		key = string(f.Interface.([]byte))
	case zapcore.Complex128Type:
		key = fmt.Sprint(f.Interface.(complex128))
	case zapcore.Complex64Type:
		key = fmt.Sprint(f.Interface.(complex64))
	case zapcore.DurationType:
		key = fmt.Sprintf("%fs", time.Duration(f.Integer).Seconds())
	case zapcore.Float64Type:
		key = fmt.Sprintf("%f", math.Float64frombits(uint64(f.Integer)))
	case zapcore.Float32Type:
		key = fmt.Sprintf("%f", math.Float32frombits(uint32(f.Integer)))
	case zapcore.Int64Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.Int32Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.Int16Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.Int8Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.StringType:
		key = f.String
	case zapcore.TimeType:
		if f.Interface != nil {
			key = fmt.Sprint(time.Unix(0, f.Integer).In(f.Interface.(*time.Location)))
		} else {
			// Fall back to UTC if location is nil.
			key = fmt.Sprint(time.Unix(0, f.Integer))
		}
	case zapcore.TimeFullType:
		key = f.Interface.(time.Time).Format("2006-01-02 15:04:05")
	case zapcore.Uint64Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.Uint32Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.Uint16Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.Uint8Type:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.UintptrType:
		key = fmt.Sprintf("%d", f.Integer)
	case zapcore.ReflectType:
		key = fmt.Sprintf("%v", f.Interface)
	}
	return key
}

// putLogs pops the oldest CloudWatchLogEventList off the queue, then
// writes it to CloudWatch.
func (c *cloudWatchCore) putLogs(ctx context.Context, logEvents []types.InputLogEvent) error {
	input := &cloudwatchlogs.PutLogEventsInput{
		LogEvents:     logEvents,
		LogGroupName:  aws.String(c.opts.logGroupName),
		LogStreamName: aws.String(time.Now().Format("2006-01-02")),
		SequenceToken: c.nextSequenceToken,
	}
	output, err := c.c.PutLogEvents(ctx, input)
	if err != nil {
		var dataAlreadyAccepted *types.DataAlreadyAcceptedException
		var invalidSequenceToken *types.InvalidSequenceTokenException
		if errors.As(err, &dataAlreadyAccepted) {
			input.SequenceToken = dataAlreadyAccepted.ExpectedSequenceToken
			_, err = c.c.PutLogEvents(ctx, input)
			if err != nil {
				return err
			}
		} else if errors.As(err, &invalidSequenceToken) {
			input.SequenceToken = invalidSequenceToken.ExpectedSequenceToken
			_, err = c.c.PutLogEvents(ctx, input)
			if err != nil {
				return err
			}
		}
		return err
	}
	c.nextSequenceToken = output.NextSequenceToken
	return nil
}

// checkLogGroup checks if the log group exists in CloudWatch.
// If it doesn't it will be created.
func (c *cloudWatchCore) checkLogGroup(ctx context.Context) error {
	input := &cloudwatchlogs.DescribeLogGroupsInput{
		LogGroupNamePrefix: aws.String(c.opts.logGroupName),
	}
	output, err := c.c.DescribeLogGroups(ctx, input)
	if err != nil {
		return err
	}
	if output.LogGroups != nil {
		for _, logGroup := range output.LogGroups {
			if *logGroup.LogGroupName == c.opts.logGroupName {
				return c.checkLogStream(ctx)
			}
		}
	}
	return c.createLogGroup(ctx)
}

// createLogGroup creates a log group in CloudWatch.
func (c *cloudWatchCore) createLogGroup(ctx context.Context) error {
	input := &cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: aws.String(c.opts.logGroupName),
	}
	_, err := c.c.CreateLogGroup(ctx, input)
	if err != nil {
		return err
	}
	return c.checkLogStream(ctx)
}

// checkLogStream checks if the log stream exists in CloudWatch.
// If it doesn't it will be created.
func (c *cloudWatchCore) checkLogStream(ctx context.Context) error {
	input := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(c.opts.logGroupName),
	}
	output, err := c.c.DescribeLogStreams(ctx, input)
	if err != nil {
		return err
	}
	if output.LogStreams != nil {
		for _, logStream := range output.LogStreams {
			if *logStream.LogStreamName == time.Now().Format("2006-01-02") {
				c.nextSequenceToken = output.NextToken
				return nil
			}
		}
	}
	return c.createLogStream(ctx)
}

// createLogStream creates a log stream in CloudWatch.
func (c *cloudWatchCore) createLogStream(ctx context.Context) error {
	input := &cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  aws.String(c.opts.logGroupName),
		LogStreamName: aws.String(time.Now().Format("2006-01-02")),
	}
	_, err := c.c.CreateLogStream(ctx, input)
	return err
}
