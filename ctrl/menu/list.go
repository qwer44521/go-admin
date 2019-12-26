package menu

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-admin/conf"
	"go-admin/models"
	"go-admin/modules/request"
	"go-admin/modules/response"
	"io/ioutil"
	"strconv"
	"time"
)

type Role struct {
	Key         string        `form:"key" json:"key"`
	Name        string        `form:"name" json:"name"`
	Description string        `form:"description" json:"description"`
	Routes      []interface{} `form:"routes" json:"routes"`
}

func List(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get(conf.Cfg.Token)
	if v == nil {
		response.ShowError(c, "fail")
		return
	}
	uid := session.Get(v)
	user := models.SystemUser{Id: uid.(int)}
	has := user.GetRow()
	if !has {
		response.ShowError(c, "fail")
		return
	}

	menu := models.SystemMenu{}
	if user.Nickname == "admin" {
		menuArr, err := menu.GetAll()
		if err != nil {
			response.ShowError(c, "fail")
			return
		}
		jsonArr := tree(menuArr)
		response.ShowData(c, jsonArr)
		return
	} else {
		menuArr := menu.GetRouteByUid(uid)
		jsonArr := tree(menuArr)
		response.ShowData(c, jsonArr)
		return
	}
}
func tree(menuArr []models.SystemMenu) ([]interface{}) {
	role := models.SystemRole{}
	mrArr := role.GetRowMenu()
	var menuMap = make(map[int][]models.SystemMenu, 0)
	for _, value := range menuArr {
		menuMap[value.Pid] = append(menuMap[value.Pid], value)
	}
	var jsonArr []interface{}

	mainMenu, ok := menuMap[0]
	if !ok {
		return nil
	}
	for _, value := range mainMenu {
		var item = make(map[string]interface{})
		item["path"] = value.Path
		item["component"] = value.Component
		if value.Redirect != "" {
			item["redirect"] = value.Redirect
		}
		if value.Alwaysshow == 1 {
			item["alwaysShow"] = true
		}
		if value.Hidden == 1 {
			item["hidden"] = true
		}
		var meta = make(map[string]interface{})
		_, ok := mrArr[value.Id]
		if ok {
			meta["roles"] = mrArr[value.Id]
		}
		if value.MetaTitle != "" {
			meta["title"] = value.MetaTitle
		}
		if value.MetaIcon != "" {
			meta["icon"] = value.MetaIcon
		}
		if value.MetaAffix == 1 {
			meta["affix"] = true
		}
		if value.MetaNocache == 1 {
			meta["noCache"] = true
		}
		if len(meta) > 0 {
			item["meta"] = meta
		}
		if _, ok := menuMap[value.Id]; ok {
			item["children"] = treeChilden(menuMap[value.Id], mrArr)
		}
		jsonArr = append(jsonArr, item)
	}
	return jsonArr

}
func treeChilden(menuArr []models.SystemMenu, mrArr map[int][]string) []interface{} {
	var jsonArr []interface{}
	for _, value := range menuArr {
		var item = make(map[string]interface{})
		item["path"] = value.Path
		item["component"] = value.Component
		if value.Redirect != "" {
			item["redirect"] = value.Redirect
		}
		if value.Alwaysshow == 1 {
			item["alwaysShow"] = true
		}
		if value.Hidden == 1 {
			item["hidden"] = true
		}
		var meta = make(map[string]interface{})
		_, ok := mrArr[value.Id]
		if ok {
			meta["roles"] = mrArr[value.Id]
		}
		if value.MetaTitle != "" {
			meta["title"] = value.MetaTitle
		}
		if value.MetaIcon != "" {
			meta["icon"] = value.MetaIcon
		}
		if value.MetaAffix == 1 {
			meta["affix"] = true
		}
		if value.MetaNocache == 1 {
			meta["noCache"] = true
		}
		if len(meta) > 0 {
			item["meta"] = meta
		}
		jsonArr = append(jsonArr, item)
	}
	return jsonArr
}
func treeMenuChilden(menuArr []models.SystemMenu, mrArr map[int][]string) []interface{} {
	var jsonArr []interface{}
	for _, value := range menuArr {
		var item = make(map[string]interface{})
		item["path"] = value.Path
		item["component"] = value.Component
		if value.Redirect != "" {
			item["redirect"] = value.Redirect
		}
		if value.Alwaysshow == 1 {
			item["alwaysShow"] = true
		}
		if value.Hidden == 1 {
			item["hidden"] = true
		} else {
			item["hidden"] = false
		}
		var meta = make(map[string]interface{})
		_, ok := mrArr[value.Id]
		if ok {
			meta["roles"] = mrArr[value.Id]
		}
		if value.MetaTitle != "" {
			meta["title"] = value.MetaTitle
		}
		if value.MetaIcon != "" {
			meta["icon"] = value.MetaIcon
		}
		if value.MetaAffix == 1 {
			meta["affix"] = true
		}
		if value.MetaNocache == 1 {
			meta["noCache"] = true
		}
		if len(meta) > 0 {
			item["meta"] = meta
		}
		item["pid"] = value.Pid
		item["id"] = value.Id
		item["url"] = value.Url
		item["name"] = value.Name

		jsonArr = append(jsonArr, item)
	}
	return jsonArr
}
func treeMenu(menuArr []models.SystemMenu) ([]interface{}) {
	role := models.SystemRole{}
	mrArr := role.GetRowMenu()
	var menuMap = make(map[int][]models.SystemMenu, 0)
	for _, value := range menuArr {
		menuMap[value.Pid] = append(menuMap[value.Pid], value)
	}
	var jsonArr []interface{}
	mainMenu, ok := menuMap[0]
	if !ok {
		return nil
	}
	for _, value := range mainMenu {
		var item = make(map[string]interface{})
		item["path"] = value.Path
		item["component"] = value.Component
		if value.Redirect != "" {
			item["redirect"] = value.Redirect
		}
		if value.Alwaysshow == 1 {
			item["alwaysShow"] = true
		}
		if value.Hidden == 1 {
			item["hidden"] = true
		} else {
			item["hidden"] = false
		}

		var meta = make(map[string]interface{})
		_, ok := mrArr[value.Id]
		if ok {
			meta["roles"] = mrArr[value.Id]
		}
		if value.MetaTitle != "" {
			meta["title"] = value.MetaTitle
		}
		if value.MetaIcon != "" {
			meta["icon"] = value.MetaIcon
		}
		if value.MetaAffix == 1 {
			meta["affix"] = true
		}
		if value.MetaNocache == 1 {
			meta["noCache"] = true
		}
		if value.Status == 1 {
			meta["status"] = true
		}

		if len(meta) > 0 {
			item["meta"] = meta
		}
		if _, ok := menuMap[value.Id]; ok {
			item["children"] = treeMenuChilden(menuMap[value.Id], mrArr)
		}
		item["pid"] = value.Pid
		item["id"] = value.Id
		item["url"] = value.Url
		item["name"] = value.Name
		jsonArr = append(jsonArr, item)
	}
	return jsonArr

}
func Roles(c *gin.Context) {
	model := models.SystemRole{}
	menu := models.SystemMenu{}
	roleArr := model.GetAll()
	var roleMenu []Role
	for _, value := range roleArr {
		r := Role{}
		r.Key = value.Name
		r.Name = value.Name
		r.Description = value.Description
		menuArr := menu.GetRouteByRole(value.Id)
		r.Routes = tree(menuArr)
		roleMenu = append(roleMenu, r)
	}
	response.ShowData(c, roleMenu)
	return
}

func Dashboard(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get(conf.Cfg.Token)
	if v == nil {
		response.ShowError(c, "fail")
		return
	}
	uid := session.Get(v)
	user := models.SystemUser{Id: uid.(int)}
	has := user.GetRow()
	if !has {
		response.ShowError(c, "fail")
		return
	}
	menu := models.SystemMenu{}
	if user.Nickname == "admin" {
		menuArr, err := menu.GetAll()
		if err != nil {
			response.ShowError(c, "fail")
			return
		}
		jsonArr := treeMenu(menuArr)
		response.ShowData(c, jsonArr)
		return
	} else {
		menuArr := menu.GetRouteByUid(uid)
		jsonArr := treeMenu(menuArr)
		response.ShowData(c, jsonArr)
		return
	}
	//
	//roleMenu:="{\"menuList\": [{ 		\"children\": [{ 			\"menu_type\": \"M\", 			\"children\": [{ 				\"menu_type\": \"C\", 				\"parent_id\": 73, 				\"menu_name\": \"人员通讯录\", 				\"icon\": null, 				\"order_num\": 1, 				\"menu_id\": 74, 				\"url\": \"/system/book/person\" 			}], 			\"parent_id\": 1, 			\"menu_name\": \"通讯录管理\", 			\"icon\": \"fafa-address-book-o\", 			\"perms\": null, 			\"order_num\": 1, 			\"menu_id\": 73, 			\"url\": \"#\" 		}], 		\"parent_id\": 0, 		\"menu_name\": \"系统管理\", 		\"icon\": \"fafa-adjust\", 		\"perms\": null, 		\"order_num\": 2, 		\"menu_id\": 1, 		\"url\": \"#\" 	}], 	\"user\": { 		\"login_name\": \"admin1\", 		\"user_id\": 1, 		\"user_name\": \"管理员\", 		\"dept_id\": 1 	} }"
	//var data map[string]interface{}
	//err := json.Unmarshal([]byte(roleMenu), &data)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 20000,
	//	"data":  data,
	//})
	//return
}

func Index(c *gin.Context) {
	//session := sessions.Default(c)
	//v := session.Get(conf.Cfg.Token)
	//if v == nil {
	//	response.ShowError(c, "fail")
	//	return
	//}
	menu := models.SystemMenu{}
	menuArr, err := menu.GetAll()
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	var menuMap = make(map[int][]models.SystemMenu, 0)
	for _, value := range menuArr {
		menuMap[value.Pid] = append(menuMap[value.Pid], value)
	}

	//response.ShowData(c, menuMap)
	//return
	var menuNewArr []models.SystemMenu
	menuNewArr =TreeNode(menuMap,0)
	response.ShowData(c, menuNewArr)
	return
}
func TreeNode(menuMap map[int][]models.SystemMenu,pid int) []models.SystemMenu {
	var menuNewArr []models.SystemMenu
	if _,ok:=menuMap[pid];ok{
		for _, v := range menuMap[pid] {
			menuNewArr=append(menuNewArr,v)
			menuNewArr=append(menuNewArr,TreeNode(menuMap,v.Id)...)
		}
	}else{
		return menuNewArr
	}
	return menuNewArr
}
func Add(c *gin.Context) {
	jsonstr, _ := ioutil.ReadAll(c.Request.Body)
	var data map[string]interface{}
	err := json.Unmarshal(jsonstr, &data)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	//fmt.Println(data)
	if _, ok := data["name"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["path"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["component"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["url"]; !ok {
		response.ShowError(c, "fail")
		return
	}

	menu := models.SystemMenu{}
	menu.Path = data["path"].(string)
	if menu.Path=="" {
		response.ShowError(c, "fail")
		return
	}
	has:=menu.GetRow()
	if has && menu.Path!="#" {
		response.ShowError(c, "path不可重复")
		return
	}
	menu.Name=data["name"].(string)
	menu.Path=data["path"].(string)
	menu.MetaTitle=data["name"].(string)
	menu.Component=data["component"].(string)
	menu.Url=data["url"].(string)
	menu.Redirect=data["redirect"].(string)
	menu.MetaIcon=data["meta_icon"].(string)
	if data["alwaysshow"].(bool){
		menu.Alwaysshow=1
	}
	if data["hidden"].(bool){
		menu.Hidden=1
	}
	if data["status"].(bool){
		menu.Status=1
	}
	menu.Ctime=time.Now()

	menu.Sort,_=strconv.Atoi(data["sort"].(string))
	menu.Pid = int(data["pid"].(float64))
	if menu.Pid==0 {
		menu.Level=menu.Pid
	}else{
		pidMenuModel := models.SystemMenu{Id: menu.Pid}
		_=pidMenuModel.GetRow()
		menu.Level=pidMenuModel.Level+1
	}
	_,err=menu.Add()
	if err!=nil {
		response.ShowError(c,"fail")
		return
	}
	response.ShowData(c,menu)
	return
}
func Update(c *gin.Context) {
	data,err:=request.GetJson(c)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["id"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	menu:=models.SystemMenu{}
	menu.Id=int(data["id"].(float64))
	has:=menu.GetRow()
	if !has {
		response.ShowError(c, "要修改数据不存在")
		return
	}
	if _, ok := data["name"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["path"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["component"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["url"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	menuModel:=models.SystemMenu{}
	menuModel.Path = data["path"].(string)
	if menuModel.Path=="" {
		response.ShowError(c, "fail")
		return
	}
	has=menuModel.GetRow()
	if has && menuModel.Path!="#" && menuModel.Id!=menu.Id {
		response.ShowError(c, "path不可重复")
		return
	}
	menu.Name=data["name"].(string)
	menu.Path=data["path"].(string)
	menu.MetaTitle=data["name"].(string)
	menu.Component=data["component"].(string)
	menu.Url=data["url"].(string)
	menu.Redirect=data["redirect"].(string)
	menu.MetaIcon=data["meta_icon"].(string)
	if data["alwaysshow"].(bool){
		menu.Alwaysshow=1
	}
	if data["hidden"].(bool){
		menu.Hidden=1
	}
	if data["status"].(bool){
		menu.Status=1
	}
	menu.Sort,_=strconv.Atoi(data["sort"].(string))
	menu.Pid = int(data["pid"].(float64))
	err=menu.Update()
	if err!=nil {
		response.ShowError(c,"fail")
		return
	}
	response.ShowData(c,"success")
	return
}
func Delete(c *gin.Context){
	str, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	id :=string(str)
	if id=="" {
		response.ShowError(c, "fail")
		return
	}
	menu :=models.SystemMenu{}
	menu.Id,_=strconv.Atoi(id)
	fmt.Println(menu)
	err=menu.Delete()
	if err!=nil {
		response.ShowError(c,"fail")
		return
	}
	response.ShowData(c,"success")
	return
}