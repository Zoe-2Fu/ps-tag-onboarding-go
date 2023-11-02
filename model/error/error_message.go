package errs

var (
	ErrorAgeMinimum          = "User does not meet minimum age requirement"
	ErrorEmailFormatT        = "User email must be properly formatted"
	ErrorEmailRequired       = "User email required"
	ErrorNameRequired        = "User first/last names required"
	ErrorNameUnique          = "User with the same first and last name already exists"
	ResponseUserNotFound     = "User not found"
	ResponseValidationFailed = "User did not pass validation"
)

type ErrorMessage struct {
	Error   string   `json:"error"`
	Details []string `json:"details"`
}
