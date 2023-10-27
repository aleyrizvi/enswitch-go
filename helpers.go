package enswitch

import "net/url"

func PointerTo[T any](v T) *T {
	return &v
}

func urlValuesEqual(v1, v2 url.Values) bool {
	if len(v1) != len(v2) {
		return false
	}

	for key, values := range v1 {
		if len(values) != len(v2[key]) {
			return false
		}

		for i, val := range values {
			if val != v2[key][i] {
				return false
			}
		}
	}

	return true
}
