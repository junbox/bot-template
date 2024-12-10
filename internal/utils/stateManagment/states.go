package statemanagment

type UserState struct {
	State string
	Step  string
}

type StateRepo struct {
	States map[int]*UserState
	Vars   map[string]map[int]string
}

func (repo *StateRepo) SetState(userID int, state string, step string) {
	repo.States[userID] = &UserState{
		State: state,
		Step:  step,
	}
}

func (repo *StateRepo) GetState(userID int) *UserState {
	state := repo.States[userID]
	return state
}

func (repo *StateRepo) SetDefault(userID int) {
	repo.States[userID] = &UserState{
		State: "default",
		Step:  "default",
	}
	val := make(map[int]string)
	val[userID] = "true"
	repo.Vars["started"] = val
}

func (repo *StateRepo) SetVar(userID int, key string, value string) {
	val := make(map[int]string)
	val[userID] = value
	repo.Vars[key] = val
}

func (repo *StateRepo) GetVar(userID int, key string) string {
	return repo.Vars[key][userID]
}
