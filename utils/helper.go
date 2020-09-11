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


type LinkedListItem struct {
	Val interface{}
	Prev *LinkedListItem
	Next *LinkedListItem
}

type LinkedListBuilder struct {}

func (builder *LinkedListBuilder) Build (onNext func() interface{}) *LinkedListItem {
	list := &LinkedListItem{}

	return builder.buildList(list, onNext)
}

func (builder *LinkedListBuilder) buildList (list *LinkedListItem, onNext func() interface{}) *LinkedListItem {
	nextValue := onNext()

	if nextValue == nil {
		return list
	}

	list.Val = &LinkedListItem{ Val: nextValue }

	return builder.buildList(list.Next, onNext)
}
