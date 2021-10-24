package model

/*
* 与 User 结构体相关的请求
 */

type CreateUserRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
}

type CreateUserResponse struct {
	Username  string `json:"username"`
	UserID    int64  `json:"user_id"`
}

type CreateUserParams struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

type LoginParams struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type LoginRequest struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToekn string             `json:"access_token"`
	User        CreateUserResponse `json:"user"`
}

type UpdateUserParams struct {
	UserID          int64  `json:"user_id" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Gender          int8   `json:"gender" binding:"required"`
	StudentID       string `json:"student_id" binding:"required"`
	Class           string `json:"class" binding:"required"`
}

type UserDetailResponse struct {
	UserID          int64  `json:"userID"`
	AcceptCount     int64  `json:"acceptCount"`
	SubmissionCount int64  `json:"submissionCount"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Gender          int8   `json:"gender"`
	StudentID       string `json:"student_id"`
	Class           string `json:"class"`
}

/*
* 与 Problem 相关的请求参数
 */

type CreateProblemRequest struct {
	Title         string                 `json:"title" binding:"required"`
	TimeLimit     uint                   `json:"time_limit" binding:"required"`
	MemoryLimit   uint                   `json:"memory_limit" binding:"required"`
	Author        string                 `json:"author" binding:"required"`
	Source        string                 `json:"source"`
	Background    string                 `json:"background"`
	Statement     string                 `json:"statement"`
	Input         string                 `json:"input"`
	Output        string                 `json:"output"`
	ExamplesIn    string                 `json:"examples_in"`
	ExamplesOut   string                 `json:"examples_out"`
	Hint          string                 `json:"hint"`
	IsOpen        bool                   `json:"is_open"`
	Tags          string                 `json:"tags"`
}

type ListProblemRequest struct {
	PageNum  int `form:"pageNum" binding:"required,min=1"`
	PageSize int `form:"pageSize" binding:"required,min=5,max=10"`
}

// DeleteProblem params
type DeleteProblemParams struct {
	ProblemID int64 `form:"problemID"`
}

/*
* 与 Submission 相关的请求参数
 */

type RunCodeParams struct {
	ProblemID int64  `json:"problemID"`
	Language  string `json:"language"`
	Code      string `json:"code"`
}

type SubmitResult struct {
	AnswerCode int   `json:"answer_code"`
	Time       int   `json:"time"`
	Score      int   `json:"score"`
	Memory     int64 `json:"memory"`
}

/*
*	与文章相关的请求参数
 */

type CreateCategoryParams struct {
	ProblemID int64  `json:"problemID"`
	Content   string `json:"content" binding:"required"`
}

type GetCategoryByProblemParams struct {
	ProblemID int64 `form:"problemID"`
}

// GetAllCategories Params
type GetAllCategoriesParams struct {
	PageNum  int `form:"pageNum" binding:"required,min=1"`
	PageSize int `form:"pageSize" binding:"required,min=5,max=10"`
}

// DeleteCategory And GetCategoryDetails
type CategoryParams struct {
	CategoryID int64 `form:"categoryID"`
}

/*
*	与比赛相关的请求参数
 */
type CreateContextParams struct {
	Title     string `json:"title"`
	StartTime Time   `json:"startTime"`
	EndTime   Time   `json:"endTime"`
	Author    string `json:"author"`
}

type GetContextListParams struct {
	PageNum  int `form:"pageNum" binding:"required,min=1"`
	PageSize int `form:"pageSize" binding:"required,min=5,max=10"`
}

type GetContextParams struct {
	ContextID int64 `params:"contextID"`
}

type DeleteContextParams struct {
	ContextID int64 `form:"contextID"`
}

type UpdateContextParams struct {
	ID        int64  `json:"contextID"`
	Title     string `json:"title"`
	StartTime Time   `json:"startTime"`
	EndTime   Time   `json:"endTime"`
	Author    string `json:"author"`
}

type AddProblemParams struct {
	ProblemID int64  `json:"problemID" binding:"required"`
	ContextID int64  `json:"contextID" binding:"required"`
	Title     string `json:"title"`
}

type PorblemInContextParams struct {
	ProblemID int64 `form:"problemID" binding:"required"`
	ContextID int64 `form:"contextID" binding:"required"`
}

type DeleteProblemInContext struct {
	ProblemID int64 `form:"problemID" binding:"required"`
	ContextID int64 `form:"contextID" binding:"required"`
}

type ContextProblemParams struct {
	ContextID int64 `form:"contextID"`
	PageNum   int   `form:"pageNum" binding:"required,min=1"`
	PageSize  int   `form:"pageSize" binding:"required,min=5,max=10"`
}

type ContextProblemResponse struct {
	InContext       bool   `json:"inContext"`
	TimeLimit       int    `json:"time_limit"`
	MemoryLimit     int    `json:"memory_limit"`
	ProblemID       int64  `json:"problemID"`
	ProblemName     string `json:"problemName"`
	Author          string `json:"author"`
	DifficultyLevel string `json:"difficulty_level"`
}

type GetContextProblemParams struct {
	ContextID int64 `form:"contextID"`
}
