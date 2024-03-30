package client

type Endpoints struct {
	BaseURL string `mapstructure:"baseUrl"`

	CoursesList        string `mapstructure:"coursesList"`
	CoursesListStarted string `mapstructure:"coursesListStarted"`
	CourseInfo         string `mapstructure:"courseInfo"`
	StartCourse        string `mapstructure:"startCourse"`
	AnswerQuestion     string `mapstructure:"answerQuestion"`

	StartArcade    string `mapstructure:"startArcade"`
	CompleteArcade string `mapstructure:"completeArcade"`
}
