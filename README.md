# google-tasks-cli


A simple (non official) CLI to manage your lists of tasks in your Google Tasks application.


Here are available subcommands :
~~~bash
google-tasks-cli done <ListName> <TaskName> # sets task status to 'completed'
google-tasks-cli add <ListName> <TaskName> --due-date <YYYY-MM-DD> --notes <someNotes> # adds a task in an existing list
google-tasks-cli addList <ListName> # creates a new list
google-tasks-cli rmList <ListName> # removes an existing list
google-tasks-cli renameList <old-name> <new-name> # renames a list
google-tasks-cli lists # prints all existing lists
google-tasks-cli get <ListName> # prints all tasks from an existing list
google-tasks-cli rm <ListName> <TaskName> # removes a task from an existing list
~~~


## Prerequisites

Before using the app, you need to follow [these steps](https://developers.google.com/tasks/firstapp#register) so as to be able to send requests to the Google Tasks API.

## TODO

- [ ] update a task (flags: notes + date) # updates an existing task
- [ ] deleteAllTasks <ListName> # removes all tasks from an existing list
- [ ] use ID instead of Title in keys because listname/taskname aren't unique
- [ ] google-tasks-cli deleteCompletedTasks <ListName> # removes all completed tasks
