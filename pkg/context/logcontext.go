package context

import (
	"context"

	"github.com/cuida-me/mvp-backend/pkg/maps"
)

type LogCtxKey struct{}

type CtxBody = map[string]interface{}

func CtxWithValues(ctx context.Context, values CtxBody) context.Context {
	m, _ := ctx.Value(LogCtxKey{}).(CtxBody)
	return context.WithValue(ctx, LogCtxKey{}, mergeMaps(m, values))
}

// GetCtxValues extracts the CtxBody currently stored on the input ctx.
func GetCtxValues(ctx context.Context) CtxBody {
	m, _ := ctx.Value(LogCtxKey{}).(CtxBody)
	if m == nil {
		m = CtxBody{}
	}
	m["request_id"] = GetRequestIDFromContext(ctx)
	return m
}

func mergeMaps(bodies ...CtxBody) CtxBody {
	body := CtxBody{}
	maps.Merge(&body, bodies...)
	return body
}

func GetFields(ctx context.Context, key string, index int) string {
	keys := GetCtxValues(ctx)[key].([]string)
	return keys[index]
}

func GetField(ctx context.Context, key string) string {
	return GetCtxValues(ctx)[key].(string)
}
