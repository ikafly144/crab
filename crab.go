package crab

import "fmt"

func New蟹(名前 string) 蟹 {
	return 蟹{
		名前: 名前,
		鍋:  false,
	}
}

type 蟹 struct {
	名前 string
	鍋  bool
}

func (this 蟹) String() string {
	return this.名前
}

func (this *蟹) Make鍋() *鍋[蟹] {
	this.鍋 = true
	return &鍋[蟹]{具材: this}
}

type 鍋[具材 fmt.Stringer] struct {
	具材 *具材
}

func (this 鍋[具材]) String() string {
	return fmt.Sprintf("鍋: 具材(%s)", this.具材)
}
