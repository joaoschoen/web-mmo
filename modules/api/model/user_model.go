package model

// TYPES
type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// // RESPONSE OBJECTS
// type GetUserResponse struct {
// 	Data SafeUser
// }

// type GetUserListResponse struct {
// 	Data   []SafeUser
// 	Paging Paging
// }

// type PostUserResponse struct {
// 	ID string
// }

// type PutUserResponse struct {
// 	Data SafeUser
// }

// type DeleteUserResponse struct {
// 	ID string
// }
