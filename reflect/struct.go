package reflect

import (
	
	"github.com/duke-git/lancet/v2/structs"
)


func ToMap(v any) (map[string]any, error) {
	return structs.ToMap(v)
}