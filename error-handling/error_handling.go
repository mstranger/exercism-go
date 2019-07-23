package erratum

// Use opens a resourse, calls Frob func and then closes the that resource.
// Handle all errors and panics.
func Use(o ResourceOpener, input string) (err error) {
	var data Resource

	data, err = o()
	if err != nil {
		// if TransientError, then keep trying to open
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
		// just return an error otherwise
		return err
	}

	defer data.Close()

	// handle panic
	defer func() {
		if r := recover(); r != nil {
			// if error is FrobError
			if f, ok := r.(FrobError); ok {
				data.Defrob(f.defrobTag)
			}
			// any other error
			err = r.(error)
		}
	}()

	data.Frob(input)

	return err
}
