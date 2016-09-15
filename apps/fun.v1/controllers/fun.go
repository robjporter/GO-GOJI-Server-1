package controllers

import (
	"net/http"

	"../../../render"
	"../../../system"
	"github.com/zenazn/goji/web"
)

func TMP(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := render.GetBaseTemplates(c)
	test := c.Env["Settings"].(*system.Settings)
	test.Count += 1
	templates = append(templates, "apps/fun.v1/views/home.html")
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LogoGithub(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/logo/github.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LogoGreenApple(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/logo/greenapple.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LogoCisco1(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/logo/cisco1.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalAnimals(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/animals.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalCat(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/cat.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalBear1(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/bear1.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalBear2(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/bear2.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalDog(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/dog.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalElephant(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/elephant.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalLion(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/lion.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalMonkey(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/monkey.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalPenguin(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/penguin.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalSheep(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/sheep.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalSimpleOwl(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/simpleowl.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ObjectClouds(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/objects/clouds.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ObjectMacKeyboard(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/objects/mackeyboard.object.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersGoofy(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/goofy.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersHomer2(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/homer2.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersMinions(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/minions.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersCoder(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/coder.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersApu(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/apu.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersBart(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/bart.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersComic(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/comicbookguy.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersHomer(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/homer.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersItchy(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/itchy.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersKrusty(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/krusty.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersLisa(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/lisa.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersMaggie(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/maggie.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersMarge(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/marge.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersMrBurns(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/mrburns.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersNedFlanders(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/nedflanders.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersRalphWiggum(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/ralphwiggum.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersSmithers(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/smithers.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
