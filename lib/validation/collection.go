package validation

type Collection struct {
	err Error
}

func NewCollection() *Collection {
	return &Collection{
		err: make(Error),
	}
}

func (c *Collection) Add(key string, validator Validator) {
	err := validator.getError()

	if len(err) > 0 {
		c.err[key] = err
	}
}

func (c *Collection) Resolve() error {
	if len(c.err) == 0 {
		return nil
	}

	return c.err
}
