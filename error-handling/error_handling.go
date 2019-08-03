package erratum

// Use opens a resourse, calls Frob func and then closes the that resource.
// Handle all errors and panics.
func Use(o ResourceOpener, input string) (err error) {
	var data Resource

	data, err = o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		// keep trying to open
		data, err = o()
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
