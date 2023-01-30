package aws

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/go-kratos/kratos/v2/errors"
)

type cloudWatchWriteSyncer struct {
	opts *options
	c    *cloudwatchlogs.Client
	//logGroupStream    *string
	nextSequenceToken *string
	queue             chan *types.InputLogEvent
	wg                sync.WaitGroup
}

func newCloudWatchWriteSyncer(opts *options) (ws *cloudWatchWriteSyncer, err error) {
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
	ws = &cloudWatchWriteSyncer{
		opts:  opts,
		c:     cloudwatchlogs.NewFromConfig(cfg),
		queue: make(chan *types.InputLogEvent, queueCap),
	}
	if err = ws.checkLogGroup(context.TODO()); err != nil {
		return nil, err
	}
	ws.wg.Add(1)
	go func() {
		defer func() {
			ws.wg.Done()
		}()
		tk := time.Tick(time.Second / 2)
		for {
			<-tk
			if err := ws.checkLogGroup(context.TODO()); err != nil {
				fmt.Printf("err: %v", err)
			}
			var inputLogEvents []types.InputLogEvent
		handleGet:
			ev, isOpen := <-ws.queue
			if isOpen {
				inputLogEvents = append(inputLogEvents, *ev)
			}
			if len(ws.queue) > 0 && len(inputLogEvents) < queueCap {
				goto handleGet
			}
			if len(inputLogEvents) > 0 {
				if err := ws.putLogs(context.TODO(), inputLogEvents); err != nil {
					fmt.Printf("failed to send logs to CloudWatch: %v", err)
				}
			}
			if !isOpen {
				fmt.Printf("cloudwatch exit\n")
				break
			}
		}
	}()
	return ws, nil
}

func (ws *cloudWatchWriteSyncer) Write(ev *types.InputLogEvent) {
	ws.queue <- ev
}

func (ws *cloudWatchWriteSyncer) Sync() error {
	return nil
}

func (ws *cloudWatchWriteSyncer) Close() {
	close(ws.queue)
	ws.wg.Wait()
}

// putLogs pops the oldest CloudWatchLogEventList off the queue, then
// writes it to CloudWatch.
func (ws *cloudWatchWriteSyncer) putLogs(ctx context.Context, logEvents []types.InputLogEvent) error {
	input := &cloudwatchlogs.PutLogEventsInput{
		LogEvents:     logEvents,
		LogGroupName:  aws.String(ws.opts.logGroupName),
		LogStreamName: aws.String(time.Now().Format("2006-01-02")),
		SequenceToken: ws.nextSequenceToken,
	}
	output, err := ws.c.PutLogEvents(ctx, input)
	if err != nil {
		var dataAlreadyAccepted *types.DataAlreadyAcceptedException
		var invalidSequenceToken *types.InvalidSequenceTokenException
		if errors.As(err, &dataAlreadyAccepted) {
			input.SequenceToken = dataAlreadyAccepted.ExpectedSequenceToken
			_, err = ws.c.PutLogEvents(ctx, input)
			if err != nil {
				return err
			}
		} else if errors.As(err, &invalidSequenceToken) {
			input.SequenceToken = invalidSequenceToken.ExpectedSequenceToken
			_, err = ws.c.PutLogEvents(ctx, input)
			if err != nil {
				return err
			}
		}
		return err
	}
	ws.nextSequenceToken = output.NextSequenceToken
	return nil
}

// checkLogGroup checks if the log group exists in CloudWatch.
// If it doesn't it will be created.
func (ws *cloudWatchWriteSyncer) checkLogGroup(ctx context.Context) error {
	input := &cloudwatchlogs.DescribeLogGroupsInput{
		LogGroupNamePrefix: aws.String(ws.opts.logGroupName),
	}
	output, err := ws.c.DescribeLogGroups(ctx, input)
	if err != nil {
		return err
	}
	if output.LogGroups != nil {
		for _, logGroup := range output.LogGroups {
			if *logGroup.LogGroupName == ws.opts.logGroupName {
				return ws.checkLogStream(ctx)
			}
		}
	}
	return ws.createLogGroup(ctx)
}

// createLogGroup creates a log group in CloudWatch.
func (ws *cloudWatchWriteSyncer) createLogGroup(ctx context.Context) error {
	input := &cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: aws.String(ws.opts.logGroupName),
	}
	_, err := ws.c.CreateLogGroup(ctx, input)
	if err != nil {
		return err
	}
	return ws.checkLogStream(ctx)
}

// checkLogStream checks if the log stream exists in CloudWatch.
// If it doesn't it will be created.
func (ws *cloudWatchWriteSyncer) checkLogStream(ctx context.Context) error {
	input := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(ws.opts.logGroupName),
	}
	output, err := ws.c.DescribeLogStreams(ctx, input)
	if err != nil {
		return err
	}
	if output.LogStreams != nil {
		for _, logStream := range output.LogStreams {
			if *logStream.LogStreamName == time.Now().Format("2006-01-02") {
				ws.nextSequenceToken = output.NextToken
				return nil
			}
		}
	}
	return ws.createLogStream(ctx)
}

// createLogStream creates a log stream in CloudWatch.
func (ws *cloudWatchWriteSyncer) createLogStream(ctx context.Context) error {
	input := &cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  aws.String(ws.opts.logGroupName),
		LogStreamName: aws.String(time.Now().Format("2006-01-02")),
	}
	_, err := ws.c.CreateLogStream(ctx, input)
	return err
}
