package repository

// PRWithReviewer представляет открытый PR с ревьювером из команды
type PRWithReviewer struct {
	PullRequestID string
	ReviewerID    string
	AuthorID      string
}
