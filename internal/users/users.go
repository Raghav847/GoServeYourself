package users

import (
	"errors"
	"fmt"
	"net/mail"
)

var ErrNoResultsFound = errors.New("no results found")

type User struct {
	FirstName string
	LastName  string
	Email     mail.Address
}

type Manager struct {
	users []User
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) AddUser(firstName string, lastName string, email string) error {
	if firstName == "" {
		return fmt.Errorf("invalid first name: %q", firstName)
	}
	if lastName == "" {
		return fmt.Errorf("invalid last name: %q", lastName)
	}

	existingUser, err := m.GetUserByName(firstName, lastName)
	if err != nil && !errors.Is(err, ErrNoResultsFound) {
		return fmt.Errorf("err checking if the user is present: %v", err)
	}

	if existingUser != nil {
		return errors.New("user already exists")
	}

	parsedAddress, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("invalid email: %s", email)
	}

	newUser := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     *parsedAddress,
	}

	m.users = append(m.users, newUser)

	return nil

}

func (m *Manager) GetUserByName(first string, last string) (*User, error) {
	for i, user := range m.users {
		if user.FirstName == first && user.LastName == last {
			result := m.users[i]
			return &result, nil
		}
	}
	return nil, ErrNoResultsFound
}
