// Identity server generates distributed ID for other services
// like user service, message service, etc.
package identity_server

// IDGenerator generates distributed ID with specified algorithm.
type IDGenerator interface {
	Generate() Unique
}

// Unique means an ID is unique all over the system.
type Unique interface {
	// Compare returns true or false when the given unique variable
	// is greater than or less than this unique variable respectively.
	Compare(Unique) bool
	String() string
}

// NewIDGenerator creates a new ID generator using the specified
// algorithm.
func NewIDGenerator(algorithm string) IDGenerator {
	switch algorithm {
	default:
		return nil
	}
}
