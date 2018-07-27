package player

// Player describes player structure
type Player struct {
	Name  string
	Age   string
	Teams []string
}

// AddTeamUnique method is used to add team name to a player.
// If team name is present, it will be skipped.
func (p *Player) AddTeamUnique(name string) {
	// don't add duplicates
	for _, v := range p.Teams {
		if v == name {
			return
		}
	}

	p.Teams = append(p.Teams, name)
}

// PlayersList is a sortable list of Players
type PlayersList []*Player

func (pl PlayersList) Len() int           { return len(pl) }
func (pl PlayersList) Swap(i, j int)      { pl[i], pl[j] = pl[j], pl[i] }
func (pl PlayersList) Less(i, j int) bool { return pl[i].Name < pl[j].Name }
