// Code generated by hertz generator.

package order

import (
	"github.com/chhz0/go-mall-kitex/app/frontend/middleware"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.Auth()}
}

func _orderlistMw() []app.HandlerFunc {
	// your code...
	return nil
}