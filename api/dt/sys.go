package dt

import (
	app "github.com/zhanghup/go-framework"
	"github.com/zhanghup/go-framework/pkg/xorm"
)

type SysService struct {
	DB   *xorm.Engine
	Base *BaseService
}

func (this *SysService) DictItems(code string) ([]app.DictItem, error) {
	result := make([]app.DictItem, 0)
	err := this.DB.Table(result).Where("code = ?", code).OrderBy("weight").Find(&result)
	return result, err
}
func (this *SysService) DictList() ([]app.Dict, error) {
	tdict := this.DB.TableInfo(new(app.Dict)).Name
	tdictItem := this.DB.TableInfo(new(app.DictItem)).Name

	data := make([]struct {
		*app.Dict     `xorm:"extends"`
		*app.DictItem `xorm:"extends"`
	}, 0)
	err := this.DB.Table(tdict).Alias("d").Join("left", []string{tdictItem, "di"}, "di.code = d.id").OrderBy("d.weight,di.weight").Find(&data)
	if err != nil {
		return nil, err
	}

	result := make([]app.Dict, 0)
	var dict *app.Dict
	for i, obj := range data {
		if dict != nil && *dict.Id != *obj.Dict.Id {
			result = append(result, *dict)
		}

		if dict != nil {
			if obj.DictItem.Id != nil {
				dict.Values = append(dict.Values, *data[i].DictItem)
			}
		}

		if dict == nil || *dict.Id != *obj.Dict.Id {
			dict = data[i].Dict
			if obj.DictItem.Id != nil {
				dict.Values = []app.DictItem{*data[i].DictItem}
			}
		}
	}
	result = append(result, *dict)
	return result, nil
}

var sysService *SysService

func NewSysService(e *xorm.Engine) *SysService {
	if sysService != nil {
		return sysService
	}
	sysService = &SysService{DB: e, Base: NewBaseService(e)}
	return sysService
}
