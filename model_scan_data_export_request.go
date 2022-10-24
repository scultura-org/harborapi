/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The criteria to select the scan data to export.
type ScanDataExportRequest struct {
	// Name of the scan data export job
	JobName string `json:"job_name,omitempty"`
	// A list of one or more projects for which to export the scan data, currently only one project is supported due to performance concerns, but define as array for extension in the future.
	Projects []int64 `json:"projects,omitempty"`
	// A list of one or more labels for which to export the scan data, defaults to all if empty
	Labels []int64 `json:"labels,omitempty"`
	// A list of repositories for which to export the scan data, defaults to all if empty
	Repositories string `json:"repositories,omitempty"`
	// CVE-IDs for which to export data. Multiple CVE-IDs can be specified by separating using ',' and enclosed between '{}'. Defaults to all if empty
	CveIds string `json:"cveIds,omitempty"`
	// A list of tags enclosed within '{}'. Defaults to all if empty
	Tags string `json:"tags,omitempty"`
}
