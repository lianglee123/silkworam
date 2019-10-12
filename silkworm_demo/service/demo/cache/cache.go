package cache

import (
	"silkwormDemo/model"
	_redis "silkwormDemo/internal/redis"
	"github.com/go-redis/redis"
	"time"
	"fmt"
	"encoding/json"
	"log"
)

type DemoCache interface {
	SetDemo(demo *model.Demo, duration time.Duration) error
	GetDemo(id uint64) *model.Demo
	ClearDemo(id uint64) error
}

type DemoCacheImpl struct {
	keyPrefix string
}

func (d *DemoCacheImpl) client() *redis.Client {
	return _redis.GetClient()
}
func (d *DemoCacheImpl) SetDemo(demo *model.Demo, duration time.Duration) error {
	if duration < time.Second {
		return nil
	}
	key := fmt.Sprintf("%s:demo:%s", d.keyPrefix, demo.Id)
	data, _ := json.Marshal(demo)
	if _, err := d.client().Set(key, data, duration).Result(); err != nil {
		return err
	}
	return nil
}

func (d *DemoCacheImpl) GetDemo(id uint64) *model.Demo {
	key := fmt.Sprintf("%s:demo:%s", d.keyPrefix, id)
	cmd := d.client().Get(key)
	if cmd.Err() != nil {
		if cmd.Err() != redis.Nil {
			log.Fatal("get client status from cache err: ", cmd.Err())
		}
		return nil
	}
	var demo = new(model.Demo)
	err := json.Unmarshal([]byte(cmd.Val()), demo)
	if err != nil {
		log.Fatal("unmarshal token fail: ", err)
		return nil
	}
	return nil
}

func (d *DemoCacheImpl) ClearDemo(id uint64) error {
	key := fmt.Sprintf("%s:demo:%s", d.keyPrefix, id)
	intCmd := d.client().Del(key)
	if intCmd.Err() != nil && intCmd.Err() != nil {
		return intCmd.Err()
	}
	return nil
}
