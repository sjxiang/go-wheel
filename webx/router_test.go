package webx

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)


// 批注：这段设计错误，注册路由和查找路由，比较冲突
func Test_router_addRoute(t *testing.T) {
	
	tests := []struct {
		method     string
		path       string
	}{
		{
			method: http.MethodGet,
			path: "/",
		},
		{
			method: http.MethodPost,
			path: "/",
		},
		{
			method: http.MethodGet,
			path: "/user",
		},
		{
			method: http.MethodGet,
			path: "/user/1/profile",
		},
	}

	var handleFunc HandleFunc = nil

	wantRouter := &router {
		trees: map[string]*node {
			http.MethodPost: &node {
				path: "/",
				handler: handleFunc,
			},
			http.MethodGet: &node {
				path: "/",
				handler: handleFunc,
				children: map[string]*node {
					"user": &node {
						path: "user",
						handler: handleFunc,
						children: map[string]*node {
							"1": &node {
								path: "1",
								children: map[string]*node {
									"profile": &node {
										path: "profile",
										handler: handleFunc,
									},
								},
							},
						},
					},
				},
			},
			
		},
	}

	res := &router{
		trees: map[string]*node {},
	}

	for _, tc := range tests {
		res.addRoute(tc.method, tc.path, handleFunc)
	}

	assert.Equal(t, wantRouter, res)


	findCases := []struct {
		name     string
		method   string
		path     string

		found    bool
		wantPath string
	}{
		{
			name:     "/",
			method:   http.MethodGet,
			path:     "/",
			found:    true,
			wantPath: "/",
		},
		{
			name:     "/user/1/profile",
			method:   http.MethodGet,
			path:     "/user/1/profile",
			found:    true,
			wantPath: "profile",
		},
	}

	for _, tc := range findCases {
		t.Run(tc.name, func(t *testing.T) {
			n, ok := res.findRoute(tc.method, tc.path)
			assert.Equal(t, tc.found, ok)
			if !ok {
				return
			}	
			assert.Equal(t, tc.wantPath, n.path)
		})
	}

} 