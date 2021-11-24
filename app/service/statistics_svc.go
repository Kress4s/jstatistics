package service

import (
	"fmt"
	"js_statistics/app/models"
	"js_statistics/app/repositories"
	"js_statistics/app/response"
	"js_statistics/commom/drivers/database"
	"js_statistics/commom/tools"
	"js_statistics/constant"
	"js_statistics/exception"
	"sync"
	"time"

	"gorm.io/gorm"
)

var (
	stcServiceInstance StcService
	stcOnce            sync.Once
)

type stcServiceImpl struct {
	db      *gorm.DB
	repo    repositories.StcRepo
	jscRepo repositories.JscRepo
	jsRepo  repositories.JsmRepo
	rmtRepo repositories.RmRepo
}

func GetStcService() StcService {
	stcOnce.Do(func() {
		stcServiceInstance = &stcServiceImpl{
			db:      database.GetDriver(),
			repo:    repositories.GetStcRepo(),
			jscRepo: repositories.GetJscRepo(),
			jsRepo:  repositories.GetJsmRepo(),
			rmtRepo: repositories.GetRmRepo(),
		}
	})
	return stcServiceInstance
}

type StcService interface {
	Process(sign, agent, origin, ip, cookie string, visitType int) (string, exception.Exception)
}

func (ssi *stcServiceImpl) Process(sign, agent, origin, ip, cookie string, visitType int) (string, exception.Exception) {
	/*
		8. 跳转管理的地址(保证开关是开启可用)
		6. 脚本的封禁小时(记录保存时间) 替代为开关
		4. 判断是pc端、移动端(安卓、ios)
		7. 来源判断(无，关键词、搜索引擎 判断) // TODO
		9. 跳转方式 --> 输出跳转代码
		10. 条件不满足 --> 伪装内容设置 --> 空白页
	*/
	js, ex := ssi.jsRepo.GetBySign(ssi.db, sign)
	if ex != nil {
		// TODO 日志记录
		return "", ex
	}
	// 脚本开启状态 脚本的封禁小时(记录保存时间)
	if !js.Status {
		// TODO 日志记录
		return "", ex
	}
	// 判断是pc端、移动端 是否合法
	clientType := tools.GetClintType(agent)
	if !tools.IsInRuleClient(int64(clientType), js.ClientType) {
		// TODO 日志记录
		return "", ex
	}

	// TODO 跳转次数，目前待定

	// TODO 封禁小时，目前待定 目前是根据开关

	// TODO 来源
	switch js.FromMode {
	case constant.FromTypeNone:
		fmt.Println("来源无")
	case constant.FromTypeKey:
		fmt.Println("关键词")
	case constant.FromTypeEngine:
		isExist, engineType := tools.GetEngineType(agent)
		if !isExist {
			return "", ex
		}
		isInRule := false
		for i := range js.SearchEngines {
			if js.SearchEngines[i] == engineType {
				isInRule = true
				break
			}
		}
		if !isInRule {
			return "", ex
		}
	}

	// TODO 跳转代码 TOP/Windows 未定

	// 跳转信息
	redirectInfo, ex := ssi.rmtRepo.Get(ssi.db, js.CategoryID)
	if ex != nil {
		// TODO 日志记录
		return "", ex
	}
	// 是否开启
	if !redirectInfo.Status {
		return "", ex
	}
	// 跳转时间区间是否合理
	now := time.Now()
	if !(now.Before(redirectInfo.OFF) && now.After(redirectInfo.ON)) {
		return "", ex
	}

	var redirectURL string
	deviceType := tools.GetDeviceType(agent)
	switch {
	case deviceType == constant.AndroidRedirectType:
		redirectURL = redirectInfo.Android
	case deviceType == constant.IOSRedirectType:
		redirectURL = redirectInfo.IOS
	default:
		redirectURL = redirectInfo.PC
	}

	jp, ex := ssi.jscRepo.Get(ssi.db, js.CategoryID)
	if ex != nil {
		return "", ex
	}

	// 记录入库
	tx := ssi.db.Begin()
	defer tx.Rollback()
	if tx.Error != nil {
		return "", exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	if visitType == constant.IPVisit {
		if ex := ssi.repo.CreateIPStatistics(tx, &models.IPStatistics{
			IP:         ip,
			JsID:       js.ID,
			CategoryID: js.CategoryID,
			PrimaryID:  jp.PrimaryID,
			VisitTime:  time.Now(),
		}); ex != nil {
			return "", ex
		}
	} else {
		if ex := ssi.repo.CreateUVStatistics(tx, &models.UVStatistics{
			IP:         ip,
			JsID:       js.ID,
			CategoryID: js.CategoryID,
			PrimaryID:  jp.PrimaryID,
			Cookie:     cookie,
			VisitTime:  time.Now(),
		}); ex != nil {
			return "", ex
		}
	}
	ipLocation, ex := tools.OriginIPLocation(ip)
	if ex != nil {
		return "", ex
	}
	region, ok := ipLocation.Subdivisions[0].Names["zh-CN"]
	if !ok {
		region = ""
	}
	if ex := ssi.repo.CreateIPRecode(tx, &models.IPRecode{
		IP:         ip,
		CategoryID: js.CategoryID,
		PrimaryID:  jp.PrimaryID,
		FromURL:    origin,
		ToURL:      redirectURL,
		RegionCode: "0",
		Region:     region,
		VisitType:  visitType,
		VisitTime:  time.Now(),
	}); ex != nil {
		return "", ex
	}
	if res := tx.Commit(); res.Error != nil {
		return "", exception.Wrap(response.ExceptionDatabase, tx.Error)
	}

	if js.WaitTime > 0 {
		time.Sleep(time.Duration(js.WaitTime))
	}
	// 判断跳转方式
	// TODO 返回内容未取伪装内容
	switch js.RedirectMode {
	case constant.Direct:
		return fmt.Sprintf(constant.RedirectPage, redirectURL), nil
	case constant.Nested:
		return fmt.Sprintf(constant.NestingRedirect, redirectURL), nil
	case constant.Screen:
		return fmt.Sprintf(constant.ScreenRedirect, redirectURL), nil
	default:
		return fmt.Sprintf(constant.HrefRedirect, redirectURL), nil
	}
}
