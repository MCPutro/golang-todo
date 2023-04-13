package model

import (
	"encoding/json"
	"github.com/MCPutro/golang-todo/helper"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_model_activity(t *testing.T) {

	now := "2023-04-12 12:39:36"

	timeNow, _ := time.Parse(helper.FORMAT_DATE, now)

	activity := Activity{
		Activity_id: 1,
		Title:       "makan",
		Email:       "makan@makan.com",
		Created_at:  timeNow,
		Updated_at:  timeNow,
	}

	expectJson := `{"id":1,"title":"makan","email":"makan@makan.com","createdAt":"2023-04-12T12:39:36Z","updatedAt":"2023-04-12T12:39:36Z"}`

	result, _ := json.Marshal(activity)

	assert.Equal(t, expectJson, string(result))

}

func Test_model_todo(t *testing.T) {

	now := "2023-04-12 12:39:36"

	timeNow, _ := time.Parse(helper.FORMAT_DATE, now)

	todo := Todo{
		Todo_id:           1,
		Activity_group_id: 4,
		Title:             "beli sayur",
		Is_active:         true,
		Priority:          "low",
		Created_at:        timeNow,
		Updated_at:        timeNow,
	}

	expectJson := `{"id":1,"activity_group_id":4,"title":"beli sayur","is_active":true,"priority":"low","createdAt":"2023-04-12T12:39:36Z","updatedAt":"2023-04-12T12:39:36Z"}`

	result, _ := json.Marshal(todo)

	assert.Equal(t, expectJson, string(result))

}
