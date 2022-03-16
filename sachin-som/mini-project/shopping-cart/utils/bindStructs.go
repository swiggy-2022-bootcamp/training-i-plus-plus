package utils

import "github.com/gin-gonic/gin"

func BindStructs(instance *interface{}, gctx *gin.Context) error {
	if err := gctx.ShouldBindJSON(instance); err != nil {
		return err
	}
	return nil
}
