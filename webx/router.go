package webx

import (
	"strings"
)

// 代表路由
type router struct {
	// trees 代表的是森林，HTTP method => 树的根节点
	trees  map[string]*node
}

func newRouter() router {
	return router{
		trees: map[string]*node {},
	}
}

func (r *router) addRoute(method string, path string, handleFunc HandleFunc) {
	root, ok := r.trees[method]
	// 这是一个全新的 HTTP 方法，创建根节点
	if !ok {
		// 创建根节点
		root = &node{path: "/"}
		r.trees[method] = root
	}
	if path == "/" {
		root.handler = handleFunc
		return
	}

	cur := root
	// 切割 url 路径
	segs := strings.Split(path[1:], "/")

	// 遍历每一段 path，填充节点
	for _, seg := range segs {
		cur = cur.childOrCreate(seg)
	}

	cur.handler = handleFunc
}


// findRouter
func (r *router) findRoute(method string, path string) (*node, bool) {
	// 1. 先根据 method 找根节点
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}
	if path == "/" {
		// e.g. 如果查找 "/" 正好
		return root, true
	}

	// 2. 遍历节点
	cur := root
	segs := strings.Split(path[1:], "/")
	for _, seg := range segs {
		// 路由树没有子节点，那毙掉，不用接着找了 
		if cur.children == nil {
			// 考虑参数路径
			if cur.paramChild != nil {
				return cur.paramChild, true
			}

			// 考虑通配符
			return cur.starChild, cur.starChild != nil 
			// return nil, false
		}
		// 如果有这个路径，那就接着迭代
		child, ok := cur.children[seg]
		if !ok {
			if cur.paramChild != nil {
				return cur.paramChild, true
			}
			
			return cur.starChild, cur.starChild != nil  
			// return nil, false
		}	
		cur = child
	}

	

	return cur, true
}


// 命中节点，执行业务逻辑
type node struct {
	// /a/b/c 中 b 这一段
	path    string

	handler HandleFunc

	// path => 到子节点的映射
	// 静态路由
	children map[string]*node

	// 通配符
	starChild *node 

	// 参数路径
	paramChild *node

}


// childOrCreate 查找子节点，如果子节点不存在就创建一个
// 并且将子节点放回到 children 中
func (n *node) childOrCreate(path string) *node {

	// 支持通配符
	if path == "*" {
		if n.starChild == nil {
			n.starChild = &node {
				path: path,
			}
		}
		return n.starChild
	}

	// 支持参数路径
	if path[0] == ':'{
		if n.paramChild == nil {
			n.paramChild = &node{
				path: path,
			}		
		}
		return n.paramChild
	}

	if n.children == nil {
		n.children = make(map[string]*node)
	}

	child, ok := n.children[path]
	if !ok {
		child = &node{path: path}
		n.children[path] = child
	}

	return child
}
