package validation

type Validator interface {
	Validate() error
	getError() ValidationError
}
