package models

import statemanagment "github.com/lex3man/playground/internal/utils/stateManagment"

type User struct {
	ID        int
	Name      string
	Username  string
	City      string
	Login     string
	PswdHache string
	Profile   *statemanagment.Profile
}

func (u User) Init(id int, username string) User {
	prof := statemanagment.Profile{}
	prof.Init()
	return User{
		ID:        id,
		Name:      username,
		Username:  username,
		City:      "Unknown",
		Login:     username,
		PswdHache: "000",
		Profile:   &prof,
	}
}

func (u *User) SetAsAdmin() {
	u.Profile.IsAdmin = true
}

func (u *User) SetStatus(status string) {
	u.Profile.Statuses = append(u.Profile.Statuses, status)
}

func (u *User) AddAchivment(caption string) {
	newTag := statemanagment.Tag{
		Caption: caption,
	}
	u.Profile.Achives = append(u.Profile.Achives, newTag)
}

func (u *User) GetStatuses() []string {
	return u.Profile.Statuses
}

func (u *User) GetAchives() []string {
	achivements := make([]string, 0)
	for _, a := range u.Profile.Achives {
		achivements = append(achivements, a.Caption)
	}
	return achivements
}
