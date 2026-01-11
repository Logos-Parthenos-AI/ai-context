package config

import (
"os"
"strings"
"sync"
)

var (
defaultModel string
once sync.Once
)

// DefaultModel returns the default AI model configured in config.yaml.
// It reads config.yaml from the repository root and looks for a "default_model:"
// entry. If not found or on error, it returns an empty string.
func DefaultModel() string {
once.Do(func() {
data, err := os.ReadFile("config.yaml")
if err != nil {
defaultModel = ""
return
}
for _, line := range strings.Split(string(data), "\n") {
line = strings.TrimSpace(line)
if strings.HasPrefix(line, "default_model:") {
parts := strings.SplitN(line, ":", 2)
if len(parts) == 2 {
val := strings.TrimSpace(parts[1])
val = strings.Trim(val, `"'`)
defaultModel = val
return
}
}
}
// no explicit value found
defaultModel = ""
})
return defaultModel
}
