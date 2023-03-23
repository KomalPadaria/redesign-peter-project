package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/entities"
	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
	"github.com/pkg/errors"
)

type RequestBodyType interface {
	entities.AnswerRequestBody | entities.QuestionnairesStatus | entities.Answer | entities.QuestionnaireAnswerFeedback
}

func decodeBodyFromRequest[T RequestBodyType](req *T, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return errors.WithMessage(httpError.ErrBadRequestBody, "Bad request body")
	}

	defer r.Body.Close()

	return nil
}

func decodeBodyFromMulipartRequest[T RequestBodyType](req *T, key string, r *http.Request) error {
	jsonData := r.FormValue(key)

	err := json.NewDecoder(strings.NewReader(jsonData)).Decode(&req)
	if err != nil {
		return errors.WithMessage(httpError.ErrBadRequestBody, "Bad request body")
	}

	return nil
}

func decodeGetQuestionnairesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	req := &entities.GetQuestionnairesRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
	}

	return req, nil
}

func decodeGetQuestionnairesByCategoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	category := params["category"]
	if strings.TrimSpace(category) == "" {
		return nil, httpError.NewErrBadOrInvalidPathParameter("category")
	}

	req := &entities.GetQuestionnairesByCategoryRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		Category:    category,
	}

	return req, nil
}

func decodeAddAnswerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	var answerRequestBody entities.AnswerRequestBody
	err = decodeBodyFromRequest(&answerRequestBody, r)
	if err != nil {
		return nil, err
	}

	req := &entities.PostAnswerRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		Answers:     answerRequestBody,
	}

	return req, nil
}

func decodeAddEngineerFeedbackRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	questionnaireUUID, err := uuid.Parse(params["questionnaire_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("questionnaire_id")
	}

	answerUUID, err := uuid.Parse(params["answer_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("answer_id")
	}

	var requestBody entities.QuestionnaireAnswerFeedback
	err = decodeBodyFromRequest(&requestBody, r)
	if err != nil {
		return nil, err
	}

	req := &entities.EngineerFeedbackRequest{
		CompanyUuid:        compUUID,
		UserUuid:           userUUID,
		QuestionnairesUuid: questionnaireUUID,
		AnswerUuid:         answerUUID,
		Feedback:           requestBody,
	}

	return req, nil
}

func decodeSubmitQuestionnairesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	var requestBody entities.QuestionnairesStatus
	err = decodeBodyFromRequest(&requestBody, r)

	req := &entities.SubmitQuestionnairesRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		Status:      requestBody,
	}

	return req, nil
}

func decodeUploadEvidencesRequest(_ context.Context, r *http.Request) (interface{}, error) {

	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	questionnaireUUID, err := uuid.Parse(params["questionnaire_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("questionnaire_id")
	}

	var answerRequestBody entities.Answer
	err = decodeBodyFromMulipartRequest(&answerRequestBody, "data", r)
	if err != nil {
		return nil, err
	}

	err = r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		return nil, err
	}

	req := &entities.AnswerWithEvidenceRequest{
		CompanyUuid:        compUUID,
		UserUuid:           userUUID,
		QuestionnairesUuid: questionnaireUUID,
		Answer:             answerRequestBody,
		Files:              r.MultipartForm.File["files"],
	}

	return req, nil
}

func decodeUpdateAnswerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	questionnaireUUID, err := uuid.Parse(params["questionnaire_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("questionnaire_id")
	}

	answerUUID, err := uuid.Parse(params["answer_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("answer_id")
	}

	var answerRequestBody entities.Answer
	err = decodeBodyFromMulipartRequest(&answerRequestBody, "data", r)
	if err != nil {
		return nil, err
	}

	err = r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		return nil, err
	}

	req := &entities.UpdateAnswerWithEvidenceRequest{
		CompanyUuid:        compUUID,
		UserUuid:           userUUID,
		QuestionnairesUuid: questionnaireUUID,
		AnswerUuid:         answerUUID,
		Answer:             answerRequestBody,
		Files:              r.MultipartForm.File["files"],
	}

	return req, nil
}

func decodeDownloadEvidenceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	questionnaireUUID, err := uuid.Parse(params["questionnaire_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("questionnaire_id")
	}

	answerUUID, err := uuid.Parse(params["answer_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("answer_id")
	}

	file_id := params["file_id"]
	if file_id == "" {
		return nil, httpError.NewErrBadOrInvalidPathParameter("file_id")
	}

	req := &entities.DownloadEvidenceRequest{
		CompanyUuid:        compUUID,
		UserUuid:           userUUID,
		QuestionnairesUuid: questionnaireUUID,
		AnswerUuid:         answerUUID,
		FileId:             file_id,
	}

	return req, nil
}

func decodeDeleteEvidenceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	questionnaireUUID, err := uuid.Parse(params["questionnaire_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("questionnaire_id")
	}

	answerUUID, err := uuid.Parse(params["answer_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("answer_id")
	}

	file_id := params["file_id"]
	if file_id == "" {
		return nil, httpError.NewErrBadOrInvalidPathParameter("file_id")
	}

	req := &entities.DeleteEvidenceRequest{
		CompanyUuid:        compUUID,
		UserUuid:           userUUID,
		QuestionnairesUuid: questionnaireUUID,
		AnswerUuid:         answerUUID,
		FileId:             file_id,
	}

	return req, nil
}
