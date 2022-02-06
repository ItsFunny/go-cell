package channel


type IData interface {
	ID() interface{}
}
type IChan interface {
	Take() IData
	Push(task IData) (int, error)
}


type ChannelID string

type ChannelShim struct {
}
type ChannelDescriptor struct {
	ID                  byte
	Priority            int
	SendQueueCapacity   int
	RecvMessageCapacity int
	RecvBufferCapacity  int
	MaxSendBytes        uint
}

type Channel struct {
	Id ChannelID
	Ch chan IData
}

func (this *Channel) Close() {
	close(this.Ch)
}

type Envelope struct {
	ChannelId ChannelID
	Data      IData
}
