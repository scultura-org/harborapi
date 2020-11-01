/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package harborapi

import (
	"time"
)

type Tag struct {
	// The ID of the tag
	Id int64 `json:"id,omitempty"`
	// The ID of the repository that the tag belongs to
	RepositoryId int64 `json:"repository_id,omitempty"`
	// The ID of the artifact that the tag attached to
	ArtifactId int64 `json:"artifact_id,omitempty"`
	// The name of the tag
	Name string `json:"name,omitempty"`
	// The push time of the tag
	PushTime time.Time `json:"push_time,omitempty"`
	// The latest pull time of the tag
	PullTime time.Time `json:"pull_time,omitempty"`
	// The immutable status of the tag
	Immutable bool `json:"immutable,omitempty"`
	// The attribute indicates whether the tag is signed or not
	Signed bool `json:"signed,omitempty"`
}
