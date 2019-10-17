package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestLoadConfigEmptyConfig(t *testing.T) {
	t.Log(`Test loading empty config file`)
	tmpf, err := createTempFile()
	if err != nil {
		t.Fatalf(`Failed to create temporary config file: %s`, err.Error())
		return
	}
	tmpf.Close()
	defer os.Remove(tmpf.Name())

	var router = newRouter(tmpf.Name())

	if err = router.LoadConfig(); err != nil {
		fail(t, "Failed to load config file", err.Error(), nil)
	}
}

func TestLoadConfigEmptyLines(t *testing.T) {
	t.Log(`Test skipping empty lines in the config file`)
	tmpf, err := createTempFile()
	if err != nil {
		t.Fatalf(`Failed to create temporary config file: %s`, err.Error())
		return
	}

	var (
		router   = newRouter(tmpf.Name())
		testcase = "a=b\n\n\nb=c\n"
	)

	tmpf.WriteString(testcase)
	tmpf.Close()
	defer os.Remove(tmpf.Name())

	if err = router.LoadConfig(); err != nil {
		fail(t, "Failed to load config file", err, nil)
	}
}

func TestLoadConfigInvalidKVPair(t *testing.T) {
	t.Log(`Test detection of invalid key-value pairs`)
	tmpf, err := createTempFile()
	if err != nil {
		t.Fatalf(`Failed to create temporary config file: %s`, err.Error())
		return
	}

	var (
		router   = newRouter(tmpf.Name())
		testcase = "invalid.rule here"
		expected = fmt.Errorf(
			`Invalid rule: %s, reason: Invalid key-value pair`, testcase)
	)

	tmpf.WriteString(testcase)
	tmpf.Close()
	defer os.Remove(tmpf.Name())

	if err = router.LoadConfig(); err.Error() != expected.Error() {
		fail(t, "Failed to detect invalid key-value pairs", err, expected)
	}
}

func TestLoadConfigBrokenPath(t *testing.T) {
	t.Log(`Test detection of broken paths in the rules`)
	tmpf, err := createTempFile()
	if err != nil {
		t.Fatalf(`Failed to create temporary config file: %s`, err.Error())
		return
	}

	var (
		router   = newRouter(tmpf.Name())
		testcase = "a. b. c.d=server1"
		expected = fmt.Errorf(`Invalid rule: %s, reason: Broken path`, testcase)
	)

	tmpf.WriteString(testcase)
	tmpf.Close()
	defer os.Remove(tmpf.Name())

	if err = router.LoadConfig(); err.Error() != expected.Error() {
		fail(t, "Failed to detect broken path in the rule", err, expected)
	}
}

func TestLoadConfigTrailingConstraints(t *testing.T) {
	t.Log(`Test detection of trailing constraints in the rules`)
	tmpf, err := createTempFile()
	if err != nil {
		t.Fatalf(`Failed to create temporary config file: %s`, err.Error())
		return
	}

	var (
		router   = newRouter(tmpf.Name())
		testcase = "*.a.b.c=server1"
		expected = fmt.Errorf(
			`Invalid rule: %s, reason: Trailing route constraint after "*"`, testcase)
	)

	tmpf.WriteString(testcase)
	tmpf.Close()
	defer os.Remove(tmpf.Name())

	if err = router.LoadConfig(); err.Error() != expected.Error() {
		fail(t, "Failed to detect invalid rules", err, expected)
	}
}

func TestLoadConfigSample(t *testing.T) {
	t.Log(`Test the sample config file`)
	tmpf, err := createTempFile()
	if err != nil {
		t.Fatalf(`Failed to create temporary config file: %s`, err.Error())
		return
	}

	var (
		router   = newRouter(tmpf.Name())
		testcase = map[string]string{
			"customer1.us.ca.*":   "server1",
			"customer2.us.*.*":    "server3",
			"customer2.*.*.*":     "server4",
			"*.*.*.*":             "server5",
			"customer1.us.ca.sjc": "server2",
		}
	)

	if err := createTestConfig(tmpf, testcase); err != nil {
		t.Fatal(err.Error())
		return
	}
	tmpf.Close()
	defer os.Remove(tmpf.Name())

	if err = router.LoadConfig(); err != nil {
		fail(t, "Failed to load config file", err, nil)
	}
}

func TestFindRoute(t *testing.T) {
	t.Log(`Test findRoute() function with varying inputs`)
	rules := map[string]string{
		"customer1.us.ca.*":   "server1",
		"customer2.us.*.*":    "server3",
		"customer2.*.*.*":     "server4",
		"*.*.*.*":             "server5",
		"customer1.us.ca.sjc": "server2",
	}
	tmpf, err := createTempFile()
	if err != nil {
		t.Fatalf(`Failed to create temporary config file: %s`, err.Error())
		return
	}
	if err := createTestConfig(tmpf, rules); err != nil {
		t.Fatal(err.Error())
		return
	}
	tmpf.Close()
	defer os.Remove(tmpf.Name())

	var (
		router    = newRouter(tmpf.Name())
		testcases = map[string]string{
			"":                     "",
			"   ":                  "",
			"\t  ":                 "",
			"\n":                   "",
			"a.b.c":                "",
			"a.b.c.d.e":            "",
			"a.b.c.d":              "server5",
			"customer1.us.ca.sfo":  "server1",
			"customer1.us.ca.sjc":  "server2",
			"customer2.us.tx.dfw":  "server3",
			"customer2.cn.tw.tai":  "server4",
			"customer10.us.ny.nyc": "server5",
		}
	)

	router.LoadConfig()
	for input, expected := range testcases {
		t.Logf(`Testing "%s" ...`, input)
		got, err := router.FindRoute(input)
		if err != nil {
			t.Log(err.Error())
		}

		if got != expected {
			fail(t, "Route is incorrect", got, expected)
		}
	}

}

func TestFindRouteBrokenInput(t *testing.T) {
	t.Log(`Test broken query in the input`)
	rules := map[string]string{
		"customer1.us.ca.*":   "server1",
		"customer2.us.*.*":    "server3",
		"customer2.*.*.*":     "server4",
		"*.*.*.*":             "server5",
		"customer1.us.ca.sjc": "server2",
	}
	tmpf, err := createTempFile()
	if err != nil {
		t.Fatalf(`Failed to create temporary config file: %s`, err.Error())
		return
	}
	if err := createTestConfig(tmpf, rules); err != nil {
		t.Fatal(err.Error())
		return
	}
	tmpf.Close()
	defer os.Remove(tmpf.Name())

	var (
		router    = newRouter(tmpf.Name())
		testcases = []string{
			"customer1    .us .\tca.sjc",
		}
	)

	router.LoadConfig()
	for _, input := range testcases {
		t.Logf(`Testing "%s" ...`, input)
		expected := fmt.Errorf(`Invalid input: %s, reason: Broken input`, input)
		_, err := router.FindRoute(input)
		if err.Error() != expected.Error() {
			fail(t, "Failed to detect ambiguous input", err, expected)
		}
	}
}

func TestFindRouteAmbiguousInput(t *testing.T) {
	t.Log(`Test ambiguous input`)
	rules := map[string]string{
		"customer1.us.ca.*":   "server1",
		"customer2.us.*.*":    "server3",
		"customer2.*.*.*":     "server4",
		"*.*.*.*":             "server5",
		"customer1.us.ca.sjc": "server2",
	}
	tmpf, err := createTempFile()
	if err != nil {
		t.Fatalf(`Failed to create temporary config file: %s`, err.Error())
		return
	}
	if err := createTestConfig(tmpf, rules); err != nil {
		t.Fatal(err.Error())
		return
	}
	tmpf.Close()
	defer os.Remove(tmpf.Name())

	var (
		router    = newRouter(tmpf.Name())
		testcases = []string{
			"customer1.us.*.sjc",
			"*.*.*.*",
		}
	)

	router.LoadConfig()
	for _, input := range testcases {
		t.Logf(`Testing "%s" ...`, input)
		expected := fmt.Errorf(
			`Invalid input: %s, reason: Ambiguous input with wildcard "*"`, input)
		_, err := router.FindRoute(input)
		if err.Error() != expected.Error() {
			fail(t, "Failed to detect ambiguous input", err, expected)
		}
	}
}

// utility functions
func createTempFile() (*os.File, error) {
	return ioutil.TempFile("", "router-test")
}

func createTestConfig(f *os.File, rules map[string]string) error {
	var (
		ruleArr []string
		err     error
	)

	for k, v := range rules {
		ruleArr = append(ruleArr, fmt.Sprintf(`%s=%s`, k, v))
	}

	if _, err = f.WriteString(strings.Join(ruleArr, "\n")); err != nil {
		return fmt.Errorf(`Failed to write config file: %s`, err.Error())
	}

	return nil
}

func fail(t *testing.T, msg string, got, expected interface{}) {
	t.Errorf(`%s, got: %v, expected: %v`, msg, got, expected)
}
