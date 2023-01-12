package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	"github.com/gofor-little/ts"
	"go.uber.org/zap/zapcore"
)

type cloudWatchCore struct {
	opts              *options
	currentDay        int
	cloudWatchLogs    *cloudwatchlogs.Client
	logEventsList     *ts.LinkedList
	logGroupName      *string
	nextSequenceToken *string
	mutex             sync.RWMutex
}

type options struct {
	l            zapcore.Level
	accessKey    string
	accessSecret string
	endpoint     string
	//project      string
	//logstore     string
	profile      string
	region       string
	logGroupName string
}

func defaultOptions() *options {
	return &options{
		profile: "projectName",
		region:  "app",
	}
}

func Level(l zapcore.Level) Option {
	return func(c *options) {
		c.l = l
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

func WithEndpoint(endpoint string) Option {
	return func(alc *options) {
		alc.endpoint = endpoint
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
		config.WithSharedConfigProfile(opts.profile),
		config.WithRegion("us-east-2"),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     opts.accessKey,
				SecretAccessKey: opts.accessSecret,
				SessionToken:    "SESSION",
			},
		}),
	)
	if err != nil {
		return
	}

	logger := &cloudWatchCore{
		currentDay:     time.Now().Day(),
		cloudWatchLogs: cloudwatchlogs.NewFromConfig(cfg),
		logEventsList:  &ts.LinkedList{},
		logGroupName:   aws.String(opts.logGroupName),
	}
	if err = logger.checkLogGroup(context.TODO()); err != nil {
		return nil, err
	}
	go func() {
		ticker := time.NewTicker(time.Second / 5)

		for {
			<-ticker.C
			if err := logger.putLogs(context.TODO()); err != nil {
				log.Fatalf("failed to send logs to CloudWatch: %v", err)
			}
		}
	}()
	return logger, nil
}

func (c *cloudWatchCore) Close() {
	c.putLogs(context.TODO())
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
	//buf := ent.Level.String()
	//levelTitle := "level"
	//
	//for key, value := range c.globalFields {
	//	fields[key] = value
	//}
	//
	//fields["level"] = level

	data, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	messages := [][]byte{}

	// Check if the data is larger than the max input log event size.
	// If so, split it into a slice so the data can be added over multiple
	// events. This may break the JSON structure of very large amounts of
	// data as it will be split between multiple log events.
	for {
		if len(data) <= maxInputLogEventSize {
			messages = append(messages, data)
			break
		}

		messages = append(messages, data[:maxBatchInputLogEventSize])
		data = data[maxBatchInputLogEventSize:]
	}

	// Lock the mutex so we can queue our messages.
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Range over the messages and push them to the event list.
	for _, m := range messages {
		var tail *CloudWatchLogEventSlice

		// Fetch the tail from the event list. If the message can be added to the
		// tail add it. Otherwise push to the event list and add to the new tail.
		if !c.logEventsList.IsEmpty() && c.logEventsList.GetTail().(*CloudWatchLogEventSlice).canAdd(m) {
			tail = c.logEventsList.GetTail().(*CloudWatchLogEventSlice)
		} else {
			tail = &CloudWatchLogEventSlice{}
			c.logEventsList.Push(tail)
		}

		if err := tail.add(m); err != nil {
			return err
		}
	}

	return err
}

func (c *cloudWatchCore) Sync() error {
	return nil
}

// toString 任意类型转string
func toString(f *zapcore.Field) string {
	var key string
	if f == nil {
		return key
	}
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

// createLogGroup creates a log group in CloudWatch.
func (c *cloudWatchCore) createLogGroup(ctx context.Context) error {
	input := &cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: c.logGroupName,
	}

	_, err := c.cloudWatchLogs.CreateLogGroup(ctx, input)
	return err
}

// putLogs pops the oldest CloudWatchLogEventList off the queue, then
// writes it to CloudWatch.
func (c *cloudWatchCore) putLogs(ctx context.Context) error {
	if c.logEventsList.IsEmpty() {
		return nil
	}

	if err := c.checkLogStream(ctx); err != nil {
		return err
	}

	elements := c.logEventsList.Pop().(*CloudWatchLogEventSlice).logEvents.GetElements()
	inputLogEvents := make([]types.InputLogEvent, len(elements))

	for i, e := range elements {
		inputLogEvents[i] = e.(types.InputLogEvent)
	}

	input := &cloudwatchlogs.PutLogEventsInput{
		LogEvents:     inputLogEvents,
		LogGroupName:  c.logGroupName,
		LogStreamName: aws.String(time.Now().Format("2006-01-02")),
		SequenceToken: c.nextSequenceToken,
	}

	output, err := c.cloudWatchLogs.PutLogEvents(ctx, input)
	if err != nil {
		var dataAlreadyAccepted *types.DataAlreadyAcceptedException
		var invalidSequenceToken *types.InvalidSequenceTokenException

		if errors.As(err, &dataAlreadyAccepted) {
			input.SequenceToken = dataAlreadyAccepted.ExpectedSequenceToken
			_, err = c.cloudWatchLogs.PutLogEvents(ctx, input)
			if err != nil {
				return err
			}
		} else if errors.As(err, &invalidSequenceToken) {
			input.SequenceToken = invalidSequenceToken.ExpectedSequenceToken
			_, err = c.cloudWatchLogs.PutLogEvents(ctx, input)
			if err != nil {
				return err
			}
		}

		return err
	}

	c.nextSequenceToken = output.NextSequenceToken

	return nil
}

// createLogStream creates a log stream in CloudWatch.
func (c *cloudWatchCore) createLogStream(ctx context.Context) error {
	input := &cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  c.logGroupName,
		LogStreamName: aws.String(time.Now().Format("2006-01-02")),
	}

	_, err := c.cloudWatchLogs.CreateLogStream(ctx, input)
	return err
}

// checkLogGroup checks if the log group exists in CloudWatch.
// If it doesn't it will be created.
func (c *cloudWatchCore) checkLogGroup(ctx context.Context) error {
	logGroupExists, err := c.logGroupExists(ctx)
	if err != nil {
		return err
	}

	if logGroupExists {
		return nil
	}

	return c.createLogGroup(ctx)
}

// checkLogStream checks if the log stream exists in CloudWatch.
// If it doesn't it will be created.
func (c *cloudWatchCore) checkLogStream(ctx context.Context) error {
	logStreamExists, err := c.logStreamExists(ctx)
	if err != nil {
		return err
	}

	if logStreamExists {
		return nil
	}

	return c.createLogStream(ctx)
}

// logGroupExists checks if the log group exists in CloudWatch.
func (c *cloudWatchCore) logGroupExists(ctx context.Context) (bool, error) {
	input := &cloudwatchlogs.DescribeLogGroupsInput{
		LogGroupNamePrefix: c.logGroupName,
	}

	output, err := c.cloudWatchLogs.DescribeLogGroups(ctx, input)
	if err != nil {
		return false, err
	}

	if output.LogGroups != nil {
		for _, logGroup := range output.LogGroups {
			if *logGroup.LogGroupName == *c.logGroupName {
				return true, nil
			}
		}
	}

	return false, nil
}

// logStreamExists checks if the log stream exists in CloudWatch.
func (c *cloudWatchCore) logStreamExists(ctx context.Context) (bool, error) {
	input := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: c.logGroupName,
	}

	output, err := c.cloudWatchLogs.DescribeLogStreams(ctx, input)
	if err != nil {
		return false, nil
	}

	if output.LogStreams != nil {
		for _, logStream := range output.LogStreams {
			if *logStream.LogStreamName == time.Now().Format("2006-01-02") {
				return true, nil
			}
		}
	}

	return false, nil
}
