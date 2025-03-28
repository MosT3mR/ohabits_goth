package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"ohabits.com/internal/api"
	"ohabits.com/internal/handlers"
	"ohabits.com/internal/util"
)

type Response struct {
	Message string `json:"message"`
}

// Handler functions

func getHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(Response{Message: fmt.Sprintf("Get request for %v", params)})
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(Response{Message: fmt.Sprintf("Post reques for %v", params)})
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(Response{Message: fmt.Sprintf("Delete reques for %v", params)})
}

func Server() {

	r := mux.NewRouter()

	// API routes (JSON)

	// Login & Register (No Auth)
	r.HandleFunc("/api/register", api.Register).Methods("POST")
	r.HandleFunc("/api/login", api.Login).Methods("POST")

	// Protected Routes (Auth)
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(util.AuthMiddleware)

	// Habits (i need to fix the date in /api/habits_completed so it can accept (2025-03-22) insted of (2025-03-22T00:00:00Z) )
	protected.HandleFunc("/habits", api.GetHabits).Methods("GET")
	protected.HandleFunc("/habits", api.PostHabits).Methods("POST")
	protected.HandleFunc("/habits", api.PutHabits).Methods("PUT")
	protected.HandleFunc("/habits", api.DeleteHabit).Methods("DELETE")

	protected.HandleFunc("/habits_completed/{date}", api.GetHabitsCompleted).Methods("GET")
	protected.HandleFunc("/habits_completed", api.PostHabitCompleted).Methods("POST")
	protected.HandleFunc("/habits_completed", api.PutHabitCompleted).Methods("PUT")

	// Workout ( need to add the monthley view )
	protected.HandleFunc("/workout", api.GetWorkouts).Methods("GET")
	protected.HandleFunc("/workout/{id}", api.GetWorkouts).Methods("GET")
	protected.HandleFunc("/workout", api.PostWorkout).Methods("POST")
	protected.HandleFunc("/workout/{id}", api.PutWorkout).Methods("PUT")
	protected.HandleFunc("/workout/{id}", api.DeleteWorkout).Methods("DELETE")

	protected.HandleFunc("/workout_logs/{date}", api.GetWorkoutLog).Methods("GET")
	protected.HandleFunc("/workout_logs/{date}", api.PostWorkoutLog).Methods("POST")

	// Todo
	protected.HandleFunc("/todo/{date}", api.GetTodosByDate).Methods("GET")
	protected.HandleFunc("/todo", api.PostTodo).Methods("POST")
	protected.HandleFunc("/todo/{id}", api.PutTodo).Methods("PUT")
	protected.HandleFunc("/todo/{id]", api.DeleteTodo).Methods("DELETE")

	// Note
	protected.HandleFunc("/note/{date}", api.GetNoteByDate).Methods("GET")
	protected.HandleFunc("/note", api.PostNote).Methods("POST")
	protected.HandleFunc("/note/{id}", api.PutNote).Methods("PUT")
	protected.HandleFunc("/note/{id}", api.DeleteNote).Methods("DELETE")

	// Rate
	protected.HandleFunc("/rate/{date}", api.GetRate).Methods("GET")
	protected.HandleFunc("/rate", api.PostRate).Methods("POST")
	protected.HandleFunc("/rate/{id}", api.PutRate).Methods("PUT")

	// View mode
	protected.HandleFunc("/view/{month}", api.GetView).Methods("GET")

	// User
	protected.HandleFunc("/user", api.GetUser).Methods("GET")
	protected.HandleFunc("/user", api.PutUser).Methods("PUT")

	// Frontend routes (HTML)
	r.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	r.HandleFunc("/habits_completions", handlers.HabitsCompletedByDate).Methods("GET")
	r.HandleFunc("/todos", handlers.Todos).Methods("GET")
	r.HandleFunc("/notes", handlers.Notes).Methods("GET")
	r.HandleFunc("/mood_rating", handlers.Mood).Methods("GET")
	r.HandleFunc("/workout_loging", handlers.WorkoutLoging).Methods("GET")
	r.HandleFunc("/workout", handlers.GetWorkoutExercises).Methods("GET")

	// Serve static files (css, js, etc.)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
