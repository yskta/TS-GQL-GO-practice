// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewProject struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	UserID      string  `json:"userId"`
}

type NewTask struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Status      string  `json:"status"`
	ProjectID   string  `json:"projectId"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Project struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	User        *User   `json:"user"`
	Tasks       []*Task `json:"tasks,omitempty"`
}

type Query struct {
}

type Task struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description *string  `json:"description,omitempty"`
	Status      string   `json:"status"`
	Project     *Project `json:"project"`
}

type User struct {
	ID       string     `json:"id"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Projects []*Project `json:"projects,omitempty"`
}
