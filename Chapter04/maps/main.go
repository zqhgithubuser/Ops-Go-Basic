package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	users := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
	}
	// 添加元素
	users["073"] = "Tracy"
	return users
}

func getUser(id string) (string, bool) {
	users := getUsers()
	user, exists := users[id]
	return user, exists
}

func deleteUser(id string) {
	users := getUsers()
	// 删除元素
	delete(users, id)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}

	userID := os.Args[1]
	name, exists := getUser(userID)
	if !exists {
		fmt.Printf("Passed user ID (%v) not found.\nUsers:\n", userID)
		// 遍历 map
		for key, value := range getUsers() {
			fmt.Println("  ID:", key, "Name:", value)
		}
		os.Exit(1)
	}
	deleteUser(userID)
	fmt.Printf("User %s is deleted.\n", name)
}
