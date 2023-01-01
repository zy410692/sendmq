package Lib

import "fmt"

func UserInit() error {
	mq := NewMq()
	if mq == nil {
		return fmt.Errorf("mq init error")
	}
	defer mq.Channel.Close()

	err := mq.Channel.ExchangeDeclare(EXCHANGE_USER, "direct", false, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("Exchange error", err)
	}
	qs := fmt.Sprintf("%s,%s", QUEUE_NEWUSER, QUEUE_NEWUSER_UNION)
	err = mq.DecQueueAndBind(qs, ROUTER_KEY_USERREG, EXCHANGE_USER)
	if err != nil {
		return fmt.Errorf("decqueue and bind error", err)
	}
	return nil
}
