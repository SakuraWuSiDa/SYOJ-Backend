package dao

import (
	"github.com/XGHXT/SYOJ-Backend/model"
	"github.com/XGHXT/SYOJ-Backend/pkg/mysql"
)

func GetAllSubmissionsByIdAndProblem(userID, problemID int64) ([]*model.Submission, error) {
	submissions := []*model.Submission{}
	err := mysql.DB.Where("user_id = ? && problem_id = ?", userID, problemID).Find(&submissions).Error
	return submissions, err
}

type Submissions struct {
	ID          int64  `json:"submissionID"`
	UserID      int64  `json:"userID"`
	QuestionsID int64  `json:"questionID"`
	Language    string `json:"langueage"`
	Score       int    `json:"score"`
	Result      string `json:"result"`
	Code        string `json:"code"`
}

func CreateSubmission(req *model.RunCodeParams, userID int64, sub *model.SubmitResult) (*model.Submission, error) {
	submission := &model.Submission{
		UserID:    userID,
		ProblemID: req.ProblemID,
		Language:  req.Language,
		Code:      req.Code,
		//Result:    judge.GetAnswerMsg(sub.AnswerCode),
		Score:     sub.Score,
		Time:      sub.Time,
		Memory:    sub.Memory,
	}

	err := mysql.DB.Select("UserID", "ProblemID", "Language", "Code", "Score", "Result", "Time", "Memory").Create(submission).Error

	return submission, err
}

func GetPersonSolved(userID int64) (int64, error) {
	var count int64
	err := mysql.DB.Model(&model.Submission{}).Where("user_id = ? AND result = ?", userID, "accept").Distinct("problem_id").Count(&count).Error
	return count, err
}

// GetPersonSubmission
func GetPersonSubmission(userID int64) (int64, error) {
	var count int64
	err := mysql.DB.Model(&model.Submission{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}