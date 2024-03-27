package main

import (
	"log"
	"time"
)

func processCourse(client *AssetDash, id string) error {
	course, err := client.Course(id)
	if err != nil {
		return err
	}

	log.Println(" -- Processing course", course.Course.ID, course.Course.Title)

	client.StartCourse(course.Course.ID)

	for _, question := range course.Course.Questions {
		_, err := client.AnswerQuestion(question.CourseID, question.ID, question.CorrectAnswer)
		if err != nil {
			return err
		}
	}

	log.Println(" -- Course proccessed", course.Course.ID, course.Course.Title)
	return nil
}

func main() {
	setupConfig()

	cfg, err := NewAssetDashConfig()
	if err != nil {
		log.Fatal(err)
	}

	client := New(cfg.Token)

	for {
		r, err := client.Courses(1, 30)
		if err != nil {
			log.Fatal(err)
		}

		if r == nil || len(r.Courses) == 0 {
			log.Println("ENDED")
			break
		}

		for _, course := range r.Courses {
			err = processCourse(client, course.ID)
			if err != nil {
				log.Fatal("Fatal on ", course.ID, err)
			}
			time.Sleep(1 * time.Second)
		}
	}

	// testing
	now := time.Now().Unix()
	session, err := client.StartFlappyBirdArcade(now)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("session: %v", session)

	time.Sleep(10 * time.Second)
	now = time.Now().Unix()
	data, err := flappyBirdEncryptedData(session.ArcadeGameSession.ID, 500, 500)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("data: %v", session)

	sessionEnd, err := client.CompleteFlappyBirdArcade(now, session.ArcadeGameSession.ID, data)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("sessionEnd: %v", sessionEnd)

}
