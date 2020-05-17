package container

// Array interface.
type Array interface {

	// Add insert item into array at the position by index,
	// preliminarily moved all items from this position (including this position) right for one cell.
	// Returns false if index is greater than current Len().
	Add(item interface{}, index int) bool

	// Append item to array, same as Add into index=Len().
	Append(item interface{})

	// Set item in array by this index, change existing item.
	// Returns false if index is out of bounds (greater or equal Len()).
	Set(item interface{}, index int) bool

	// Get item from index.
	// If index is out of bounds returns nil
	Get(index int) interface{}

	// Get existed item from index
	// If item by this index exists (is not out of bounds) will return true by second parameter, otherwise false.
	// Might be useful in case if nils saved in array intentionally.
	GetExisted(index int) (interface{}, bool)

	// Len returns number of items in array.
	Len() int

	// Capacity returns number of allocated places in array for this moment.
	Cap() int

	// Remove removes item by index, return this item and successful or not this operation was.
	// If index is out of range item returned will be nil.
	Remove(index int) (interface{}, bool)
}

// Convert converts our implementation of array to GO slice - helper to tests.
func Convert(array Array) []interface{} {
	arrayLen := array.Len()
	slice := make([]interface{}, arrayLen)

	for i := 0; i < arrayLen; i++ {
		slice[i] = array.Get(i)
	}

	return slice
}
