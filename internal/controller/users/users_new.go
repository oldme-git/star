// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package users

import (
    "star/api/users"
    userLogic "star/internal/logic/users"
)

type ControllerV1 struct {
    users *userLogic.Users
}

func NewV1() users.IUsersV1 {
    return &ControllerV1{
        users: userLogic.New(),
    }
}
