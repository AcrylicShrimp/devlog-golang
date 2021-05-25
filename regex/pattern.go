package regex

import "regexp"

var Whitespaces = regexp.MustCompile("\\s+")
var Slug = regexp.MustCompile("^[a-z\\d](?:[-a-z\\d]*[a-z\\d])?$")
