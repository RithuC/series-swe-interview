package gildedrose

import (
	"testing"
)

/***
Rithu handoff notes:
	- The original gildedrose package import was giving me issues,
		so I changed the package name to use the Item defined locally
		(hopefully this is okay since I didn't modify Item).
	- With more time, I would've liked to split out the test into different fns
		to correspond to the new update fns but hopefully at least the grouping with
		'name' should provide some extra clarity. For example, failing tests should be
		prefixed with grouping name:
			--- FAIL: TestUpdateQuality/Conjured#01 (0.00s)
        				gildedrose_test.go:83: SellIn: Expected -1 but got 0
        				gildedrose_test.go:86: Quality: Expected 6 but got 0
***/

func TestUpdateQuality(t *testing.T) {
	var tests = []struct {
		name  string
		input []*Item
		want  []*Item
	}{
		{"Normal items",
			[]*Item{{"Standard", 5, 50},
				{"Past Sell Date", -1, 50}, // Quality decreases twice as fast
				{"Quality Zero", 7, 0},     // Quality doesn't go below 0
				{"Quality Zero, Past Sell Date", -3, 0}}, // Quality doesn't go below 0
			[]*Item{{"Standard", 4, 49},
				{"Past Sell Date", -2, 48},
				{"Quality Zero", 6, 0},
				{"Quality Zero, Past Sell Date", -4, 0}}},

		{"Aged Brie",
			[]*Item{{"Aged Brie", 7, 30}, // Quality increases
				{"Aged Brie", 7, 50}, // Increase doesn't go over 50
				{"Aged Brie", -5, 10}}, // Quality increases twice as fast past sell date
			[]*Item{{"Aged Brie", 6, 31},
				{"Aged Brie", 6, 50},
				{"Aged Brie", -6, 12}}},

		{"Sulfuras",
			[]*Item{{"Sulfuras, Hand of Ragnaros", 7, 80}, // Expect no change
				{"Sulfuras, Hand of Ragnaros", -10, 80}},
			[]*Item{{"Sulfuras, Hand of Ragnaros", 7, 80}, // Expect no change
				{"Sulfuras, Hand of Ragnaros", -10, 80}}},

		{"Backstage passes",
			[]*Item{{"Backstage passes to a TAFKAL80ETC concert", 15, 10}, // Quality increases
				{"Backstage passes to a TAFKAL80ETC concert", 10, 10}, // Quality increases twice as fast
				{"Backstage passes to a TAFKAL80ETC concert", 7, 10},  // Quality increases twice as fast
				{"Backstage passes to a TAFKAL80ETC concert", 5, 10},  // Quality increases thrice as fast
				{"Backstage passes to a TAFKAL80ETC concert", 3, 10},  // Quality increases thrice as fast
				{"Backstage passes to a TAFKAL80ETC concert", 1, 10},  // Quality increases thrice as fast
				{"Backstage passes to a TAFKAL80ETC concert", -1, 40}, // Quality is 0
				{"Backstage passes to a TAFKAL80ETC concert", 0, 30},  // Quality is 0
				{"Backstage passes to a TAFKAL80ETC concert", -1, 0}}, // Quality doesn't go below 0
			[]*Item{{"Backstage passes to a TAFKAL80ETC concert", 14, 11},
				{"Backstage passes to a TAFKAL80ETC concert", 9, 12},
				{"Backstage passes to a TAFKAL80ETC concert", 6, 12},
				{"Backstage passes to a TAFKAL80ETC concert", 4, 13},
				{"Backstage passes to a TAFKAL80ETC concert", 2, 13},
				{"Backstage passes to a TAFKAL80ETC concert", 0, 13},
				{"Backstage passes to a TAFKAL80ETC concert", -2, 0},
				{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
				{"Backstage passes to a TAFKAL80ETC concert", -2, 0}}},

		{"Conjured",
			[]*Item{{"Conjured", 5, 50}, // Quality decreases twice as fast
				{"Conjured", 1, 0}, // Quality doesn't go below 0
				{"Conjured", 0, 10}}, // Quality decreases twice as fast x 2
			[]*Item{{"Conjured", 4, 48},
				{"Conjured", 0, 0},
				{"Conjured", -1, 6}}},
	}

	for _, test := range tests {
		UpdateQuality(test.input)
		for i := 0; i < len(test.input); i++ {
			t.Run(test.name, func(t *testing.T) {
				if test.input[i].Name != test.want[i].Name {
					t.Errorf("Name: Expected %s but got %s ", test.want[i].Name, test.input[i].Name)
				}
				if test.input[i].SellIn != test.want[i].SellIn {
					t.Errorf("SellIn: Expected %d but got %d ", test.want[i].SellIn, test.input[i].SellIn)
				}
				if test.input[i].Quality != test.want[i].Quality {
					t.Errorf("Quality: Expected %d but got %d ", test.want[i].Quality, test.input[i].Quality)
				}
			})
		}
	}
}
