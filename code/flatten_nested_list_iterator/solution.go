package flatten_nested_list_iterator

type NestedIterator struct {
	stack []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	n := len(nestedList)
	stack := make([]*NestedInteger, 0, n)
	for i := n - 1; i >= 0; i-- {
		stack = append(stack, nestedList[i])
	}
	return &NestedIterator{stack}
}

func (itr *NestedIterator) Next() int {
	n := len(itr.stack)
	result := itr.stack[n-1]
	itr.stack = itr.stack[:n-1]
	return result.GetInteger()
}

func (itr *NestedIterator) HasNext() bool {
	for len(itr.stack) > 0 {
		n := len(itr.stack)
		top := itr.stack[n-1]
		if top.IsInteger() {
			return true
		}

		itr.stack = itr.stack[:n-1]
		list := top.GetList()
		for i := len(list) - 1; i >= 0; i-- {
			itr.stack = append(itr.stack, list[i])
		}
	}
	return false
}

type NestedInteger struct {
	isInteger   bool
	intValue    int
	nestedValue []*NestedInteger
}

func (ni *NestedInteger) IsInteger() bool {
	return ni.isInteger
}

func (ni *NestedInteger) GetInteger() int {
	return ni.intValue
}

func (ni *NestedInteger) SetInteger(value int) {
	ni.intValue = value
}

func (ni *NestedInteger) Add(intItem *NestedInteger) {
	ni.nestedValue = append(ni.nestedValue, intItem)
}

func (ni *NestedInteger) GetList() []*NestedInteger {
	return ni.nestedValue
}
