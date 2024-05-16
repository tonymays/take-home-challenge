package fstree

// SentinelError allows the declaration of constant errors
type SentinelError string

// Error satisfies the interface definition
func (e SentinelError) Error() string {
	return string(e)
}
