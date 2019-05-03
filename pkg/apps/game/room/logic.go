package room

import (
	"github.com/user/2019_1_newTeam2/models"
)

func (r *Room) CreateTask() models.GameQuestion{
	r.Logger.Log("CreateTaskStart")
	answer, err := r.DB.CreateTask()
	if err != nil {
		r.Logger.Log("CreateTask: ", err)
		return models.GameQuestion{}
	}
	return answer
}