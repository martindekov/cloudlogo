package function

// Handle a serverless request
func Handle(req []byte) string {
	//return fmt.Sprintf("Hello, Go. You said: %s", string(req))
	html := `<h1><div style="width:800px; margin:0 auto;">WAT</div></h1>`
	return html
}
