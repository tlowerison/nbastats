package model

import (
  "fmt"
  "io"
  "net/http"
  "regexp"
  "strconv"
  "strings"
)

type DataSource struct {
  BaseUrl     string
  BaseHeaders map[string]string
}

type Result struct {
  Headers    []string        `json:"headers"`
  Rows       [][]interface{} `json:"rowSet"`
  HideHeader bool
}

func (r *Result) String() string {
  if r == nil {
    return ""
  }
  output := ""
  separator := "\t"
  if !r.HideHeader {
    output = strings.Join(r.Headers, separator) + "\n"
  }
  strRows := make([]string, len(r.Rows))
  for i, row := range r.Rows {
    strRow := make([]string, len(row))
    for j, entry := range row {
      strRow[j] = fmt.Sprintf("%v", entry)
    }
    strRows[i] = strings.Join(strRow, separator)
  }
  output += strings.Join(strRows, "\n")
  if len(output) > 0 && string(output[len(output) - 1]) != "\n" {
    output += "\n"
  }
  return output
}

type FetchFunc func(fields map[string]interface{}) (*Result, error)

type Client struct {
  DataSources map[string]DataSource
  FetchFuncs  map[string]FetchFunc
  Fetch       func(endpoint string, fields map[string]interface{}) (*Result, error)
  ShouldLog   bool
}

type ArgDateTime struct {
  Name    string
  Default string
}

func (a ArgDateTime) Assert(value string, required bool) error {
  if value == "" && !required {
    return nil
  }
  nums := strings.Split(value, "-")
  retErr := fmt.Errorf("Invalid %s (type date): Expected to be in format YYYY-MM-dd, received %s", a.Name, value)
  if len(nums) != 3 {
    return retErr
  }
  var err error
  _, err = strconv.Atoi(nums[0]); if err != nil { return retErr }
  _, err = strconv.Atoi(nums[1]); if err != nil { return retErr }
  _, err = strconv.Atoi(nums[2]); if err != nil { return retErr }
  return nil
}

func (a ArgDateTime) FromVal(value string) string {
  return value
}

func (a ArgDateTime) FromPtr(ptr *string) string {
  if ptr == nil {
    return ""
  }
  return *ptr
}

type ArgEnum struct {
  Name    string
  Default string
  Options []string
}

func (a ArgEnum) Assert(value string, required bool) error {
  if value == "" && !required {
    return nil
  }
  for _, option := range a.Options {
    if value == option {
      return nil
    }
  }
  return fmt.Errorf("Invalid %s (type enum): Expected to be in {%s}, received %s", a.Name, strings.Join(a.Options, ","), value)
}

func (a ArgEnum) FromVal(value string) string {
  return value
}

func (a ArgEnum) FromValAsInt(value string) string {
  if value == "" {
    return "0"
  }
  return a.FromVal(value)
}

func (a ArgEnum) FromPtr(ptr *string) string {
  if ptr == nil {
    return ""
  }
  return *ptr
}

func (a ArgEnum) FromPtrAsInt(ptr *string) string {
  if ptr == nil {
    return "0"
  }
  return a.FromValAsInt(*ptr)
}

type ArgInt struct {
  Name    string
  Default int
  Min     *int
  Max     *int
}

func (a ArgInt) Assert(value int, required bool) error {
  if value == 0 && !required {
    return nil
  }
  if a.Min != nil && value < *a.Min {
    return fmt.Errorf("Invalid %s (type int): Expected >= %d, received %d", a.Name, *a.Min, value)
  }
  if a.Max != nil && value > *a.Max {
    return fmt.Errorf("Invalid %s (type int): Expected <= %d, received %d", a.Name, *a.Max, value)
  }
  return nil
}

func (a ArgInt) FromVal(value int) string {
  return strconv.Itoa(value)
}

func (a ArgInt) FromValAsStr(value int) string {
  if value == 0 {
    return ""
  }
  return a.FromVal(value)
}

func (a ArgInt) FromPtr(ptr *int) string {
  if ptr == nil {
    return "0"
  }
  return strconv.Itoa(*ptr)
}

func (a ArgInt) FromPtrAsStr(ptr *int) string {
  if ptr == nil {
    return ""
  }
  return a.FromValAsStr(*ptr)
}

type ArgString struct {
  Name    string
  Default string
  // Assumes strict match (i.e. Regexp = "abc" is matched for "^abc$"
  Regexp  *string
}

func (a ArgString) Assert(value string, required bool) error {
  if value == "" && !required {
    return nil
  }
  if a.Regexp == nil {
    return nil
  }
  pattern := fmt.Sprintf("^%s$", *a.Regexp)
  _, err := regexp.Match(pattern, []byte(value))
  if err != nil {
    return fmt.Errorf("Invalid %s (type string): Expected to match pattern %s, received %s", a.Name, pattern, value)
  }
  return nil
}

func (a ArgString) FromVal(value string) string {
  return value
}

func (a ArgString) FromValAsInt(value string) string {
  if value == "" {
    return "0"
  }
  return a.FromVal(value)
}

func (a ArgString) FromPtr(ptr *string) string {
  if ptr == nil {
    return ""
  }
  return *ptr
}

func (a ArgString) FromPtrAsInt(ptr *string) string {
  if ptr == nil {
    return "0"
  }
  return a.FromValAsInt(*ptr)
}

type FetchConfig struct {
  DataSource string
  Endpoint   string
  Fields     *map[string]string
  Headers    *map[string]string
  Params     *map[string]string
}

func removeSpaces(value string, newSpace string) string {
  return strings.ReplaceAll(value, " ", newSpace)
}

func (c *Client) Get(config FetchConfig) ([]byte, error) {
  if _, ok := c.DataSources[config.DataSource]; !ok {
    return nil, fmt.Errorf("Invalid datasource, received %s", config.DataSource)
  }
  dataSource := c.DataSources[config.DataSource]

  url := fmt.Sprintf("%s%s", dataSource.BaseUrl, config.Endpoint)

  if config.Params != nil {
    for key, value := range *config.Params {
      // Not sure if we want to replace space in params with +
      url = strings.ReplaceAll(url, fmt.Sprintf(":%s", key), removeSpaces(value, "+"))
    }
  }

  if config.Fields != nil && len(*config.Fields) > 0 {
    pairs := []string{}
    for key, value := range *config.Fields {
      pairs = append(pairs, fmt.Sprintf("%s=%s", key, removeSpaces(value, "+")))
    }
    url = fmt.Sprintf("%s?%s", url, strings.Join(pairs, "&"))
  }

  mergedHeaders := map[string][]string{}
  for key, value := range dataSource.BaseHeaders {
    mergedHeaders[key] = []string{value}
  }
  if config.Headers != nil {
    for key, value := range *config.Headers {
      mergedHeaders[key] = []string{value}
    }
  }

  req, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
    return nil, err
  }
  req.Header = mergedHeaders

  if c.ShouldLog && req.URL != nil {
    fmt.Printf("GET %s\n", req.URL.String())
  }

  res, err := http.DefaultClient.Do(req)
  if err != nil {
    return nil, err
  }

  data, err := io.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    return nil, err
  }

  if c.ShouldLog && data != nil {
    fmt.Println(string(data))
  }

  return data, nil
}

func IntPtr(value int) *int       { return &value }
func StrPtr(value string) *string { return &value }
