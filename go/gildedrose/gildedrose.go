package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const (
	AgedBrie        string = "Aged Brie"
	Sulfuras        string = "Sulfuras, Hand of Ragnaros"
	BackstagePasses string = "Backstage passes to a TAFKAL80ETC concert"
	Conjured        string = "Conjured"
)

const (
	MaxQuality          int = 50
	MinQuality          int = 0
	MaxQualityLegendary int = 80
	MinSellIn           int = 0
)

func UpdateQuality(items []*Item) {
	for _, item := range items {

		// Nothing changes for a legendary item
		if item.Name == Sulfuras {
			continue
		}

		switch item.Name {
		case Conjured:
			updateQualityConjured(item)
		case AgedBrie:
			updateQualityBrie(item)
		case BackstagePasses:
			updateQualityBackstagePasses(item)
		default:
			updateQuality(item)
		}

		updateSellIn(item)

		if item.SellIn < MinSellIn {
			switch item.Name {
			case Conjured:
				updateQualityConjured(item)
			case AgedBrie:
				updateQualityBrie(item)
			case BackstagePasses:
				updateQualityBackstagePasses(item)
			default:
				updateQuality(item)
			}
		}
	}
}

func updateSellIn(item *Item) {
	item.SellIn--
}

func updateQuality(item *Item) {
	if item.Quality > MinQuality {
		item.Quality--
	}
}

func updateQualityConjured(item *Item) {
	if item.Quality > MinQuality {
		item.Quality -= 2
	}
}

func updateQualityBrie(item *Item) {
	if item.Quality < MaxQuality {
		item.Quality++
	}
}

func updateQualityBackstagePasses(item *Item) {
	if item.SellIn < MinSellIn {
		item.Quality = MinQuality
	} else if item.Quality < MaxQuality {
		if item.SellIn < 11 {
			item.Quality++
		}
		if item.SellIn < 6 {
			item.Quality++
		}
		item.Quality++
	}
}
