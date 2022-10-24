/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// The response for robot account creation.
type RobotCreated struct {
	// The ID of the robot
	Id int64 `json:"id,omitempty"`
	// The name of the tag
	Name string `json:"name,omitempty"`
	// The secret of the robot
	Secret string `json:"secret,omitempty"`
	// The creation time of the robot.
	CreationTime time.Time `json:"creation_time,omitempty"`
	// The expiration data of the robot
	ExpiresAt int64 `json:"expires_at,omitempty"`
}
