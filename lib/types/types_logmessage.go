package types

// LogTags are attachables for LogMessages to help describe them further.
type LogTag struct {
	// The display title of the log tag
	Title string `json:"title"`
	// The actual content description of the tag
	Content string `json:"content"`
}

// Log messages received from servers
type LogMessage struct {
	// The time that the message has been recorded at
	TimeStamp int64 `json:"timestamp"`
	// The content that has been captured from the message
	Content string `json:"content"`
	// A slice of LogTags to hold extra information
	Tags []LogTag `json:"tags"`
}
