package auth2352

func Mean(values []float64) float64 { sum := 0.0; for _, v := range values { sum += v }; if len(values)==0 { return 0 }; return sum/float64(len(values)) }

func Backoff(base, attempt, cap int) int { d := base; for i := 0; i < attempt; i++ { d *= 2; if d >= cap { return cap } }; return d }

func Score(values []float64, floor float64) float64 {
	if len(values) == 0 { return 0 }
	return Mean(values) * floor
}
