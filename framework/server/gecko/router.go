package gecko

import (
	"fmt"
)

// Router is a prefix tree.
type Router struct {
	root *node
}

// HTTPMethod represents method of http request.
type HTTPMethod int

type node struct {
	pattern      string
	part         string
	isWild       bool
	wildChildren *node
	children     map[string]*node
	handler      []HandlerFunc
	groups       []*GroupRouter
}

// HTTPMethod types.
const (
	Get = iota
	Post
	Put
	Delete
	Patch
	Option
	Head
)

// NewRouter returns a new router.
func NewRouter() *Router {
	return &Router{}
}

// SearchHandler searches through the router tree, and find handler function if possible.
func (g *Gecko) SearchHandler(method HTTPMethod, pattern string) (HandlerFunc, ParamMap, []*GroupRouter) {
	if g.router.root == nil {
		return nil, nil, nil
	}
	return g.router.root.search(method, pattern)
}

func (r *Router) insert(method HTTPMethod, pat string, handler HandlerFunc, gr *GroupRouter) {
	if r.root == nil {
		// lazy init
		r.root = newNode("/", "/", 0, nil, false)
	}
	pat = gr.prefix + pat // add GroupRouter's prefix
	r.root.insert(method, pat, pat, handler, gr)
}

func newNode(pattern string, part string, method HTTPMethod, handler HandlerFunc, isWild bool) *node {
	ret := &node{part: part,
		pattern:  pattern,
		children: make(map[string]*node),
		handler:  make([]HandlerFunc, 8),
		isWild:   isWild,
		groups:   make([]*GroupRouter, 0),
	}
	ret.handler[method] = handler
	return ret
}

func (r *node) insert(method HTTPMethod, pat string, cpat string, handler HandlerFunc, gr *GroupRouter) {
	cpatlen := len(cpat)
	if cpatlen == 0 {
		// Already finished. Set method.
		r.handler[method] = handler
		r.pattern = pat
		r.groups = append(r.groups, gr)
		return
	}
	old := r
	var split int
	wildcard := false
	if cpat[1] == ':' {
		wildcard = true
	} else if cpat[1] == '*' {
		// All match. Can only occur at the very end of a route.
		if old.wildChildren != nil {
			panic(fmt.Sprintf("Ambiguos route. Cannot have double parameter wildcard under one parent. Err route is %s", pat))
		}
		old.wildChildren = newNode(pat, cpat, method, handler, true)
	}
	for split = 1; split < cpatlen && cpat[split] != '/'; split++ {
	}
	part := cpat[:split]
	next := r.children[part]
	nextCpat := cpat[split:]
	if next == nil {
		// Not exist. Create a new node.
		n := newNode("", part, method, nil, wildcard)
		if wildcard {
			// Can only have one wild children.
			if old.wildChildren == nil {
				old.wildChildren = n
			} else {
				panic(fmt.Sprintf("Ambiguos route. Cannot have double parameter wildcard under one parent. Err route is %s", pat))
			}
		} else {
			old.children[part] = n
		}
		n.insert(method, pat, nextCpat, handler, gr)
		return
	}
	// Already Exist.
	next.insert(method, pat, nextCpat, handler, gr)
	return
}

func (r *node) search(method HTTPMethod, pattern string) (HandlerFunc, ParamMap, []*GroupRouter) {
	var split int = 0
	var record int = 0
	prev := r
	wildParams := make(ParamMap)
	for {
		split++ // Skip the next '/' delimiter.
		for ; split < len(pattern) && pattern[split] != '/'; split++ {
		}
		cpart := pattern[record:split]
		// Priority 1: exact match.
		if prev.children[cpart] == nil {
			// Priority 2: go wild.
			if prev.wildChildren != nil {
				// Record wild params.
				switch prev.wildChildren.part[1] {
				case ':':
					// If :, go wild.
					wildParams[prev.wildChildren.part[2:]] = cpart[1:] // get rid of '/', ':', '*' symbol
					prev = prev.wildChildren
				case '*':
					// If *, stop matching and return result immediately.
					wildParams[prev.wildChildren.part[2:]] = pattern[1:] // Not cpart. Need to include all the rest of the string.
					return prev.wildChildren.handler[method], wildParams, prev.wildChildren.groups
				default:
					panic("Unexpected route search result.")
				}

			} else {
				return nil, nil, nil // Not match.
			}
		} else {
			prev = prev.children[cpart]
		}
		if split >= len(pattern) {
			// Search end.
			return prev.handler[method], wildParams, prev.groups
		}
		record = split
	}
}
