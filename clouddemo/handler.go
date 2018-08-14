package function

// Handle a serverless request
func Handle(req []byte) string {
	//return fmt.Sprintf("Hello, Go. You said: %s", string(req))
	html := `<h1><center>WAT</center></h1>`
	return html
}
