package core

func MTProtoPrivateID(id int64) int64 {
	// Telegram Bot API looks like -1001129109101
	return -(id % -1000000000000)
}
