package profile

import "goyave.dev/goyave/v4/validation"

var DestroyLikedAttractionsRequest = validation.RuleSet{
	"attractionIds": validation.List{"required","array:integer"},
}
