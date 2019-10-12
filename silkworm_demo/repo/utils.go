package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"silkwormDemo/model/params"
)

func LimitDB(db *gorm.DB, limit, offset uint64) *gorm.DB {
	if limit != 0 {
		db = db.Limit(limit)
	}
	if offset != 0 {
		db = db.Offset(offset)
	}
	return db
}

func DefaultListDB(db *gorm.DB, req *params.ListRequest) *gorm.DB {
	if req == nil || db == nil {
		return db
	}
	if req.PartnerId != 0 {
		db = db.Where("partner_id = ?", req.PartnerId)
	}
	if req.Id != 0 {
		db = db.Where("id = ?", req.Id)
	}
	if req.SearchValue != "" {
		var sqlStr string
		for _, field := range req.SearchFields {
			sql := fmt.Sprintf("(%s ilike '%%%s%%')", field, req.SearchValue)
			if sqlStr == "" {
				sqlStr = sql
			} else {
				sqlStr = sqlStr + " or " + sql
			}
		}
		db = db.Where(sqlStr)
	}
	if len(req.Ids) != 0 {
		db = db.Where("id in (?)", req.Ids)
	}
	if req.OrderField != "" {
		orderType := req.OrderType
		if orderType == "" {
			orderType = "asc"
		}
		db = db.Order(req.OrderField + " " + orderType)
	} else {
		db = db.Order("created_at asc")
	}
	if req.Exclude.Id != 0 {
		db = db.Not("id = ?", req.Exclude.Id)
	}
	if len(req.Exclude.Ids) != 0 {
		db = db.Not("id in (?)", req.Exclude.Ids)
	}
	return db
}

func LogError(err error, when string) {
	if err != nil {
		log.Error(when + " : " + err.Error())
	}
}

