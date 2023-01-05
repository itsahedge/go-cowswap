package subgraph

import (
	"fmt"
	"strings"
)

func buildVariables(vars map[string]interface{}) string {
	var args []string
	for key, value := range vars {
		switch key {
		case "block":
			blockVars := value.(map[string]interface{})
			var blockArgs []string
			for blockKey, blockValue := range blockVars {
				blockArgs = append(blockArgs, fmt.Sprintf("%s: %v", blockKey, blockValue))
			}
			args = append(args, fmt.Sprintf("block: {%s}", strings.Join(blockArgs, ", ")))
		case "orderBy", "orderDirection":
			args = append(args, fmt.Sprintf("%s: \"%s\"", key, value.(string)))
		default:
			args = append(args, fmt.Sprintf("%s: %v", key, value))
		}
	}
	return strings.Join(args, ", ")
}
