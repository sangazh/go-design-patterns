package flyweight

import (
	"time"
)

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

const (
	TeamA = iota
	TeamB
)

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func NewTeamFactory() *teamFlyweightFactory{
	return &teamFlyweightFactory{
		createdTeams: make(map[int]*Team, 0),
	}
}

func (t *teamFlyweightFactory) GetTeam(teamId int) *Team {
	if v, ok := t.createdTeams[teamId]; ok {
		return v
	}
	t.createdTeams[teamId] = getTeamFactory(teamId)

	return t.createdTeams[teamId]
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}

func getTeamFactory(teamId int) *Team {
	switch teamId {
	case TeamA:
		return &Team{
			ID: 1, Name: "TEAM_A",
		}
	case TeamB:
		return &Team{
			ID: 2, Name: "TEAM_B",
		}
	}
	return nil
}
