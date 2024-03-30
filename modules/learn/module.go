package learn

import (
	"fmt"
	"strconv"

	"github.com/leenzstra/assetdash-go/client"
	"github.com/leenzstra/assetdash-go/modules"
)

var _ modules.Module = (*LearnModule)(nil)

type LearnModule struct {
	*client.AssetDashClient
}

func New(client *client.AssetDashClient) *LearnModule {
	return &LearnModule{client}
}

func (m *LearnModule) Name() string {
	return "Learn Module"
}

func (m *LearnModule) Courses(page int, limit int) (*CoursesResponse, error) {
	endpoint, err := m.BuildURL(m.Endpoints.CoursesList)
	if err != nil {
		return nil, err
	}

	response := &CoursesResponse{}
	r, err := m.Client.R().SetQueryParams(map[string]string{
		"page":  strconv.Itoa(page),
		"limit": strconv.Itoa(limit),
	}).SetResult(response).Get(endpoint)

	if err != nil || r.IsError() {
		return nil, fmt.Errorf("%s %s", r.Error(), err)
	}

	return response, nil
}

func (m *LearnModule) StartedCourses(page int, limit int) (*CoursesResponse, error) {
	endpoint, err := m.BuildURL(m.Endpoints.CoursesListStarted)
	if err != nil {
		return nil, err
	}

	response := &CoursesResponse{}
	r, err := m.Client.R().SetQueryParams(map[string]string{
		"page":  strconv.Itoa(page),
		"limit": strconv.Itoa(limit),
	}).SetResult(response).Get(endpoint)

	if err != nil || r.IsError() {
		return nil, fmt.Errorf("%s %s", r.Error(), err)
	}

	return response, nil
}

func (m *LearnModule) CourseInfo(id string) (*CourseInfoResponse, error) {
	endpoint, err := m.BuildURL(m.Endpoints.CourseInfo)
	if err != nil {
		return nil, err
	}

	response := &CourseInfoResponse{}
	r, err := m.Client.R().SetQueryParams(map[string]string{
		"course_id":  id,
	}).SetResult(response).Get(endpoint)

	if err != nil || r.IsError() {
		return nil, fmt.Errorf("%s %s", r.Error(), err)
	}

	return response, nil
}

func (m *LearnModule) StartCourse(id string) (*CourseStartResponse, error) {
	endpoint, err := m.BuildURL(m.Endpoints.StartCourse)
	if err != nil {
		return nil, err
	}

	response := &CourseStartResponse{}
	r, err := m.Client.R().SetBody(map[string]interface{}{
		"course_id": id,
	}).SetResult(response).Post(endpoint)

	if err != nil || r.IsError() {
		fmt.Println(string(r.Body()))
		return nil, fmt.Errorf("%s %s", r.Error(), err)
	}

	return response, nil
}

func (m *LearnModule) AnswerQuestion(course_id, question_id, answer string) (*AnswerQuestionResponse, error) {
	endpoint, err := m.BuildURL(m.Endpoints.AnswerQuestion)
	if err != nil {
		return nil, err
	}

	response := &AnswerQuestionResponse{}
	r, err := m.Client.R().SetBody(map[string]interface{}{
		"course_id": course_id,
		"course_question_id": question_id,
		"answer": answer,
	}).SetResult(response).Post(endpoint)

	if err != nil || r.IsError() {
		fmt.Println(string(r.Body()))
		return nil, fmt.Errorf("%s %s", r.Error(), err)
	}

	return response, nil
}