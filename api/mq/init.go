package mq

var (
	mq *redisMQ
)

// Init 初始化
func Init() {
	mq = &redisMQ{}
	mq.Init()
	// go StartConsumeAdRequestLog()
	// go StartConsumeDeviceInfo()
}
