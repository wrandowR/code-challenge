package main

/*import (
	"fmt"
	"log"
	"os"
)

const (
	FILE_NAME      = "csv/students.csv"
	MAX_GOROUTINES = 10
)

func ProcessFile() {
	f, err := os.Open(FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	users := scanFile(f)

	// sequential processing
	// sequentialProcessing(users)

	// concurrent processing
	concurrentProcessing(users)
}

func concurrentProcessing(users []*User) {
	usersCh := make(chan []*User)
	unvisitedUsers := make(chan *User)
	go func() {
		usersCh <- users
	}()
	initializeWorkers(unvisitedUsers, usersCh, users)
	processUsers(unvisitedUsers, usersCh, len(users))
}

func initializeWorkers(unvisitedUsers <-chan *User, usersCh chan []*User, users []*User) {
	for i := 0; i < MAX_GOROUTINES; i++ {
		go func() {
			for user := range unvisitedUsers {
				sendSmsNotification(user)
				go func(user *User) {
					friendIds := user.FriendIds
					friends := []*User{}
					for _, friendId := range friendIds {
						friend, err := findUserById(friendId, users)
						if err != nil {
							fmt.Printf("Error %v\n", err)
							continue
						}
						friends = append(friends, friend)
					}

					_, ok := <-usersCh
					if ok {
						usersCh <- friends
					}
				}(user)
			}
		}()
	}
}
*/
