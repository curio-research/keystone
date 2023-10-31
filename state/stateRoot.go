package state

import (
	"crypto/sha256"
	"fmt"
	"sort"
)

// calculate a deterministic state root for a game world
// in the future: used to check if a world is executed correctly by peers

func CalculateWorldStateRootHash(w *GameWorld) string {

	worldStateRoot := ""

	// get all table names and sort names alphabetically
	tableNames := make([]string, 0, len(w.Tables))
	for k := range w.Tables {
		tableNames = append(tableNames, k)
	}

	// sort tables alphabetically
	tableNames = SortStrings(tableNames)

	// sort all entities within each table alphabetically
	for _, tableName := range tableNames {
		tableRootHash := ""

		table := w.Tables[tableName]

		// get and sort all entities
		sortedEntities := SortInts(table.All())

		for _, entity := range sortedEntities {
			entityValue, _ := table.Get(entity)
			strVal := ConvertStructToString(entityValue)

			newHash := HashStringsTogether(tableRootHash, strVal)
			tableRootHash = newHash
		}

		worldStateRoot = HashStringsTogether(worldStateRoot, tableRootHash)
	}

	return worldStateRoot
}

func HashStringsTogether(str1, str2 string) string {
	// Create a new SHA-256 hasher
	hasher := sha256.New()

	// Concatenate the two strings together and hash the result
	data := []byte(str1 + str2)
	hasher.Write(data)

	// Get the hash sum as a byte slice
	hashBytes := hasher.Sum(nil)

	// Convert the hash to a hexadecimal string
	hashString := fmt.Sprintf("%x", hashBytes)

	return hashString
}

func ConvertStructToString(inputStruct interface{}) string {
	// Use fmt.Sprintf to format the struct as a string
	return fmt.Sprintf("%#v", inputStruct)
}

// deep copy and sort array of strings
func SortStrings(input []string) []string {
	// Make a copy of the original array
	originalCopy := make([]string, len(input))
	copy(originalCopy, input)

	// Sort the original array in place
	sort.Strings(originalCopy)

	return originalCopy
}

func SortInts(input []int) []int {
	// Make a copy of the original array
	originalCopy := make([]int, len(input))
	copy(originalCopy, input)

	// Sort the original array in place
	sort.Ints(originalCopy)

	return originalCopy
}
