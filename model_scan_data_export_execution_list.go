/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The list of scan data export executions
type ScanDataExportExecutionList struct {
	// The list of scan data export executions
	Items []ScanDataExportExecution `json:"items,omitempty"`
}
