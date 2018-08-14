package function

// Handle a serverless request
func Handle(req []byte) string {
	html := `
		<body style="background-color:#B0D0D3">
			<div style="display: block;margin-left: auto;margin-right: auto;width:50%;overflow:hidden">
				<img src="https://www.openfaas.com/images/openfaas/openfaas.png" style="display:block;margin:auto">
			</div>
			<div style="display: block;margin-left: auto;margin-right: auto;width:50%;overflow:hidden">
				<img src="https://www.openfaas.com/images/openfaas/logo.png" style="display:block;margin:auto">
				<img src="https://irp-cdn.multiscreensite.com/0e5c5b66/dms3rep/multi/tablet/animated-cloud-right-to-left-350x211-350x211.png" style="z-index:5;height:250px;position:relative;top:-70px;left:-10px;display:block;margin:auto">
			</div>
		</body>
		`
	return html
}
