/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type RetentionExecution struct {
	Id        int64  `json:"id,omitempty"`
	PolicyId  int64  `json:"policy_id,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
	Status    string `json:"status,omitempty"`
	Trigger   string `json:"trigger,omitempty"`
	DryRun    bool   `json:"dry_run,omitempty"`
}