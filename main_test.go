package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCurrentDate_FormatCorrect(t *testing.T) {
	returnedDate := currentDate();
	_, err := time.Parse("2006-01-02", returnedDate);
	assert.Nil(t, err);
}