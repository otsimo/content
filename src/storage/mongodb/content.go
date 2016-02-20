package mongodb

import (
	"models"

	"github.com/otsimo/api/apipb"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (d *MongoDBDriver) GetById(id bson.ObjectId) (*apipb.Content, error) {
	c := d.Session.DB("").C(ContentCollection)
	var doc apipb.Content
	err := c.FindId(id).One(&doc)
	if err == mgo.ErrNotFound {
		return nil, models.ErrorNotFound
	}
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

func (d *MongoDBDriver) GetBySlug(slug string) (*apipb.Content, error) {
	c := d.Session.DB("").C(ContentCollection)
	var doc apipb.Content
	err := c.Find(bson.M{"slug": slug}).One(&doc)
	if err == mgo.ErrNotFound {
		return nil, models.ErrorNotFound
	}
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

func (d *MongoDBDriver) Put(mc *apipb.Content) error {
	c := d.Session.DB("").C(ContentCollection)
	return c.Insert(mc)
}

func (d *MongoDBDriver) Update(mc *apipb.Content) error {
	c := d.Session.DB("").C(ContentCollection)
	return c.Update(bson.M{"slub": mc.Slug}, bson.M{"$set": mc})
}

func (d *MongoDBDriver) ChangeStatus(title string, stat apipb.Content_Status) error {
	c := d.Session.DB("").C(ContentCollection)
	return c.Update(bson.M{"title": title}, bson.M{"$set": bson.M{"status": stat}})
}

func (d *MongoDBDriver) List(q apipb.ContentListRequest) ([]*apipb.Content, error) {
	var result []*apipb.Content

	/*	c := d.Session.DB("").C(ContentCollection)

		query := bson.M{}
		if q.Limit == 0 {
			q.Limit = 100
		}
		switch q.Status {
		case apipb.CatalogListRequest_ONLY_DRAFT:
			query["status"] = apipb.CatalogStatus_DRAFT
		case apipb.CatalogListRequest_ONLY_APPROVED:
			query["status"] = apipb.CatalogStatus_APPROVED
		}
		if q.Time > 0 {
			query["visible_at"] = bson.M{"$lte": q.Time}
			if q.HideExpired {
				m := int64(0)
				now := models.MillisecondsNow()
				if q.Time > now {
					m = q.Time
				} else {
					m = now
				}
				query["expires_at"] = bson.M{"$gt": m}
			} else {
				query["expires_at"] = bson.M{"$gt": q.Time}
			}
		} else if q.HideExpired {
			query["expires_at"] = bson.M{"$gt": models.MillisecondsNow()}
		}

		iter := c.Find(query).Limit(int(q.Limit)).Sort("-visible_at").Iter()
		err := iter.All(&result)
		if err != nil {
			return nil, err
		}
	*/
	return result, nil
}
