package configuration

var (
	CLAMAV = ClamAvModel{
		SCAN_EXECUTION_PATH: GetEnvOrDefaultString("CLAMAV_EXECUTION_PATH", "clamdscan"),
	}

	REST = RESTModel{
		SCANFILE_FORM_KEY: GetEnvOrDefaultString("REST_SCANFILE_FORM_KEY", "file"),
	}
)
