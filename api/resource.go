package api

import (
	"errors"
	"fmt"
	app "github.com/zhanghup/go-framework"
	"github.com/zhanghup/go-framework/ctx/cfg"
	"github.com/zhanghup/go-framework/pkg/gin"
	"github.com/zhanghup/go-framework/pkg/xorm"
	"github.com/zhanghup/go-framework/tools"
	"io/ioutil"
	"net/http"
	"strings"
)

func Resource(public, auth *gin.RouterGroup, e *xorm.Engine) {
	auth.Action("/upload", func(c *gin.Context, p interface{}) (obj interface{}, err error) {
		prefix := cfg.GetCfg().Gin.Prefix
		//得到上传的文件
		file, header, err := c.Request.FormFile("zfile") //image这个是uplaodify参数定义中的   'fileObjName':'image'
		if err != nil {
			return nil, errors.New("Bad request")
		}
		datas, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		ff := new(app.Resource)
		md := tools.MD5(datas)
		ok, err := e.Table(ff).Where("md5 = ?", md).Cols("id", "type").Get(ff)
		if err != nil {
			return nil, err
		}
		if ok {
			return prefix + "/upload/" + *ff.Id + ff.Type, nil
		}

		str := header.Filename
		//文件的名称

		ff.Id = tools.ObjectString()
		ff.Status = tools.PtrOfInt(1)
		ff.Name = header.Filename
		ff.ContentType = c.GetHeader("content-type")
		ff.Size = header.Size
		ff.Type = str[strings.LastIndex(str, "."):]

		ff.MD5 = md
		ff.Datas = datas
		_, err = e.Insert(ff)

		return prefix + "/upload/" + *ff.Id + ff.Type, nil
	})

	public.GET("/upload/:id", func(c *gin.Context) {
		idstr := c.Param("id")
		id := idstr[:strings.Index(idstr, ".")]
		ff := new(app.Resource)
		ok, err := e.Table(ff).Where("id = ?", id).Get(ff)
		if err != nil {
			return
		}
		if ok {
			c.Writer.Header().Set("Content-Length", fmt.Sprintf("%v", ff.Size))
			c.Writer.Write(ff.Datas)
		}
	})

	public.StaticFS("/static", http.Dir("resource/static/"))
	public.StaticFile("/favicon.ico", "resource/favicon.ico")
}
