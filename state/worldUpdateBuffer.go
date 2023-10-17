package state

type WorldUpdateBuffer struct {
	bufferWorld  IWorld
	currentWorld IWorld
	deletionMap  map[int]string
}

// TODO we can also make the register tables function a part of IWorld
func NewWorldUpdateBuffer(currentWorld *GameWorld) *WorldUpdateBuffer {
	bufferWorld := NewWorld()
	bufferWorld.entityManager = currentWorld.entityManager
	for _, table := range currentWorld.Tables {
		bufferWorld.AddTable(table.TableInterface())
	}

	return &WorldUpdateBuffer{
		currentWorld: currentWorld,
		bufferWorld:  bufferWorld,
		deletionMap:  make(map[int]string),
	}
}

func (w *WorldUpdateBuffer) Add(obj interface{}, tableName string) int {
	entity := w.bufferWorld.Add(obj, tableName)
	delete(w.deletionMap, entity)
	return entity
}

func (w *WorldUpdateBuffer) AddSpecific(entity int, obj interface{}, tableName string) int {
	entity = w.bufferWorld.AddSpecific(entity, obj, tableName)
	delete(w.deletionMap, entity)
	return entity
}

func (w *WorldUpdateBuffer) Get(entity int, tableName string) (any, bool) {
	if _, ok := w.deletionMap[entity]; ok {
		return nil, false
	}

	val, found := w.bufferWorld.Get(entity, tableName)
	if !found {
		val, found = w.currentWorld.Get(entity, tableName)
	}

	return val, found
}

func (w *WorldUpdateBuffer) Set(entity int, obj interface{}, tableName string) interface{} {
	val := w.bufferWorld.Set(entity, obj, tableName)
	delete(w.deletionMap, entity)
	return val
}

func (w *WorldUpdateBuffer) Delete(entity int, tableName string) int {
	w.deletionMap[entity] = tableName
	return entity
}

func (w *WorldUpdateBuffer) Filter(filter interface{}, fieldNames []string, tableName string) []int {
	endResultsMap := map[int]struct{}{}
	bufferWorldResults := w.bufferWorld.Filter(filter, fieldNames, tableName)
	for _, res := range bufferWorldResults {
		endResultsMap[res] = struct{}{}
	}

	// if it is in buffer world but we don't get the result from the buffer world filter, we don't want it
	currentWorldResults := w.currentWorld.Filter(filter, fieldNames, tableName)
	for _, res := range currentWorldResults {
		if _, ok := w.deletionMap[res]; !ok {
			if _, ok := w.bufferWorld.Get(res, tableName); !ok {
				endResultsMap[res] = struct{}{}
			}
		}
	}

	var endRes []int
	for i := range endResultsMap {
		endRes = append(endRes, i)
	}

	return endRes
}

func (w *WorldUpdateBuffer) Entities(tableName string) []int {
	currEnt := w.currentWorld.Entities(tableName)
	bufferEnt := w.bufferWorld.Entities(tableName)

	u := union(currEnt, bufferEnt)
	for v := range w.deletionMap {
		delete(u, v)
	}

	var ret []int
	for v := range u {
		ret = append(ret, v)
	}

	return ret
}

func (w *WorldUpdateBuffer) ApplyUpdates() {
	tableUpdates := w.bufferWorld.GetTableUpdates()
	for _, update := range tableUpdates {
		switch update.OP {
		case UpdateOP:
			w.currentWorld.Set(update.Entity, update.Value, update.Table)
		}
	}

	for entityToDelete, tableName := range w.deletionMap {
		w.currentWorld.Delete(entityToDelete, tableName)
	}
}

func (w *WorldUpdateBuffer) GetTableUpdates() TableUpdateArray {
	return w.bufferWorld.GetTableUpdates() // should this panic?
}

func union(arr1 []int, arr2 []int) map[int]struct{} {
	m := make(map[int]struct{})
	for _, v := range arr1 {
		m[v] = struct{}{}
	}
	for _, v := range arr2 {
		m[v] = struct{}{}
	}

	return m
}
