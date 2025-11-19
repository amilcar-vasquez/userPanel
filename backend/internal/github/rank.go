package github

// RankInfo represents the developer rank and progress information
type RankInfo struct {
	Rank              string `json:"rank"`
	Score             int    `json:"score"`
	NextRank          string `json:"next_rank,omitempty"`
	NextRankThreshold int    `json:"next_rank_threshold,omitempty"`
	ProgressPercent   int    `json:"progress_percent"`
}

// Rank tier thresholds
const (
	ThresholdSPlus = 2000
	ThresholdS     = 1000
	ThresholdAPlus = 500
	ThresholdA     = 200
	ThresholdBPlus = 100
	ThresholdB     = 50
)

// CalculateRank calculates a developer rank based on GitHub statistics
func CalculateRank(stats UserProfileStats) RankInfo {
	// Define weights for each metric
	const (
		WeightCommits     = 2
		WeightPRs         = 3
		WeightIssues      = 1
		WeightReviews     = 2
		WeightStarsEarned = 4
		WeightFollowers   = 1
	)

	// Calculate total score
	totalScore := (stats.TotalCommits * WeightCommits) +
		(stats.TotalPullRequests * WeightPRs) +
		(stats.TotalIssues * WeightIssues) +
		(stats.TotalReviews * WeightReviews) +
		(stats.TotalStarsEarned * WeightStarsEarned) +
		(stats.Followers * WeightFollowers)

	// Determine rank tier and next rank
	var rank, nextRank string
	var nextThreshold int

	switch {
	case totalScore >= ThresholdSPlus:
		rank = "S+"
		nextRank = ""
		nextThreshold = 0
	case totalScore >= ThresholdS:
		rank = "S"
		nextRank = "S+"
		nextThreshold = ThresholdSPlus
	case totalScore >= ThresholdAPlus:
		rank = "A+"
		nextRank = "S"
		nextThreshold = ThresholdS
	case totalScore >= ThresholdA:
		rank = "A"
		nextRank = "A+"
		nextThreshold = ThresholdAPlus
	case totalScore >= ThresholdBPlus:
		rank = "B+"
		nextRank = "A"
		nextThreshold = ThresholdA
	case totalScore >= ThresholdB:
		rank = "B"
		nextRank = "B+"
		nextThreshold = ThresholdBPlus
	default:
		rank = "C"
		nextRank = "B"
		nextThreshold = ThresholdB
	}

	// Calculate progress percentage to next rank
	var progressPercent int
	if nextThreshold > 0 {
		// Find the current threshold
		var currentThreshold int
		switch rank {
		case "S":
			currentThreshold = ThresholdS
		case "A+":
			currentThreshold = ThresholdAPlus
		case "A":
			currentThreshold = ThresholdA
		case "B+":
			currentThreshold = ThresholdBPlus
		case "B":
			currentThreshold = ThresholdB
		default:
			currentThreshold = 0
		}

		// Calculate percentage within current tier
		tierRange := nextThreshold - currentThreshold
		progressInTier := totalScore - currentThreshold
		progressPercent = (progressInTier * 100) / tierRange

		// Clamp between 0 and 100
		if progressPercent < 0 {
			progressPercent = 0
		} else if progressPercent > 100 {
			progressPercent = 100
		}
	} else {
		// Maximum rank achieved
		progressPercent = 100
	}

	return RankInfo{
		Rank:              rank,
		Score:             totalScore,
		NextRank:          nextRank,
		NextRankThreshold: nextThreshold,
		ProgressPercent:   progressPercent,
	}
}
