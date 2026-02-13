package utils

func IsCompositeChannel(channelCode string) bool {
	if channelCode == "" {
		return false
	}
	return channelCode[0] == '9'
}
