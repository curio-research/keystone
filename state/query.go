package state

type QueryContext struct {
	mappingCache map[int]bool
}

func NewQueryContext() *QueryContext {
	return &QueryContext{
		mappingCache: make(map[int]bool),
	}
}

// Clears a query context
func (ctx *QueryContext) Clear() {
	ctx.mappingCache = make(map[int]bool)
}

// Checks if query context includes a number
func (ctx *QueryContext) Includes(num int) bool {
	return ctx.mappingCache[num]
}

func (ctx *QueryContext) Add(num int) {
	ctx.mappingCache[num] = true
}

// Get all values as int array
func (ctx *QueryContext) GetAll() []int {
	var nums []int
	for num := range ctx.mappingCache {
		if ctx.mappingCache[num] {
			nums = append(nums, num)
		}
	}
	return nums
}

func ArrayIntersectionWithContext(ctx *QueryContext, nums1, nums2 []int) []int {

	// first clear query context
	ctx.Clear()

	for _, num := range nums1 {
		ctx.mappingCache[num] = true
	}

	// find the intersection
	var intersection []int
	for _, num := range nums2 {
		if ctx.mappingCache[num] {
			intersection = append(intersection, num)
		}
	}

	return intersection

}
