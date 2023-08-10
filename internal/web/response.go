package web

type MemberResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	StudentId string `json:"student_id"`
}
