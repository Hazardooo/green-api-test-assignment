package tests

import "green-api-test-assignment/internal/utils"

var (
	_, IdInstanceMock = utils.ReadEnv("IDINSTANCE")
	_, ApiTokenMock   = utils.ReadEnv("APITOKENINSTANCE")
)
