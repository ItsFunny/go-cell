package types

func CreateNoopSignEnvelope(
	protocol string, sequenceId string,
	data []byte) *Envelope {

	header := &EnvelopeHeader{
		Flag:       0,
		Length:     int64(len(data)),
		Protocol:   protocol,
		SequenceId: sequenceId,
	}
	pHeader := &Header{}
	payLoad := &Payload{
		Header: pHeader,
		Data:   data,
	}
	ret := &Envelope{
		Header:  header,
		Payload: payLoad,
	}

	return ret
}
