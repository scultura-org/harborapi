/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The namespace of registry
type Namespace struct {
	// The name of namespace
	Name string `json:"name,omitempty"`
	// The metadata of namespace
	Metadata interface{} `json:"metadata,omitempty"`
}