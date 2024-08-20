package recommendation

import "context"

type GetRecommendation func(ctx context.Context) (string, error)
