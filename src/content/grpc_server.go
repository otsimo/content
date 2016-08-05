package content

import (
	"errors"
	"sort"

	apipb "github.com/otsimo/otsimopb"
	"golang.org/x/net/context"
)

type contentGrpcServer struct {
	server *Server
}
type contentSorter struct {
	contents     []*apipb.Content
	orderAsc     bool
	sortByWeight bool
	category     bool
}

func (slice contentSorter) Len() int {
	return len(slice.contents)
}

func (slice contentSorter) Less(i, j int) bool {
	if slice.sortByWeight {
		if slice.category {
			if slice.orderAsc {
				return slice.contents[i].CategoryWeight < slice.contents[j].CategoryWeight
			} else {
				return slice.contents[i].CategoryWeight > slice.contents[j].CategoryWeight
			}
		} else {
			if slice.orderAsc {
				return slice.contents[i].Weight < slice.contents[j].Weight
			} else {
				return slice.contents[i].Weight > slice.contents[j].Weight
			}
		}
	} else {
		if slice.orderAsc {
			return slice.contents[i].Date < slice.contents[j].Date
		} else {
			return slice.contents[i].Date > slice.contents[j].Date
		}
	}
}

func (slice contentSorter) Swap(i, j int) {
	slice.contents[i], slice.contents[j] = slice.contents[j], slice.contents[i]
}

func (w *contentGrpcServer) List(_ context.Context, query *apipb.ContentListRequest) (*apipb.ContentListResponse, error) {
	var contents []*apipb.Content
	if query.Language == "" {
		query.Language = w.server.Config.DefaultLanguage
	}
	for _, c := range w.server.Content.contents {
		if c.Language != query.Language {
			continue
		}
		if query.Status == apipb.ContentListRequest_ONLY_APPROVED && c.Draft {
			continue
		}
		if query.Status == apipb.ContentListRequest_ONLY_DRAFT && !c.Draft {
			continue
		}
		if query.Category != "" && query.Category != c.Category {
			continue
		}

		if len(query.Categories) > 0 {
			founded := false
			for _, cat := range query.Categories {
				if cat == c.Category {
					founded = true
					break
				}
			}
			if !founded {
				continue
			}
		}

		if len(query.ExceptCategories) > 0 {
			founded := false
			for _, cat := range query.ExceptCategories {
				if cat == c.Category {
					founded = true
					break
				}
			}
			if founded {
				continue
			}
		}

		cp := *c
		cp.Markdown = []byte{}
		contents = append(contents, &cp)
	}
	sorter := &contentSorter{
		contents:     contents,
		orderAsc:     (query.Order == apipb.ContentListRequest_ASC),
		sortByWeight: (query.Sort == apipb.ContentListRequest_WEIGHT),
		category:     (query.Category != ""),
	}
	sort.Sort(sorter)
	return &apipb.ContentListResponse{
		AssetVersion: w.server.Content.assetVersion,
		Contents:     sorter.contents,
	}, nil
}

func (w *contentGrpcServer) Get(_ context.Context, in *apipb.ContentGetRequest) (*apipb.Content, error) {
	for _, c := range w.server.Content.contents {
		if c.Slug == in.Slug {
			return c, nil
		}
	}
	return nil, errors.New("not found")
}
