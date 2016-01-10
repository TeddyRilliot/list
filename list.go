package list

// Interface exposes methods to navigate through a list.
//
// Each type that implements Interface can have its own sorting mechanism
// (e.g.: sequential or random results).
type Interface interface {
	// Next returns the following item. If no item can be found, nil
	// is returned.
	Next() interface{}

	// Reset sets the container to its initial state.
	//
	// For example, a sequential list, if Reset, will return its first
	// element when calling Next.
	Reset()
}
