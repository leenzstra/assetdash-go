package arcade

type ArcadeGameSession struct {
	ID           string `json:"id"`
	Created      string `json:"created"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	ArcadeGameID string `json:"arcade_game_id"`
	SpentCoins   int    `json:"spent_coins"`
	EarnedCoins  int    `json:"earned_coins"`
}

type ArcadeData struct {
	Data        string `json:"data"`
	SessionHash string `json:"session_hash"`
	Error       bool   `json:"error"`
	Msg         string `json:"msg"`
}
