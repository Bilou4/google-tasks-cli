package api

import (
	"errors"
	"fmt"

	. "github.com/bilou4/google-tasks-cli/constants"
	"google.golang.org/api/tasks/v1"
)

func GetLists() (map[string]*tasks.TaskList, error) {
	listId := make(map[string]*tasks.TaskList)

	r, err := Srv.Tasklists.List().Do()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to retrieve task lists. %v", err.Error()))
	}
	if len(r.Items) > 0 {
		for _, i := range r.Items {
			listId[i.Title] = i
		}
	} else {
		return nil, errors.New("No task lists found.")
	}
	return listId, nil
}

func GetTasks(listId string) (map[string]*tasks.Task, error) {
	taskId := make(map[string]*tasks.Task)

	tasks, err := Srv.Tasks.List(listId).Do()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to retrieve tasks from lists: %v. %v", listId, err))
	}
	for _, i := range tasks.Items {
		taskId[i.Title] = i

	}
	return taskId, nil
}

func AddList(listname string) error {
	t := &tasks.TaskList{
		Title: listname,
	}
	_, err := Srv.Tasklists.Insert(t).Do()

	return err
}

func RemoveList(listId string) error {
	return Srv.Tasklists.Delete(listId).Do()
}

func AddTask(tasklistId, taskName, dueDate, notes string) error {
	if dueDate != "" {
		dueDate += "T00:00:00.00Z"
	}
	task := &tasks.Task{
		Title: taskName,
		Due:   dueDate,
		Notes: notes,
	}
	_, err := Srv.Tasks.Insert(tasklistId, task).Do()
	return err
}

func RemoveTask(tasklistId, taskId string) error {
	return Srv.Tasks.Delete(tasklistId, taskId).Do()
}

func RenameList(listId, newName string) error {
	list := &tasks.TaskList{
		Title: newName,
		Id:    listId,
	}

	_, err := Srv.Tasklists.Update(listId, list).Do()
	if err != nil {
		return err
	}
	return nil
}

func TaskDone(listId, taskId string) error {
	task, err := Srv.Tasks.Get(listId, taskId).Do()
	if err != nil {
		return err
	}

	task.Status = "completed"
	_, err = Srv.Tasks.Update(listId, taskId, task).Do()
	if err != nil {
		return err
	}
	return nil
}
