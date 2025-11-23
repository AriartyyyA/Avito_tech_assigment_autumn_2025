package transport

import (
	"fmt"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
)

func InvalidRequest(item string) models.ErrorResponse {

	errStr := fmt.Sprintf("%s is required", item)

	if item == "" {

		errStr = "invalid request"

	}

	error := models.NewErrorDetail(

		models.ErrorCodeInvalidRequest,

		errStr,
	)

	respErr := models.ErrorResponse{

		Error: error,
	}

	return respErr

}

func TeamExists() models.ErrorResponse {

	error := models.NewErrorDetail(

		models.ErrorCodeTeamExists,

		"team_name already exists",
	)

	respErr := models.ErrorResponse{

		Error: error,
	}

	return respErr

}

func NotFound(code models.ErrorCode) models.ErrorResponse {

	error := models.NewErrorDetail(

		code,

		"resource not found",
	)

	respErr := models.ErrorResponse{

		Error: error,
	}

	return respErr

}

func PRExists() models.ErrorResponse {

	error := models.NewErrorDetail(

		models.ErrorCodePRExists,

		"PR id already exists",
	)

	respErr := models.ErrorResponse{

		Error: error,
	}

	return respErr

}

func PRMerged() models.ErrorResponse {

	error := models.NewErrorDetail(

		models.ErrorCodePRMerged,

		"cannot reassign on merged PR",
	)

	respErr := models.ErrorResponse{

		Error: error,
	}

	return respErr

}

func NotAssigned() models.ErrorResponse {
	error := models.NewErrorDetail(
		models.ErrorCodeNotAssigned,
		"reviewer is not assigned to this PR",
	)

	respErr := models.ErrorResponse{
		Error: error,
	}

	return respErr

}

func NoCandidate() models.ErrorResponse {
	error := models.NewErrorDetail(
		models.ErrorCodeNoCandidate,
		"no active replacement candidate in team",
	)

	respErr := models.ErrorResponse{
		Error: error,
	}

	return respErr

}

func InternalError() models.ErrorResponse {
	error := models.NewErrorDetail(
		models.ErrorCodeInternal,
		"internal error",
	)

	respErr := models.ErrorResponse{
		Error: error,
	}

	return respErr
}
