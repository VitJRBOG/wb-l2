package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

func ExecuteFacadeExample() {
	u := NewUserInterface("login", "Some text for an article")
	result := u.PublishNewArticle()
	fmt.Println(result)
}

type UserInterface struct {
	Profile *Profile
	Article *Article
}

func NewUserInterface(userLogin, articleText string) *UserInterface {
	return &UserInterface{
		Profile: NewProfile(userLogin),
		Article: NewArticle(articleText),
	}
}

func (u *UserInterface) PublishNewArticle() string {
	authorizationStatus := u.Profile.Authorization()
	if authorizationStatus != "Authorization has been succeded" {
		return authorizationStatus
	}

	publicationStatus := u.Article.PublishArticle()
	if publicationStatus != "Article has been successfully published" {
		return publicationStatus
	}

	return "Operation has been succeded"
}

type Profile struct {
	Login string
}

func NewProfile(login string) *Profile {
	return &Profile{
		Login: login,
	}
}

func (a *Profile) Authorization() string {
	loginStatus := a.checkLogin()
	if loginStatus != "Login is correct" {
		return "Error while authorization: login is not correct"
	}

	return "Authorization has been succeded"
}

func (a *Profile) checkLogin() string {
	return "Login is correct"
}

type Article struct {
	Text string
}

func NewArticle(text string) *Article {
	return &Article{
		Text: text,
	}
}

func (a *Article) PublishArticle() string {
	articleStatus := a.checkArticle()
	if articleStatus != "Article is unic" {
		return "Error while publishing article: article is not unic"
	}

	titleStatus := a.composeTitle()
	if titleStatus != "Title has been composed from text of the article" {
		return "Error while publishing article: coudn't compose title"
	}

	return "Article has been successfully published"
}

func (a *Article) checkArticle() string {
	return "Article is unic"
}

func (a *Article) composeTitle() string {
	return "Title has been composed from text of the article"
}
