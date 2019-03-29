package user

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
	"github.com/rmitsubayashi/bodyweight-server/src/model/server"
)

func serverToClientUser(su server.User) client.User {
	return client.User {
		Points: su.Points,
		Cat1Level: su.Cat1Level,
		Cat2Level: su.Cat2Level,
		Cat3Level: su.Cat3Level,
		Cat4Level: su.Cat4Level,
		Cat5Level: su.Cat5Level,
		Cat6Level: su.Cat6Level,
	}
}