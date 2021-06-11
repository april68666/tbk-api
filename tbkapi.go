package tbkapi

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"


	"github.com/guonaihong/gout"

)

const gw = "https://eco.taobao.com/router/rest"

type Api struct {
	appKey    string
	appSecret string
	session   string
}

func NewTbkApi(appKey, appSecret, session string) *Api {
	return &Api{
		appKey:    appKey,
		appSecret: appSecret,
		session:   session,
	}
}

// TbkItemInfoGet 淘宝客-公用-淘宝客商品详情查询(简版)
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.aec0669aumR0CN&source=search&docId=24518&docType=2
func (a *Api) TbkItemInfoGet(param ItemInfoGetParams) (*TbkItemInfoGetResp, error) {
	p, err := a.parameterCombination("taobao.tbk.item.info.get", param, false)
	if err != nil {
		return nil, err
	}
	resp := TbkItemInfoGetResp{}
	err = gout.POST(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PCouponConvert 淘宝客-推广者-单品券高效转链
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.aec0669aumR0CN&source=search&docId=29289&docType=2
func (a *Api) PCouponConvert(param CouponConvertParams) (*CouponConvertResp, error) {
	p, err := a.parameterCombination("taobao.tbk.coupon.convert", param, false)
	if err != nil {
		return nil, err
	}
	resp := CouponConvertResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SCouponConvert 淘宝客-服务商-单品券高效转链
// https://open.taobao.com/api.htm?docId=28625&docType=2&scopeId=12403
func (a *Api) SCouponConvert(param CouponConvertParams) (*CouponConvertResp, error) {
	p, err := a.parameterCombination("taobao.tbk.privilege.get", param, true)
	if err != nil {
		return nil, err
	}
	resp := CouponConvertResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// TextTPwdCreate 淘宝客-公用-淘口令生成
// https://open.taobao.com/api.htm?docId=31127&docType=2&scopeId=11655
func (a *Api) TextTPwdCreate(param TextTPwdCreateParams) (*TextTPwdCreateResp, error) {
	p, err := a.parameterCombination("taobao.tbk.tpwd.create", param, false)
	if err != nil {
		return nil, err
	}
	resp := TextTPwdCreateResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PTPwdConvert 淘宝客-推广者-淘口令解析&转链
// https://open.taobao.com/api.htm?docId=32932&docType=2&scopeId=16290
func (a *Api) PTPwdConvert(param TPwdConvertParams) (*TPwdConvertResp, error) {
	p, err := a.parameterCombination("taobao.tbk.tpwd.convert", param, false)
	if err != nil {
		return nil, err
	}
	resp := TPwdConvertResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// STPwdConvert 淘宝客-服务商-淘口令解析&转链
// https://open.taobao.com/api.htm?docId=43873&docType=2&scopeId=16401
func (a *Api) STPwdConvert(param TPwdConvertParams) (*TPwdConvertResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.tpwd.convert", param, true)
	if err != nil {
		return nil, err
	}
	resp := TPwdConvertResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SPunishOrderGet 淘宝客-服务商-处罚订单查询
// https://open.taobao.com/api.htm?docId=41942&docType=2&scopeId=15738
func (a *Api) SPunishOrderGet(param PunishOrderGetParams) (*PunishOrderGetResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.punish.order.get", param, true)
	if err != nil {
		return nil, err
	}
	resp := PunishOrderGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PPunishOrderGet 淘宝客-推广者-处罚订单查询
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.5d97669aNPczDi&source=search&docId=42050&docType=2
func (a *Api) PPunishOrderGet(param PunishOrderGetParams) (*PunishOrderGetResp, error) {
	p, err := a.parameterCombination("taobao.tbk.dg.punish.order.get", param, false)
	if err != nil {
		return nil, err
	}
	resp := PunishOrderGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SRelationRefund 淘宝客-服务商-维权退款订单查询
// https://open.taobao.com/api.htm?docId=43874&docType=2&scopeId=16322
func (a *Api) SRelationRefund(param RelationRefundParams) (*RelationRefundResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.relation.refund", param, true)
	if err != nil {
		return nil, err
	}
	resp := RelationRefundResp{}
	err = gout.GET(gw).Debug(true).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PRelationRefund 淘宝客-推广者-维权退款订单查询
// https://open.taobao.com/api.htm?docId=40121&docType=2&scopeId=16175
func (a *Api) PRelationRefund(param RelationRefundParams) (*RelationRefundResp, error) {
	p, err := a.parameterCombination("taobao.tbk.relation.refund", param, false)
	if err != nil {
		return nil, err
	}
	resp := RelationRefundResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// POrderDetailsGet 淘宝客-推广者-所有订单查询
// https://open.taobao.com/api.htm?docId=43328&docType=2&scopeId=16175
func (a *Api) POrderDetailsGet(param OrderDetailsGetParams) (*OrderDetailsGetResp, error) {
	p, err := a.parameterCombination("taobao.tbk.order.details.get", param, false)
	if err != nil {
		return nil, err
	}
	resp := OrderDetailsGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SOrderDetailsGet 淘宝客-服务商-所有订单查询
// https://open.taobao.com/api.htm?docId=43755&docType=2&scopeId=16322
func (a *Api) SOrderDetailsGet(param OrderDetailsGetParams) (*OrderDetailsGetResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.order.details.get", param, true)
	if err != nil {
		return nil, err
	}
	resp := OrderDetailsGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PVegasTljCreate 淘宝客-推广者-淘礼金创建
// https://open.taobao.com/api.htm?docId=40173&docType=2&scopeId=15029
func (a *Api) PVegasTljCreate(param VegasTljCreateParams) (*VegasTljCreateResp, error) {
	p, err := a.parameterCombination("taobao.tbk.dg.vegas.tlj.create", param, false)
	if err != nil {
		return nil, err
	}
	resp := VegasTljCreateResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SVegasTljCreate 淘宝客-服务商-淘礼金创建
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.68a1669aGNlNb6&source=search&docId=40172&docType=2
func (a *Api) SVegasTljCreate(param ScVegasTljCreateParams) (*VegasTljCreateResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.vegas.tlj.create", param, true)
	if err != nil {
		return nil, err
	}
	resp := VegasTljCreateResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PTljInstanceReport 淘宝客-推广者-淘礼金发放及使用报表
// https://open.taobao.com/api.htm?docId=43317&docType=2&scopeId=15029
func (a *Api) PTljInstanceReport(param TljInstanceReportParams) (*TljInstanceReportResp, error) {
	p, err := a.parameterCombination("taobao.tbk.dg.vegas.tlj.instance.report", param, false)
	if err != nil {
		return nil, err
	}
	resp := TljInstanceReportResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PublisherInfoSave 淘宝客-公用-私域用户备案
// https://open.taobao.com/api.htm?docId=37988&docType=2&scopeId=14474
func (a *Api) PublisherInfoSave(param PublisherInfoSaveParams) (*PublisherInfoSaveResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.publisher.info.save", param, false)
	if err != nil {
		return nil, err
	}
	resp := PublisherInfoSaveResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PublisherInfoGet 淘宝客-公用-私域用户备案信息查询
// https://open.taobao.com/api.htm?docId=37989&docType=2&scopeId=14474
func (a *Api) PublisherInfoGet(param PublisherInfoGetParams) (*PublisherInfoGetResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.publisher.info.get", param, false)
	if err != nil {
		return nil, err
	}
	resp := PublisherInfoGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InviteCodeGet 淘宝客-公用-私域用户邀请码生成
// https://open.taobao.com/api.htm?docId=38046&docType=2&scopeId=14474
func (a *Api) InviteCodeGet(param InviteCodeGetParams) (*InviteCodeGetResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.invitecode.get", param, false)
	if err != nil {
		return nil, err
	}
	resp := InviteCodeGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SMaterialOptional 淘宝客-服务商-物料搜索
// https://open.taobao.com/api.htm?docId=35263&docType=2&scopeId=13991
func (a *Api) SMaterialOptional(param MaterialOptionalParams) (*MaterialOptionalResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.material.optional", param, true)
	if err != nil {
		return nil, err
	}
	resp := MaterialOptionalResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PMaterialOptional 淘宝客-推广者-物料搜索
// https://open.taobao.com/api.htm?docId=35896&docType=2&scopeId=16516
func (a *Api) PMaterialOptional(param MaterialOptionalParams) (*MaterialOptionalResp, error) {
	p, err := a.parameterCombination("taobao.tbk.dg.material.optional", param, false)
	if err != nil {
		return nil, err
	}
	resp := MaterialOptionalResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PShopGet 淘宝客-推广者-店铺搜索
// https://open.taobao.com/api.htm?docId=24521&docType=2&scopeId=16516
func (a *Api) PShopGet(param ShopGetParams) (*ShopGetResp, error) {
	p, err := a.parameterCombination("taobao.tbk.shop.get", param, false)
	if err != nil {
		return nil, err
	}
	resp := ShopGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ItemClickExtract 淘宝客-公用-链接解析出商品id
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.3925669aFLMnkJ&source=search&docId=28156&docType=2
func (a *Api) ItemClickExtract(param ItemClickExtractParams) (*ItemClickExtractResp, error) {
	p, err := a.parameterCombination("taobao.tbk.item.click.extract", param, false)
	if err != nil {
		return nil, err
	}
	resp := ItemClickExtractResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SpreadGet 淘宝客-公用-长链转短链
// https://open.taobao.com/api.htm?docId=27832&docType=2&scopeId=12340
func (a *Api) SpreadGet(param SpreadGetParams) (*SpreadGetResp, error) {
	p, err := a.parameterCombination("taobao.tbk.spread.get", param, false)
	if err != nil {
		return nil, err
	}
	resp := SpreadGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PActivityInfoGet 淘宝客-推广者-官方活动转链
// https://open.taobao.com/api.htm?docId=48340&docType=2&scopeId=18294
func (a *Api) PActivityInfoGet(param ActivityInfoGetParams) (*ActivityInfoGetResp, error) {
	p, err := a.parameterCombination("taobao.tbk.activity.info.get", param, false)
	if err != nil {
		return nil, err
	}
	resp := ActivityInfoGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SActivityInfoGet 淘宝客-服务商-官方活动转链
// https://open.taobao.com/api.htm?docId=48417&docType=2&scopeId=18353
func (a *Api) SActivityInfoGet(param ActivityInfoGetParams) (*ActivityInfoGetResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.activity.info.get", param, true)
	if err != nil {
		return nil, err
	}
	resp := ActivityInfoGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SVegasSendReport  淘宝客-服务商-查询超级红包发放个数
// https://open.taobao.com/api.htm?docId=47590&docType=2&scopeId=17875
func (a *Api) SVegasSendReport(param VegasSendReportParams) (*VegasSendReportResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.vegas.send.report", param, true)
	if err != nil {
		return nil, err
	}
	resp := VegasSendReportResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PVegasSendReport  淘宝客-推广者-查询超级红包发放个数
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.78e5669aWq31Sg&source=search&docId=47593&docType=2
func (a *Api) PVegasSendReport(param VegasSendReportParams) (*VegasSendReportResp, error) {
	p, err := a.parameterCombination("taobao.tbk.dg.vegas.send.report", param, false)
	if err != nil {
		return nil, err
	}
	resp := VegasSendReportResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PJuItemsSearch  淘宝客-推广者-聚划算商品获取
// https://open.taobao.com/api.htm?docId=28762&docType=2&scopeId=16517
func (a *Api) PJuItemsSearch(param JuItemsSearchParams) (*JuItemsSearchResp, error) {
	p, err := a.parameterCombination("taobao.ju.items.search", param, false)
	if err != nil {
		return nil, err
	}
	resp := JuItemsSearchResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// POptimusMaterial 淘宝客-推广者-物料精选
// https://open.taobao.com/api.htm?docId=33947&docType=2&scopeId=16518
func (a *Api) POptimusMaterial(param OptimusMaterialParams) (*OptimusMaterialResp, error) {
	p, err := a.parameterCombination("taobao.tbk.dg.optimus.material", param, false)
	if err != nil {
		return nil, err
	}
	resp := OptimusMaterialResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SOptimusMaterial 淘宝客-服务商-物料精选
// https://open.taobao.com/api.htm?docId=37884&docType=2&scopeId=16287
func (a *Api) SOptimusMaterial(param OptimusMaterialParams) (*OptimusMaterialResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.optimus.material", param, true)
	if err != nil {
		return nil, err
	}
	resp := OptimusMaterialResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SOptimusPromotion 淘宝客-服务商-权益物料精选
// https://open.taobao.com/api.htm?docId=52701&docType=2&scopeId=16287
func (a *Api) SOptimusPromotion(param OptimusPromotionParams) (*OptimusPromotionResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.optimus.promotion", param, true)
	if err != nil {
		return nil, err
	}
	resp := OptimusPromotionResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// POptimusPromotion 淘宝客-推广者-权益物料精选
// https://open.taobao.com/api.htm?docId=52700&docType=2&scopeId=16518
func (a *Api) POptimusPromotion(param OptimusPromotionParams) (*OptimusPromotionResp, error) {
	p, err := a.parameterCombination("taobao.tbk.dg.optimus.promotion", param, false)
	if err != nil {
		return nil, err
	}
	resp := OptimusPromotionResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PNewUserOrderGet 淘宝客-推广者-新用户订单明细查询
// https://open.taobao.com/api.htm?docId=33892&docType=2&scopeId=16188
func (a *Api) PNewUserOrderGet(param NewUserOrderGetParams) (*NewUserOrderGetResp, error) {
	p, err := a.parameterCombination("taobao.tbk.dg.newuser.order.get", param, false)
	if err != nil {
		return nil, err
	}
	resp := NewUserOrderGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SNewUserOrderGet 淘宝客-服务商-新用户订单明细查询
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.6969669as9PC5i&source=search&docId=33897&docType=2
func (a *Api) SNewUserOrderGet(param NewUserOrderGetParams) (*NewUserOrderGetResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.newuser.order.get", param, true)
	if err != nil {
		return nil, err
	}
	resp := NewUserOrderGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PNewUserOrderSum 淘宝客-推广者-拉新活动对应数据查询
// https://open.taobao.com/api.htm?docId=36836&docType=2&scopeId=16188
func (a *Api) PNewUserOrderSum(param NewUserOrderSumParams) (*NewUserOrderSumResp, error) {
	p, err := a.parameterCombination("taobao.tbk.dg.newuser.order.sum", param, false)
	if err != nil {
		return nil, err
	}
	resp := NewUserOrderSumResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SNewUserOrderSum 淘宝客-服务商-拉新活动数据查询
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.464b669aDVh5GZ&source=search&docId=36837&docType=2
func (a *Api) SNewUserOrderSum(param NewUserOrderSumParams) (*NewUserOrderSumResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.newuser.order.sum", param, true)
	if err != nil {
		return nil, err
	}
	resp := NewUserOrderSumResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ShopRecommendGet 淘宝客-公用-店铺关联推荐
// https://open.taobao.com/api.htm?docId=24522&docType=2&scopeId=16292
func (a *Api) ShopRecommendGet(param ShopRecommendGetParams) (*ShopRecommendGetResp, error) {
	p, err := a.parameterCombination("taobao.tbk.shop.recommend.get", param, false)
	if err != nil {
		return nil, err
	}
	resp := ShopRecommendGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// TbkCouponGet 淘宝客-公用-阿里妈妈推广券详情查询
// https://open.taobao.com/api.htm?docId=31106&docType=2&scopeId=16189
func (a *Api) TbkCouponGet(param TbkCoUponGetParams) (*TbkCoUonGetResp, error) {
	p, err := a.parameterCombination("taobao.tbk.coupon.get", param, false)
	if err != nil {
		return nil, err
	}
	resp := TbkCoUonGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SShopConvert 淘宝客-服务商-店铺链接转换
// https://open.taobao.com/api.htm?docId=43878&docType=2&scopeId=16403
func (a *Api) SShopConvert(param ShopConvertParams) (*ShopConvertResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.shop.convert", param, true)
	if err != nil {
		return nil, err
	}
	resp := ShopConvertResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PShopConvert 淘宝客-推广者-店铺链接转换
// https://open.taobao.com/api.htm?docId=24523&docType=2&scopeId=11653
func (a *Api) PShopConvert(param ShopConvertParams) (*ShopConvertResp, error) {
	p, err := a.parameterCombination("taobao.tbk.shop.convert", param, false)
	if err != nil {
		return nil, err
	}
	resp := ShopConvertResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PItemConvert 淘宝客-推广者-商品链接转换
// https://open.taobao.com/api.htm?docId=24516&docType=2&scopeId=11653
func (a *Api) PItemConvert(param ItemConvertParams) (*ItemConvertResp, error) {
	p, err := a.parameterCombination("taobao.tbk.item.convert", param, false)
	if err != nil {
		return nil, err
	}
	resp := ItemConvertResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PVegasSendStatus 淘宝客-推广者-超级红包领取状态查询
// https://open.taobao.com/api.htm?docId=52958&docType=2&scopeId=21226
func (a *Api) PVegasSendStatus(param VegasSendStatusParams) (*VegasSendStatusResp, error) {
	p, err := a.parameterCombination("taobao.tbk.dg.vegas.send.status", param, false)
	if err != nil {
		return nil, err
	}
	resp := VegasSendStatusResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SVegasSendStatus 淘宝客-服务商-超级红包领取状态查询
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.3b54669aEtxGpd&source=search&docId=52957&docType=2
func (a *Api) SVegasSendStatus(param VegasSendStatusParams) (*VegasSendStatusResp, error) {
	p, err := a.parameterCombination("taobao.tbk.sc.vegas.send.status", param, true)
	if err != nil {
		return nil, err
	}
	resp := VegasSendStatusResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// TpwdQuery 查询解析淘口令
// https://open.taobao.com/api.htm?docId=32461&docType=2&scopeId=11998
func (a *Api) TpwdQuery(param TPwdQueryParams) (*TPwdQueryResp, error) {
	p, err := a.parameterCombination("taobao.wireless.share.tpwd.query", param, false)
	if err != nil {
		return nil, err
	}
	resp := TPwdQueryResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// TaoBaoTimeGet 获取淘宝系统当前时间
// https://open.taobao.com/api.htm?docId=120&docType=2&scopeId=381
func (a *Api) TaoBaoTimeGet() (*TaoBaoTimeGetResp, error) {
	p, err := a.parameterCombination("taobao.time.get", nil, false)
	if err != nil {
		return nil, err
	}
	resp := TaoBaoTimeGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// KeywordSearch 关键词过滤匹配
// https://open.taobao.com/api.htm?docId=10385&docType=2&scopeId=381
func (a *Api) KeywordSearch(param KeywordSearchParams) (*KeywordSearchResp, error) {
	p, err := a.parameterCombination("taobao.kfc.keyword.search", param, false)
	if err != nil {
		return nil, err
	}
	resp := KeywordSearchResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AppIPGet 获取ISV发起请求服务器IP
// https://open.taobao.com/api.htm?docId=21784&docType=2&scopeId=381
func (a *Api) AppIPGet() (*AppIPGetResp, error) {
	p, err := a.parameterCombination("taobao.appip.get", nil, false)
	if err != nil {
		return nil, err
	}
	resp := AppIPGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// OpenUIDChange 淘宝openUid 转换
// https://open.taobao.com/api.htm?docId=23831&docType=2&scopeId=381
func (a *Api) OpenUIDChange(param OpenUIDChangeParams) (*OpenUIDChangeResp, error) {
	p, err := a.parameterCombination("taobao.openuid.change", param, false)
	if err != nil {
		return nil, err
	}
	resp := OpenUIDChangeResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AuthTokenRefresh 刷新Access Token
// https://open.taobao.com/api.htm?docId=25387&docType=2&scopeId=381
func (a *Api) AuthTokenRefresh(param AuthTokenRefreshParams) (*AuthTokenRefreshResp, error) {
	p, err := a.parameterCombination("taobao.top.auth.token.refresh", param, false)
	if err != nil {
		return nil, err
	}
	resp := AuthTokenRefreshResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AuthTokenCreate 获取Access Token
// https://open.taobao.com/api.htm?docId=25388&docType=2&scopeId=381
func (a *Api) AuthTokenCreate(param AuthTokenCreateParams) (*AuthTokenCreateResp, error) {
	p, err := a.parameterCombination("taobao.top.auth.token.create", param, false)
	if err != nil {
		return nil, err
	}
	resp := AuthTokenCreateResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// HttpDnsGet TOPDNS配置
// https://open.taobao.com/api.htm?docId=25414&docType=2&scopeId=381
func (a *Api) HttpDnsGet() (*HttpDnsGetResp, error) {
	p, err := a.parameterCombination("taobao.httpdns.get", nil, false)
	if err != nil {
		return nil, err
	}
	resp := HttpDnsGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// TopIPOutGet 获取开放平台出口IP段
// https://open.taobao.com/api.htm?docId=25441&docType=2&scopeId=381
func (a *Api) TopIPOutGet() (*TopIPOutGetResp, error) {
	p, err := a.parameterCombination("taobao.top.ipout.get", nil, false)
	if err != nil {
		return nil, err
	}
	resp := TopIPOutGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// TopSecretGet 获取TOP通道解密秘钥
// https://open.taobao.com/api.htm?docId=26567&docType=2&scopeId=381
func (a *Api) TopSecretGet(param TopSecretGetParams) (*TopSecretGetResp, error) {
	p, err := a.parameterCombination("taobao.top.secret.get", param, false)
	if err != nil {
		return nil, err
	}
	resp := TopSecretGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// TopSecretRegisterParams 注册加密账号
//// https://open.taobao.com/api.htm?docId=27385&docType=2&scopeId=381
func (a *Api) TopSecretRegisterParams(param TopSecretRegisterParams) (*TopSecretRegisterResp, error) {
	p, err := a.parameterCombination("taobao.top.secret.register", param, false)
	if err != nil {
		return nil, err
	}
	resp := TopSecretRegisterResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SdkFeedbackUpload sdk信息回调
// https://open.taobao.com/api.htm?docId=27512&docType=2&scopeId=381
func (a *Api) SdkFeedbackUpload(param SdkFeedbackUploadParams) (*SdkFeedbackUploadResp, error) {
	p, err := a.parameterCombination("taobao.top.sdk.feedback.upload", param, false)
	if err != nil {
		return nil, err
	}
	resp := SdkFeedbackUploadResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FilesGet 业务文件获取
// https://open.taobao.com/api.htm?docId=32298&docType=2&scopeId=381
func (a *Api) FilesGet(param FilesGetParams) (*FilesGetResp, error) {
	p, err := a.parameterCombination("taobao.files.get", param, false)
	if err != nil {
		return nil, err
	}
	resp := FilesGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// OpenUIDGet 获取授权账号对应的OpenUid
// https://open.taobao.com/api.htm?docId=33220&docType=2&scopeId=381
func (a *Api) OpenUIDGet() (*OpenUIDGetResp, error) {
	p, err := a.parameterCombination("taobao.openuid.get", nil, false)
	if err != nil {
		return nil, err
	}
	resp := OpenUIDGetResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// OpenUIDGetByTrade 通过订单获取对应买家的openUID
// https://open.taobao.com/api.htm?docId=33221&docType=2&scopeId=381
func (a *Api) OpenUIDGetByTrade(param OpenUIDGetByTradeParams) (*OpenUIDGetByTradeResp, error) {
	p, err := a.parameterCombination("taobao.openuid.get.bytrade", param, false)
	if err != nil {
		return nil, err
	}
	resp := OpenUIDGetByTradeResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// OpenUIDGetByMixNick 通过mixnick转换openuid
// https://open.taobao.com/api.htm?docId=33223&docType=2&scopeId=381
func (a *Api) OpenUIDGetByMixNick(param OpenUIDGetByMixNickParams) (*OpenUIDGetByMixNickResp, error) {
	p, err := a.parameterCombination("taobao.openuid.get.bymixnick", param, false)
	if err != nil {
		return nil, err
	}
	resp := OpenUIDGetByMixNickResp{}
	err = gout.GET(gw).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (a *Api) parameterCombination(method string, param interface{}, session bool) (map[string]interface{}, error) {
	p := a.getPublicParameters(session)
	p["method"] = method

	_p, err := Struct2Map(param)
	if err != nil {
		return nil, err
	}

	for k, v := range _p {
		if _, ok := v.(map[string]interface{}); ok {
			dataType, _ := Marshal(v)
			dataString := string(dataType)
			p[k] = dataString
		} else if _, ok := v.([]interface{}); ok {
			dataType, _ := Marshal(v)
			dataString := string(dataType)
			p[k] = dataString
		} else {
			p[k] = v
		}
	}
	sign := a.getSign(p)
	p["sign"] = sign
	return p, nil
}

func (a *Api) getSign(param map[string]interface{}) string {
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	str := a.appSecret
	for _, key := range keys {
		str += key
		str += fmt.Sprint(param[key])
	}
	str += a.appSecret
	return GetMD5Encode(str)
}

func (a *Api) getPublicParameters(session bool) map[string]interface{} {
	if session {
		return map[string]interface{}{
			"app_key":     a.appKey,
			"format":      "json",
			"timestamp":   time.Now().Format("2006-01-02 15:04:05"),
			"v":           "2.0",
			"sign_method": "md5",
			"session":     a.session,
		}
	}
	return map[string]interface{}{
		"app_key":     a.appKey,
		"format":      "json",
		"timestamp":   time.Now().Format("2006-01-02 15:04:05"),
		"v":           "2.0",
		"sign_method": "md5",
	}

}


func GetMD5Encode(data string) string {
	h := md5.New()
	_, _ = h.Write([]byte(data))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Struct2Map(obj interface{}) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	j, _ := json.Marshal(obj)
	err := json.Unmarshal(j, &data)
	return data, err
}