package generator

type Workflow struct {
	Name         string `survey:"name"`
	Description  string `survey:"description"`
	Category     string `survey:"category"`
	BundleID     string `survey:"bundle_id"`
	CreatedBy    string `survey:"created_by"`
	WebAddress   string `survey:"web_address"`
	Version      string `survey:"version"`
	GoModPackage string `survey:"go_mod_package"`
}
