package model

import (
  "fmt"
  "io"
  "net/http"
  "regexp"
  "strconv"
  "strings"
)

type Client struct {
  Name        string
  BaseUrl     string
  BaseHeaders map[string]string
}

type ArgDateTime struct {
  Name    string
  Default *string
}

func (a ArgDateTime) Assert(value string) error {
  if value == "" {
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

type ArgEnum struct {
  Name    string
  Default *string
  Options []string
}

func (a ArgEnum) Assert(value string) error {
  for _, option := range a.Options {
    if value == option {
      return nil
    }
  }
  return fmt.Errorf("Invalid %s (type enum): Expected to be in {%s}, received %s", a.Name, strings.Join(a.Options, ","), value)
}

type ArgInt struct {
  Name    string
  Default *int
  Min     *int
  Max     *int
}

func (a ArgInt) Assert(value int) error {
  if a.Min != nil && value < *a.Min {
    return fmt.Errorf("Invalid %s (type int): Expected >= %d, received %d", a.Name, *a.Min, value)
  }
  if a.Max != nil && value > *a.Max {
    return fmt.Errorf("Invalid %s (type int): Expected <= %d, received %d", a.Name, *a.Max, value)
  }
  return nil
}

type ArgString struct {
  Name    string
  Default *string
  // Assumes strict match (i.e. Regexp = "abc" is matched for "^abc$"
  Regexp  *string
}

func (a ArgString) Assert(value string) error {
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

type FetchConfig struct {
  Endpoint string
  Fields   *map[string]string
  Headers  *map[string]string
  Params   *map[string]string
}

func (c *Client) Fetch(config FetchConfig) ([]byte, error) {
  url := fmt.Sprintf("%s%s", c.BaseUrl, config.Endpoint)

  if config.Params != nil {
    for key, value := range *config.Params {
      url = strings.ReplaceAll(url, fmt.Sprintf(":%s", key), value)
    }
  }

  if config.Fields != nil && len(*config.Fields) > 0 {
    pairs := []string{}
    for key, value := range *config.Fields {
      pairs = append(pairs, fmt.Sprintf("%s=%s", key, value))
    }
    url = fmt.Sprintf("%s?%s", url, strings.Join(pairs, "&"))
  }

  mergedHeaders := map[string][]string{}
  for key, value := range c.BaseHeaders {
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

  res, err := http.DefaultClient.Do(req)
  if err != nil {
    return nil, err
  }

  data, err := io.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    return nil, err
  }

  return data, nil
}

func IntPtr(value int) *int       { return &value }
func StrPtr(value string) *string { return &value }
