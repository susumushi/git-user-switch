package profile

import "testing"

func TestProfile(t *testing.T) {
	t.Run("entry", func(t *testing.T) {
		ps := Profiles{}
		ps.Load()
		ps.Set(Profile{
			Name:     "example.testuser2",
			Email:    "test2@example.com",
			NickName: "etu2",
		})
		ps.Save()
		ps.Flush()
	})
}
