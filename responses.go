package main

type CoursesResponse struct {
	Courses []Course `json:"courses"`
}

type CourseInfoResponse struct {
	Course     CourseInfo `json:"course"`
	CourseUser any        `json:"course_user"`
	ShareURL   string     `json:"share_url"`
}

type CourseStartResponse struct {
	CourseUser CourseUser `json:"course_user"`
}

type AnswerQuestionResponse struct {
	CourseUser CourseUser `json:"course_user"`
}

type ArcadeGameSessionResponse struct {
	ArcadeGameSession ArcadeGameSession `json:"arcade_game_session"`
}
