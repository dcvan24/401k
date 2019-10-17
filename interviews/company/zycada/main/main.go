package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Router -- an routing application
type Router interface {
	// LoadConfig() assumes every route is destined to a *SINGLE* server. If
	// duplicates exist, last route will always win and overwrite the conflict ones,
	// if any, appeared in the front of the config files.
	LoadConfig() error
	// FindRoute() finds the server that matches the input string. If more than one
	// rule matches, the most specific one wins.
	FindRoute(string) (string, error)
}

type routerImpl struct {
	routes     map[string]interface{}
	configPath string
}

func newRouter(path string) Router {
	return &routerImpl{
		routes:     make(map[string]interface{}),
		configPath: path,
	}
}

func (r *routerImpl) LoadConfig() error {
	f, err := os.Open(r.configPath)
	if err != nil {
		return err
	}
	defer f.Close()

	var (
		scanner = bufio.NewScanner(f)
		l       string
	)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if l = strings.TrimSpace(scanner.Text()); len(l) == 0 { // skip empty lines
			continue
		}

		if err := r.validateRule(l); err != nil {
			return fmt.Errorf(`Invalid rule: %s, reason: %s`, l, err.Error())
		}

		var (
			kv   = strings.Split(l, "=")
			hops = strings.Split(strings.TrimSpace(kv[0]), ".")
			prev = r.routes
		)

		for i, hop := range hops {
			if _, ok := prev[hop]; !ok {
				prev[hop] = make(map[string]interface{})
			}
			if i == len(hops)-1 {
				prev[hop] = strings.TrimSpace(kv[1])
			} else {
				prev = prev[hop].(map[string]interface{})
			}
		}
	}

	return nil
}

func (r *routerImpl) FindRoute(input string) (server string, err error) {
	if input = strings.TrimSpace(input); len(input) == 0 {
		return server, nil
	}
	if err := r.validateInput(input); err != nil {
		return server,
			fmt.Errorf(`Invalid input: %s, reason: %s`, input, err.Error())
	}
	var (
		hops = strings.Split(input, ".")
		prev = r.routes
		ok   bool
	)
	for i, hop := range hops {
		var cur interface{}

		if cur, ok = prev[hop]; !ok {
			if cur, ok = prev["*"]; !ok {
				break
			}
		}
		if i < len(hops)-1 {
			if prev, ok = cur.(map[string]interface{}); !ok {
				break
			}
		} else {
			server, ok = cur.(string)
		}
	}
	return server, nil
}

// validateRule() validates non-empty rules in the config
func (r *routerImpl) validateRule(rule string) error {
	// validate key-value pairs
	var kv []string
	if kv = strings.Split(rule, "="); len(kv) != 2 {
		return fmt.Errorf(`Invalid key-value pair`)
	}

	// validate if there are spaces in the path
	if len(strings.Fields(kv[0])) > 1 {
		return fmt.Errorf(`Broken path`)
	}

	// validate if there are trailing route constraint after "*"
	var (
		hops     = strings.Split(kv[0], ".")
		trailing bool
	)

	for i := len(hops) - 1; i > -1; i-- {
		if hops[i] != "*" {
			trailing = true
		} else if hops[i] == "*" && trailing {
			return fmt.Errorf(`Trailing route constraint after "*"`)
		}
	}

	return nil
}

// validateInput() validates non-empty input for locating servers
func (r *routerImpl) validateInput(input string) error {
	// validate if there are spaces in the input
	if len(strings.Fields(input)) > 1 {
		return fmt.Errorf("Broken input")
	}
	// validate if the input contains the wildcard "*" at any hop, which may lead
	// to multiple servers, making the input ambiguous
	for _, hop := range strings.Split(input, ".") {
		if hop == "*" {
			return fmt.Errorf(`Ambiguous input with wildcard "*"`)
		}
	}
	return nil
}
