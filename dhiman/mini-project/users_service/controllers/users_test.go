package controllers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{"Try creating", args{nil}},
		{},
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateUser(tt.args.c)
		})
	}
}
