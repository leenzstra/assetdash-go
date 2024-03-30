package learn

type Course struct {
	ID                       string `json:"id"`
	Created                  string `json:"created"`
	Updated                  string `json:"updated"`
	OrganizationID           string `json:"organization_id"`
	PartnerName              string `json:"partner_name"`
	PartnerLogoURL           string `json:"partner_logo_url"`
	Title                    string `json:"title"`
	Description              string `json:"description"`
	ImageURL                 string `json:"image_url"`
	BannerImageURL           string `json:"banner_image_url"`
	Status                   string `json:"status"`
	IsBoosted                bool   `json:"is_boosted"`
	IsCtaEnabled             bool   `json:"is_cta_enabled"`
	UtilizationUnavailable   bool   `json:"utilization_unavailable"`
	Reward                   Reward `json:"reward"`
	NumViews                 int    `json:"num_views"`
	IsLocked                 bool   `json:"is_locked"`
	WeeklyCompletionsPercent any    `json:"weekly_completions_percent"`
	WeeklyUnlockTime         string `json:"weekly_unlock_time"`
	Slug                     string `json:"slug"`
}

type Reward struct {
	ID       string  `json:"id"`
	Created  string  `json:"created"`
	Updated  string  `json:"updated"`
	CourseID string  `json:"course_id"`
	CashBack float64 `json:"cash_back"`
	Tickets  int     `json:"tickets"`
	Coins    int     `json:"coins"`
}

type Step struct {
	ID       string `json:"id"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
	CourseID string `json:"course_id"`
	Order    int    `json:"order"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	ImageURL string `json:"image_url"`
	Duration int    `json:"duration"`
}

type Question struct {
	ID            string `json:"id"`
	Created       string `json:"created"`
	Updated       string `json:"updated"`
	CourseID      string `json:"course_id"`
	Order         int    `json:"order"`
	Question      string `json:"question"`
	AnswerA       string `json:"answer_a"`
	AnswerB       string `json:"answer_b"`
	AnswerC       string `json:"answer_c"`
	AnswerD       string `json:"answer_d"`
	CorrectAnswer string `json:"correct_answer"`
}

type Cta struct {
	ID          string `json:"id"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	CourseID    string `json:"course_id"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	ImageURL    string `json:"image_url"`
	CtaURL      string `json:"cta_url"`
	CtaCourseID any    `json:"cta_course_id"`
	CtaDealID   any    `json:"cta_deal_id"`
}

type CourseInfo struct {
	Course
	Steps                 []Step     `json:"steps"`
	Questions             []Question `json:"questions"`
	Cta                   Cta        `json:"cta"`
	IsStandardUserAllowed bool       `json:"is_standard_user_allowed"`
}

type CourseUser struct {
	CourseID         string `json:"course_id"`
	CourseQuestionID string `json:"course_question_id"`
}