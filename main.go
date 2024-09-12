package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var actions = []string{
	"logged in",
	"logged out",
	"create record",
	"delete record",
	"update record",
}

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getActivityInfo() string {
	out := fmt.Sprintf("[ID] %d | [EMAIL] %s\n [ACTIVITY LOG]\n", u.id, u.email)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s \n", i, item.action, item.timestamp)
	}
	return out
}

func generateUsers(count int) []User {
	users := make([]User, count)

	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@mail.ru", i+1),
			logs:  generateLogs(rand.Intn(1000)),
		}
	}
	return users
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)
	for i := range logs {
		logs[i] = logItem{
			timestamp: time.Now(),
			action:    actions[rand.Intn(len(actions)-1)],
		}
	}
	return logs
}

func saveUserInfo(user User, wg *sync.WaitGroup) error {

	defer wg.Done()

	time.Sleep(time.Microsecond * 10)

	fmt.Printf("WRITTING FILE FOR USER ID: %d\n", user.id)

	filename := fmt.Sprintf("logs/uid_%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(user.getActivityInfo())
	return err
}

func main() {
	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup

	t := time.Now()

	users := generateUsers(1000)

	wg.Add(1000)

	for _, user := range users {
		go saveUserInfo(user, &wg)
	}
	wg.Wait()

	fmt.Println("TIME ELAPSED:", time.Since(t).String())
}
