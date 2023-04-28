package message_queue

var TempUrlChan = make(chan map[string]string)
var TempRespDataChan = make(chan map[string]interface{}, 100)
