package routes

import (
	"net/http"

	"../controllers"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func Prefetch() {

}

func Routes(m *web.Mux) {
	goji.Handle("/api/v1/fun/*", m)
	goji.Get("/api/v1/fun", http.RedirectHandler("/api/v1/fun/", 301))

	m.Get("/api/v1/fun/logo/github", controllers.LogoGithub)
	m.Get("/api/v1/fun/logo/greenapple", controllers.LogoGreenApple)
	m.Get("/api/v1/fun/logo/cisco1", controllers.LogoCisco1)

	m.Get("/api/v1/fun/animal/simpleowl", controllers.AnimalSimpleOwl)
	m.Get("/api/v1/fun/animal/animals", controllers.AnimalAnimals)
	m.Get("/api/v1/fun/animal/cat", controllers.AnimalCat)
	m.Get("/api/v1/fun/animal/bear1", controllers.AnimalBear1)
	m.Get("/api/v1/fun/animal/bear2", controllers.AnimalBear2)
	m.Get("/api/v1/fun/animal/dog", controllers.AnimalDog)
	m.Get("/api/v1/fun/animal/elephant", controllers.AnimalElephant)
	m.Get("/api/v1/fun/animal/lion", controllers.AnimalLion)
	m.Get("/api/v1/fun/animal/monkey", controllers.AnimalMonkey)
	m.Get("/api/v1/fun/animal/penguin", controllers.AnimalPenguin)
	m.Get("/api/v1/fun/animal/sheep", controllers.AnimalSheep)

	m.Get("/api/v1/fun/object/clouds", controllers.ObjectClouds)
	m.Get("/api/v1/fun/object/mackeyboard", controllers.ObjectMacKeyboard)

	m.Get("/api/v1/fun/characters/minions", controllers.CharactersMinions)
	m.Get("/api/v1/fun/characters/goofy", controllers.CharactersGoofy)
	m.Get("/api/v1/fun/characters/coder", controllers.CharactersCoder)

	m.Get("/api/v1/fun/characters/simpsons/apu", controllers.CharactersApu)
	m.Get("/api/v1/fun/characters/simpsons/bart", controllers.CharactersBart)
	m.Get("/api/v1/fun/characters/simpsons/comicbookguy", controllers.CharactersComic)
	m.Get("/api/v1/fun/characters/simpsons/homer", controllers.CharactersHomer)
	m.Get("/api/v1/fun/characters/simpsons/homer2", controllers.CharactersHomer2)
	m.Get("/api/v1/fun/characters/simpsons/itchy", controllers.CharactersItchy)
	m.Get("/api/v1/fun/characters/simpsons/krusty", controllers.CharactersKrusty)
	m.Get("/api/v1/fun/characters/simpsons/lisa", controllers.CharactersLisa)
	m.Get("/api/v1/fun/characters/simpsons/maggie", controllers.CharactersMaggie)
	m.Get("/api/v1/fun/characters/simpsons/marge", controllers.CharactersMarge)
	m.Get("/api/v1/fun/characters/simpsons/mrburns", controllers.CharactersMrBurns)
	m.Get("/api/v1/fun/characters/simpsons/ned", controllers.CharactersNedFlanders)
	m.Get("/api/v1/fun/characters/simpsons/ralph", controllers.CharactersRalphWiggum)
	m.Get("/api/v1/fun/characters/simpsons/smithers", controllers.CharactersSmithers)
}
