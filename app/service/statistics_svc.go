package service

import (
	"fmt"
	"js_statistics/app/models"
	"js_statistics/app/repositories"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/commom/drivers/database"
	"js_statistics/commom/tools"
	"js_statistics/constant"
	"js_statistics/exception"
	"strings"
	"sync"
	"time"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

var (
	stcServiceInstance StcService
	stcOnce            sync.Once
)

type stcServiceImpl struct {
	db        *gorm.DB
	repo      repositories.StcRepo
	blackRepo repositories.BlackIPRepo
	wipRepo   repositories.IPRepo
	jscRepo   repositories.JscRepo
	jsRepo    repositories.JsmRepo
	rmtRepo   repositories.RmRepo
	fakerRepo repositories.FakerRepo
	cdnRepo   repositories.CdnRepo
}

func GetStcService() StcService {
	stcOnce.Do(func() {
		stcServiceInstance = &stcServiceImpl{
			db:        database.GetDriver(),
			repo:      repositories.GetStcRepo(),
			blackRepo: repositories.GetBlackIPRepo(),
			wipRepo:   repositories.GetIPRepo(),
			jscRepo:   repositories.GetJscRepo(),
			jsRepo:    repositories.GetJsmRepo(),
			rmtRepo:   repositories.GetRmRepo(),
			fakerRepo: repositories.GetFakerRepo(),
			cdnRepo:   repositories.GetCdnRepo(),
		}
	})
	return stcServiceInstance
}

type StcService interface {
	ProcessJsRequest(ctx iris.Context)
}

func (ssi *stcServiceImpl) ProcessJsRequest(ctx iris.Context) {
	// ip := tools.GetRemoteAddr(ctx)
	ip := ctx.RemoteAddr()
	// ip := "121.41.38.13"
	isBlack, ex := ssi.blackRepo.IsExistByIP(ssi.db, ip)
	if ex != nil {
		if ex.Type() != response.ExceptionRecordNotFound {
			ctx.Application().Logger().Error(ex.Error())
			tools.ErrorResponse(ctx, ex)
			return
		}
	}
	sign := ctx.Params().Get("sign")
	js, ex := ssi.jsRepo.GetBySign(ssi.db, sign)
	if ex != nil {
		ctx.Application().Logger().Error(ex.Error())
		tools.ErrorResponse(ctx, ex)
		return
	}
	// 伪装内容
	faker, ex := ssi.GetFakerRedirectInfoByJsID(js.ID)
	if ex != nil {
		if ex.Type() != response.ExceptionRecordNotFound {
			ctx.Application().Logger().Error(ex.Error())
			tools.ErrorResponse(ctx, ex)
			return
		}
	}
	// 黑名单
	if isBlack {
		tools.BeyondRuleRedirect(ctx, faker, js)
		return
	}
	agent := strings.ToLower(ctx.Request().UserAgent())
	origin := ctx.GetHeader("Origin")
	// origin = "www.baidu.com/45/34"
	if len(origin) == 0 {
		origin = ctx.GetHeader("Referer")
	}
	// 白名单
	isWhite, ex := ssi.wipRepo.IsExistByIP(ssi.db, ip)
	if ex != nil {
		if ex.Type() != response.ExceptionRecordNotFound {
			ctx.Application().Logger().Error(ex.Error())
			tools.ErrorResponse(ctx, ex)
			return
		}
	}
	visitType, cookie := tools.GetVisitType(ctx)
	if isWhite {
		//返回输出代码
		ssi.GetRedirectInfo(ctx, js, faker, sign, agent, ip, cookie, origin, visitType)
		return
	}
	isCDN, ex := ssi.cdnRepo.IsExistByIP(ssi.db, ip)
	if ex != nil {
		if ex.Type() != response.ExceptionRecordNotFound {
			ctx.Application().Logger().Error(ex.Error())
			tools.ErrorResponse(ctx, ex)
			return
		}
	}
	if isCDN {
		// 过滤掉cdn的ip，直接返回
		tools.BeyondRuleRedirect(ctx, faker, js)
		return
	}
	// js判断条件
	pass := ssi.JSJudgeMent(ctx, js, faker, ip, sign, agent, origin)
	if !pass {
		return
	}
	ssi.GetRedirectInfo(ctx, js, faker, sign, agent, ip, cookie, origin, visitType)
}

func (ssi *stcServiceImpl) JSJudgeMent(ctx iris.Context, js *models.JsManage, faker *vo.FakerResp, ip, sign, agent, origin string) bool {
	// 默认屏蔽国外、香港、澳门、台湾IP
	if !ssi.IsValidLocation(ip) {
		// 伪装内容
		tools.BeyondRuleRedirect(ctx, faker, js)
		return false
	}
	// 国内屏蔽地区
	if len(js.ShieldArea) > 0 {
		loc, ex := tools.OriginIPLocation(ip)
		if ex != nil {
			ctx.Application().Logger().Error(ex.Error())
			tools.ErrorResponse(ctx, ex)
			return false
		}
		shieldAreas := strings.Split(js.ShieldArea, "-")
		region, ok := loc.City.Names["zh-CN"]
		if !ok {
			ctx.Application().Logger().Error("get ip location failed")
			tools.ErrorResponse(ctx, ex)
			return false
		}
		for i := range shieldAreas {
			if strings.Contains(shieldAreas[i], region) {
				//  伪装内容
				tools.BeyondRuleRedirect(ctx, faker, js)
				return false
			}
		}
	}

	if !js.Status {
		// 伪装内容
		tools.BeyondRuleRedirect(ctx, faker, js)
		return false
	}

	// 判断是pc端、移动端 是否合法
	clientType := tools.GetClintType(agent)
	if !tools.IsInRuleClient(int64(clientType), js.ClientType) {
		tools.BeyondRuleRedirect(ctx, faker, js)
		return false
	}

	// 封禁小时 和 次数，为0不跳转
	if js.RedirectCount == 0 {
		// 伪装内容
		tools.BeyondRuleRedirect(ctx, faker, js)
		return false
	}
	// 规定时间内，跳转，次数减一，为0不跳转
	// 0: 无限制
	if js.ReleaseTime > 0 {
		if time.Since(js.UpdateAt) > time.Duration(js.ReleaseTime*int(time.Hour)) {
			// 伪装内容
			tools.BeyondRuleRedirect(ctx, faker, js)
			return false
		}
		if ex := ssi.jsRepo.DecreaseRedirectCount(ssi.db, js.ID); ex != nil {
			ctx.Application().Logger().Error("get ip location failed")
			tools.ErrorResponse(ctx, ex)
			return false
		}
	}

	switch js.FromMode {
	case constant.FromTypeNone:
		fmt.Println("来源无")
	case constant.FromTypeKey:
		// 判断origin是否匹配
		keyWord := strings.Split(js.KeyWord, ",")
		isInKey := false
		for i := range keyWord {
			if strings.Contains(origin, keyWord[i]) {
				isInKey = true
				break
			}
		}
		if !isInKey {
			// 伪装内容
			tools.BeyondRuleRedirect(ctx, faker, js)
			return false
		}
	case constant.FromTypeEngine:
		isExist, engineType := tools.GetEngineType(agent)
		if !isExist {
			tools.BeyondRuleRedirect(ctx, faker, js)
			return false
		}
		isInRule := false
		for i := range js.SearchEngines {
			if js.SearchEngines[i] == engineType {
				isInRule = true
				break
			}
		}
		if !isInRule {
			tools.BeyondRuleRedirect(ctx, faker, js)
			return false
		}
	}
	return true
}

func (ssi *stcServiceImpl) GetRedirectInfo(ctx iris.Context, js *models.JsManage, faker *vo.FakerResp, sign, agent,
	ip, cookie, origin string, visitType int) {
	redirectInfo, ex := ssi.rmtRepo.GetUsefulByCategoryID(ssi.db, js.CategoryID)
	if ex != nil {
		if ex.Type() == response.ExceptionRecordNotFound {
			ctx.Application().Logger().Error(ex.Error())
			tools.DefaultBlackCode(ctx)
			return
		}
		ctx.Application().Logger().Error(ex.Error())
		tools.ErrorResponse(ctx, ex)
		return
	}
	// 跳转时间区间是否合理
	if *redirectInfo.ON != "" && *redirectInfo.OFF != "" {
		if res, err := tools.IsInRedirectOnOff(*redirectInfo.ON, *redirectInfo.OFF); err != nil {
			ctx.Application().Logger().Error(err.Error())
			tools.ErrorResponse(ctx, ex)
			return
		} else if !res {
			// 跳转管理时间区间不符合
			tools.BeyondRuleRedirect(ctx, faker, js)
			return
		}
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
		ctx.Application().Logger().Error(ex.Error())
		tools.ErrorResponse(ctx, ex)
		return
	}
	// 记录入库
	tx := ssi.db.Begin()
	defer tx.Rollback()
	if tx.Error != nil {
		ctx.Application().Logger().Error("get ip location failed")
		tools.ErrorResponse(ctx, exception.Wrap(response.ExceptionDatabase, tx.Error))
		return
	}
	if visitType == constant.IPVisit {
		if ex := ssi.repo.CreateIPStatistics(tx, &models.IPStatistics{
			IP:         ip,
			JsID:       js.ID,
			CategoryID: js.CategoryID,
			PrimaryID:  jp.PrimaryID,
			VisitTime:  time.Now(),
		}); ex != nil {
			ctx.Application().Logger().Error("get ip location failed")
			tools.ErrorResponse(ctx, ex)
			return
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
			ctx.Application().Logger().Error("get ip location failed")
			tools.ErrorResponse(ctx, ex)
			return
		}
	}
	ipLocation, ex := tools.OriginIPLocation(ip)
	if ex != nil {
		ctx.Application().Logger().Error(ex.Error())
		tools.ErrorResponse(ctx, ex)
		return
	}
	region, ok := ipLocation.Subdivisions[0].Names["zh-CN"]
	if !ok {
		region = ""
	}
	if ex := ssi.repo.CreateIPRecode(tx, &models.IPRecode{
		IP:         ip,
		CategoryID: js.CategoryID,
		PrimaryID:  jp.PrimaryID,
		JsID:       js.ID,
		FromURL:    origin,
		ToURL:      redirectURL,
		RegionCode: "0",
		Region:     region,
		VisitType:  visitType,
		VisitTime:  time.Now(),
	}); ex != nil {
		ctx.Application().Logger().Error(ex.Error())
		tools.ErrorResponse(ctx, ex)
		return
	}
	if res := tx.Commit(); res.Error != nil {
		ctx.Application().Logger().Error(res.Error)
		tools.ErrorResponse(ctx, exception.Wrap(response.ExceptionDatabase, res.Error))
		return
	}

	if js.WaitTime > 0 {
		time.Sleep(time.Duration(js.WaitTime))
	}
	// 判断跳转方式
	switch js.RedirectMode {
	case constant.Direct:
		if js.RedirectCode == 0 {
			tools.DirectWindowsRedirect(ctx, redirectURL)
		} else {
			tools.DirectTopRedirect(ctx, redirectURL)
		}
	case constant.Nested:
		tools.NestedRedirect(ctx, redirectURL)
	case constant.Screen:
		tools.ScreenRedirect(ctx, redirectURL)
	default:
		// id 为动态参数
		tools.HrefRedirect(ctx, redirectURL+"/"+strings.ReplaceAll(js.HrefID, ",", "/"))
	}
}

func (ssi *stcServiceImpl) IsValidLocation(ip string) bool {
	location, ex := tools.OriginIPLocation(ip)
	if ex != nil {
		return false
	}
	return location.Country.IsoCode == constant.CN_ISO_CODE
}

func (ssi *stcServiceImpl) GetFakerRedirectInfoByJsID(jsID int64) (*vo.FakerResp, exception.Exception) {
	faker, ex := ssi.fakerRepo.GetByJsID(ssi.db, jsID)
	if ex != nil {
		return nil, ex
	}
	return vo.NewFakerResponse(faker), nil
}

// func(ssi *stcServiceImpl)
