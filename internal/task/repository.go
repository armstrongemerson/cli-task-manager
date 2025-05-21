package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Repository manages task storage
type Repository struct {
	tasks  []Task
	nextID int
	file   string
}

// NewRepository creates a new task repository
func NewRepository(file string) (*Repository, error) {
	repo := &Repository{
		tasks:  make([]Task, 0),
		nextID: 1,
		file:   file,
	}

	// Try to load existing tasks
	if err := repo.Load(); err != nil {
		// If file doesn't exist, that's OK
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	return repo, nil
}

// Add creates a new task
func (r *Repository) Add(title, description string) Task {
	task := Task{
		ID:          r.nextID,
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
	}

	r.tasks = append(r.tasks, task)
	r.nextID++
	return task
}

// Get returns a task by ID
func (r *Repository) Get(id int) (Task, error) {
	for _, task := range r.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, fmt.Errorf("task with ID %d not found", id)
}

// List returns all tasks
func (r *Repository) List() []Task {
	return r.tasks
}

// Update modifies an existing task
func (r *Repository) Update(task Task) error {
	for i, t := range r.tasks {
		if t.ID == task.ID {
			r.tasks[i] = task
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", task.ID)
}

// Save writes tasks to file
func (r *Repository) Save() error {
	data, err := json.Marshal(r.tasks)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(r.file, data, 0644)
}

// Load reads tasks from file
func (r *Repository) Load() error {
	data, err := ioutil.ReadFile(r.file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &r.tasks); err != nil {
		return err
	}

	// Find the highest ID to set nextID correctly
	for _, task := range r.tasks {
		if task.ID >= r.nextID {
			r.nextID = task.ID + 1
		}
	}

	return nil
}
