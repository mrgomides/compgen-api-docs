package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Document struct {
	ID           int64         `orm:"auto;column(id)" json:"ID"`
	Name         string        `orm:"size(128)" json:"Name"`
	Status       string        `orm:"size(128)" json:"Status"`
	PublicCopy   string        `orm:"size(128)" json:"PublicCopy"`
	PrivateCopy  string        `orm:"size(128)" json:"PrivateCopy"`
	Type         string        `orm:"size(128)" json:"Type"`
	User         *User         `orm:"column(user);rel(fk)" json:"User"`
	ServiceOrder *ServiceOrder `orm:"column(service_order);rel(fk)" json:"ServiceOrder"`
}

func init() {
	orm.RegisterModel(new(Document))
}

func (a *Document) TableName() string {
	return "Document"
}

// AddDocument insert a new Document into database and returns
// last inserted ID on success.
func AddDocument(m *Document) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDocumentByID retrieves Document by ID. Returns error if
// ID doesn't exist
func GetDocumentByID(id int64) (v *Document, err error) {
	o := orm.NewOrm()
	v = &Document{ID: id}
	if err = o.QueryTable(new(Document)).Filter("ID", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDocument retrieves all Document matches certain condition. Returns empty list if
// no records exist
func GetAllDocument(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Document))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Document
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateDocument updates Document by ID and returns error if
// the record to be updated doesn't exist
func UpdateDocumentByID(m *Document) (err error) {
	o := orm.NewOrm()
	v := Document{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDocument deletes Document by ID and returns error if
// the record to be deleted doesn't exist
func DeleteDocument(id int64) (err error) {
	o := orm.NewOrm()
	v := Document{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Document{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
