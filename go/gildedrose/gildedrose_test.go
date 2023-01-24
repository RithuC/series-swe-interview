package gildedrose

import (
	"testing"
)

/*
- Once the sell by date has passed, Quality degrades twice as fast - DONE
- The Quality of an item is never negative - DONE
- "Aged Brie" actually increases in Quality the older it gets - DONE
- The Quality of an item is never more than 50 - DONE
- "Sulfuras", being a legendary item, never has to be sold or decreases in Quality - DONE
- "Backstage passes", like aged brie, increases in Quality as its SellIn value approaches;
 			- Quality increases by 2 when there are 10 days or less and
			- by 3 when there are 5 days or less but
			- Quality drops to 0 after the concert
*/

//var items = []*Item{
//	{"Standard", 5, 50},
//	{"Past Sell Date", -1, 50},
//	{"Quality Zero", 7, 0},
//	{"Aged Brie", 7, 30},
//	{"Aged Brie", 7, 50},                                  // Can't go over 50 Quality
//	{"Aged Brie", -5, 10},                                 // Past sell date quality increase double
//	{"Sulfuras, Hand of Ragnaros", 7, 80},                 // Doesn't change
//	{"Sulfuras, Hand of Ragnaros", -10, 80},               // Doesn't change
//	{"Backstage passes to a TAFKAL80ETC concert", 15, 10}, // Normal quality decrease
//	{"Backstage passes to a BTS concert", 10, 10},         // Double quality increase
//	{"Backstage passes to a DPR concert", 7, 10},          // Double quality increase
//	{"Backstage passes to a DEAN concert", 5, 10},         // Triple quality increase
//	{"Backstage passes to a Mitski concert", 3, 10},       // Triple quality increase
//	{"Backstage passes to a Tabber concert", 1, 10},       // Triple quality increase
//	{"Backstage passes to the Lumineers concert", -1, 50}, // Zero Quality
//	{"Conjured", 5, 50},
//	{"Conjured", 1, 0},
//}

//var expected = []*Item{
//	{"Standard", 4, 49},
//	{"Past Sell Date", -2, 48},
//	{"Quality Zero", 6, 0},
//	{"Aged Brie", 6, 31},
//	{"Aged Brie", 6, 50},                                  // Can't go over 50 Quality
//	{"Aged Brie", -6, 12},                                 // Past sell date quality increase double
//	{"Sulfuras, Hand of Ragnaros", 7, 80},                 // Doesn't change
//	{"Sulfuras, Hand of Ragnaros", -10, 80},               // Doesn't change
//	{"Backstage passes to a TAFKAL80ETC concert", 14, 11}, // Normal quality decrease
//	{"Backstage passes to a BTS concert", 9, 12},          // Double quality increase
//	{"Backstage passes to a DPR concert", 6, 12},          // Double quality increase
//	{"Backstage passes to a DEAN concert", 4, 13},         // Triple quality increase
//	{"Backstage passes to a Mitski concert", 2, 13},       // Triple quality increase
//	{"Backstage passes to a Tabber concert", 0, 13},       // Triple quality increase
//	{"Backstage passes to the Lumineers concert", -2, 0},  // Zero Quality
//	{"Conjured", 4, 48},
//	{"Conjured", 0, 0},
//}

func Test_Foo(t *testing.T) {
	var tests = []struct {
		name  string
		input []*Item
		want  []*Item
	}{
		{"Normal items",
			[]*Item{{"Standard", 5, 50},
				{"Past Sell Date", -1, 50},
				{"Quality Zero", 7, 0}},
			[]*Item{{"Standard", 4, 49},
				{"Past Sell Date", -2, 48},
				{"Quality Zero", 6, 0}}},

		{"Aged Brie",
			[]*Item{{"Aged Brie", 7, 30},
				{"Aged Brie", 7, 50},
				{"Aged Brie", -5, 10}},
			[]*Item{{"Aged Brie", 6, 31},
				{"Aged Brie", 6, 50},
				{"Aged Brie", -6, 12}}},

		{"Sulfuras",
			[]*Item{{"Sulfuras, Hand of Ragnaros", 7, 80},
				{"Sulfuras, Hand of Ragnaros", -10, 80}},
			[]*Item{{"Sulfuras, Hand of Ragnaros", 7, 80},
				{"Sulfuras, Hand of Ragnaros", -10, 80}}},

		{"Backstage passes",
			[]*Item{{"Backstage passes to a TAFKAL80ETC concert", 15, 10},
				{"Backstage passes to a TAFKAL80ETC concert", 10, 10},
				{"Backstage passes to a TAFKAL80ETC concert", 7, 10},
				{"Backstage passes to a TAFKAL80ETC concert", 5, 10},
				{"Backstage passes to a TAFKAL80ETC concert", 3, 10},
				{"Backstage passes to a TAFKAL80ETC concert", 1, 10},
				{"Backstage passes to a TAFKAL80ETC concert", -1, 50}},
			[]*Item{{"Backstage passes to a TAFKAL80ETC concert", 14, 11},
				{"Backstage passes to a TAFKAL80ETC concert", 9, 12},
				{"Backstage passes to a TAFKAL80ETC concert", 6, 12},
				{"Backstage passes to a TAFKAL80ETC concert", 4, 13},
				{"Backstage passes to a TAFKAL80ETC concert", 2, 13},
				{"Backstage passes to a TAFKAL80ETC concert", 0, 13},
				{"Backstage passes to a TAFKAL80ETC concert", -2, 0}}},

		{"Conjured",
			[]*Item{{"Conjured", 5, 50},
				{"Conjured", 1, 0}},
			[]*Item{{"Conjured", 4, 48},
				{"Conjured", 0, 0}}},
	}

	//for i := 0; i < len(items); i++ {
	//	if items[i].Name != expected[i].Name {
	//		t.Errorf("Name: Expected %s but got %s ", expected[i].Name, items[i].Name)
	//	}
	//	if items[i].SellIn != expected[i].SellIn {
	//		t.Errorf("Name: Expected %d but got %d ", expected[i].SellIn, items[i].SellIn)
	//	}
	//	if items[i].Quality != expected[i].Quality {
	//		t.Errorf("Name: Expected %d but got %d ", expected[i].Quality, items[i].Quality)
	//	}
	//}

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
