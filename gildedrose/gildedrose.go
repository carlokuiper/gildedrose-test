// Package gildedrose contains a possible solution to the gilded rose refactoring kata
// The exercise can be found at https://github.com/emilybache/GildedRose-Refactoring-Kata
package gildedrose

const (
	minQuality = 0
	maxQuality = 50
)

// Item is an item sold by Gilded Rose.
type Item struct {
	Name            string
	SellIn, Quality int
}

func (i *Item) updateNormal() {
	qualityChange := -2
	if i.SellIn > 0 {
		qualityChange = -1
	}
	i.Quality = max(i.Quality+qualityChange, minQuality)
}

func (i *Item) updateAgedBrie() {
	qualityChange := 2
	if i.SellIn > 0 {
		qualityChange = 1
	}
	i.Quality = min(i.Quality+qualityChange, maxQuality)
}

func (i *Item) updateBackstage() {
	var qualityChange int
	switch {
	case i.SellIn > 10:
		qualityChange = 1
	case i.SellIn > 5:
		qualityChange = 2
	case i.SellIn > 0:
		qualityChange = 3
	default:
		qualityChange = -i.Quality
	}
	i.Quality = min(i.Quality+qualityChange, maxQuality)
}

func (i *Item) updateConjured() {
	qualityChange := -4
	if i.SellIn > 0 {
		qualityChange = -2
	}
	i.Quality = max(i.Quality+qualityChange, minQuality)
}

// UpdateQuality updates the quality of all items.
func UpdateQuality(items ...*Item) {
	for _, item := range items {
		switch item.Name {
		case "Sulfuras, Hand of Ragnaros":
			continue
		case "Aged Brie":
			item.updateAgedBrie()
		case "Backstage passes to a TAFKAL80ETC concert":
			item.updateBackstage()
		case "Conjured Mana Cake":
			item.updateConjured()
		default:
			item.updateNormal()
		}
		item.SellIn--
	}
}
