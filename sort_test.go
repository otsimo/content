package content

import (
	"sort"
	"testing"

	apipb "github.com/otsimo/otsimopb"
)

func TestAscWeightSorting(t *testing.T) {
	contents := []*apipb.Content{
		&apipb.Content{Slug: "10", Weight: 10},
		&apipb.Content{Slug: "9", Weight: 9},
		&apipb.Content{Slug: "12", Weight: 12},
		&apipb.Content{Slug: "5", Weight: 5},
		&apipb.Content{Slug: "15", Weight: 15},
	}
	sorter := &contentSorter{
		contents:     contents,
		orderAsc:     true,
		sortByWeight: true,
		category:     false,
	}
	sort.Sort(sorter)

	if len(sorter.contents) != len(contents) {
		t.Fatalf("len(sorter.contents) want=%d got=%d", len(contents), len(sorter.contents))
	}

	if sorter.contents[0].Slug != "5" {
		t.Fatalf("sorter.contents[0].Slug want=%s got=%s", "5", sorter.contents[0].Slug)
	}
	if sorter.contents[4].Slug != "15" || sorter.contents[4].Weight != 15 {
		t.Fatalf("sorter.contents[4].Slug want=%s got=%s", "15", sorter.contents[4].Slug)
	}
}

func TestDscWeightSorting(t *testing.T) {
	contents := []*apipb.Content{
		&apipb.Content{Slug: "10", Weight: 10},
		&apipb.Content{Slug: "9", Weight: 9},
		&apipb.Content{Slug: "12", Weight: 12},
		&apipb.Content{Slug: "5", Weight: 5},
		&apipb.Content{Slug: "15", Weight: 15},
	}
	sorter := &contentSorter{
		contents:     contents,
		orderAsc:     false,
		sortByWeight: true,
		category:     false,
	}
	sort.Sort(sorter)

	if len(sorter.contents) != len(contents) {
		t.Fatalf("len(sorter.contents) want=%d got=%d", len(contents), len(sorter.contents))
	}

	if sorter.contents[0].Slug != "15" {
		t.Fatalf("sorter.contents[0].Slug want=%s got=%s", "15", sorter.contents[0].Slug)
	}
	if sorter.contents[4].Slug != "5" {
		t.Fatalf("sorter.contents[4].Slug want=%s got=%s", "5", sorter.contents[4].Slug)
	}
}

func TestDscCategoryWeightSorting(t *testing.T) {
	contents := []*apipb.Content{
		&apipb.Content{Slug: "10", Weight: 1023, CategoryWeight: 10},
		&apipb.Content{Slug: "9", Weight: 512, CategoryWeight: 9},
		&apipb.Content{Slug: "12", Weight: 2, CategoryWeight: 12},
		&apipb.Content{Slug: "5", Weight: 17, CategoryWeight: 5},
		&apipb.Content{Slug: "15", Weight: 12, CategoryWeight: 15},
	}
	sorter := &contentSorter{
		contents:     contents,
		orderAsc:     false,
		sortByWeight: true,
		category:     true,
	}
	sort.Sort(sorter)

	if len(sorter.contents) != len(contents) {
		t.Fatalf("len(sorter.contents) want=%d got=%d", len(contents), len(sorter.contents))
	}

	if sorter.contents[0].Slug != "15" {
		t.Fatalf("sorter.contents[0].Slug want=%s got=%s", "15", sorter.contents[0].Slug)
	}
	if sorter.contents[4].Slug != "5" {
		t.Fatalf("sorter.contents[4].Slug want=%s got=%s", "5", sorter.contents[4].Slug)
	}
}

func TestAscCategoryWeightSorting(t *testing.T) {
	contents := []*apipb.Content{
		&apipb.Content{Slug: "10", Weight: 1023, CategoryWeight: 10},
		&apipb.Content{Slug: "9", Weight: 512, CategoryWeight: 9},
		&apipb.Content{Slug: "12", Weight: 2, CategoryWeight: 12},
		&apipb.Content{Slug: "5", Weight: 17, CategoryWeight: 5},
		&apipb.Content{Slug: "15", Weight: 12, CategoryWeight: 15},
	}
	sorter := &contentSorter{
		contents:     contents,
		orderAsc:     true,
		sortByWeight: true,
		category:     true,
	}
	sort.Sort(sorter)

	if len(sorter.contents) != len(contents) {
		t.Fatalf("len(sorter.contents) want=%d got=%d", len(contents), len(sorter.contents))
	}

	if sorter.contents[0].Slug != "5" {
		t.Fatalf("sorter.contents[0].Slug want=%s got=%s", "5", sorter.contents[0].Slug)
	}
	if sorter.contents[4].Slug != "15" {
		t.Fatalf("sorter.contents[4].Slug want=%s got=%s", "15", sorter.contents[4].Slug)
	}
}

func TestAscTimeSorting(t *testing.T) {
	contents := []*apipb.Content{
		&apipb.Content{Slug: "10", Date: 10},
		&apipb.Content{Slug: "9", Date: 9},
		&apipb.Content{Slug: "12", Date: 12},
		&apipb.Content{Slug: "5", Date: 5},
		&apipb.Content{Slug: "15", Date: 15},
	}
	sorter := &contentSorter{
		contents:     contents,
		orderAsc:     true,
		sortByWeight: false,
		category:     false,
	}
	sort.Sort(sorter)

	if len(sorter.contents) != len(contents) {
		t.Fatalf("len(sorter.contents) want=%d got=%d", len(contents), len(sorter.contents))
	}

	if sorter.contents[0].Slug != "5" {
		t.Fatalf("sorter.contents[0].Slug want=%s got=%s", "5", sorter.contents[0].Slug)
	}
	if sorter.contents[4].Slug != "15" {
		t.Fatalf("sorter.contents[4].Slug want=%s got=%s", "15", sorter.contents[4].Slug)
	}
}

func TestDscTimeSorting(t *testing.T) {
	contents := []*apipb.Content{
		&apipb.Content{Slug: "10", Date: 10},
		&apipb.Content{Slug: "9", Date: 9},
		&apipb.Content{Slug: "12", Date: 12},
		&apipb.Content{Slug: "5", Date: 5},
		&apipb.Content{Slug: "15", Date: 15},
	}
	sorter := &contentSorter{
		contents:     contents,
		orderAsc:     false,
		sortByWeight: false,
		category:     true,
	}
	sort.Sort(sorter)

	if len(sorter.contents) != len(contents) {
		t.Fatalf("len(sorter.contents) want=%d got=%d", len(contents), len(sorter.contents))
	}

	if sorter.contents[0].Slug != "15" {
		t.Fatalf("sorter.contents[0].Slug want=%s got=%s", "15", sorter.contents[0].Slug)
	}
	if sorter.contents[4].Slug != "5" {
		t.Fatalf("sorter.contents[4].Slug want=%s got=%s", "5", sorter.contents[4].Slug)
	}
}
