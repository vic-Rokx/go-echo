package db

import "example/api_gateway/models"

var Users = []models.User{
	{
		ID:       "1",
		Password: "securePassword123",
		Name:     "John Doe",
		Email:    "johndoe@example.com",
	},
	{
		ID:       "2",
		Password: "alicePass789",
		Name:     "Alice Blue",
		Email:    "alice.blue@example.com",
	},
	{
		ID:       "3",
		Password: "b0bS3cur3!",
		Name:     "Bob Smith",
		Email:    "bob.smith@example.net",
	},
	{
		ID:       "4",
		Password: "charlieP@ssw0rd",
		Name:     "Charlie Brown",
		Email:    "charlie.brown@example.org",
	},
}
