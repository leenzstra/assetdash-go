package main

import (
	"fmt"
	"log"
	"time"

	"github.com/leenzstra/assetdash-go/client"
	"github.com/leenzstra/assetdash-go/config"
	"github.com/leenzstra/assetdash-go/modules/arcade"
	"github.com/leenzstra/assetdash-go/modules/learn"
)

func processCourse(m *learn.LearnModule, id string) error {
	course, err := m.CourseInfo(id)
	if err != nil {
		return err
	}

	log.Println(" -- Processing course", course.Course.ID, course.Course.Title)

	_, err = m.StartCourse(course.Course.ID)
	if err != nil {
		return err
	}

	for _, question := range course.Course.Questions {
		_, err := m.AnswerQuestion(question.CourseID, question.ID, question.CorrectAnswer)
		if err != nil {
			return err
		}
	}

	log.Println(" -- Course proccessed", course.Course.ID, course.Course.Title)

	return nil
}

func processCourses(m *learn.LearnModule) error {
	for {
		r, err := m.Courses(1, 30)
		if err != nil {
			return err
		}

		if r == nil || len(r.Courses) == 0 {
			break
		}

		for _, course := range r.Courses {
			err = processCourse(m, course.ID)
			if err != nil {
				return fmt.Errorf("Fatal on %s %s", course.ID, err)
			}

			time.Sleep(1 * time.Second)
		}
	}

	return nil
}

func processFlappyBirdArcade(m *arcade.ArcadeModule, targetScore int) error {
	gameId := "14ebdfca-1154-49e4-bb4c-29f3c0da5bd8"
	startT := time.Now()

	score := targetScore
	coins := score * 3

	totalSessionTime := time.Duration(int64(score*3) * int64(time.Second))
	endT := startT.Add(totalSessionTime)

	session, err := m.StartArcade(gameId, startT.Unix())
	log.Printf("session: %+v", session)
	if err != nil {
		return err
	}

	// time.Sleep(totalSessionTime)

	data, err := m.ComputeSessionData(session.ArcadeGameSession.ID, coins, score)
	log.Printf("data: %+v", data)
	if err != nil {
		return err
	}

	session, err = m.CompleteArcade(endT.Unix(), gameId, session.ArcadeGameSession.ID, data)
	log.Printf("sessionEnd: %v", session)
	if err != nil {
		return err
	}

	return nil

}

func main() {
	config.SetupConfig()

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	client := client.New(cfg.Token, cfg.Endpoints)
	
	// testing learn
	lm := learn.New(client)
	if processCourses(lm) != nil {
		log.Fatal(err)
	} else {
		log.Println("ENDED")
	}

	// testing arcade
	am := arcade.New(client)
	if processFlappyBirdArcade(am, 10) != nil {
		log.Fatal(err)
	} else {
		log.Println("ENDED")
	}

}
