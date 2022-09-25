package db_utils

import "BeQuestPrep/internal/rest_utils"

type DeleteAPI interface {
	DeleteAnswer(answer rest_utils.AnswerData) error
}

func DeleteAnswer(api DeleteAPI, answer rest_utils.AnswerData) error {
	return api.DeleteAnswer(answer)
}