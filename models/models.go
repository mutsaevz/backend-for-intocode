package models

import "gorm.io/gorm"

// PaymentStatus
type PaymentStatus string

const (
	Paid    PaymentStatus = "paid"
	Unpaid  PaymentStatus = "unpaid"
	Partial PaymentStatus = "partial"
)

// StudyStatus
type StudyStatus string

const (
	Learning  StudyStatus = "learning"
	JobSearch StudyStatus = "job_search"
	Offer     StudyStatus = "offer"
	Working   StudyStatus = "working"
)

// Структры Student
type Student struct {
	gorm.Model
	FullName      string        `json:"full_name"`
	Email         string        `json:"email"`
	Telegram      string        `json:"telegram"`
	GroupID       uint          `json:"group_id"`
	TuitionTotal  int           `json:"tuition_total"`
	TuitionPaid   int           `json:"tuition_paid"`
	PaymentStatus PaymentStatus `json:"payment_status"`
	StudyStatus   StudyStatus   `json:"study_status"`
}

// Для POST/PATCH запросов
type StudentRequest struct {
	FullName      string        `json:"full_name"`
	Email         string        `json:"email"`
	Telegram      string        `json:"telegram"`
	GroupID       uint          `json:"group_id"`
	TuitionTotal  int           `json:"tuition_total"`
	TuitionPaid   int           `json:"tuition_paid"`
	PaymentStatus PaymentStatus `json:"payment_status"`
	StudyStatus   StudyStatus   `json:"study_status"`
}

// Структуры Group
type Group struct {
	gorm.Model
	Title       string `json:"title"`
	CurrentWeek int    `json:"current_week"`
	TotalWeeks  int    `json:"total_weeks"`
	IsFinished  bool   `json:"is_finished"`
}

type GroupEmpty struct {
	Title       string `json:"title"`
	CurrentWeek int    `json:"current_week"`
	TotalWeeks  int    `json:"total_weeks"`
	IsFinished  bool   `json:"is_finished"`
}

// Структуры Note
type Note struct {
	gorm.Model
	StudentID uint   `json:"student_id"`
	Author    string `json:"author"`
	Text      string `json:"text"`
}

type NoteEmpty struct {
	StudentID uint   `json:"student_id"`
	Author    string `json:"author"`
	Text      string `json:"text"`
}