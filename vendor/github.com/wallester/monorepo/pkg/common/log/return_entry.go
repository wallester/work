package log

type ReturnEntry struct {
	*Entry
}

func (e *ReturnEntry) Error(err error) error {
	e.Entry.Error(err)
	return err
}

func (e *ReturnEntry) Fatal(err error) error {
	e.Entry.Fatal(err)
	return err
}

func (e *ReturnEntry) Panic(err error) error {
	e.Entry.Panic(err)
	return err
}

func (e *ReturnEntry) Warn(err error) error {
	e.Entry.Warn(err)
	return err
}
