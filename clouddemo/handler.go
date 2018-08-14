package function

// Handle a serverless request
func Handle(req []byte) string {
	//return fmt.Sprintf("Hello, Go. You said: %s", string(req))
	/*
		html := `
				<body style="background-color:#B0D0D3">
					<div style="display: block;margin-left: auto;margin-right: auto;width:50%;position=absolute;">
						<div style="position:relative;margin:50%">
							<img src="https://www.openfaas.com/images/openfaas/logo.png">
							<img src="https://irp-cdn.multiscreensite.com/0e5c5b66/dms3rep/multi/tablet/animated-cloud-right-to-left-350x211-350x211.png" style="z-index:5;height:250px;position:relative;top:-70px;left:-20px;">
						</div>
					</div>
				</body>
				`
	*/
	html := `
		<body style="background-color:#B0D0D3">
			<div style="display: block;margin-left: auto;margin-right: auto;width:50%;overflow:hidden">
				<img src="https://www.openfaas.com/images/openfaas/logo.png" style="display:block">
				<img src="https://irp-cdn.multiscreensite.com/0e5c5b66/dms3rep/multi/tablet/animated-cloud-right-to-left-350x211-350x211.png" style="z-index:5;height:250px;position:relative;top:-70px;left:-20px;display:block">
			</div>
		</body>
		`
	return html
}
