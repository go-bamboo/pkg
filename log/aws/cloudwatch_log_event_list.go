package aws

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/gofor-little/ts"
)

// These values come from CloudWatch, check cloudwatchlogs.CloudWatchLogs.PutLogEvents()
// comments for more information on the following values.
const (
	// The offset in bytes that is added to each InputLogEvent that is pushed to CloudWatch.
	inputLogEventOffset = 28
	// The max size in bytes a single InputLogEvent's message can be. Approx 256KB.
	maxInputLogEventSize = 256000
	// The max size in bytes that a batch of InputLogEvents can be. Approx 1MB.
	maxBatchInputLogEventSize = 1048576
	// The max amount of InputLogEvents that can be in a single batch.
	maxBatchInputLogEvents = 10000
)

// CloudWatchLogEventSlice stores a thread safe slice of InputLogEvents.
// size is the current size in bytes of the log messages taking into
// account the inputLogEventOffet.
type CloudWatchLogEventSlice struct {
	logEvents *ts.Slice
	size      int
}

// add adds a new InputLogEvent to the logEvents slice.
func (c *CloudWatchLogEventSlice) add(message []byte) error {
	if len(message) > maxInputLogEventSize {
		return fmt.Errorf("event size: %v is larger than the max event size: %v", len(message), maxInputLogEventSize)
	}

	if !c.canAdd(message) {
		return fmt.Errorf("max put size of %v exceeded or will be exceeded by adding to it", maxBatchInputLogEventSize)
	}

	// Build and validate the InputLogEvent with a timestamp.
	inputLogEvent := types.InputLogEvent{
		Message:   aws.String(string(message)),
		Timestamp: aws.Int64(time.Now().UnixNano() / int64(time.Millisecond)),
	}

	// Adjust the size to accounting for the inputLogEventOffet.
	// And add the InputLogEvent.
	c.size += len(message) + inputLogEventOffset
	c.logEvents.Add(inputLogEvent)

	return nil
}

// canAdd checks if message can be added to the logEvents slice.
func (c *CloudWatchLogEventSlice) canAdd(message []byte) bool {
	// Check if the event list is full.
	if c.isFull() {
		return false
	}

	// Check if the message including the inputLogEventOffer exceeeds the maxBatchInputLogEventSize.
	if c.size+len(message)+inputLogEventOffset > maxBatchInputLogEventSize {
		return false
	}

	return true
}

// isFull checks if the maxBatchInputLogEvents or maxBatchInputLogEventSize limits have been reached.
func (c *CloudWatchLogEventSlice) isFull() bool {
	if c.logEvents == nil {
		c.logEvents = &ts.Slice{}
	}

	if c.logEvents.Length() >= maxBatchInputLogEvents || c.size >= maxBatchInputLogEventSize {
		return true
	}

	return false
}
