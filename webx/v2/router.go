package webx

import "strings"

// 代表路由
type router struct {
	// trees 代表的是森林，HTTP method => 树的根节点 rootNode
	// 如，GET => home
	trees map[string]*node
}


func newRouter() router {
	return router{
		trees: map[string]*node {},
	}
}

// addRoute 注册路由
func (r *router) addRoute(method string, path string, handleFunc HandleFunc) {

	// 校验（筛掉不合法路由，过于繁琐，删了）

	root, ok := r.trees[method]
	// 这是一个全新的 HTTP 方法，创建根节点

	if !ok {
		root = &node{segment: "/"}
		r.trees[method] = root 
	}

	if path == "/" {
		// 特例 e.g. localhost:8080/
		root.handler = handleFunc
		return 
	}

	cur := root
	segs := strings.Split(strings.Trim(path, "/"), "/")

	// 遍历 segment
	for _, seg := range segs {
		// 有对应 Node，开路先锋，那就走到底；没有，那就补上
		cur = cur.childOrCreate(seg)
	}

	// 走到底，最终节点（绑定业务逻辑）
	cur.handler = handleFunc

}



// findRoute 查找对应的节点
// 注意，返回的 node 内部 HandleFunc 不为 nil 才算是注册了路由
func (r *router) findRoute(method string, path string) (*matchInfo, bool) {
	
	// 先 HTTP method 匹配，找出根节点 root Node
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}
	if path == "/" {
		return &matchInfo{n: root}, true
	}

	cur := root
	urlPath := strings.Trim(path, "/")
	segs := strings.Split(urlPath, "/")

	// 遍历到底，就找到了
	for _, seg := range segs {
		if cur.children == nil {
			if cur.paramChild != nil {
				return &matchInfo{
							n:             cur.paramChild, 
							segmentParams: map[string]string{cur.paramChild.segment[1:]: seg},
					}, true
			}

			return &matchInfo{n: cur.starChild}, cur.starChild != nil
		}

		child, ok := cur.children[seg]
		if !ok {
			if cur.paramChild != nil {
				return &matchInfo{
							n:             cur.paramChild, 
							segmentParams: map[string]string{cur.paramChild.segment[1:]: seg},
					}, true
			}

			return &matchInfo{n: cur.starChild}, cur.starChild != nil
		}

		cur = child
	}

	return &matchInfo{n: cur}, true
}



// node 代表路由树的节点
// 路由树的匹配顺序是：
// 1. 静态完全匹配
// 2. 动态路由匹配
// 2.1 路径参数匹配：形式 :param_name
// 2.2 通配符匹配：*
type node struct {

	// 表示路由中某个段的字符串；例如，/a/b/c 中的 b 这一段
	segment  string      
	
	// 表示业务逻辑
	handler  HandleFunc  // 代表这个节点中包含的业务逻辑，用于最终加载调用
	
	// 静态路由匹配
	children map[string]*node  // 代表这个节点下的子节点；例如，user 节点下，应该有 register、login  
	
	// 动态路由匹配 - 参数路径
	paramChild *node

	// 动态路由匹配 - 通配符
	starChild *node 

}



// childOrCreate 查找子节点
// 如果子节点不存在就创建一个，并且将子节点放回到 children 中
func (n *node) childOrCreate(segment string) *node {
	
	// 支持动态路由匹配 - 通配符
	if segment == "*" {
		if n.starChild == nil {
			n.starChild = &node{segment: segment}
		}
		return n.starChild
	}

	// 支持动态路由匹配 - 路径参数
	if segment[0] == ':'{
		if n.paramChild == nil {
			n.paramChild = &node{segment: segment}		
		}
		return n.paramChild
	}

	// map 赋值前要初始化
	if n.children == nil {
		n.children = make(map[string]*node)
	}
	
	child, ok := n.children[segment]
	if !ok {
		child = &node{segment: segment}
		n.children[segment] = child
	}

	return child
}



// 支持动态路由匹配 - 参数路径提取参数
type matchInfo struct {
	n             *node
	segmentParams map[string]string
}

func (m *matchInfo) addValue(key, value string) {
	if m.segmentParams == nil {
		// 大多数情况下，参数路径只有一段
		m.segmentParams = map[string]string{key: value}
	}

	m.segmentParams[key] = value
}