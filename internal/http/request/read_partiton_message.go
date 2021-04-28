package request

type ReadMessagesQuery struct {
	Topic     string `param:"topic"`
	ClusterId string `param:"cluster_id"`
	Partition int32  `param:"partition"`
}

type ReadMessagesForm struct {
}
