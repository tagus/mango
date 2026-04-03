package mango

import (
	"fmt"
	"strings"
)

func ParseOrderBy(orderBy string, allowedFields []string) ([]string, error) {
	if strings.TrimSpace(orderBy) == "" {
		return nil, nil
	}

	allowed := NewSet(allowedFields...)
	clauses := make([]string, 0)
	for _, raw := range strings.Split(orderBy, ",") {
		raw = strings.TrimSpace(raw)
		if raw == "" {
			continue
		}

		parts := strings.Fields(raw)
		if len(parts) > 2 {
			return nil, fmt.Errorf("invalid order by clause %q", raw)
		}

		field := parts[0]
		if !allowed.Has(field) {
			return nil, fmt.Errorf("invalid order by field %q", field)
		}

		direction := "desc"
		if len(parts) == 2 {
			switch strings.ToLower(parts[1]) {
			case "asc", "desc":
				direction = strings.ToLower(parts[1])
			default:
				return nil, fmt.Errorf("invalid order by direction %q", parts[1])
			}
		}

		clauses = append(clauses, fmt.Sprintf("%s %s", field, direction))
	}

	if len(clauses) == 0 {
		return nil, nil
	}

	return clauses, nil
}
