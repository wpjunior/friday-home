package api

type Profile struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	YoutubeChannels []string `json:"youtubeChannels"`
}

var Profiles = []Profile{
	{
		ID:   "off",
		Name: "Desligado",
	},
	{
		ID:   "vibe-leidy",
		Name: "Vibe Leidy",
		YoutubeChannels: []string{
			"https://www.youtube.com/watch?v=VIWVfkF2IeI",
		},
	},
	{
		ID:   "metal",
		Name: "Metal Júnior",
		YoutubeChannels: []string{
			"https://www.youtube.com/watch?v=tKJMSKmueGk",
		},
	},
}

func GetProfile(id string) *Profile {
	for _, profile := range Profiles {
		if profile.ID == id {
			return &profile
		}
	}
	return nil
}
