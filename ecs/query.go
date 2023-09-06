package ecs

type QueryContext struct {
	mappingCache map[int]bool
}

func NewQueryContext() QueryContext {
	return QueryContext{
		mappingCache: make(map[int]bool),
	}
}

func (ctx QueryContext) Clear() {
	ctx.mappingCache = nil
	ctx.mappingCache = make(map[int]bool)

}

func (ctx QueryContext) Terminate() {
	ctx.mappingCache = nil
}

func ArrayIntersectionWithContext(ctx QueryContext, nums1, nums2 []int) []int {

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
