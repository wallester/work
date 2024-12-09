package errors

// ICauser is interface for getting original error cause.
// Can be implemented by different error packages.
type ICauser interface {
	Cause() error
}
