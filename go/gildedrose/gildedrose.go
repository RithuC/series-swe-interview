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
)

func isSpecialProduct(name string) bool {
	switch name {
	case AgedBrie, Sulfuras, BackstagePasses, Conjured:
		return true
	}
	return false
}

func UpdateQuality(items []*Item) {
	for _, item := range items {

		// Nothing changes for a legendary item
		if item.Name == Sulfuras {
			continue
		}
		//} else if !isSpecialProduct(item.Name) {
		//	if item.Quality > MinQuality {
		//		item.Quality--
		//	}
		//}
		//
		//if item.Name == Conjured && item.Quality > MinQuality {
		//	item.Quality = item.Quality - 2
		//}
		//
		//if item.Name == AgedBrie || item.Name == BackstagePasses {
		//	if item.Quality < 50 {
		//		item.Quality = item.Quality + 1
		//		if item.Name == BackstagePasses {
		//			if item.SellIn < 11 {
		//				if item.Quality < 50 {
		//					item.Quality = item.Quality + 1
		//				}
		//			}
		//			if item.SellIn < 6 {
		//				if item.Quality < 50 {
		//					item.Quality = item.Quality + 1
		//				}
		//			}
		//		}
		//	}
		//}

		if item.Name != AgedBrie && item.Name != BackstagePasses {
			if item.Quality > 0 {
				item.Quality = item.Quality - 1
				if item.Name == Conjured {
					item.Quality = item.Quality - 1
				}
			}
		} else {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
				if item.Name == BackstagePasses {
					if item.SellIn < 11 {
						if item.Quality < 50 {
							item.Quality = item.Quality + 1
						}
					}
					if item.SellIn < 6 {
						if item.Quality < 50 {
							item.Quality = item.Quality + 1
						}
					}
				}
			}
		}

		item.SellIn--

		if item.SellIn < 0 {
			if item.Name != AgedBrie {
				if item.Name != BackstagePasses {
					if item.Quality > 0 {
						item.Quality = item.Quality - 1
					}
				} else {
					item.Quality = item.Quality - item.Quality
				}
			} else {
				if item.Quality < 50 {
					item.Quality = item.Quality + 1
				}
			}
		}
	}
}

func updateSellIn(item *Item) {
	item.SellIn--
}
