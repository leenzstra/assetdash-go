package arcade

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/leenzstra/assetdash-go/client"
	"github.com/leenzstra/assetdash-go/modules"
)

var _ modules.Module = (*ArcadeModule)(nil)

type ArcadeModule struct {
	*client.AssetDashClient
}

func New(client *client.AssetDashClient) *ArcadeModule {
	return &ArcadeModule{client}
}

func (m *ArcadeModule) Name() string {
	return "Arcade Module"
}

func (m *ArcadeModule) StartArcade(gameId string, timestamp int64) (*ArcadeGameSessionResponse, error) {
	endpoint, err := m.BuildURL(m.Endpoints.StartArcade)
	if err != nil {
		return nil, err
	}

	response := &ArcadeGameSessionResponse{}
	r, err := m.Client.R().
		SetQueryParams(map[string]string{
			"timestamp": fmt.Sprint(timestamp),
		}).
		SetBody(map[string]interface{}{
			"arcade_game_id": gameId,
		}).SetResult(response).Post(endpoint)

	if err != nil || r.IsError() {
		return nil, fmt.Errorf("%s %s", r.Error(), err)
	}

	return response, nil
}

func (m *ArcadeModule) CompleteArcade(timestamp int64, gameId, sessionId string, arcadeData *ArcadeData) (*ArcadeGameSessionResponse, error) {
	endpoint, err := m.BuildURL(m.Endpoints.CompleteArcade)
	if err != nil {
		return nil, err
	}

	response := &ArcadeGameSessionResponse{}
	r, err := m.Client.R().
		SetQueryParams(map[string]string{
			"timestamp": fmt.Sprint(timestamp),
		}).
		SetBody(map[string]interface{}{
			"arcade_game_id":         gameId,
			"arcade_game_session_id": sessionId,
			"session_hash":           arcadeData.SessionHash,
			"data":                   arcadeData.Data,
		}).SetResult(response).Post(endpoint)

	if err != nil || r.IsError() {
		return nil, fmt.Errorf("%s %s", r.Error(), err)
	}

	return response, nil
}

func (m *ArcadeModule) ComputeSessionData(sessionId string, coins, score int) (*ArcadeData, error) {
	out, err := exec.Command("python", "scripts/_bcrypt.py", "-i", sessionId, "-c", fmt.Sprint(coins), "-s", fmt.Sprint(score)).Output()
	if err != nil {
		return nil, err
	}

	data := &ArcadeData{}
	err = json.Unmarshal(out, data)
	if err != nil || data.Error {
		return nil, fmt.Errorf("error: %s %s", data.Msg, err)
	}

	return data, nil
}
