package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type transport struct {
	headers map[string]string
	base    http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range t.headers {
		req.Header.Add(k, v)
	}
	base := t.base
	if base == nil {
		base = http.DefaultTransport
	}
	return base.RoundTrip(req)
}

type AssetDash struct {
	baseUrl string
	token   string
	client  *http.Client
}

func New(token string, opts ...AssetDashOption) *AssetDash {
	ad := &AssetDash{
		baseUrl: "https://mobile-api.assetdash.com/",
		token:   token,
		client: &http.Client{
			Transport: &transport{
				headers: map[string]string{
					"Authorization": "Bearer " + token,
					"User-Agent":    "Mozilla/5.0 (Windows; U; Windows NT 10.0; WOW64; en-US) AppleWebKit/533.33 (KHTML, like Gecko) Chrome/53.0.1715.150 Safari/534.9 Edge/14.90437",
				},
			},
		},
	}

	for _, option := range opts {
		option(ad)
	}

	return ad
}

func (ad *AssetDash) Courses(page int, limit int) (*CoursesResponse, error) {
	query := url.Values{}
	query.Add("page", fmt.Sprint(page))
	query.Add("limit", fmt.Sprint(limit))
	u, err := prepareUrl(ad.baseUrl, coursesEndpoint, query)
	if err != nil {
		return nil, err
	}

	resp, err := ad.client.Get(u.String())
	if err := responseErrorHandler(resp, err); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := &CoursesResponse{}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (ad *AssetDash) Course(id string) (*CourseInfoResponse, error) {
	query := url.Values{}
	query.Add("course_id", id)
	u, err := prepareUrl(ad.baseUrl, courseEndpoint, query)
	if err != nil {
		return nil, err
	}

	log.Println("looking for course", id)

	resp, err := ad.client.Get(u.String())
	if err := responseErrorHandler(resp, err); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := &CourseInfoResponse{}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (ad *AssetDash) StartCourse(id string) (*CourseStartResponse, error) {
	u, err := prepareUrl(ad.baseUrl, courseStartEndpoint, nil)
	if err != nil {
		return nil, err
	}

	log.Println("starting course", id)

	payload := map[string]interface{}{"course_id": id}
	body, _ := json.Marshal(payload)

	resp, err := ad.client.Post(u.String(), "application/json", bytes.NewBuffer(body))
	if err := responseErrorHandler(resp, err); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := &CourseStartResponse{}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (ad *AssetDash) AnswerQuestion(c_id, q_id, answer string) (*AnswerQuestionResponse, error) {
	u, err := prepareUrl(ad.baseUrl, answerQuestionEndpoint, nil)
	if err != nil {
		return nil, err
	}

	log.Println("answer question", q_id, "variant", answer)

	payload := map[string]interface{}{"course_id": c_id, "course_question_id": q_id, "answer": answer}
	body, _ := json.Marshal(payload)

	resp, err := ad.client.Post(u.String(), "application/json", bytes.NewBuffer(body))
	if err := responseErrorHandler(resp, err); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := &AnswerQuestionResponse{}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (ad *AssetDash) StartFlappyBirdArcade(timestamp int64) (*ArcadeGameSessionResponse, error) {
	query := url.Values{}
	query.Add("timestamp", fmt.Sprint(timestamp))
	u, err := prepareUrl(ad.baseUrl, "api/api_v4/games/arcade/games/start_session", query)
	if err != nil {
		return nil, err
	}

	payload := map[string]interface{}{"arcade_game_id": "14ebdfca-1154-49e4-bb4c-29f3c0da5bd8"}
	body, _ := json.Marshal(payload)

	resp, err := ad.client.Post(u.String(), "application/json", bytes.NewBuffer(body))
	if err := responseErrorHandler(resp, err); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := &ArcadeGameSessionResponse{}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (ad *AssetDash) CompleteFlappyBirdArcade(timestamp int64, sessionId string, arcadeData *FlappyBirdArcadeData) (*ArcadeGameSessionResponse, error) {
	query := url.Values{}
	query.Add("timestamp", fmt.Sprint(timestamp))
	u, err := prepareUrl(ad.baseUrl, "api/api_v4/games/arcade/games/complete_session", query)
	if err != nil {
		return nil, err
	}

	payload := map[string]interface{}{
		"arcade_game_id":         "14ebdfca-1154-49e4-bb4c-29f3c0da5bd8",
		"arcade_game_session_id": sessionId,
		"session_hash":           arcadeData.SessionHash,
		"data":                   arcadeData.Data}
	body, _ := json.Marshal(payload)

	resp, err := ad.client.Post(u.String(), "application/json", bytes.NewBuffer(body))
	if err := responseErrorHandler(resp, err); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := &ArcadeGameSessionResponse{}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
