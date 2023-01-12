package Support

import (
	"net/http"
)

var Inserted int = http.StatusCreated
var Success int = http.StatusOK
var NotFound int = http.StatusBadRequest
var InvalidFormat int = http.StatusUnprocessableEntity
var Error int = http.StatusBadRequest
var Exsits int = http.StatusForbidden
var Accepted int = http.StatusAccepted
