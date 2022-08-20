package auth

func (helper *Helper) CheckServiceKey(key string) bool {
	return helper.serviceKey == key
}
