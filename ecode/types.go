package ecode

// InternalServerMd new InternalServer error that is mapped to a 500 response.
func InternalServerMd(reason, message string, md map[string]string) error {
	return NewMetadataf(500, reason, md, message)
}
