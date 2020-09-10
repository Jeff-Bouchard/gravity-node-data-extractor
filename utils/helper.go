package utils



func ContainsInt(list []int, item int) bool {
	return Contains(list, func (value interface{}) bool {
		return value == item
	})
}

func ContainsString(list []string, item string) bool {
	return Contains(list, func (value interface{}) bool {
		return value == item
	})
}

func Contains (
	list interface{},
	callback func (interface{}) bool,
) bool {

	for _, value := range list.([]interface{}) {
		if callback(value) { return true }
	}

	return false
}
