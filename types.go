package tbkapi

// 淘宝客-淘宝客商品详情查询(简版)
type TbkItemInfoGetResp struct {
	ErrorResponse struct {
		SubMsg  string `json:"sub_msg"`
		Code    int    `json:"code"`
		SubCode string `json:"sub_code"`
		Msg     string `json:"msg"`
	} `json:"error_response"`
	TbkItemInfoGetResponse struct {
		Results struct {
			NTbkItem []struct {
				CatName     string `json:"cat_name"` //一级类目名称
				NumIid      int64  `json:"num_iid"`  //商品ID
				Title       string `json:"title"`    //商品标题
				PictURL     string `json:"pict_url"` //商品主图
				SmallImages struct {
					String []string `json:"string"` //商品小图列表
				} `json:"small_images"`
				ReservePrice               string `json:"reserve_price"`                  //商品一口价格
				ZkFinalPrice               string `json:"zk_final_price"`                 //折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价
				UserType                   int    `json:"user_type"`                      //卖家类型，0表示集市，1表示商城
				ProvCity                   string `json:"provcity"`                       //商品所在地
				ItemURL                    string `json:"item_url"`                       //商品链接
				SellerID                   int64  `json:"seller_id"`                      //卖家id
				Volume                     int    `json:"volume"`                         //30天销量
				Nick                       string `json:"nick"`                           //店铺名称
				CatLeafName                string `json:"cat_leaf_name"`                  //叶子类目名称
				FreeShipment               bool   `json:"free_shipment"`                  //是否包邮
				MaterialLibType            string `json:"material_lib_type"`              //商品库类型，支持多库类型输出，以英文逗号分隔“,”分隔，1:营销商品主推库，如果值为空则不属于1这种商品类型
				PresaleDiscountFeeText     string `json:"presale_discount_fee_text"`      //预售商品-商品优惠信息
				PresaleTailEndTime         int64  `json:"presale_tail_end_time"`          //预售商品-付定金结束时间（毫秒）
				PresaleTailStartTime       int64  `json:"presale_tail_start_time"`        //预售商品-付尾款开始时间（毫秒）
				PresaleEndTime             int64  `json:"presale_end_time"`               //预售商品-付定金结束时间（毫秒）
				PresaleStartTime           int64  `json:"presale_start_time"`             //预售商品-付定金开始时间（毫秒）
				PresaleDeposit             string `json:"presale_deposit"`                //预售商品-定金（元）
				TmallPlayActivityEndTime   int64  `json:"tmall_play_activity_end_time"`   //天猫限时抢可售 -结束时间（毫秒）
				TmallPlayActivityStartTime int64  `json:"tmall_play_activity_start_time"` //天猫限时抢可售 -开始时间（毫秒）
				JuOnlineStartTime          string `json:"ju_online_start_time"`           //聚划算信息-聚淘开始时间（毫秒）
				JuOnlineEndTime            string `json:"ju_online_end_time"`             //聚划算信息-聚淘结束时间（毫秒）
				JuPreShowStartTime         string `json:"ju_pre_show_start_time"`         //聚划算信息-商品预热开始时间（毫秒）
				JuPreShowEndTime           string `json:"ju_pre_show_end_time"`           //聚划算信息-商品预热结束时间（毫秒）
				SuperiorBrand              string `json:"superior_brand"`                 //是否品牌精选，0不是，1是
				IsPrepay                   bool   `json:"is_prepay"`                      //是否加入消费者保障
				ShopDsr                    int    `json:"shop_dsr"`                       //店铺dsr 评分
				RateSum                    int    `json:"ratesum"`                        //卖家等级
				IRfdRate                   bool   `json:"i_rfd_rate"`                     //退款率是否低于行业均值
				HGoodRate                  bool   `json:"h_good_rate"`                    //好评率是否高于行业均值
				HPayRate30                 bool   `json:"h_pay_rate30"`                   //成交转化是否高于行业均值
				JuPlayEndTime              int64  `json:"ju_play_end_time"`               //聚划算满减 -结束时间（毫秒）
				JuPlayStartTime            int64  `json:"ju_play_start_time"`             //聚划算满减 -开始时间（毫秒）
				PlayInfo                   string `json:"play_info"`                      //1聚划算满减：满N件减X元，满N件X折，满N件X元） 2天猫限时抢：前N分钟每件X元，前N分钟满N件每件X元，前N件每件X元）
				SalePrice                  string `json:"sale_price"`                     //活动价
				KuaDianPromotionInfo       string `json:"kuadian_promotion_info"`         //跨店满减信息
			} `json:"n_tbk_item"`
		} `json:"results"`
	} `json:"tbk_item_info_get_response"`
}

// 淘宝客-公用-淘宝客商品详情查询(简版)
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.aec0669aumR0CN&source=search&docId=24518&docType=2
type ItemInfoGetParams struct {
	NumIds   string `json:"num_iids,omitempty"` //商品ID串，用,分割，最大40个
	Platform int    `json:"platform,omitempty"` //链接形式：1：PC，2：无线，默认：１
	IP       string `json:"ip,omitempty"`       //ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
}

// CouponConvert 淘宝客-单品券高效转链
type CouponConvertResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkCouponConvertResponse struct {
		Result struct {
			ResultCode int `json:"result_code"`
			Results    struct {
				CouponType          int    `json:"coupon_type"`             //优惠券(商品优惠券推广链接中的券)类型，1 公开券，2 私有券，3 妈妈券
				CampaignType        int    `json:"campaign_type"`           //计划类型，0代表通用计划，1代表定向计划，2代表鹊桥计划，3代表营销计划
				CategoryID          int    `json:"category_id"`             //后台一级类目
				CouponClickURL      string `json:"coupon_click_url"`        //商品优惠券推广链接
				CouponEndTime       string `json:"coupon_end_time"`         //优惠券结束时间
				CouponInfo          string `json:"coupon_info"`             //优惠券面额
				CouponRemainCount   int    `json:"coupon_remain_count"`     //优惠券剩余量
				CouponStartTime     string `json:"coupon_start_time"`       //优惠券开始时间
				CouponTotalCount    int    `json:"coupon_total_count"`      //优惠券总量
				ItemID              int64  `json:"item_id"`                 //商品id
				ItemURL             string `json:"item_url"`                //商品淘客链接
				MaxCommissionRate   string `json:"max_commission_rate"`     //当不入参special_id、relation_id、external_id时，展示常规佣金率(%)
				RewardInfo          int    `json:"reward_info"`             //比价场景专用，当系统检测到入参消费者ID购买当前商品会获得《天天开彩蛋》玩法的彩蛋时，该字段显示1，否则为0
				YsYlClickURL        string `json:"ysyl_click_url"`          //预售有礼-推广链接
				YsYlTljFace         string `json:"ysyl_tlj_face"`           //预售有礼-预估淘礼金（元）
				YsYlTljSendTime     string `json:"ysyl_tlj_send_time"`      //预售有礼-淘礼金发放时间
				YsYlTljUseStartTime string `json:"ysyl_tlj_use_start_time"` //预售有礼-淘礼金使用开始时间
				YsYlCommissionRate  string `json:"ysyl_commission_rate"`    //预售有礼-佣金比例（%）（ 预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=9334376 ）
				YsYlTljUseEndTime   string `json:"ysyl_tlj_use_end_time"`   //预售有礼-淘礼金使用结束时间
				MinCommissionRate   string `json:"min_commission_rate"`
			} `json:"results"` // 推广者接口返回的数据
		} `json:"result"`
		RequestID string `json:"request_id"`
	} `json:"tbk_coupon_convert_response"`
	TbkPrivilegeGetResponse struct {
		Result struct {
			Data struct {
				CategoryID          int    `json:"category_id"`             //后台一级类目
				CouponClickURL      string `json:"coupon_click_url"`        //商品优惠券推广链接
				CouponEndTime       string `json:"coupon_end_time"`         //优惠券结束时间
				CouponInfo          string `json:"coupon_info"`             //优惠券面额
				CouponStartTime     string `json:"coupon_start_time"`       //优惠券开始时间
				ItemID              int64  `json:"item_id"`                 //商品id
				MaxCommissionRate   string `json:"max_commission_rate"`     //当不入参special_id、relation_id、external_id时，展示常规佣金率(%)
				CouponTotalCount    int    `json:"coupon_total_count"`      //优惠券总量
				CouponRemainCount   int    `json:"coupon_remain_count"`     //优惠券剩余量
				CouponType          int    `json:"coupon_type"`             //优惠券(商品优惠券推广链接中的券)类型，1 公开券，2 私有券，3 妈妈券
				ItemURL             string `json:"item_url"`                //商品淘客链接
				YsYlClickURL        string `json:"ysyl_click_url"`          //预售有礼-推广链接
				YsYlTljFace         string `json:"ysyl_tlj_face"`           //预售有礼-预估淘礼金（元）
				YsYlTljSendTime     string `json:"ysyl_tlj_send_time"`      //预售有礼-淘礼金发放时间
				YsYlCommissionRate  string `json:"ysyl_commission_rate"`    //预售有礼-佣金比例（%）（ 预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=9334376 ）
				YsYlTljUseStartTime string `json:"ysyl_tlj_use_start_time"` //预售有礼-淘礼金使用开始时间
				YsYlTljUseEndTime   string `json:"ysyl_tlj_use_end_time"`   //预售有礼-淘礼金使用结束时间
				MinCommissionRate   string `json:"min_commission_rate"`     //当入参special_id、relation_id、external_id时，该字段展示预估最低佣金率(%)
				RewardInfo          int    `json:"reward_info"`             //比价场景专用，当系统检测到入参消费者ID购买当前商品会获得《天天开彩蛋》玩法的彩蛋时，该字段显示1，否则为0
			} `json:"data"`
		} `json:"result"`
	} `json:"tbk_privilege_get_response"` //服务商接口返回的数据
}

// 淘宝客-服务商-单品券高效转链
// https://open.taobao.com/api.htm?docId=28625&docType=2&scopeId=12403
// 淘宝客-推广者-单品券高效转链
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.aec0669aumR0CN&source=search&docId=29289&docType=2
type CouponConvertParams struct {
	ItemID     string `json:"item_id,omitempty"`     //	淘客商品id
	AdZoneID   string `json:"adzone_id,omitempty"`   //来源广告位ID(pid中mm_1_2_3)中第3位
	Platform   int    `json:"platform,omitempty"`    //	1：PC，2：无线，默认：１
	SiteID     string `json:"site_id,omitempty"`     //源广告位ID(pid中mm_1_2_3)中第2位(服务商必须传入)
	ExternalId string `json:"external_id,omitempty"` //	淘宝客外部用户标记，如自身系统账户ID；微信ID等
	SpecialID  string `json:"special_id,omitempty"`  //	会员运营ID
	RelationID string `json:"relation_id,omitempty"` //	渠道管理ID
	XID        string `json:"xid,omitempty"`         //	团长与下游渠道合作的特殊标识，用于统计渠道推广效果
}

// 淘宝客-淘口令生成
type TextTPwdCreateResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkTPwdCreateResponse struct {
		Data struct {
			//针对苹果ios14及以上版本的苹果设备，手淘将按照示例值信息格式读取淘口令(需包含：数字+羊角符+url，识别规则可能根据ios情况变更)。
			//如需更改淘口令内文案、url等内容，请务必先验证更改后的淘口令在手淘可被识别打开！
			Model string `json:"model"`
			//非苹果ios14以上版本的设备（即其他ios版本、Android系统等），可以用此淘口令正常在复制到手淘打开
			PasswordSimple string `json:"password_simple"`
		} `json:"data"`
		RequestID string `json:"request_id"`
	} `json:"tbk_tpwd_create_response"`
}

// 淘宝客-公用-淘口令生成
// https://open.taobao.com/api.htm?docId=31127&docType=2&scopeId=11655
type TextTPwdCreateParams struct {
	UserID string `json:"user_id,omitempty"` //生成口令的淘宝用户ID
	Text   string `json:"text,omitempty"`    //口令弹框内容
	URL    string `json:"url,omitempty"`     //口令跳转目标页
	Logo   string `json:"logo,omitempty"`    //口令弹框logoURL
}

// 淘宝客-淘口令解析&转链 返回参数
type TPwdConvertResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkTPwdConvertResponse struct {
		Data struct {
			NumIid    string `json:"num_iid"`    //	商品Id
			ClickURL  string `json:"click_url"`  //商品淘客转链
			SellerID  string `json:"seller_id"`  //店铺卖家ID
			OriginURL string `json:"origin_url"` //入参淘口令对应原始链接
			OriginPid string `json:"origin_pid"` //入参淘口令推广链接中的pid，如果不属于当前调用的推广者则展示“0”
		} `json:"data"`
		RequestID string `json:"request_id"`
	} `json:"tbk_tpwd_convert_response"` //推广者接口数据
	TbkScTPwdConvertResponse struct {
		Data struct {
			NumIid    string `json:"num_iid"`    //	商品Id
			ClickURL  string `json:"click_url"`  //商品淘客转链
			SellerID  string `json:"seller_id"`  //店铺卖家ID
			OriginURL string `json:"origin_url"` //入参淘口令对应原始链接
			OriginPid string `json:"origin_pid"` //入参淘口令推广链接中的pid，如果不属于当前调用的推广者则展示“0”
		} `json:"data"`
	} `json:"tbk_sc_tpwd_convert_response"` //服务商接口数据
}

// 淘宝客-服务商-淘口令解析&转链
// https://open.taobao.com/api.htm?docId=43873&docType=2&scopeId=16401
// 淘宝客-推广者-淘口令解析&转链
// https://open.taobao.com/api.htm?docId=32932&docType=2&scopeId=16290
type TPwdConvertParams struct {
	PasswordContent string `json:"password_content,omitempty"` //需要解析的淘口令
	AdZoneID        string `json:"adzone_id,omitempty"`        //来源广告位ID(pid中mm_1_2_3)中第3位
	Dx              string `json:"dx,omitempty"`               //1表示商品转通用计划链接，其他值或不传表示优先转营销计划链接
	SiteID          string `json:"site_id,omitempty"`          //源广告位ID(pid中mm_1_2_3)中第2位(服务商必须传入)
}

// 淘宝客-处罚订单查询 返回信息
type PunishOrderGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkDgPunishOrderGetResponse struct {
		Result struct {
			Data struct {
				PageNo     int `json:"page_no"`
				PageSize   int `json:"page_size"`
				TotalCount int `json:"total_count"`
				Results    struct {
					Result []struct {
						RelationID        int64  `json:"relation_id"`
						SettleMonth       int    `json:"settle_month"`
						TkTradeCreateTime string `json:"tk_trade_create_time"`
						TbTradeParentID   int    `json:"tb_trade_parent_id"`
						SpecialID         int64  `json:"special_id"`
						UnionID           string `json:"union_id"`
						PunishStatus      string `json:"punish_status"`
						ViolationType     string `json:"violation_type"`
						TbTradeID         int64  `json:"tb_trade_id"`
						TkAdZoneID        int64  `json:"tk_adzone_id"`
						TkSiteID          int    `json:"tk_site_id"`
						TkPubID           string `json:"tk_pub_id"`
					} `json:"result"`
				} `json:"results"`
			} `json:"data"`
			BizErrorDesc string `json:"biz_error_desc"`
			BizErrorCode int    `json:"biz_error_code"`
			ResultMsg    string `json:"result_msg"`
			ResultCode   int    `json:"result_code"`
		} `json:"result"`
	} `json:"tbk_dg_punish_order_get_response"` //推广者接口数据
	TbkScPunishOrderGetResponse struct {
		Result struct {
			Data struct {
				PageNo     int `json:"page_no"`     //翻页的pageno
				PageSize   int `json:"page_size"`   //翻页的pagesie
				TotalCount int `json:"total_count"` //一共能查询出来的结果总数
				Results    struct {
					Result []struct {
						RelationID        int64  `json:"relation_id"`          //渠道关系id
						SettleMonth       int    `json:"settle_month"`         //结算月份
						TkTradeCreateTime string `json:"tk_trade_create_time"` //淘客订单创建时间
						TbTradeParentID   int    `json:"tb_trade_parent_id"`   //父订单号（该字段不再支持）
						SpecialID         int64  `json:"special_id"`           //会员运营id（该字段不再支持）
						UnionID           string `json:"union_id"`             //淘宝联盟unionid（该字段不再支持）
						PunishStatus      string `json:"punish_status"`        //处罚状态。0 冻结，1 解冻
						ViolationType     string `json:"violation_type"`       //处罚类型，目前包括 1:店铺淘宝客 2:订单虚假交易
						TbTradeID         int64  `json:"tb_trade_id"`          //子订单号
						TkAdZoneID        int64  `json:"tk_adzone_id"`         //pid里的adzoneid
						TkSiteID          int    `json:"tk_site_id"`           //pid里的siteid
						TkPubID           int    `json:"tk_pub_id"`            //pid里的pubid
					} `json:"result"`
				} `json:"results"`
			} `json:"data"`
			BizErrorDesc string `json:"biz_error_desc"` //业务出错的描述
			BizErrorCode int    `json:"biz_error_code"` //业务出错的状态码
			ResultMsg    string `json:"result_msg"`     //执行结果
			ResultCode   int    `json:"result_code"`    //执行结果状态码 200成功
		} `json:"result"`
	} `json:"tbk_sc_punish_order_get_response"` // 服务商接口数据
}

// 淘宝客-服务商-处罚订单查询
// https://open.taobao.com/api.htm?docId=41942&docType=2&scopeId=15738
// 淘宝客-推广者-处罚订单查询
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.5d97669aNPczDi&source=search&docId=42050&docType=2
type PunishOrderGetParams struct {
	AfOrderOption AfOrderOption `json:"af_order_option"`
}

type AfOrderOption struct {
	SiteID          string `json:"site_id,omitempty"`            //	源广告位ID(pid中mm_1_2_3)中第2位(服务商必须传入)
	Span            int    `json:"span,omitempty"`               //	查询时间跨度，不超过30天，单位是天
	RelationID      string `json:"relation_id,omitempty"`        //	渠道关系id
	TbTradeID       int    `json:"tb_trade_id,omitempty"`        //	子订单号
	TbTradeParentID int    `json:"tb_trade_parent_id,omitempty"` //	此参数不再使用，请勿入参
	PageSize        int    `json:"page_size,omitempty"`          //	pagesize
	PageNo          int    `json:"page_no,omitempty"`            //	pageNo
	StartTime       string `json:"start_time,omitempty"`         //2018-11-11 00:01:01	查询开始时间，以taoke订单创建时间开始
	AdZoneID        string `json:"adzone_id,omitempty"`          //	来源广告位ID(pid中mm_1_2_3)中第3位
}

// 淘宝客-维权退款订单查询 返回信息
type RelationRefundResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkRelationRefundResponse struct {
		Result struct {
			BizErrorCode int    `json:"biz_error_code"` //业务错误码 101, 102,103
			BizErrorDesc string `json:"biz_error_desc"` //业务错误信息
			Data         struct {
				TotalCount string `json:"total_count"`
				PageSize   string `json:"page_size"`
				PageNo     string `json:"page_no"`
				Results    struct {
					Result []struct {
						TbTradeParentID             int    `json:"tb_trade_parent_id"`              //淘宝订单编号
						RelationID                  int64  `json:"relation_id"`                     //渠道关系id
						Tk3RdPubID                  int    `json:"tk3rd_pub_id"`                    //第三方推广者memberid
						TkPubID                     int    `json:"tk_pub_id"`                       //推广者memberid
						TkSubsidyFeeRefund3RdPub    string `json:"tk_subsidy_fee_refund3rd_pub"`    //第三方应该返还的补贴
						TkCommissionFeeRefund3RdPub string `json:"tk_commission_fee_refund3rd_pub"` //第三方应该返还的佣金
						TkSubsidyFeeRefundPub       string `json:"tk_subsidy_fee_refund_pub"`       //第二方应该返还的补贴(不包括技术服务费)
						TkCommissionFeeRefundPub    string `json:"tk_commission_fee_refund_pub"`    //第二方应该返还的佣金(不包括技术服务费)
						TkRefundSuitTime            string `json:"tk_refund_suit_time"`             //维权完成时间
						TkRefundTime                string `json:"tk_refund_time"`                  //维权创建时间
						EarningTime                 string `json:"earning_time"`                    //订单结算时间
						TbTradeCreateTime           string `json:"tb_trade_create_time"`            //订单创建时间
						//维权创建(淘客结算回执) 4,维权成功(淘客结算回执) 2,维权失败(淘客结算回执) 3,发生多次维权，待处理 11,
						//从淘客处补扣（钱已结给淘客） 等待扣款 12,从淘客处补扣（钱已结给淘客） 扣款成功 13,
						//从卖家处补扣（钱已结给卖家） 等待扣款 14,从卖家处补扣（钱已结给卖家） 扣款成功 15
						RefundStatus          int    `json:"refund_status"`
						TbAuctionTitle        string `json:"tb_auction_title"`          //宝贝标题
						TbTradeID             int    `json:"tb_trade_id"`               //淘宝子订单编号
						SpecialID             int64  `json:"special_id"`                //会员关系id
						RefundFee             string `json:"refund_fee"`                //维权金额
						TbTradeFinishPrice    string `json:"tb_trade_finish_price"`     //结算金额
						TkPubShowReturnFee    string `json:"tk_pub_show_return_fee"`    //应返商家金额(二方)
						Tk3RdPubShowReturnFee string `json:"tk3rd_pub_show_return_fee"` //应返商家金额(三方)
						RefundType            int    `json:"refund_type"`               //1 表示2方，2表示3方
						AlScPid               string `json:"alsc_pid"`                  //（口碑订单）口碑父订单号
						AlScID                string `json:"alsc_id"`                   //（口碑订单）口碑子订单号
					} `json:"result"`
				} `json:"results"`
			} `json:"data"`
			ResultCode int    `json:"result_code"` //接口返回值信息，跟rpc架构保持一致
			ResultMsg  string `json:"result_msg"`  //返回信息
		} `json:"result"`
	} `json:"tbk_relation_refund_response"` //推广者接口数据
	TbkScRelationRefundResponse struct {
		Result struct {
			BizErrorCode int    `json:"biz_error_code"` //业务错误码 101, 102,103
			BizErrorDesc string `json:"biz_error_desc"` //业务错误信息
			Data         struct {
				TotalCount int `json:"total_count"`
				PageSize   int `json:"page_size"`
				PageNo     int `json:"page_no"`
				Results    struct {
					Result []struct {
						TbTradeParentID             int    `json:"tb_trade_parent_id"`              //淘宝订单编号
						RelationID                  int64  `json:"relation_id"`                     //渠道关系id
						Tk3RdPubID                  int    `json:"tk3rd_pub_id"`                    //第三方推广者memberid
						TkPubID                     int    `json:"tk_pub_id"`                       //推广者memberid
						TkSubsidyFeeRefund3RdPub    string `json:"tk_subsidy_fee_refund3rd_pub"`    //第三方应该返还的补贴
						TkCommissionFeeRefund3RdPub string `json:"tk_commission_fee_refund3rd_pub"` //第三方应该返还的佣金
						TkSubsidyFeeRefundPub       string `json:"tk_subsidy_fee_refund_pub"`       //第二方应该返还的补贴(不包括技术服务费)
						TkCommissionFeeRefundPub    string `json:"tk_commission_fee_refund_pub"`    //第二方应该返还的佣金(不包括技术服务费)
						TkRefundSuitTime            string `json:"tk_refund_suit_time"`             //维权完成时间
						TkRefundTime                string `json:"tk_refund_time"`                  //维权创建时间
						EarningTime                 string `json:"earning_time"`                    //订单结算时间
						TbTradeCreateTime           string `json:"tb_trade_create_time"`            //订单创建时间
						//维权创建(淘客结算回执) 4,维权成功(淘客结算回执) 2,维权失败(淘客结算回执) 3,发生多次维权，待处理 11,
						//从淘客处补扣（钱已结给淘客） 等待扣款 12,从淘客处补扣（钱已结给淘客） 扣款成功 13,
						//从卖家处补扣（钱已结给卖家） 等待扣款 14,从卖家处补扣（钱已结给卖家） 扣款成功 15
						RefundStatus          int    `json:"refund_status"`
						TbAuctionTitle        string `json:"tb_auction_title"`          //宝贝标题
						TbTradeID             int    `json:"tb_trade_id"`               //淘宝子订单编号
						SpecialID             int64  `json:"special_id"`                //会员关系id
						RefundFee             string `json:"refund_fee"`                //维权金额
						TbTradeFinishPrice    string `json:"tb_trade_finish_price"`     //结算金额
						TkPubShowReturnFee    string `json:"tk_pub_show_return_fee"`    //应返商家金额(二方)
						Tk3RdPubShowReturnFee string `json:"tk3rd_pub_show_return_fee"` //应返商家金额(三方)
						RefundType            int    `json:"refund_type"`               //1 表示2方，2表示3方
						AlScPid               string `json:"alsc_pid"`                  //（口碑订单）口碑父订单号
						AlScID                string `json:"alsc_id"`                   //（口碑订单）口碑子订单号
					} `json:"result"`
				} `json:"results"`
			} `json:"data"`
			ResultCode int    `json:"result_code"` //接口返回值信息，跟rpc架构保持一致
			ResultMsg  string `json:"result_msg"`  //返回信息
		} `json:"result"`
	} `json:"tbk_sc_relation_refund_response"` //服务商接口数据
}

// 淘宝客-服务商-维权退款订单查询
// https://open.taobao.com/api.htm?docId=43874&docType=2&scopeId=16322
// 淘宝客-推广者-维权退款订单查询
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.7a86669aRBEZUg&source=search&docId=40121&docType=2
type RelationRefundParams struct {
	SearchOption SearchOption `json:"search_option"`
}

type SearchOption struct {
	PageSize   int    `json:"page_size,omitempty"`   //	pagesize
	SearchType int    `json:"search_type,omitempty"` //	1-维权发起时间，2-订单结算时间（正向订单），3-维权完成时间，4-订单创建时间
	RefundType int    `json:"refund_type"`           //	1 表示2方，2表示3方，0表示不限
	StartTime  string `json:"start_time,omitempty"`  //	2018-10-10 00:00:00	开始时间
	PageNo     int    `json:"page_no,omitempty"`     //	pagenumber
	BizType    int    `json:"biz_type,omitempty"`    //	1代表渠道关系id，2代表会员关系id
}

//淘宝客-所有订单查询
type OrderDetailsGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkOrderDetailsGetResponse struct {
		Data struct {
			Results struct {
				PublisherOrderDto []struct {
					TbPaidTime       string `json:"tb_paid_time"`       //订单在淘宝拍下付款的时间
					TkPaidTime       string `json:"tk_paid_time"`       //订单付款的时间，该时间同步淘宝，可能会略晚于买家在淘宝的订单创建时间
					PayPrice         string `json:"pay_price"`          //买家确认收货的付款金额（不包含运费金额）
					PubShareFee      string `json:"pub_share_fee"`      //结算预估收入=结算金额*提成。以买家确认收货的付款金额为基数，预估您可能获得的收入。因买家退款、您违规推广等原因，可能与您最终收入不一致。最终收入以月结后您实际收到的为准
					TradeID          string `json:"trade_id"`           //买家通过购物车购买的每个商品对应的订单编号，此订单编号并未在淘宝买家后台透出
					TkOrderRole      int    `json:"tk_order_role"`      //二方：佣金收益的第一归属者； 三方：从其他淘宝客佣金中进行分成的推广者
					TkEarningTime    string `json:"tk_earning_time"`    //订单确认收货后且商家完成佣金支付的时间
					AdZoneID         int64  `json:"adzone_id"`          //来源广告位ID(pid中mm_1_2_3)中第3位
					PubShareRate     string `json:"pub_share_rate"`     //从结算佣金中分得的收益比率
					UnID             string `json:"unid"`               //unid(本字段不对外开放)
					RefundTag        int    `json:"refund_tag"`         //维权标签，0 含义为非维权 1 含义为维权订单
					SubsidyRate      string `json:"subsidy_rate"`       //平台给与的补贴比率，如天猫、淘宝、聚划算等
					TkTotalRate      string `json:"tk_total_rate"`      //提成=收入比率*分成比率。指实际获得收益的比率
					ItemCategoryName string `json:"item_category_name"` //商品所属的一级类目名称
					SellerNick       string `json:"seller_nick"`        //掌柜旺旺
					PubID            int    `json:"pub_id"`             //推广者的会员id
					AliMaMaRate      string `json:"alimama_rate"`       //推广者赚取佣金后支付给阿里妈妈的技术服务费用的比率
					SubsidyType      string `json:"subsidy_type"`       //平台出资方，如天猫、淘宝、或聚划算等
					ItemImg          string `json:"item_img"`           //商品图片
					PubSharePreFee   string `json:"pub_share_pre_fee"`  //付款预估收入=付款金额*提成。指买家付款金额为基数，预估您可能获得的收入。因买家退款等原因，可能与结算预估收入不一致
					AlipayTotalPrice string `json:"alipay_total_price"` //买家拍下付款的金额（不包含运费金额）
					ItemTitle        string `json:"item_title"`         //商品标题
					SiteName         string `json:"site_name"`          //媒体管理下的对应ID的自定义名称
					ItemNum          int    `json:"item_num"`           //商品数量
					SubsidyFee       string `json:"subsidy_fee"`        //补贴金额=结算金额*补贴比率
					AliMaMaShareFee  string `json:"alimama_share_fee"`  //技术服务费=结算金额*收入比率*技术服务费率。推广者赚取佣金后支付给阿里妈妈的技术服务费用
					TradeParentID    string `json:"trade_parent_id"`    //买家在淘宝后台显示的订单编号
					OrderType        string `json:"order_type"`         //订单所属平台类型，包括天猫、淘宝、聚划算等
					TkCreateTime     string `json:"tk_create_time"`     //订单创建的时间，该时间同步淘宝，可能会略晚于买家在淘宝的订单创建时间
					FlowSource       string `json:"flow_source"`        //产品类型
					TerminalType     string `json:"terminal_type"`      //成交平台
					ClickTime        string `json:"click_time"`         //通过推广链接达到商品、店铺详情页的点击时间
					//已付款：指订单已付款，但还未确认收货 已收货：指订单已确认收货，但商家佣金未支付 已结算：指订单已确认收货，且商家佣金已支付成功
					//已失效：指订单关闭/订单佣金小于0.01元，订单关闭主要有：1）买家超时未付款； 2）买家付款前，买家/卖家取消了订单；3）订单付款后发起售中退款成功；3：订单结算，12：订单付款， 13：订单失效，14：订单成功
					TkStatus                           int      `json:"tk_status"`
					ItemPrice                          string   `json:"item_price"`                               //商品单价
					ItemID                             int64    `json:"item_id"`                                  //商品id
					AdZoneName                         string   `json:"adzone_name"`                              //推广位管理下的自定义推广位名称
					TotalCommissionRate                string   `json:"total_commission_rate"`                    //佣金比率
					ItemLink                           string   `json:"item_link"`                                //商品链接
					SiteID                             int      `json:"site_id"`                                  //媒体管理下的ID，同时也是pid=mm_1_2_3中的“2”这段数字
					SellerShopTitle                    string   `json:"seller_shop_title"`                        //店铺名称
					IncomeRate                         string   `json:"income_rate"`                              //订单结算的佣金比率+平台的补贴比率
					TotalCommissionFee                 string   `json:"total_commission_fee"`                     //佣金金额=结算金额*佣金比率
					TkCommissionPreFeeForMediaPlatform string   `json:"tk_commission_pre_fee_for_media_platform"` //预估内容专项服务费：内容场景专项技术服务费，内容推广者在内容场景进行推广需要支付给阿里妈妈专项的技术服务费用。专项服务费＝付款金额＊专项服务费率。
					TkCommissionFeeForMediaPlatform    string   `json:"tk_commission_fee_for_media_platform"`     //结算内容专项服务费：内容场景专项技术服务费，内容推广者在内容场景进行推广需要支付给阿里妈妈专项的技术服务费用。专项服务费＝结算金额＊专项服务费率。
					TkCommissionRateForMediaPlatform   string   `json:"tk_commission_rate_for_media_platform"`    //内容专项服务费率：内容场景专项技术服务费率，内容推广者在内容场景进行推广需要按结算金额支付一定比例给阿里妈妈作为内容场景专项技术服务费，用于提供与内容平台实现产品技术对接等服务。
					SpecialID                          int64    `json:"special_id"`                               //会员运营id
					RelationID                         int64    `json:"relation_id"`                              //渠道关系id
					DepositPrice                       string   `json:"deposit_price"`                            //预售时期，用户对预售商品支付的定金金额
					TbDepositTime                      string   `json:"tb_deposit_time"`                          //预售时期，用户对预售商品支付定金的付款时间
					TkDepositTime                      string   `json:"tk_deposit_time"`                          //预售时期，用户对预售商品支付定金的付款时间，可能略晚于在淘宝付定金时间
					AlScID                             string   `json:"alsc_id"`                                  //口碑子订单号
					AlScPid                            string   `json:"alsc_pid"`                                 //口碑父订单号
					ServiceFeeDtoList                  struct { //服务费信息
						ServiceFeeDto []struct {
							TkShareRoleType   int    `json:"tk_share_role_type"`  //专项服务费来源，122-渠道
							ShareRelativeRate string `json:"share_relative_rate"` //专项服务费率
							ShareFee          string `json:"share_fee"`           //结算专项服务费
							SharePreFee       string `json:"share_pre_fee"`       //预估专项服务费
						} `json:"service_fee_dto"`
					} `json:"service_fee_dto_list"`
					LxRid string `json:"lx_rid"` //激励池对应的rid
					IsLx  string `json:"is_lx"`  //订单是否为激励池订单 1，表征是 0，表征否
				} `json:"publisher_order_dto"`
			} `json:"results"`
			HasPre        bool   `json:"has_pre"`        //是否还有上一页
			PositionIndex string `json:"position_index"` //位点字段，由调用方原样传递
			HasNext       bool   `json:"has_next"`       //是否还有下一页
			PageNo        int    `json:"page_no"`        //页码
			PageSize      int    `json:"page_size"`      //页大小
		} `json:"data"`
	} `json:"tbk_order_details_get_response"` //推广者接口数据
	TbkScOrderDetailsGetResponse struct {
		Data struct {
			Results struct {
				PublisherOrderDto []struct {
					TbPaidTime       string `json:"tb_paid_time"`       //订单在淘宝拍下付款的时间
					TkPaidTime       string `json:"tk_paid_time"`       //订单付款的时间，该时间同步淘宝，可能会略晚于买家在淘宝的订单创建时间
					PayPrice         string `json:"pay_price"`          //买家确认收货的付款金额（不包含运费金额）
					PubShareFee      string `json:"pub_share_fee"`      //结算预估收入=结算金额*提成。以买家确认收货的付款金额为基数，预估您可能获得的收入。因买家退款、您违规推广等原因，可能与您最终收入不一致。最终收入以月结后您实际收到的为准
					TradeID          string `json:"trade_id"`           //买家通过购物车购买的每个商品对应的订单编号，此订单编号并未在淘宝买家后台透出
					TkOrderRole      int    `json:"tk_order_role"`      //二方：佣金收益的第一归属者； 三方：从其他淘宝客佣金中进行分成的推广者
					TkEarningTime    string `json:"tk_earning_time"`    //订单确认收货后且商家完成佣金支付的时间
					AdZoneID         int64  `json:"adzone_id"`          //推广位管理下的推广位名称对应的ID，同时也是pid=mm_1_2_3中的“3”这段数字
					PubShareRate     string `json:"pub_share_rate"`     //从结算佣金中分得的收益比率
					UnID             string `json:"unid"`               //unid(本字段不对外开放)
					RefundTag        int    `json:"refund_tag"`         //维权标签，0 含义为非维权 1 含义为维权订单
					SubsidyRate      string `json:"subsidy_rate"`       //平台给与的补贴比率，如天猫、淘宝、聚划算等
					TkTotalRate      string `json:"tk_total_rate"`      //提成=收入比率*分成比率。指实际获得收益的比率
					ItemCategoryName string `json:"item_category_name"` //商品所属的一级类目名称
					SellerNick       string `json:"seller_nick"`        //掌柜旺旺
					PubID            int    `json:"pub_id"`             //推广者的会员id
					AliMaMaRate      string `json:"alimama_rate"`       //推广者赚取佣金后支付给阿里妈妈的技术服务费用的比率
					SubsidyType      string `json:"subsidy_type"`       //平台出资方，如天猫、淘宝、或聚划算等
					ItemImg          string `json:"item_img"`           //商品图片
					PubSharePreFee   string `json:"pub_share_pre_fee"`  //付款预估收入=付款金额*提成。指买家付款金额为基数，预估您可能获得的收入。因买家退款等原因，可能与结算预估收入不一致
					AlipayTotalPrice string `json:"alipay_total_price"` //买家拍下付款的金额（不包含运费金额）
					ItemTitle        string `json:"item_title"`         //商品标题
					SiteName         string `json:"site_name"`          //媒体管理下的对应ID的自定义名称
					ItemNum          int    `json:"item_num"`           //商品数量
					SubsidyFee       string `json:"subsidy_fee"`        //补贴金额=结算金额*补贴比率
					AliMaMaShareFee  string `json:"alimama_share_fee"`  //技术服务费=结算金额*收入比率*技术服务费率。推广者赚取佣金后支付给阿里妈妈的技术服务费用
					TradeParentID    string `json:"trade_parent_id"`    //买家在淘宝后台显示的订单编号
					OrderType        string `json:"order_type"`         //订单所属平台类型，包括天猫、淘宝、聚划算等
					TkCreateTime     string `json:"tk_create_time"`     //订单创建的时间，该时间同步淘宝，可能会略晚于买家在淘宝的订单创建时间
					FlowSource       string `json:"flow_source"`        //产品类型
					TerminalType     string `json:"terminal_type"`      //成交平台
					ClickTime        string `json:"click_time"`         //通过推广链接达到商品、店铺详情页的点击时间
					//已付款：指订单已付款，但还未确认收货 已收货：指订单已确认收货，但商家佣金未支付 已结算：指订单已确认收货，且商家佣金已支付成功
					//已失效：指订单关闭/订单佣金小于0.01元，订单关闭主要有：1）买家超时未付款； 2）买家付款前，买家/卖家取消了订单；3）订单付款后发起售中退款成功；3：订单结算，12：订单付款， 13：订单失效，14：订单成功
					TkStatus                           int      `json:"tk_status"`
					ItemPrice                          string   `json:"item_price"`                               //商品单价
					ItemID                             int64    `json:"item_id"`                                  //商品id
					AdZoneName                         string   `json:"adzone_name"`                              //推广位管理下的自定义推广位名称
					TotalCommissionRate                string   `json:"total_commission_rate"`                    //佣金比率
					ItemLink                           string   `json:"item_link"`                                //商品链接
					SiteID                             int      `json:"site_id"`                                  //媒体管理下的ID，同时也是pid=mm_1_2_3中的“2”这段数字
					SellerShopTitle                    string   `json:"seller_shop_title"`                        //店铺名称
					IncomeRate                         string   `json:"income_rate"`                              //订单结算的佣金比率+平台的补贴比率
					TotalCommissionFee                 string   `json:"total_commission_fee"`                     //佣金金额=结算金额*佣金比率
					TkCommissionPreFeeForMediaPlatform string   `json:"tk_commission_pre_fee_for_media_platform"` //预估内容专项服务费：内容场景专项技术服务费，内容推广者在内容场景进行推广需要支付给阿里妈妈专项的技术服务费用。专项服务费＝付款金额＊专项服务费率。
					TkCommissionFeeForMediaPlatform    string   `json:"tk_commission_fee_for_media_platform"`     //结算内容专项服务费：内容场景专项技术服务费，内容推广者在内容场景进行推广需要支付给阿里妈妈专项的技术服务费用。专项服务费＝结算金额＊专项服务费率。
					TkCommissionRateForMediaPlatform   string   `json:"tk_commission_rate_for_media_platform"`    //内容专项服务费率：内容场景专项技术服务费率，内容推广者在内容场景进行推广需要按结算金额支付一定比例给阿里妈妈作为内容场景专项技术服务费，用于提供与内容平台实现产品技术对接等服务。
					SpecialID                          int64    `json:"special_id"`                               //会员运营id
					RelationID                         int64    `json:"relation_id"`                              //渠道关系id
					DepositPrice                       string   `json:"deposit_price"`                            //预售时期，用户对预售商品支付的定金金额
					TbDepositTime                      string   `json:"tb_deposit_time"`                          //预售时期，用户对预售商品支付定金的付款时间
					TkDepositTime                      string   `json:"tk_deposit_time"`                          //预售时期，用户对预售商品支付定金的付款时间，可能略晚于在淘宝付定金时间
					AlScID                             string   `json:"alsc_id"`                                  //口碑子订单号
					AlScPid                            string   `json:"alsc_pid"`                                 //口碑父订单号
					ServiceFeeDtoList                  struct { //服务费信息
						ServiceFeeDto []struct {
							TkShareRoleType   int    `json:"tk_share_role_type"`  //专项服务费来源，122-渠道
							ShareRelativeRate string `json:"share_relative_rate"` //专项服务费率
							ShareFee          string `json:"share_fee"`           //结算专项服务费
							SharePreFee       string `json:"share_pre_fee"`       //预估专项服务费
						} `json:"service_fee_dto"`
					} `json:"service_fee_dto_list"`
					LxRid string `json:"lx_rid"` //激励池对应的rid
					IsLx  string `json:"is_lx"`  //订单是否为激励池订单 1，表征是 0，表征否
				} `json:"publisher_order_dto"`
			} `json:"results"`
			HasPre        bool   `json:"has_pre"`        //是否还有上一页
			PositionIndex string `json:"position_index"` //位点字段，由调用方原样传递
			HasNext       bool   `json:"has_next"`       //是否还有下一页
			PageNo        int    `json:"page_no"`        //页码
			PageSize      int    `json:"page_size"`      //页大小
		} `json:"data"`
	} `json:"tbk_sc_order_details_get_response"` //服务商接口数据
}

// 淘宝客-推广者-所有订单查询
// https://open.taobao.com/api.htm?docId=43328&docType=2&scopeId=16175
// 淘宝客-服务商-所有订单查询
// https://open.taobao.com/api.htm?docId=43755&docType=2&scopeId=16322
type OrderDetailsGetParams struct {
	QueryType     int    `json:"query_type,omitempty"`     //查询时间类型，1：按照订单淘客创建时间查询，2:按照订单淘客付款时间查询，3:按照订单淘客结算时间查询
	PositionIndex string `json:"position_index,omitempty"` //位点，除第一页之外，都需要传递；前端原样返回。
	PageSize      int    `json:"page_size,omitempty"`      //页大小，默认20，1~100
	MemberType    int    `json:"member_type,omitempty"`    //推广者角色类型,2:二方，3:三方，不传，表示所有角色
	TkStatus      int    `json:"tk_status,omitempty"`      //淘客订单状态，12-付款，13-关闭，14-确认收货，3-结算成功;不传，表示所有状态
	EndTime       string `json:"end_time,omitempty"`       //2019-04-23 12:28:22	订单查询结束时间，订单开始时间至订单结束时间，中间时间段日常要求不超过3个小时，但如618、双11、年货节等大促期间预估时间段不可超过20分钟，超过会提示错误，调用时请务必注意时间段的选择，以保证亲能正常调用！
	StartTime     string `json:"start_time,omitempty"`     //2019-04-05 12:18:22	订单查询开始时间
	JumpType      int    `json:"jump_type,omitempty"`      //跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页
	PageNo        int    `json:"page_no,omitempty"`        //第几页，默认1，1~100
	OrderScene    int    `json:"order_scene,omitempty"`    //场景订单场景类型，1:常规订单，2:渠道订单，3:会员运营订单，默认为1
}

// 淘宝客-淘礼金创建
type VegasTljCreateResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkDgVegasTljCreateResponse struct {
		Result struct {
			Model struct {
				RightsID     string `json:"rights_id"`     //淘礼金Id
				SendURL      string `json:"send_url"`      //淘礼金领取Url
				VegasCode    string `json:"vegas_code"`    //投放code【百川商品详情页业务专用】
				AvailableFee string `json:"available_fee"` //创建完成后资金账户可用资金，单位元，保留2位小数
			} `json:"model"`
			MsgCode string `json:"msg_code"`
			MsgInfo string `json:"msg_info"`
			Success bool   `json:"success"`
		} `json:"result"`
	} `json:"tbk_dg_vegas_tlj_create_response"` //推广者接口数据
	TbkScVegasTljCreateResponse struct {
		Result struct {
			Model struct {
				RightsID string `json:"rights_id"` //淘礼金Id
				SendURL  string `json:"send_url"`  //淘礼金领取Url
			} `json:"model"`
			MsgCode string `json:"msg_code"`
			MsgInfo string `json:"msg_info"`
			Success bool   `json:"success"`
		} `json:"result"`
	} `json:"tbk_sc_vegas_tlj_create_response"` //服务商接口数据
}

// 淘宝客-服务商-淘礼金创建
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.68a1669aGNlNb6&source=search&docId=40172&docType=2
// 淘宝客-推广者-淘礼金创建
// https://open.taobao.com/api.htm?docId=40173&docType=2&scopeId=15029
type VegasTljCreateParams struct {
	AaMpAiGnType         string `json:"campaign_type,omitempty"`            //	CPS佣金计划类型
	AdZoneID             string `json:"adzone_id,omitempty"`                //	来源广告位ID(pid中mm_1_2_3)中第3位
	ItemID               int    `json:"item_id,omitempty"`                  //宝贝id
	TotalNum             int    `json:"total_num,omitempty"`                //	淘礼金总个数
	Name                 string `json:"name,omitempty"`                     //	淘礼金名称，最大10个字符
	UserTotalWinNumLimit int    `json:"user_total_win_num_limit,omitempty"` //	单用户累计中奖次数上限
	SecuritySwitch       bool   `json:"security_switch"`                    //启动安全：true；不启用安全：false安全开关，若不进行安全校验，可能放大您的资损风险，请谨慎选择
	PerFace              string `json:"per_face,omitempty"`                 //单个淘礼金面额，支持两位小数，单位元
	SendStartTime        string `json:"send_start_time,omitempty"`          //	2018-09-01 00:00:00	发放开始时间
	SendEndTime          string `json:"send_end_time,omitempty"`            //	2018-09-01 00:00:00	发放截止时间
	UseEndTime           string `json:"use_end_time,omitempty"`             //	使用结束日期。如果是结束时间模式为相对时间，时间格式为1-7直接的整数, 例如，1（相对领取时间1天）； 如果是绝对时间，格式为yyyy-MM-dd，例如，2019-01-29，表示到2019-01-29 23:59:59结束
	UseEndTimeMode       int    `json:"use_end_time_mode,omitempty"`        //	结束日期的模式,1:相对时间，2:绝对时间
	UseStartTime         string `json:"use_start_time,omitempty"`           //2019-01-29	使用开始日期。相对时间，无需填写，以用户领取时间作为使用开始时间。绝对时间，格式 yyyy-MM-dd，例如，2019-01-29，表示从2019-01-29 00:00:00 开始
	SecurityLevel        int    `json:"security_level,omitempty"`           //安全等级，0：适用于常规淘礼金投放场景；1：更高安全级别，适用于淘礼金面额偏大等安全性较高的淘礼金投放场景，可能导致更多用户拦截。security_switch为true，此字段不填写时，使用0作为默认安全级别。如果security_switch为false，不进行安全判断。
}

// 淘宝客-服务商-淘礼金创建
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.68a1669aGNlNb6&source=search&docId=40172&docType=2
type ScVegasTljCreateParams struct {
	AaMpAiGnType         string `json:"campaign_type,omitempty"`            //CPS佣金计划类型
	AdZoneID             string `json:"adzone_id,omitempty"`                //来源广告位ID(pid中mm_1_2_3)中第3位
	ItemID               int    `json:"item_id,omitempty"`                  //宝贝id
	TotalNum             int    `json:"total_num,omitempty"`                //淘礼金总个数
	Name                 string `json:"name,omitempty"`                     //淘礼金名称，最大10个字符
	UserTotalWinNumLimit int    `json:"user_total_win_num_limit,omitempty"` //单用户累计中奖次数上限
	SecuritySwitch       bool   `json:"security_switch"`                    //启动安全：true；不启用安全：false安全开关，若不进行安全校验，可能放大您的资损风险，请谨慎选择
	PerFace              string `json:"per_face,omitempty"`                 //单个淘礼金面额，支持两位小数，单位元
	SiteID               string `json:"site_id,omitempty"`                  //源广告位ID(pid中mm_1_2_3)中第2位(服务商必须传入)
}

// 淘宝客-淘礼金发放及使用报表
type TljInstanceReportResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkDgVegasTljInstanceReportResponse struct {
		Result struct {
			Model struct {
				UnfreezeAmount      string `json:"unfreeze_amount"`       //解冻金额
				UnfreezeNum         int    `json:"unfreeze_num"`          //解冻红包个数
				RefundAmount        string `json:"refund_amount"`         //失效回退金额
				RefundNum           int    `json:"refund_num"`            //失效回退红包个数
				AlipayAmount        string `json:"alipay_amount"`         //引导成交金额
				UseAmount           string `json:"use_amount"`            //红包核销金额
				UseNum              int    `json:"use_num"`               //红包核销个数
				WinAmount           string `json:"win_amount"`            //红包领取金额
				WinNum              int    `json:"win_num"`               //红包领取个数
				PreCommissionAmount string `json:"pre_commission_amount"` //引导预估佣金金额
				FpRefundAmount      string `json:"fp_refund_amount"`      //退款红包金额
				FpRefundNum         int    `json:"fp_refund_num"`         //退款红包个数
			} `json:"model"`
			MsgCode string `json:"msg_code"`
			MsgInfo string `json:"msg_info"`
			Success bool   `json:"success"` //是否成功
		} `json:"result"`
	} `json:"tbk_dg_vegas_tlj_instance_report_response"`
}

// 淘宝客-推广者-淘礼金发放及使用报表
// https://open.taobao.com/api.htm?docId=43317&docType=2&scopeId=15029
type TljInstanceReportParams struct {
	RightsID string `json:"rights_id,omitempty"` //实例ID
}

// 淘宝客-公用-私域用户备案
type PublisherInfoSaveResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkScPublisherInfoSaveResponse struct {
		Data struct {
			RelationID  int64  `json:"relation_id"`  //渠道关系ID
			AccountName string `json:"account_name"` //渠道昵称
			SpecialID   int64  `json:"special_id"`   //会员运营ID
			Desc        string `json:"desc"`         //如果重复绑定会提示：”重复绑定渠道“或”重复绑定粉丝“
		} `json:"data"`
	} `json:"tbk_sc_publisher_info_save_response"`
}

// 淘宝客-公用-私域用户备案
// https://open.taobao.com/api.htm?docId=37988&docType=2&scopeId=14474
type PublisherInfoSaveParams struct {
	RelationFrom string `json:"relation_from,omitempty"` //	渠道备案 - 来源，取链接的来源
	OfflineScene string `json:"offline_scene,omitempty"` //渠道备案 - 线下场景信息，1 - 门店，2- 学校，3 - 工厂，4 - 其他
	OnlineScene  string `json:"online_scene,omitempty"`  //	渠道备案 - 线上场景信息，1 - 微信群，2- QQ群，3 - 其他
	InviterCode  string `json:"inviter_code,omitempty"`  //	淘宝客邀请渠道或会员的邀请码
	InfoType     int    `json:"info_type,omitempty"`     //	类型，必选 默认为1:
	Note         string `json:"note,omitempty"`          //	媒体侧渠道备注
	//线下备案注册信息,字段包含: 电话号码(phoneNumber，必填),
	//省(province,必填),
	//市(city,必填),
	//区县街道(location,必填),
	//详细地址(detailAddress,必填),
	//经营类型(career,线下个人必填),
	//店铺类型(shopType,线下店铺必填),
	//店铺名称(shopName,线下店铺必填),
	//店铺证书类型(shopCertifyType,线下店铺选填),
	//店铺证书编号(certifyNumber,线下店铺选填)
	RegisterInfo RegisterInfo `json:"register_info,omitempty"` // 线下备案注册信息
}

type RegisterInfo struct {
	PhoneNumber     string `json:"phoneNumber,omitempty"`     //电话号码
	City            string `json:"city,omitempty"`            //省
	Province        string `json:"province,omitempty"`        //市
	Location        string `json:"location,omitempty"`        //区县街道
	DetailAddress   string `json:"detailAddress,omitempty"`   //详细地址
	ShopType        string `json:"shopType,omitempty"`        //店铺类型
	ShopName        string `json:"shopName,omitempty"`        //店铺名称
	ShopCertifyType string `json:"shopCertifyType,omitempty"` //店铺证书类型
	CertifyNumber   string `json:"certifyNumber,omitempty"`   //店铺证书编号
}

// 淘宝客-公用-私域用户备案信息查询
type PublisherInfoGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkScPublisherInfoGetResponse struct {
		Data struct {
			RootPidChannelList struct {
				String []string `json:"string"` //渠道专属pidList
			} `json:"root_pid_channel_list"`
			TotalCount  int `json:"total_count"` //共享字段 - 总记录数
			InviterList struct {
				MapData []struct {
					RelationApp  string `json:"relation_app"`  //共享字段 - 备案场景：common（通用备案），etao（一淘备案），minietao（一淘小程序备案），offlineShop（线下门店备案），offlinePerson（线下个人备案）
					CreateDate   string `json:"create_date"`   //共享字段 - 备案日期
					AccountName  string `json:"account_name"`  //共享字段 - 渠道/会员昵称
					RealName     string `json:"real_name"`     //共享字段 - 渠道/会员姓名
					RelationID   int64  `json:"relation_id"`   //渠道独有 - 渠道关系ID
					OfflineScene string `json:"offline_scene"` //渠道独有 - 线下场景信息，1 - 门店，2- 学校，3 - 工厂，4 - 其他
					OnlineScene  string `json:"online_scene"`  //渠道独有 - 线上场景信息，1 - 微信群，2- QQ群，3 - 其他
					Note         string `json:"note"`          //渠道独有 - 媒体侧渠道备注信息
					RootPid      string `json:"root_pid"`      //共享字段 - 渠道/会员专属pid
					RTag         string `json:"rtag"`          //共享字段 - 渠道/会员原始身份信息
					OfflineInfo  struct {
						ShopName        string `json:"shop_name"`         //渠道独有 -店铺名称
						ShopType        string `json:"shop_type"`         //渠道独有 -店铺类型
						PhoneNumber     string `json:"phone_number"`      //渠道独有 -手机号码
						DetailAddress   string `json:"detail_address"`    //渠道独有 -详细地址
						Location        string `json:"location"`          //渠道独有 -地区
						ShopCertifyType string `json:"shop_certify_type"` //渠道独有 -证件类型
						CertifyNumber   string `json:"certify_number"`    //渠道独有 -对应的证件证件类型编号
						Career          string `json:"career"`            //渠道独有 -经营类型
					} `json:"offline_info"` //线下备案专属信息
					SpecialID    int64  `json:"special_id"`    //会员独有 - 会员运营ID
					PunishStatus string `json:"punish_status"` //渠道独有 - 处罚状态
					ExternalID   string `json:"external_id"`   //淘宝客外部用户标记，如自身系统账户ID；微信ID等
				} `json:"map_data"`
			} `json:"inviter_list"` //共享字段 - 渠道或会员列表
		} `json:"data"`
	} `json:"tbk_sc_publisher_info_get_response"`
}

// 淘宝客-公用-私域用户备案信息查询
// https://open.taobao.com/api.htm?docId=37989&docType=2&scopeId=14474
type PublisherInfoGetParams struct {
	InfoType    int    `json:"info_type,omitempty"`    //	类型，必选 1:渠道信息；2:会员信息
	RelationID  int64  `json:"relation_id,omitempty"`  //	渠道独占 - 渠道关系ID
	PageNO      int    `json:"page_no,omitempty"`      //	第几页
	PageSize    int    `json:"page_size,omitempty"`    //	每页大小
	RelationApp string `json:"relation_app,omitempty"` //	common	备案的场景：common（通用备案），etao（一淘备案），minietao（一淘小程序备案），offlineShop（线下门店备案），offlinePerson（线下个人备案）。如不填默认common。查询会员信息只需填写common即可
	SpecialID   string `json:"special_id,omitempty"`   //	会员独占 - 会员运营ID
	ExternalID  string `json:"external_id,omitempty"`  //	淘宝客外部用户标记，如自身系统账户ID；微信ID等
}

// 淘宝客-公用-私域用户邀请码生成
type InviteCodeGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkScInviteCodeGetResponse struct {
		Data struct {
			InviterCode string `json:"inviter_code"` //邀请码
		} `json:"data"`
	} `json:"tbk_sc_invitecode_get_response"`
}

// 淘宝客-公用-私域用户邀请码生成
// https://open.taobao.com/api.htm?docId=38046&docType=2&scopeId=14474
type InviteCodeGetParams struct {
	RelationID  int64  `json:"relation_id,omitempty"`  //	渠道关系ID
	RelationPpp string `json:"relation_app,omitempty"` //	common	渠道推广的物料类型
	CodeType    int    `json:"code_type,omitempty"`    //	邀请码类型，1 - 渠道邀请，2 - 渠道裂变，3 -会员邀请
}

// 淘宝客-物料搜索
type MaterialOptionalResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkDgMaterialOptionalResponse struct {
		TotalResults int `json:"total_results"` //搜索到符合条件的结果总数
		ResultList   struct {
			MapData []struct {
				CouponStartTime string `json:"coupon_start_time"` //优惠券信息-优惠券开始时间
				CouponEndTime   string `json:"coupon_end_time"`   //优惠券信息-优惠券结束时间
				InfoDxJh        string `json:"info_dxjh"`         //商品信息-定向计划信息
				TkTotalSales    string `json:"tk_total_sales"`    //商品信息-淘客30天推广量
				TkTotalComMi    string `json:"tk_total_commi"`    //商品信息-月支出佣金(该字段废弃，请勿再用)
				CouponID        string `json:"coupon_id"`         //优惠券信息-优惠券id
				NumIid          int64  `json:"num_iid"`           //商品信息-宝贝id(该字段废弃，请勿再用)
				Title           string `json:"title"`             //商品信息-商品标题
				PictURL         string `json:"pict_url"`          //商品信息-商品主图
				SmallImages     struct {
					String []string `json:"string"` //商品信息-商品小图列表
				} `json:"small_images"` //商品信息-商品小图列表
				ReservePrice           string `json:"reserve_price"`             //商品信息-商品一口价格
				ZkFinalPrice           string `json:"zk_final_price"`            //折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价
				UserType               int    `json:"user_type"`                 //店铺信息-卖家类型。0表示集市，1表示天猫
				ProvCity               string `json:"provcity"`                  //商品信息-宝贝所在地
				ItemURL                string `json:"item_url"`                  //链接-宝贝地址
				IncludeMkt             string `json:"include_mkt"`               //商品信息-是否包含营销计划
				IncludeDxJh            string `json:"include_dxjh"`              //商品信息-是否包含定向计划
				CommissionRate         string `json:"commission_rate"`           //商品信息-佣金比率。1550表示15.5%
				Volume                 int    `json:"volume"`                    //商品信息-30天销量（饿了么卡券信息-总销量）
				SellerID               int64  `json:"seller_id"`                 //店铺信息-卖家id
				CouponTotalCount       int    `json:"coupon_total_count"`        //优惠券信息-优惠券总量
				CouponRemainCount      int    `json:"coupon_remain_count"`       //优惠券信息-优惠券剩余量
				CouponInfo             string `json:"coupon_info"`               //优惠券信息-优惠券满减信息
				CommissionType         string `json:"commission_type"`           //商品信息-佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划
				ShopTitle              string `json:"shop_title"`                //店铺信息-店铺名称
				URL                    string `json:"url"`                       //	链接-宝贝推广链接
				CouponShareURL         string `json:"coupon_share_url"`          //链接-宝贝+券二合一页面链接
				ShopDsr                int    `json:"shop_dsr"`                  //店铺信息-店铺dsr评分
				WhiteImage             string `json:"white_image"`               //商品信息-商品白底图
				ShortTitle             string `json:"short_title"`               //商品信息-商品短标题
				CategoryID             int    `json:"category_id"`               //商品信息-叶子类目id
				CategoryName           string `json:"category_name"`             //商品信息-叶子类目名称
				LevelOneCategoryID     int    `json:"level_one_category_id"`     //商品信息-一级类目ID
				LevelOneCategoryName   string `json:"level_one_category_name"`   //商品信息-一级类目名称
				OeTime                 string `json:"oetime"`                    //拼团专用-拼团结束时间
				OsTime                 string `json:"ostime"`                    //拼团专用-拼团开始时间
				JddNum                 int    `json:"jdd_num"`                   //拼团专用-拼团几人团
				JddPrice               string `json:"jdd_price"`                 //拼团专用-拼团拼成价，单位元
				UvSumPreSale           int    `json:"uv_sum_pre_sale"`           //预售专用-预售数量
				CouponAmount           string `json:"coupon_amount"`             //优惠券（元） 若属于预售商品，该优惠券付尾款可用，付定金不可用
				CouponStartFee         string `json:"coupon_start_fee"`          //优惠券信息-优惠券起用门槛，满X元可用。如：满299元减20元
				ItemDescription        string `json:"item_description"`          //商品信息-宝贝描述(推荐理由)
				Nick                   string `json:"nick"`                      //店铺信息-卖家昵称
				XID                    string `json:"x_id"`                      //店铺信息-卖家昵称
				OrigPrice              string `json:"orig_price"`                //拼团专用-拼团一人价（原价)，单位元
				TotalStock             int    `json:"total_stock"`               //拼团专用-拼团库存数量
				SellNum                int    `json:"sell_num"`                  //拼团专用-拼团已售数量
				Stock                  int    `json:"stock"`                     //拼团专用-拼团剩余库存
				TmaLlPlayActivityInfo  string `json:"tmall_play_activity_info"`  //营销-天猫营销玩法
				ItemID                 int64  `json:"item_id"`                   //商品信息-宝贝id
				RealPostFee            string `json:"real_post_fee"`             //商品邮费
				LockRateStartTime      int64  `json:"lock_rate_start_time"`      //锁佣开始时间
				LockRateEndTime        int64  `json:"lock_rate_end_time"`        //锁佣结束时间
				LockRate               string `json:"lock_rate"`                 //锁住的佣金率
				PresaleDiscountFeeText string `json:"presale_discount_fee_text"` //预售商品-优惠
				PresaleTailEndTime     int64  `json:"presale_tail_end_time"`     //预售商品-付尾款结束时间（毫秒）
				PresaleTailStartTime   int64  `json:"presale_tail_start_time"`   //预售商品-付尾款开始时间（毫秒）
				PresaleEndTime         int64  `json:"presale_end_time"`          //预售商品-付定金结束时间（毫秒）
				PresaleStartTime       int64  `json:"presale_start_time"`        //预售商品-付定金开始时间（毫秒）
				PresaleDeposit         string `json:"presale_deposit"`           //预售商品-定金（元）
				YsYlTljSendTime        string `json:"ysyl_tlj_send_time"`        //预售有礼-淘礼金发放时间
				YsYlClickURL           string `json:"ysyl_click_url"`            //预售有礼-推广链接
				YsYlCommissionRate     string `json:"ysyl_commission_rate"`      //预售有礼-佣金比例（ 预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=9334376 ）
				YsYlTljFace            string `json:"ysyl_tlj_face"`             //预售有礼-预估淘礼金（元）
				YsYlTljUseEndTime      string `json:"ysyl_tlj_use_end_time"`     //预售有礼-淘礼金使用结束时间
				YsYlTljUseStartTime    string `json:"ysyl_tlj_use_start_time"`   //预售有礼-淘礼金使用开始时间
				UsableShopName         string `json:"usable_shop_name"`          //本地化-可用店铺名称
				UsableShopID           string `json:"usable_shop_id"`            //本地化-可用店铺id
				Distance               string `json:"distance"`                  //本地化-到门店距离（米）
				SaleEndTime            string `json:"sale_end_time"`             //本地化-销售结束时间
				SaleBeginTime          string `json:"sale_begin_time"`           //本地化-销售开始时间
				SalePrice              string `json:"sale_price"`                //活动价
				KuaDianPromotionInfo   string `json:"kuadian_promotion_info"`    //跨店满减信息
				SuperiorBrand          string `json:"superior_brand"`            //是否品牌精选，0不是，1是
				RewardInfo             int    `json:"reward_info"`               //比价场景专用，当系统检测到入参消费者ID购买当前商品会获得《天天开彩蛋》玩法的彩蛋时，该字段显示1，否则为0
			} `json:"map_data"`
		} `json:"result_list"`
	} `json:"tbk_dg_material_optional_response"` //推广者接口数据
	TbkScMaterialOptionalResponse struct {
		TotalResults int `json:"total_results"` //搜索到符合条件的结果总数
		ResultList   struct {
			MapData []struct {
				CouponStartTime string `json:"coupon_start_time"` //优惠券信息-优惠券开始时间
				CouponEndTime   string `json:"coupon_end_time"`   //优惠券信息-优惠券结束时间
				InfoDxJh        string `json:"info_dxjh"`         //商品信息-定向计划信息
				TkTotalSales    string `json:"tk_total_sales"`    //商品信息-淘客30天推广量
				TkTotalComMi    string `json:"tk_total_commi"`    //商品信息-月支出佣金(该字段废弃，请勿再用)
				CouponID        string `json:"coupon_id"`         //优惠券信息-优惠券id
				NumIid          int64  `json:"num_iid"`           //商品信息-宝贝id(该字段废弃，请勿再用)
				Title           string `json:"title"`             //商品信息-商品标题
				PictURL         string `json:"pict_url"`          //商品信息-商品主图
				SmallImages     struct {
					String []string `json:"string"` //商品信息-商品小图列表
				} `json:"small_images"` //商品信息-商品小图列表
				ReservePrice           string `json:"reserve_price"`             //商品信息-商品一口价格
				ZkFinalPrice           string `json:"zk_final_price"`            //折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价
				UserType               int    `json:"user_type"`                 //店铺信息-卖家类型。0表示集市，1表示天猫
				ProvCity               string `json:"provcity"`                  //商品信息-宝贝所在地
				ItemURL                string `json:"item_url"`                  //链接-宝贝地址
				IncludeMkt             string `json:"include_mkt"`               //商品信息-是否包含营销计划
				IncludeDxJh            string `json:"include_dxjh"`              //商品信息-是否包含定向计划
				CommissionRate         string `json:"commission_rate"`           //商品信息-佣金比率。1550表示15.5%
				Volume                 int    `json:"volume"`                    //商品信息-30天销量（饿了么卡券信息-总销量）
				SellerID               int64  `json:"seller_id"`                 //店铺信息-卖家id
				CouponTotalCount       int    `json:"coupon_total_count"`        //优惠券信息-优惠券总量
				CouponRemainCount      int    `json:"coupon_remain_count"`       //优惠券信息-优惠券剩余量
				CouponInfo             string `json:"coupon_info"`               //优惠券信息-优惠券满减信息
				CommissionType         string `json:"commission_type"`           //商品信息-佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划
				ShopTitle              string `json:"shop_title"`                //店铺信息-店铺名称
				URL                    string `json:"url"`                       //	链接-宝贝推广链接
				CouponShareURL         string `json:"coupon_share_url"`          //链接-宝贝+券二合一页面链接
				ShopDsr                int    `json:"shop_dsr"`                  //店铺信息-店铺dsr评分
				WhiteImage             string `json:"white_image"`               //商品信息-商品白底图
				ShortTitle             string `json:"short_title"`               //商品信息-商品短标题
				CategoryID             int    `json:"category_id"`               //商品信息-叶子类目id
				CategoryName           string `json:"category_name"`             //商品信息-叶子类目名称
				LevelOneCategoryID     int    `json:"level_one_category_id"`     //商品信息-一级类目ID
				LevelOneCategoryName   string `json:"level_one_category_name"`   //商品信息-一级类目名称
				OeTime                 string `json:"oetime"`                    //拼团专用-拼团结束时间
				OsTime                 string `json:"ostime"`                    //拼团专用-拼团开始时间
				JddNum                 int    `json:"jdd_num"`                   //拼团专用-拼团几人团
				JddPrice               string `json:"jdd_price"`                 //拼团专用-拼团拼成价，单位元
				UvSumPreSale           int    `json:"uv_sum_pre_sale"`           //预售专用-预售数量
				CouponAmount           string `json:"coupon_amount"`             //优惠券（元） 若属于预售商品，该优惠券付尾款可用，付定金不可用
				CouponStartFee         string `json:"coupon_start_fee"`          //优惠券信息-优惠券起用门槛，满X元可用。如：满299元减20元
				ItemDescription        string `json:"item_description"`          //商品信息-宝贝描述(推荐理由)
				Nick                   string `json:"nick"`                      //店铺信息-卖家昵称
				XID                    string `json:"x_id"`                      //店铺信息-卖家昵称
				OrigPrice              string `json:"orig_price"`                //拼团专用-拼团一人价（原价)，单位元
				TotalStock             int    `json:"total_stock"`               //拼团专用-拼团库存数量
				SellNum                int    `json:"sell_num"`                  //拼团专用-拼团已售数量
				Stock                  int    `json:"stock"`                     //拼团专用-拼团剩余库存
				TmaLlPlayActivityInfo  string `json:"tmall_play_activity_info"`  //营销-天猫营销玩法
				ItemID                 int64  `json:"item_id"`                   //商品信息-宝贝id
				RealPostFee            string `json:"real_post_fee"`             //商品邮费
				LockRateStartTime      int64  `json:"lock_rate_start_time"`      //锁佣开始时间
				LockRateEndTime        int64  `json:"lock_rate_end_time"`        //锁佣结束时间
				LockRate               string `json:"lock_rate"`                 //锁住的佣金率
				PresaleDiscountFeeText string `json:"presale_discount_fee_text"` //预售商品-优惠
				PresaleTailEndTime     int64  `json:"presale_tail_end_time"`     //预售商品-付尾款结束时间（毫秒）
				PresaleTailStartTime   int64  `json:"presale_tail_start_time"`   //预售商品-付尾款开始时间（毫秒）
				PresaleEndTime         int64  `json:"presale_end_time"`          //预售商品-付定金结束时间（毫秒）
				PresaleStartTime       int64  `json:"presale_start_time"`        //预售商品-付定金开始时间（毫秒）
				PresaleDeposit         string `json:"presale_deposit"`           //预售商品-定金（元）
				YsYlTljSendTime        string `json:"ysyl_tlj_send_time"`        //预售有礼-淘礼金发放时间
				YsYlClickURL           string `json:"ysyl_click_url"`            //预售有礼-推广链接
				YsYlCommissionRate     string `json:"ysyl_commission_rate"`      //预售有礼-佣金比例（ 预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=9334376 ）
				YsYlTljFace            string `json:"ysyl_tlj_face"`             //预售有礼-预估淘礼金（元）
				YsYlTljUseEndTime      string `json:"ysyl_tlj_use_end_time"`     //预售有礼-淘礼金使用结束时间
				YsYlTljUseStartTime    string `json:"ysyl_tlj_use_start_time"`   //预售有礼-淘礼金使用开始时间
				UsableShopName         string `json:"usable_shop_name"`          //本地化-可用店铺名称
				UsableShopID           string `json:"usable_shop_id"`            //本地化-可用店铺id
				Distance               string `json:"distance"`                  //本地化-到门店距离（米）
				SaleEndTime            string `json:"sale_end_time"`             //本地化-销售结束时间
				SaleBeginTime          string `json:"sale_begin_time"`           //本地化-销售开始时间
				SalePrice              string `json:"sale_price"`                //活动价
				KuaDianPromotionInfo   string `json:"kuadian_promotion_info"`    //跨店满减信息
				SuperiorBrand          string `json:"superior_brand"`            //是否品牌精选，0不是，1是
				RewardInfo             int    `json:"reward_info"`               //比价场景专用，当系统检测到入参消费者ID购买当前商品会获得《天天开彩蛋》玩法的彩蛋时，该字段显示1，否则为0
			} `json:"map_data"`
		} `json:"result_list"`
	} `json:"tbk_sc_material_optional_response"` //服务商接口数据
}

// 淘宝客-服务商-物料搜索
// https://open.taobao.com/api.htm?docId=35263&docType=2&scopeId=13991
// 淘宝客-推广者-物料搜索
// https://open.taobao.com/api.htm?docId=35896&docType=2&scopeId=16516
type MaterialOptionalParams struct {
	StartDsr          int    `json:"start_dsr,omitempty"`            //	商品筛选(特定媒体支持)-店铺dsr评分。筛选大于等于当前设置的店铺dsr评分的商品0-50000之间
	PageSize          int    `json:"page_size,omitempty"`            //	页大小，默认20，1~100
	PageNo            int    `json:"page_no,omitempty"`              //	第几页，默认：１
	Platform          int    `json:"platform,omitempty"`             //	链接形式-1：PC，2：无线，默认为１
	EndTkRate         int    `json:"end_tk_rate,omitempty"`          //	商品筛选-淘客佣金比率上限。如：1234表示12.34%
	StartTkRate       int    `json:"start_tk_rate,omitempty"`        //	商品筛选-淘客佣金比率下限。如：1234表示12.34%
	EndPrice          int    `json:"end_price,omitempty"`            //	商品筛选-折扣价范围上限。单位：元
	StartPrice        int    `json:"start_price,omitempty"`          //	商品筛选-折扣价范围上限。单位：元
	IsOverseas        bool   `json:"is_overseas"`                    //	商品筛选-是否海外商品。true表示属于海外商品，false或不设置表示不限
	IsTmaLl           bool   `json:"is_tmall"`                       //	商品筛选-是否天猫商品。true表示属于天猫商品，false或不设置表示不限
	Sort              string `json:"sort,omitempty"`                 //	tk_rate_des	排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price）
	ItemLoc           string `json:"itemloc,omitempty"`              //	杭州	商品筛选-所在地
	Cat               string `json:"cat,omitempty"`                  //	商品筛选-后台类目ID。用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
	Q                 string `json:"q,omitempty"`                    //	女装	查询的关键词
	AdZoneID          string `json:"adzone_id,omitempty"`            //来源广告位ID(pid中mm_1_2_3)中第3位
	SiteID            string `json:"site_id,omitempty"`              //源广告位ID(pid中mm_1_2_3)中第2位(服务商必须传入)
	MaterialID        int    `json:"material_id,omitempty"`          //	不传时默认物料id=2836。如果直接对消费者投放，可使用官方个性化算法优化的搜索物料id=17004
	HasCoupon         bool   `json:"has_coupon"`                     //	优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限
	IP                string `json:"ip,omitempty"`                   //	13.2.33.4	ip参数影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
	IncludeRfdRate    bool   `json:"include_rfd_rate"`               //	商品筛选(特定媒体支持)-退款率是否低于行业均值。True表示大于等于，false或不设置表示不限
	IncludeGoodRate   bool   `json:"include_good_rate"`              //	商品筛选(特定媒体支持)-好评率是否高于行业均值。True表示大于等于，false或不设置表示不限
	IncludePayRate30  bool   `json:"include_pay_rate_30"`            //	商品筛选(特定媒体支持)-成交转化是否高于行业均值。True表示大于等于，false或不设置表示不限
	NeedPrepay        bool   `json:"need_prepay"`                    //	商品筛选-是否加入消费者保障。true表示加入，false或不设置表示不限
	NeedFreeShipment  bool   `json:"need_free_shipment"`             //	商品筛选-是否包邮。true表示包邮，false或不设置表示不限
	NpxLevel          int    `json:"npx_level,omitempty"`            //	商品筛选-牛皮癣程度。取值：1不限，2无，3轻微
	EndKaTkRate       int    `json:"end_ka_tk_rate,omitempty"`       //	商品筛选-KA媒体淘客佣金率上限。如：1234表示12.34%
	StartKaTkRate     int    `json:"start_ka_tk_rate,omitempty"`     //	商品筛选-KA媒体淘客佣金率下限。如：1234表示12.34%
	DeviceValue       string `json:"device_value,omitempty"`         //	智能匹配-设备号加密后的值（MD5加密需32位小写）
	DeviceEncrypt     string `json:"device_encrypt,omitempty"`       //	智能匹配-设备号加密类型：MD5
	DeviceType        string `json:"device_type,omitempty"`          //	智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
	LockRateEndTime   int    `json:"lock_rate_end_time,omitempty"`   //	1567440000000	商品筛选-锁佣结束时间
	LockRateStartTime int    `json:"lock_rate_start_time,omitempty"` //	1567440000000	商品筛选-锁佣开始时间
	SellerIds         string `json:"seller_ids,omitempty"`           //	商家筛选-商家id，仅支持饿了么卡券商家id，支持批量请求1-100以内，多个商家id使用英文逗号分隔
	CityCode          string `json:"city_code,omitempty"`            //	本地化入参-LBS信息-国标城市码，仅支持单个请求，请求饿了么卡券物料时，该字段必填。 （详细城市ID见：https://mo.m.taobao.com/page_2020010315120200508）
	Latitude          string `json:"latitude,omitempty"`             //	本地化入参-LBS信息-纬度
	Longitude         string `json:"longitude,omitempty"`            //	本地化入参-LBS信息-经度
	RelationID        string `json:"relation_id,omitempty"`          //	渠道关系ID，仅适用于渠道推广场景
	SpecialID         string `json:"special_id,omitempty"`           //	会员运营ID，仅适用于会员运营场景
}

// 淘宝客-店铺搜索
type ShopGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkShopGetResponse struct {
		Results struct {
			NTbkShop []struct {
				UserID     int    `json:"user_id"`     //卖家ID
				ShopTitle  string `json:"shop_title"`  //店铺名称
				ShopType   string `json:"shop_type"`   //店铺类型，B：天猫，C：淘宝
				SellerNick string `json:"seller_nick"` //卖家昵称
				PictURL    string `json:"pict_url"`    //店标图片
				ShopURL    string `json:"shop_url"`    //店铺地址
			} `json:"n_tbk_shop"`
		} `json:"results"`
		TotalResults int `json:"total_results"` //搜索到符合条件的结果总数
	} `json:"tbk_shop_get_response"`
}

// 淘宝客-推广者-店铺搜索
// https://open.taobao.com/api.htm?docId=24521&docType=2&scopeId=16516
type ShopGetParams struct {
	Fields              string `json:"fields,omitempty"`                //	user_id,shop_title,shop_type,seller_nick,pict_url,shop_url	需返回的字段列表
	Q                   string `json:"q,omitempty"`                     //	女装	查询词
	Sort                string `json:"sort,omitempty"`                  //	commission_rate_des	排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction）
	IsTmaLl             bool   `json:"is_tmall"`                        //是否商城的店铺，设置为true表示该是属于淘宝商城的店铺，设置为false或不设置表示不判断这个属性
	StartCredit         int    `json:"start_credit,omitempty"`          //	信用等级下限，1~20
	EndCredit           int    `json:"end_credit,omitempty"`            //	信用等级上限，1~20
	StartCommissionRate int    `json:"start_commission_rate,omitempty"` //	淘客佣金比率下限，1~10000
	EndCommissionRate   int    `json:"end_commission_rate,omitempty"`   //	淘客佣金比率上限，1~10000
	StartTotalAction    int    `json:"start_total_action,omitempty"`    //	店铺商品总数下限
	EndTotalAction      int    `json:"end_total_action,omitempty"`      //店铺商品总数上限
	StartAuctionCount   int    `json:"start_auction_count,omitempty"`   //累计推广商品下限
	EndAuctionCount     int    `json:"end_auction_count,omitempty"`     //	累计推广商品上限
	Platform            int    `json:"platform,omitempty"`              //	链接形式：1：PC，2：无线，默认：１
	PageNo              int    `json:"page_no,omitempty"`               //	第几页，默认1，1~100
	PageSize            int    `json:"page_size,omitempty"`             //	页大小，默认20，1~100
}

// 淘宝客-公用-链接解析出商品id
type ItemClickExtractResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkItemClickExtractResponse struct {
		ItemID  string `json:"item_id"`  //商品id
		OpenIid string `json:"open_iid"` //商品混淆id
	} `json:"tbk_item_click_extract_response"`
}

// 淘宝客-公用-链接解析出商品id
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.3925669aFLMnkJ&source=search&docId=28156&docType=2
type ItemClickExtractParams struct {
	ClickURL string `json:"click_url,omitempty"` //长链接或短链接
}

// 淘宝客-公用-长链转短链
type SpreadGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkSpreadGetResponse struct {
		Results struct {
			TbkSpread []struct {
				Content string `json:"content"` //播形式, 目前只支持短链接
				ErrMsg  string `json:"err_msg"` //调用错误信息；由于是批量接口，请重点关注每条请求返回的结果，如果非OK，则说明该结果对应的content不正常，请酌情处理;
			} `json:"tbk_spread"`
		} `json:"results"`
		TotalResults int `json:"total_results"`
	} `json:"tbk_spread_get_response"`
}

// 淘宝客-公用-长链转短链
// https://open.taobao.com/api.htm?docId=27832&docType=2&scopeId=12340
type SpreadGetParams struct {
	SpreadGetRequests []SpreadGetRequests `json:"requests"`
}

type SpreadGetRequests struct {
	URL string `json:"url"` //原始url, 只支持uland.taobao.com，s.click.taobao.com， ai.taobao.com，temai.taobao.com的域名转换，否则判错
}

// 淘宝客-官方活动转链
type ActivityInfoGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkActivityInfoGetResponse struct {
		Data struct {
			WxQrcodeURL       string `json:"wx_qrcode_url"`       //【本地化】微信推广二维码地址
			ClickURL          string `json:"click_url"`           //淘客推广长链
			ShortClickURL     string `json:"short_click_url"`     //淘客推广短链
			TerminalType      string `json:"terminal_type"`       //投放平台, 1-PC 2-无线
			MaterialOssURL    string `json:"material_oss_url"`    //物料素材下载地址
			PageName          string `json:"page_name"`           //会场名称
			PageStartTime     string `json:"page_start_time"`     //活动开始时间
			PageEndTime       string `json:"page_end_time"`       //活动结束时间
			WxMinIProGramPath string `json:"wx_miniprogram_path"` //【本地化】微信小程序路径
		} `json:"data"`
	} `json:"tbk_activity_info_get_response"` //推广者接口数据
	TbkScActivityInfoGetResponse struct {
		Data struct {
			WxQrcodeURL       string `json:"wx_qrcode_url"`       //【本地化】微信推广二维码地址
			ClickURL          string `json:"click_url"`           //淘客推广长链
			ShortClickURL     string `json:"short_click_url"`     //淘客推广短链
			TerminalType      string `json:"terminal_type"`       //投放平台, 1-PC 2-无线
			MaterialOssURL    string `json:"material_oss_url"`    //物料素材下载地址
			PageName          string `json:"page_name"`           //会场名称
			PageStartTime     string `json:"page_start_time"`     //活动开始时间
			PageEndTime       string `json:"page_end_time"`       //活动结束时间
			WxMinIProGramPath string `json:"wx_miniprogram_path"` //【本地化】微信小程序路径
		} `json:"data"`
	} `json:"tbk_sc_activity_info_get_response"` //服务商接口数据
}

// 淘宝客-服务商-官方活动转链
// https://open.taobao.com/api.htm?docId=48417&docType=2&scopeId=18353
// 淘宝客-推广者-官方活动转链
// https://open.taobao.com/api.htm?docId=48340&docType=2&scopeId=18294
type ActivityInfoGetParams struct {
	SiteID             string `json:"site_id,omitempty"`              //源广告位ID(pid中mm_1_2_3)中第2位(服务商必须传入)
	AdZoneID           string `json:"adzone_id,omitempty"`            //来源广告位ID(pid中mm_1_2_3)中第3位
	SubPid             string `json:"sub_pid,omitempty"`              //mm_xxx_xxx_xxx 仅三方分成场景使用
	RelationID         int64  `json:"relation_id,omitempty"`          //渠道关系id
	ActivityMaterialID string `json:"activity_material_id,omitempty"` //官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取
	UnionID            string `json:"union_id,omitempty"`             //自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
}

// 淘宝客-查询超级红包发放个数
type VegasSendReportResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkDgVegasSendReportResponse struct {
		Result struct {
			Success bool `json:"success"` //是否成功
			Model   struct {
				RelationRptList struct {
					RightsSendRelationRptDto []struct {
						BizDate    string `json:"biz_date"`    //统计日期
						RelationID int64  `json:"relation_id"` //渠道关系id
						FundNum    int    `json:"fund_num"`    //红包发放数量
					} `json:"rights_send_relation_rpt_dto"`
				} `json:"relation_rpt_list"` //渠道关系id的发放数据的数据集
			} `json:"model"`
			MsgInfo string `json:"msg_info"`
			MsgCode string `json:"msg_code"`
		} `json:"result"`
	} `json:"tbk_dg_vegas_send_report_response"` //推广者接口数据
	TbkScVegasSendReportResponse struct {
		Result struct {
			Success bool `json:"success"` //是否成功
			Model   struct {
				RelationRptList struct {
					RightsSendRelationRptDto []struct {
						BizDate    string `json:"biz_date"`    //统计日期
						RelationID int64  `json:"relation_id"` //渠道关系id
						FundNum    int    `json:"fund_num"`    //红包发放数量
					} `json:"rights_send_relation_rpt_dto"`
				} `json:"relation_rpt_list"` //渠道关系id的发放数据的数据集
			} `json:"model"`
			MsgInfo string `json:"msg_info"`
			MsgCode string `json:"msg_code"`
		} `json:"result"`
	} `json:"tbk_sc_vegas_send_report_response"` //服务商接口数据
}

// 淘宝客-服务商-查询超级红包发放个数
// https://open.taobao.com/api.htm?docId=47590&docType=2&scopeId=17875
// 淘宝客-推广者-查询超级红包发放个数
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.78e5669aWq31Sg&source=search&docId=47593&docType=2
type VegasSendReportParams struct {
	BizDate    string `json:"biz_date,omitempty"`    //	20191201	统计日期
	RelationID int64  `json:"relation_id,omitempty"` //渠道关系id
	ActivityID int    `json:"activity_id,omitempty"` //	2020双11大促活动id：1306
	PageNo     int    `json:"page_no,omitempty"`     //	页码
	PageSize   int    `json:"page_size,omitempty"`   //	每页大小
}

// 淘宝客-聚划算商品获取
type JuItemsSearchResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	JuItemsSearchResponse struct {
		Result struct {
			CurrentPage int `json:"current_page"` //当前页码
			Extend      struct {
				Empty bool `json:"empty"`
			} `json:"extend"` //扩展属性
			ModelList struct {
				Items []struct {
					UspDescList struct {
						String []string `json:"string"`
					} `json:"usp_desc_list"` //卖点描述
					TbFirstCatID    int    `json:"tb_first_cat_id"` //淘宝类目id
					OrigPrice       string `json:"orig_price"`      //原价
					ItemID          int64  `json:"item_id"`
					ShowEndTime     int64  `json:"show_end_time"`     //展示结束时间
					PcURL           string `json:"pc_url"`            //pc链接
					PlatformID      int    `json:"platform_id"`       //频道id
					JuID            int64  `json:"ju_id"`             //聚划算id
					PicURLForWL     string `json:"pic_url_for_w_l"`   //无线主图
					OnlineStartTime int64  `json:"online_start_time"` //开团时间
					CategoryName    string `json:"category_name"`     //类目名称
					ActPrice        string `json:"act_price"`         //聚划算价格，单位分
					Title           string `json:"title"`             //商品标题
					WapURL          string `json:"wap_url"`           //无线链接
					ItemUspList     struct {
						String []string `json:"string"`
					} `json:"item_usp_list"` //商品卖点
					ShowStartTime int64  `json:"show_start_time"` //开始展示时间
					OnlineEndTime int64  `json:"online_end_time"` //开团结束时间
					PicURLForPC   string `json:"pic_url_for_p_c"` //pc主图
					PriceUspList  struct {
						String []string `json:"string"`
					} `json:"price_usp_list"` //价格卖点
					PayPostage bool `json:"pay_postage"` //是否包邮
				} `json:"items"`
			} `json:"model_list"` //商品数据
			MsgCode     string `json:"msg_code"`
			MsgInfo     string `json:"msg_info"`   //错误信息
			PageSize    int    `json:"page_size"`  //一页大小
			Success     bool   `json:"success"`    //请求是否成功
			TotalItem   int    `json:"total_item"` //商品总数
			TotalPage   int    `json:"total_page"` //总页数
			TrackParams struct {
				Empty bool `json:"empty"`
			} `json:"track_params"` //埋点信息
		} `json:"result"`
	} `json:"ju_items_search_response"`
}

// 淘宝客-推广者-聚划算商品获取
// https://open.taobao.com/api.htm?docId=28762&docType=2&scopeId=16517
type JuItemsSearchParams struct {
	ParamTopItemQuery ParamTopItemQuery `json:"param_top_item_query"`
}

type ParamTopItemQuery struct {
	CurrentPage      int         `json:"current_page,omitempty"`       //页码,必传
	PageSize         int         `json:"page_size,omitempty"`          //一页大小,必传
	PID              string      `json:"pid,omitempty"`                //媒体pid,必传
	Postage          interface{} `json:"postage"`                      //是否包邮,可不传
	Status           int         `json:"status,omitempty"`             //状态，预热：1，正在进行中：2,可不传
	TaoBaoCategoryID int         `json:"taobao_category_id,omitempty"` //淘宝类目id,可不传
	Word             string      `json:"word,omitempty"`               //搜索关键词,可不传
}

// 淘宝客-物料精选
type OptimusMaterialResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkDgOptimusMaterialResponse struct {
		ResultList struct {
			MapData []struct {
				CouponAmount int `json:"coupon_amount"` //优惠券（元） 若属于预售商品，该优惠券付尾款可用，付定金不可用
				SmallImages  struct {
					String []string `json:"string"`
				} `json:"small_images"` //商品信息-商品小图列表
				ShopTitle            string `json:"shop_title"`              //店铺信息-店铺名称
				CategoryID           int    `json:"category_id"`             //商品信息-叶子类目id
				CouponStartFee       string `json:"coupon_start_fee"`        //优惠券信息-优惠券起用门槛，满X元可用。如：满299元减20元
				ItemID               int64  `json:"item_id"`                 //商品信息-宝贝id
				CouponTotalCount     int    `json:"coupon_total_count"`      //优惠券信息-优惠券总量
				UserType             int    `json:"user_type"`               //店铺信息-卖家类型，0表示集市，1表示商城
				ZkFinalPrice         string `json:"zk_final_price"`          //折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价
				CouponRemainCount    int    `json:"coupon_remain_count"`     //优惠券信息-优惠券剩余量
				CommissionRate       string `json:"commission_rate"`         //商品信息-佣金比率(%)
				CouponStartTime      string `json:"coupon_start_time"`       //优惠券信息-优惠券开始时间
				Title                string `json:"title"`                   //商品信息-商品标题
				ItemDescription      string `json:"item_description"`        //商品信息-宝贝描述（推荐理由,不一定有）
				SellerID             int    `json:"seller_id"`               //店铺信息-卖家id
				Volume               int    `json:"volume"`                  //商品信息-30天销量
				CouponEndTime        string `json:"coupon_end_time"`         //优惠券信息-优惠券结束时间
				CouponClickURL       string `json:"coupon_click_url"`        //链接-宝贝+券二合一页面链接(该字段废弃，请勿再用)
				PictURL              string `json:"pict_url"`                //商品信息-商品主图
				ClickURL             string `json:"click_url"`               //链接-宝贝推广链接
				Stock                int    `json:"stock"`                   //拼团专用-拼团剩余库存
				SellNum              int    `json:"sell_num"`                //拼团专用-拼团已售数量
				TotalStock           int    `json:"total_stock"`             //拼团专用-拼团库存数量
				OeTime               string `json:"oetime"`                  //拼团专用-拼团结束时间
				OsTime               string `json:"ostime"`                  //拼团专用-拼团开始时间
				JddNum               int    `json:"jdd_num"`                 //拼团专用-拼团几人团
				JddPrice             string `json:"jdd_price"`               //拼团专用-拼团拼成价，单位元
				OrigPrice            string `json:"orig_price"`              //拼团专用-拼团一人价（原价)，单位元
				LevelOneCategoryName string `json:"level_one_category_name"` //商品信息-一级类目名称
				LevelOneCategoryID   int    `json:"level_one_category_id"`   //商品信息-一级类目ID
				CategoryName         string `json:"category_name"`           //商品信息-叶子类目名称
				WhiteImage           string `json:"white_image"`             //商品信息-商品白底图
				ShortTitle           string `json:"short_title"`             //商品信息-商品短标题
				WordList             struct {
					WordMapData []struct {
						URL  string `json:"url"`  //链接-商品相关关联词落地页地址
						Word string `json:"word"` //商品相关的关联词
					} `json:"word_map_data"`
				} `json:"word_list"` //商品信息-商品关联词
				TmaLlPlayActivityInfo      string `json:"tmall_play_activity_info"`       //营销-天猫营销玩法
				UvSumPreSale               int    `json:"uv_sum_pre_sale"`                //商品信息-预售数量
				XID                        string `json:"x_id"`                           //物料块id(测试中请勿使用)
				NewUserPrice               string `json:"new_user_price"`                 //商品信息-新人价
				CouponInfo                 string `json:"coupon_info"`                    //优惠券信息-优惠券满减信息
				CouponShareURL             string `json:"coupon_share_url"`               //链接-宝贝+券二合一页面链接
				Nick                       string `json:"nick"`                           //店铺信息-卖家昵称
				ReservePrice               string `json:"reserve_price"`                  //商品信息-一口价
				JuOnlineEndTime            string `json:"ju_online_end_time"`             //聚划算信息-聚淘结束时间
				JuOnlineStartTime          string `json:"ju_online_start_time"`           //聚划算信息-聚淘开始时间
				MaoChaoPlayEndTime         string `json:"maochao_play_end_time"`          //猫超玩法信息-活动结束时间，精确到毫秒
				MaoChaoPlayStartTime       string `json:"maochao_play_start_time"`        //猫超玩法信息-活动开始时间，精确到毫秒
				MaoChaoPlayConditions      string `json:"maochao_play_conditions"`        //猫超玩法信息-折扣条件，价格百分数存储，件数按数量存储。可以有多个折扣条件，与折扣字段对应，','分割
				MaoChaoPlayDiscounts       string `json:"maochao_play_discounts"`         //猫超玩法信息-折扣，折扣按照百分数存储，其余按照单位分存储。可以有多个折扣，','分割
				MaoChaoPlayDiscountType    string `json:"maochao_play_discount_type"`     //猫超玩法信息-玩法类型，2:折扣(满n件折扣),5:减钱(满n元减m元)
				MaoChaoPlayFreePostFee     string `json:"maochao_play_free_post_fee"`     //猫超玩法信息-当前是否包邮，1:是，0:否
				MultiCouponZkRate          string `json:"multi_coupon_zk_rate"`           //多件券优惠比例
				PriceAfterMultiCoupon      string `json:"price_after_multi_coupon"`       //多件券件单价
				MultiCouponItemCount       string `json:"multi_coupon_item_count"`        //多件券单品件数
				LockRate                   string `json:"lock_rate"`                      //锁住的佣金率
				LockRateEndTime            int64  `json:"lock_rate_end_time"`             //锁佣结束时间
				LockRateStartTime          int64  `json:"lock_rate_start_time"`           //锁佣开始时间
				PromotionType              string `json:"promotion_type"`                 //满减满折的类型（1. 满减 2. 满折）
				PromotionInfo              string `json:"promotion_info"`                 //满减满折信息
				PromotionDiscount          string `json:"promotion_discount"`             //满减满折门槛（满2件打5折中值为2；满300减20中值为300）
				PromotionCondition         string `json:"promotion_condition"`            //满减满折优惠（满2件打5折中值为5；满300减20中值为20）
				PresaleDiscountFeeText     string `json:"presale_discount_fee_text"`      //预售商品-优惠信息
				PresaleTailEndTime         int64  `json:"presale_tail_end_time"`          //预售商品-付尾款结束时间（毫秒）
				PresaleTailStartTime       int64  `json:"presale_tail_start_time"`        //预售商品-付尾款开始时间（毫秒）
				PresaleEndTime             int64  `json:"presale_end_time"`               //预售商品-付定金结束时间（毫秒）
				PresaleStartTime           int64  `json:"presale_start_time"`             //预售商品-付定金开始时间（毫秒）
				PresaleDeposit             string `json:"presale_deposit"`                //预售商品-定金（元）
				YsYlTljUseStartTime        string `json:"ysyl_tlj_use_start_time"`        //预售有礼-淘礼金使用开始时间
				YsYlCommissionRate         string `json:"ysyl_commission_rate"`           //预售有礼-佣金比例（ 预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=9334376 ）
				YsYlTljSendTime            string `json:"ysyl_tlj_send_time"`             //预售有礼-淘礼金发放时间
				YsYlTljFace                string `json:"ysyl_tlj_face"`                  //预售有礼-预估淘礼金（元）
				YsYlClickURL               string `json:"ysyl_click_url"`                 //预售有礼-推广链接
				YsYlTljUseEndTime          string `json:"ysyl_tlj_use_end_time"`          //预售有礼-淘礼金使用结束时间
				JuPlayEndTime              int64  `json:"ju_play_end_time"`               //聚划算满减 -结束时间（毫秒）
				JuPlayStartTime            int64  `json:"ju_play_start_time"`             //聚划算满减 -开始时间（毫秒）
				PlayInfo                   string `json:"play_info"`                      //1聚划算满减：满N件减X元，满N件X折，满N件X元） 2天猫限时抢：前N分钟每件X元，前N分钟满N件每件X元，前N件每件X元）
				TmaLlPlayActivityEndTime   int64  `json:"tmall_play_activity_end_time"`   //天猫限时抢可售 -结束时间（毫秒）
				TmaLlPlayActivityStartTime int64  `json:"tmall_play_activity_start_time"` //天猫限时抢可售 -开始时间（毫秒）
				JuPreShowEndTime           string `json:"ju_pre_show_end_time"`           //聚划算信息-商品预热开始时间（毫秒）
				JuPreShowStartTime         string `json:"ju_pre_show_start_time"`         //聚划算信息-商品预热结束时间（毫秒）
				FavoritesInfo              struct {
					TotalCount    int `json:"total_count"` //选品库总数量
					FavoritesList struct {
						FavoritesDetail []struct {
							FavoritesID    int    `json:"favorites_id"`    //选品库id
							FavoritesTitle string `json:"favorites_title"` //选品库标题
						} `json:"favorites_detail"` //选品库详细信息
					} `json:"favorites_list"`
				} `json:"favorites_info"` //选品库信息
				SalePrice            string `json:"sale_price"`             //活动价
				KuaDianPromotionInfo string `json:"kuadian_promotion_info"` //跨店满减信息
				SubTitle             string `json:"sub_title"`              //商品子标题
				JhsPriceUspList      string `json:"jhs_price_usp_list"`     //聚划算商品价格卖点描述
				TqgOnlineEndTime     string `json:"tqg_online_end_time"`    //淘抢购商品专用-结束时间
				TqgOnlineStartTime   string `json:"tqg_online_start_time"`  //淘抢购商品专用-开团时间
				TqgSoldCount         int    `json:"tqg_sold_count"`         //淘抢购商品专用-已抢购数量
				TqgTotalCount        int    `json:"tqg_total_count"`        //淘抢购商品专用-总库存
				SuperiorBrand        string `json:"superior_brand"`         //是否品牌精选，0不是，1是
			} `json:"map_data"`
		} `json:"result_list"`
		IsDefault  string `json:"is_default"`  //推荐信息-是否抄底
		TotalCount int    `json:"total_count"` //商品总数-目前只有全品库商品查询有该字段
	} `json:"tbk_dg_optimus_material_response"` //推广者接口数据
	TbkScOptimusMaterialResponse struct {
		ResultList struct {
			MapData []struct {
				CouponAmount int `json:"coupon_amount"` //优惠券（元） 若属于预售商品，该优惠券付尾款可用，付定金不可用
				SmallImages  struct {
					String []string `json:"string"`
				} `json:"small_images"` //商品信息-商品小图列表
				ShopTitle            string `json:"shop_title"`              //店铺信息-店铺名称
				CategoryID           int    `json:"category_id"`             //商品信息-叶子类目id
				CouponStartFee       string `json:"coupon_start_fee"`        //优惠券信息-优惠券起用门槛，满X元可用。如：满299元减20元
				ItemID               int64  `json:"item_id"`                 //商品信息-宝贝id
				CouponTotalCount     int    `json:"coupon_total_count"`      //优惠券信息-优惠券总量
				UserType             int    `json:"user_type"`               //店铺信息-卖家类型，0表示集市，1表示商城
				ZkFinalPrice         string `json:"zk_final_price"`          //折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价
				CouponRemainCount    int    `json:"coupon_remain_count"`     //优惠券信息-优惠券剩余量
				CommissionRate       string `json:"commission_rate"`         //商品信息-佣金比率(%)
				CouponStartTime      string `json:"coupon_start_time"`       //优惠券信息-优惠券开始时间
				Title                string `json:"title"`                   //商品信息-商品标题
				ItemDescription      string `json:"item_description"`        //商品信息-宝贝描述（推荐理由,不一定有）
				SellerID             int    `json:"seller_id"`               //店铺信息-卖家id
				Volume               int    `json:"volume"`                  //商品信息-30天销量
				CouponEndTime        string `json:"coupon_end_time"`         //优惠券信息-优惠券结束时间
				CouponClickURL       string `json:"coupon_click_url"`        //链接-宝贝+券二合一页面链接(该字段废弃，请勿再用)
				PictURL              string `json:"pict_url"`                //商品信息-商品主图
				ClickURL             string `json:"click_url"`               //链接-宝贝推广链接
				Stock                int    `json:"stock"`                   //拼团专用-拼团剩余库存
				SellNum              int    `json:"sell_num"`                //拼团专用-拼团已售数量
				TotalStock           int    `json:"total_stock"`             //拼团专用-拼团库存数量
				OeTime               string `json:"oetime"`                  //拼团专用-拼团结束时间
				OsTime               string `json:"ostime"`                  //拼团专用-拼团开始时间
				JddNum               int    `json:"jdd_num"`                 //拼团专用-拼团几人团
				JddPrice             string `json:"jdd_price"`               //拼团专用-拼团拼成价，单位元
				OrigPrice            string `json:"orig_price"`              //拼团专用-拼团一人价（原价)，单位元
				LevelOneCategoryName string `json:"level_one_category_name"` //商品信息-一级类目名称
				LevelOneCategoryID   int    `json:"level_one_category_id"`   //商品信息-一级类目ID
				CategoryName         string `json:"category_name"`           //商品信息-叶子类目名称
				WhiteImage           string `json:"white_image"`             //商品信息-商品白底图
				ShortTitle           string `json:"short_title"`             //商品信息-商品短标题
				WordList             struct {
					WordMapData []struct {
						URL  string `json:"url"`  //链接-商品相关关联词落地页地址
						Word string `json:"word"` //商品相关的关联词
					} `json:"word_map_data"`
				} `json:"word_list"` //商品信息-商品关联词
				TmaLlPlayActivityInfo      string `json:"tmall_play_activity_info"`       //营销-天猫营销玩法
				UvSumPreSale               int    `json:"uv_sum_pre_sale"`                //商品信息-预售数量
				XID                        string `json:"x_id"`                           //物料块id(测试中请勿使用)
				NewUserPrice               string `json:"new_user_price"`                 //商品信息-新人价
				CouponInfo                 string `json:"coupon_info"`                    //优惠券信息-优惠券满减信息
				CouponShareURL             string `json:"coupon_share_url"`               //链接-宝贝+券二合一页面链接
				Nick                       string `json:"nick"`                           //店铺信息-卖家昵称
				ReservePrice               string `json:"reserve_price"`                  //商品信息-一口价
				JuOnlineEndTime            string `json:"ju_online_end_time"`             //聚划算信息-聚淘结束时间
				JuOnlineStartTime          string `json:"ju_online_start_time"`           //聚划算信息-聚淘开始时间
				MaoChaoPlayEndTime         string `json:"maochao_play_end_time"`          //猫超玩法信息-活动结束时间，精确到毫秒
				MaoChaoPlayStartTime       string `json:"maochao_play_start_time"`        //猫超玩法信息-活动开始时间，精确到毫秒
				MaoChaoPlayConditions      string `json:"maochao_play_conditions"`        //猫超玩法信息-折扣条件，价格百分数存储，件数按数量存储。可以有多个折扣条件，与折扣字段对应，','分割
				MaoChaoPlayDiscounts       string `json:"maochao_play_discounts"`         //猫超玩法信息-折扣，折扣按照百分数存储，其余按照单位分存储。可以有多个折扣，','分割
				MaoChaoPlayDiscountType    string `json:"maochao_play_discount_type"`     //猫超玩法信息-玩法类型，2:折扣(满n件折扣),5:减钱(满n元减m元)
				MaoChaoPlayFreePostFee     string `json:"maochao_play_free_post_fee"`     //猫超玩法信息-当前是否包邮，1:是，0:否
				MultiCouponZkRate          string `json:"multi_coupon_zk_rate"`           //多件券优惠比例
				PriceAfterMultiCoupon      string `json:"price_after_multi_coupon"`       //多件券件单价
				MultiCouponItemCount       string `json:"multi_coupon_item_count"`        //多件券单品件数
				LockRate                   string `json:"lock_rate"`                      //锁住的佣金率
				LockRateEndTime            int64  `json:"lock_rate_end_time"`             //锁佣结束时间
				LockRateStartTime          int64  `json:"lock_rate_start_time"`           //锁佣开始时间
				PromotionType              string `json:"promotion_type"`                 //满减满折的类型（1. 满减 2. 满折）
				PromotionInfo              string `json:"promotion_info"`                 //满减满折信息
				PromotionDiscount          string `json:"promotion_discount"`             //满减满折门槛（满2件打5折中值为2；满300减20中值为300）
				PromotionCondition         string `json:"promotion_condition"`            //满减满折优惠（满2件打5折中值为5；满300减20中值为20）
				PresaleDiscountFeeText     string `json:"presale_discount_fee_text"`      //预售商品-优惠信息
				PresaleTailEndTime         int64  `json:"presale_tail_end_time"`          //预售商品-付尾款结束时间（毫秒）
				PresaleTailStartTime       int64  `json:"presale_tail_start_time"`        //预售商品-付尾款开始时间（毫秒）
				PresaleEndTime             int64  `json:"presale_end_time"`               //预售商品-付定金结束时间（毫秒）
				PresaleStartTime           int64  `json:"presale_start_time"`             //预售商品-付定金开始时间（毫秒）
				PresaleDeposit             string `json:"presale_deposit"`                //预售商品-定金（元）
				YsYlTljUseStartTime        string `json:"ysyl_tlj_use_start_time"`        //预售有礼-淘礼金使用开始时间
				YsYlCommissionRate         string `json:"ysyl_commission_rate"`           //预售有礼-佣金比例（ 预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=9334376 ）
				YsYlTljSendTime            string `json:"ysyl_tlj_send_time"`             //预售有礼-淘礼金发放时间
				YsYlTljFace                string `json:"ysyl_tlj_face"`                  //预售有礼-预估淘礼金（元）
				YsYlClickURL               string `json:"ysyl_click_url"`                 //预售有礼-推广链接
				YsYlTljUseEndTime          string `json:"ysyl_tlj_use_end_time"`          //预售有礼-淘礼金使用结束时间
				JuPlayEndTime              int64  `json:"ju_play_end_time"`               //聚划算满减 -结束时间（毫秒）
				JuPlayStartTime            int64  `json:"ju_play_start_time"`             //聚划算满减 -开始时间（毫秒）
				PlayInfo                   string `json:"play_info"`                      //1聚划算满减：满N件减X元，满N件X折，满N件X元） 2天猫限时抢：前N分钟每件X元，前N分钟满N件每件X元，前N件每件X元）
				TmaLlPlayActivityEndTime   int64  `json:"tmall_play_activity_end_time"`   //天猫限时抢可售 -结束时间（毫秒）
				TmaLlPlayActivityStartTime int64  `json:"tmall_play_activity_start_time"` //天猫限时抢可售 -开始时间（毫秒）
				JuPreShowEndTime           string `json:"ju_pre_show_end_time"`           //聚划算信息-商品预热开始时间（毫秒）
				JuPreShowStartTime         string `json:"ju_pre_show_start_time"`         //聚划算信息-商品预热结束时间（毫秒）
				FavoritesInfo              struct {
					TotalCount    int `json:"total_count"` //选品库总数量
					FavoritesList struct {
						FavoritesDetail []struct {
							FavoritesID    int    `json:"favorites_id"`    //选品库id
							FavoritesTitle string `json:"favorites_title"` //选品库标题
						} `json:"favorites_detail"` //选品库详细信息
					} `json:"favorites_list"`
				} `json:"favorites_info"` //选品库信息
				SalePrice            string `json:"sale_price"`             //活动价
				KuaDianPromotionInfo string `json:"kuadian_promotion_info"` //跨店满减信息
				SubTitle             string `json:"sub_title"`              //商品子标题
				JhsPriceUspList      string `json:"jhs_price_usp_list"`     //聚划算商品价格卖点描述
				TqgOnlineEndTime     string `json:"tqg_online_end_time"`    //淘抢购商品专用-结束时间
				TqgOnlineStartTime   string `json:"tqg_online_start_time"`  //淘抢购商品专用-开团时间
				TqgSoldCount         int    `json:"tqg_sold_count"`         //淘抢购商品专用-已抢购数量
				TqgTotalCount        int    `json:"tqg_total_count"`        //淘抢购商品专用-总库存

				SuperiorBrand string `json:"superior_brand"` //是否品牌精选，0不是，1是
			} `json:"map_data"`
		} `json:"result_list"`
		IsDefault  string `json:"is_default"`  //推荐信息-是否抄底
		TotalCount int    `json:"total_count"` //商品总数-目前只有全品库商品查询有该字段
	} `json:"tbk_sc_optimus_material_response"` //服务商接口数据
}

// 淘宝客-服务商-物料精选
// https://open.taobao.com/api.htm?docId=37884&docType=2&scopeId=16287
// 淘宝客-推广者-物料精选
// https://open.taobao.com/api.htm?docId=33947&docType=2&scopeId=16518
type OptimusMaterialParams struct {
	PageSize      int    `json:"page_size,omitempty"`      //页大小，默认20，1~100
	SiteID        string `json:"site_id,omitempty"`        //源广告位ID(pid中mm_1_2_3)中第2位(服务商必须传入)
	AdZoneID      string `json:"adzone_id,omitempty"`      //来源广告位ID(pid中mm_1_2_3)中第3位
	PageMo        int    `json:"page_no,omitempty"`        //第几页，默认：1
	MaterialID    int    `json:"material_id,omitempty"`    //官方的物料Id(详细物料id见：https://market.m.taobao.com/app/qn/toutiao-new/index-pc.html#/detail/10628875?_k=gpov9a)
	DeviceValue   string `json:"device_value,omitempty"`   //智能匹配-设备号加密后的值（MD5加密需32位小写），类型为OAID时传原始OAID值
	DeviceEncrypt string `json:"device_encrypt,omitempty"` //智能匹配-设备号加密类型：MD5，类型为OAID时不传
	DeviceType    string `json:"device_type,omitempty"`    //智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
	ContentID     int    `json:"content_id,omitempty"`     //内容专用-内容详情ID
	ContentSource string `json:"content_source,omitempty"` //内容专用-内容渠道信息
	ItemID        int64  `json:"item_id,omitempty"`        //商品ID，用于相似商品推荐
	FavoritesID   string `json:"favorites_id,omitempty"`   //选品库投放id
}

// 淘宝客-新用户订单明细查询
type NewUserOrderGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkDgNewUserOrderGetResponse struct {
		Results struct {
			Data struct {
				Results struct {
					MapData []struct {
						RegisterTime    string `json:"register_time"`      //新注册时间，仅淘宝拉新适用
						BindTime        string `json:"bind_time"`          //当前活动为淘宝拉新活动时，bind_time为新激活时间； 当前活动为支付宝拉新活动时，bind_time为绑定时间。
						BuyTime         string `json:"buy_time"`           //首购时间，仅淘宝，天猫拉新适用
						Status          int    `json:"status"`             //新人状态， 当前活动为淘宝拉新活动时，1: 新注册，2:激活，3:首购，4:确认收货； 当前活动为支付宝实名活动时，1：已绑定，2：拉新成功，3：无效用户；当前活动为支付宝新登活动时，3：手淘首购，4：手淘确认收货；当前活动为天猫拉新活动时，2:已领取，3:已首购，4:已收货
						Mobile          string `json:"mobile"`             //新人手机号
						OrderTkType     int    `json:"order_tk_type"`      //订单淘客类型:1.淘客订单；2.非淘客订单，仅淘宝，天猫拉新适用
						UnionID         string `json:"union_id"`           //分享用户(unionid)，仅淘宝，天猫拉新适用
						MemberID        int    `json:"member_id"`          //来源媒体ID(pid中mm_1_2_3)中第1位
						MemberNick      string `json:"member_nick"`        //来源媒体名称
						SiteID          int    `json:"site_id"`            //来源站点ID(pid中mm_1_2_3)中第2位
						SiteName        string `json:"site_name"`          //来源站点名称
						AdZoneID        int64  `json:"adzone_id"`          //来源广告位ID(pid中mm_1_2_3)中第3位
						AdZoneName      string `json:"adzone_name"`        //来源广告位名称
						TbTradeParentID int    `json:"tb_trade_parent_id"` //淘宝订单id，仅淘宝，天猫拉新适用
						AcceptTime      string `json:"accept_time"`        //确认收货时间，仅天猫拉新适用
						ReceiveTime     string `json:"receive_time"`       //领取红包时间，仅天猫拉新适用
						SuccessTime     string `json:"success_time"`       //拉新成功时间，仅支付宝拉新适用
						ActivityType    string `json:"activity_type"`      //活动类型，taobao-淘宝 alipay-支付宝 tmall-天猫
						ActivityID      string `json:"activity_id"`        //活动id
						BizDate         string `json:"biz_date"`           //日期，格式为"20180202"
						Orders          struct {
							OrderData []struct {
								Commission         string `json:"commission"`           //预估佣金
								ConfirmReceiveTime string `json:"confirm_receive_time"` //收货时间
								PayTime            string `json:"pay_time"`             //支付时间
								OrderNo            string `json:"order_no"`             //订单号
							} `json:"order_data"`
						} `json:"orders"` //复购订单，仅适用于手淘拉新
						BindCardTime  string `json:"bind_card_time"`  //绑卡日期，仅适用于手淘拉新
						LoginTime     string `json:"login_time"`      //loginTime
						IsCardSave    int    `json:"is_card_save"`    //银行卡是否是绑定状态：1-绑定，0-未绑定
						UseRightsTime string `json:"use_rights_time"` //使用权益时间
						GetRightsTime string `json:"get_rights_time"` //领取权益时间
						RelationID    string `json:"relation_id"`     //渠道关系id
					} `json:"map_data"`
				} `json:"results"`
				PageNo   int  `json:"page_no"`   //页码
				PageSize int  `json:"page_size"` //每页大小
				HasNext  bool `json:"has_next"`  //是否有下一页
			} `json:"data"`
		} `json:"results"`
	} `json:"tbk_dg_newuser_order_get_response"` //推广者接口数据
	TbkScNewUserOrderGetResponse struct {
		Results struct {
			Data struct {
				Results struct {
					MapData []struct {
						RegisterTime    string `json:"register_time"`      //新注册时间，仅淘宝拉新适用
						BindTime        string `json:"bind_time"`          //当前活动为淘宝拉新活动时，bind_time为新激活时间； 当前活动为支付宝拉新活动时，bind_time为绑定时间。
						BuyTime         string `json:"buy_time"`           //首购时间，仅淘宝，天猫拉新适用
						Status          int    `json:"status"`             //新人状态， 当前活动为淘宝拉新活动时，1: 新注册，2:激活，3:首购，4:确认收货； 当前活动为支付宝实名活动时，1：已绑定，2：拉新成功，3：无效用户；当前活动为支付宝新登活动时，3：手淘首购，4：手淘确认收货；当前活动为天猫拉新活动时，2:已领取，3:已首购，4:已收货
						Mobile          string `json:"mobile"`             //新人手机号
						OrderTkType     int    `json:"order_tk_type"`      //订单淘客类型:1.淘客订单；2.非淘客订单，仅淘宝，天猫拉新适用
						UnionID         string `json:"union_id"`           //分享用户(unionid)，仅淘宝，天猫拉新适用
						MemberID        int    `json:"member_id"`          //来源媒体ID(pid中mm_1_2_3)中第1位
						MemberNick      string `json:"member_nick"`        //来源媒体名称
						SiteID          int    `json:"site_id"`            //来源站点ID(pid中mm_1_2_3)中第2位
						SiteName        string `json:"site_name"`          //来源站点名称
						AdZoneID        int64  `json:"adzone_id"`          //来源广告位ID(pid中mm_1_2_3)中第3位
						AdZoneName      string `json:"adzone_name"`        //来源广告位名称
						TbTradeParentID int    `json:"tb_trade_parent_id"` //淘宝订单id，仅淘宝，天猫拉新适用
						AcceptTime      string `json:"accept_time"`        //确认收货时间，仅天猫拉新适用
						ReceiveTime     string `json:"receive_time"`       //领取红包时间，仅天猫拉新适用
						SuccessTime     string `json:"success_time"`       //拉新成功时间，仅支付宝拉新适用
						ActivityType    string `json:"activity_type"`      //活动类型，taobao-淘宝 alipay-支付宝 tmall-天猫
						ActivityID      string `json:"activity_id"`        //活动id
						BizDate         string `json:"biz_date"`           //日期，格式为"20180202"
						Orders          struct {
							OrderData []struct {
								Commission         string `json:"commission"`           //预估佣金
								ConfirmReceiveTime string `json:"confirm_receive_time"` //收货时间
								PayTime            string `json:"pay_time"`             //支付时间
								OrderNo            string `json:"order_no"`             //订单号
							} `json:"order_data"`
						} `json:"orders"` //复购订单，仅适用于手淘拉新
						BindCardTime  string `json:"bind_card_time"`  //绑卡日期，仅适用于手淘拉新
						LoginTime     string `json:"login_time"`      //loginTime
						IsCardSave    int    `json:"is_card_save"`    //银行卡是否是绑定状态：1-绑定，0-未绑定
						UseRightsTime string `json:"use_rights_time"` //使用权益时间
						GetRightsTime string `json:"get_rights_time"` //领取权益时间
						RelationID    string `json:"relation_id"`     //渠道关系id
					} `json:"map_data"`
				} `json:"results"`
				PageNo   int  `json:"page_no"`   //页码
				PageSize int  `json:"page_size"` //每页大小
				HasNext  bool `json:"has_next"`  //是否有下一页
			} `json:"data"`
		} `json:"results"`
	} `json:"tbk_sc_newuser_order_get_response"` //服务商接口数据
}

// 淘宝客-服务商-新用户订单明细查询
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.6969669as9PC5i&source=search&docId=33897&docType=2
// 淘宝客-推广者-新用户订单明细查询
// https://open.taobao.com/api.htm?docId=33892&docType=2&scopeId=16188
type NewUserOrderGetParams struct {
	PageSize   int    `json:"page_size,omitempty"`   //	页大小，默认20，1~100
	SiteID     string `json:"site_id,omitempty"`     //源广告位ID(pid中mm_1_2_3)中第2位(服务商必须传入)
	AdZoneID   string `json:"adzone_id,omitempty"`   //	来源广告位ID(pid中mm_1_2_3)中第3位
	PageNo     int    `json:"page_no,omitempty"`     //	页码，默认1
	StartTime  string `json:"start_time,omitempty"`  //	2018-01-24 00:34:05	开始时间，当活动为淘宝活动，表示最早注册时间；当活动为支付宝活动，表示最早绑定时间；当活动为天猫活动，表示最早领取红包时间
	EndTime    string `json:"end_time,omitempty"`    //	2018-01-24 00:34:05	结束时间，当活动为淘宝活动，表示最晚结束时间；当活动为支付宝活动，表示最晚绑定时间；当活动为天猫活动，表示最晚领取红包的时间
	ActivityID string `json:"activity_id,omitempty"` //	活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277

}

// 淘宝客-拉新活动对应数据查询
type NewUserOrderSumResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkDgNewUserOrderSumResponse struct {
		Results struct {
			Data struct {
				PageNo   int  `json:"page_no"`   //页码
				PageSize int  `json:"page_size"` //每页大小
				HasNext  bool `json:"has_next"`  //是否有下一页
				Results  struct {
					Data []struct {
						ActivityID           string `json:"activity_id"`               //活动ID
						BizDate              string `json:"biz_date"`                  //日期
						RegUserCnt           int    `json:"reg_user_cnt"`              //新注册用户数
						LoginUserCnt         int    `json:"login_user_cnt"`            //新激活用户数
						AlipayUserCnt        int    `json:"alipay_user_cnt"`           //首购用户数
						RcvValidUserCnt      int    `json:"rcv_valid_user_cnt"`        //结算有效用户数
						RcvUserCnt           int    `json:"rcv_user_cnt"`              //确认收货用户数
						AlipayUserCpaPreAmt  string `json:"alipay_user_cpa_pre_amt"`   //结算CPA 奖励金额：仅支持member 维度的统计
						BindBuyUserCpaPreAmt string `json:"bind_buy_user_cpa_pre_amt"` //当日激活且首购结算的CPA 金额：仅适用于八天乐，仅支持member维度的统计
						BindBuyValidUserCnt  int    `json:"bind_buy_valid_user_cnt"`   //当日激活且首购的有效用户数：仅适用于八天乐，支持member，adzone维度的统计
						BindCardValidUserCnt int    `json:"bind_card_valid_user_cnt"`  //bindCardValidUserCnt
						ReBuyValidUserCnt    int    `json:"re_buy_valid_user_cnt"`     //reBuyValidUserCnt
						ValidNum             int    `json:"valid_num"`                 //validNum
						RelationID           string `json:"relation_id"`               //渠道关系id
					} `json:"data"`
				} `json:"results"`
			} `json:"data"`
		} `json:"results"`
	} `json:"tbk_dg_newuser_order_sum_response"` //推广者接口数据
	TbkScNewUserOrderSumResponse struct {
		Results struct {
			Data struct {
				PageNo   int  `json:"page_no"`   //页码
				PageSize int  `json:"page_size"` //每页大小
				HasNext  bool `json:"has_next"`  //是否有下一页
				Results  struct {
					Data []struct {
						ActivityID           string `json:"activity_id"`               //活动ID
						BizDate              string `json:"biz_date"`                  //日期
						RegUserCnt           int    `json:"reg_user_cnt"`              //新注册用户数
						LoginUserCnt         int    `json:"login_user_cnt"`            //新激活用户数
						AlipayUserCnt        int    `json:"alipay_user_cnt"`           //首购用户数
						RcvValidUserCnt      int    `json:"rcv_valid_user_cnt"`        //结算有效用户数
						RcvUserCnt           int    `json:"rcv_user_cnt"`              //确认收货用户数
						AlipayUserCpaPreAmt  string `json:"alipay_user_cpa_pre_amt"`   //结算CPA 奖励金额：仅支持member 维度的统计
						BindBuyUserCpaPreAmt string `json:"bind_buy_user_cpa_pre_amt"` //当日激活且首购结算的CPA 金额：仅适用于八天乐，仅支持member维度的统计
						BindBuyValidUserCnt  int    `json:"bind_buy_valid_user_cnt"`   //当日激活且首购的有效用户数：仅适用于八天乐，支持member，adzone维度的统计
						BindCardValidUserCnt int    `json:"bind_card_valid_user_cnt"`  //bindCardValidUserCnt
						ReBuyValidUserCnt    int    `json:"re_buy_valid_user_cnt"`     //reBuyValidUserCnt
						ValidNum             int    `json:"valid_num"`                 //validNum
						RelationID           string `json:"relation_id"`               //渠道关系id
					} `json:"data"`
				} `json:"results"`
			} `json:"data"`
		} `json:"results"`
	} `json:"tbk_sc_newuser_order_sum_response"` //服务商接口数据
}

// 淘宝客-推广者-拉新活动对应数据查询
// https://open.taobao.com/api.htm?docId=36836&docType=2&scopeId=16188
type NewUserOrderSumParams struct {
	PageSize    int    `json:"page_size,omitempty"`    //页大小，默认20，1~100
	AdZoneID    string `json:"adzone_id,omitempty"`    //来源广告位ID(pid中mm_1_2_3)中第3位
	PageNo      int    `json:"page_no,omitempty"`      //页码，默认1
	SiteID      int    `json:"site_id,omitempty"`      //源广告位ID(pid中mm_1_2_3)中第2位(服务商必须传入)
	ActivityID  string `json:"activity_id,omitempty"`  //活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
	SettleMonth string `json:"settle_month,omitempty"` //201807	结算月份
}

//  淘宝客-公用-店铺关联推荐
type ShopRecommendGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkShopRecommendGetResponse struct {
		Results struct {
			NTbkShop []struct {
				UserID     int    `json:"user_id"`     //卖家ID
				ShopTitle  string `json:"shop_title"`  //店铺名称
				ShopType   string `json:"shop_type"`   //店铺类型，B：天猫，C：淘宝
				SellerNick string `json:"seller_nick"` //卖家昵称
				PictURL    string `json:"pict_url"`    //店标图片
				ShopURL    string `json:"shop_url"`    //店铺地址
			} `json:"n_tbk_shop"`
		} `json:"results"`
	} `json:"tbk_shop_recommend_get_response"`
}

// 淘宝客-公用-店铺关联推荐
// https://open.taobao.com/api.htm?docId=24522&docType=2&scopeId=16292
type ShopRecommendGetParams struct {
	Fields   string `json:"fields,omitempty"`   //	user_id,shop_title,shop_type,seller_nick,pict_url,shop_url	需返回的字段列表
	UserID   int    `json:"user_id,omitempty"`  //	卖家Id
	Count    int    `json:"count,omitempty"`    //	返回数量，默认20，最大值40
	Platform int    `json:"platform,omitempty"` //	链接形式：1：PC，2：无线，默认：１
}

// 淘宝客-公用-阿里妈妈推广券详情查询
type TbkCoUonGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkCouponGetResponse struct {
		Data struct {
			CouponStartFee    string `json:"coupon_start_fee"`    //优惠券门槛金额
			CouponRemainCount int    `json:"coupon_remain_count"` //优惠券剩余量
			CouponTotalCount  int    `json:"coupon_total_count"`  //优惠券总量
			CouponEndTime     string `json:"coupon_end_time"`     //优惠券结束时间
			CouponStartTime   string `json:"coupon_start_time"`   //优惠券开始时间
			CouponAmount      string `json:"coupon_amount"`       //优惠券金额
			CouponSrcScene    int    `json:"coupon_src_scene"`    //券类型，1 表示全网公开券，4 表示妈妈渠道券
			CouponType        int    `json:"coupon_type"`         //券属性，0表示店铺券，1表示单品券
			CouponActivityID  string `json:"coupon_activity_id"`  //券ID
		} `json:"data"`
	} `json:"tbk_coupon_get_response"`
}

// 淘宝客-公用-阿里妈妈推广券详情查询
// https://open.taobao.com/api.htm?docId=31106&docType=2&scopeId=16189
type TbkCoUponGetParams struct {
	Me         string `json:"me,omitempty"`          //	带券ID与商品ID的加密串
	ItemID     int64  `json:"item_id,omitempty"`     //商品ID
	ActivityID string `json:"activity_id,omitempty"` //券ID
}

// 淘宝客-店铺链接转换
type ShopConvertResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkShopConvertResponse struct {
		Results struct {
			NTbkShop []struct {
				UserID   int    `json:"user_id"`   //卖家ID
				ClickURL string `json:"click_url"` //淘客地址
			} `json:"n_tbk_shop"`
		} `json:"results"`
	} `json:"tbk_shop_convert_response"` //推广者接口数据
	TbkScShopConvertResponse struct {
		Results struct {
			NTbkShop []struct {
				UserID   int    `json:"user_id"`   //卖家ID
				ClickURL string `json:"click_url"` //淘客地址
			} `json:"n_tbk_shop"`
		} `json:"results"`
	} `json:"tbk_sc_shop_convert_response"` //服务商接口数据
}

// 淘宝客-服务商-店铺链接转换
// https://open.taobao.com/api.htm?docId=43878&docType=2&scopeId=16403
//  淘宝客-推广者-店铺链接转换
//  https://open.taobao.com/api.htm?docId=24523&docType=2&scopeId=11653
type ShopConvertParams struct {
	SiteID   int    `json:"site_id,omitempty"`   //	源广告位ID(pid中mm_1_2_3)中第2位(服务商必须传入)
	Fields   string `json:"fields,omitempty"`    //	user_id,click_url	需返回的字段列表
	UserIDs  string `json:"user_ids,omitempty"`  //	卖家id串，用,分割，从taobao.tbk.shop.get接口获取user_id字段
	Platform int    `json:"platform,omitempty"`  //	链接形式-1：PC，2：无线。默认为１
	AdZoneID string `json:"adzone_id,omitempty"` //	来源广告位ID(pid中mm_1_2_3)中第3位
	UnID     string `json:"unid,omitempty"`      //	(该字段不开放)自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
}

// 淘宝客-商品链接转换
type ItemConvertResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkItemConvertResponse struct {
		Results struct {
			NTbkItem []struct {
				NumIid   int    `json:"num_iid"`   //商品ID
				ClickURL string `json:"click_url"` //淘客地址
			} `json:"n_tbk_item"`
		} `json:"results"`
	} `json:"tbk_item_convert_response"`
}

// 淘宝客-推广者-商品链接转换
// https://open.taobao.com/api.htm?docId=24516&docType=2&scopeId=11653
type ItemConvertParams struct {
	Fields   string `json:"fields,omitempty"`    //	num_iid,click_url	需返回的字段列表
	NumIIDs  string `json:"num_iids,omitempty"`  //	商品ID串，用','分割，从taobao.tbk.item.get接口获取num_iid字段，最大40个
	AdZoneID string `json:"adzone_id,omitempty"` //	来源广告位ID(pid中mm_1_2_3)中第3位
	Platform int    `json:"platform,omitempty"`  //	链接形式：1：PC，2：无线，默认：１
	UnID     string `json:"unid,omitempty"`      //	自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
	Dx       string `json:"dx,omitempty"`        //	1表示商品转通用计划链接，其他值或不传表示转营销计划链接
}

// 淘宝客-超级红包领取状态查询
type VegasSendStatusResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkDgVegasSendStatusResponse struct {
		ResultMsg string `json:"result_msg"`
		Data      struct {
			ResultList struct {
				MapData []struct {
					IsNewUser string `json:"is_new_user"` //若为超红新用户，则返回1 若非超红新用户，则返回0
				} `json:"map_data"`
			} `json:"result_list"`
		} `json:"data"`
	} `json:"tbk_dg_vegas_send_status_response"` //推广者接口数据
	TbkScVegasSendStatusResponse struct {
		ResultMsg string `json:"result_msg"`
		Data      struct {
			ResultList struct {
				MapData []struct {
					IsNewUser string `json:"is_new_user"` //若为超红新用户，则返回1 若非超红新用户，则返回0
				} `json:"map_data"`
			} `json:"result_list"`
		} `json:"data"`
	} `json:"tbk_sc_vegas_send_status_response"` //服务商接口数据
}

// 淘宝客-服务商-超级红包领取状态查询
// https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.3b54669aEtxGpd&source=search&docId=52957&docType=2
// 淘宝客-推广者-超级红包领取状态查询
// https://open.taobao.com/api.htm?docId=52958&docType=2&scopeId=21226
type VegasSendStatusParams struct {
	RelationID  string `json:"relation_id,omitempty"`  //	渠道管理id
	SpecialID   string `json:"special_id,omitempty"`   //	会员运营id
	DeviceValue string `json:"device_value,omitempty"` //加密后的值(ALIPAY_ID除外)，需用MD5加密，32位小写
	DeviceType  string `json:"device_type,omitempty"`  //	入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE 5. ALIPAY_ID
}

// 查询解析淘口令
type TPwdQueryResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	WirelessShareTPwdQueryResponse struct {
		Content     string `json:"content"`       //淘口令-文案
		Title       string `json:"title"`         //淘口令-宝贝
		Price       string `json:"price"`         //如果是宝贝，则为宝贝价格
		PicURL      string `json:"pic_url"`       //图片url
		Suc         bool   `json:"suc"`           //是否成功
		URL         string `json:"url"`           //跳转url(长链)
		NativeURL   string `json:"native_url"`    //nativeUrl
		ThumbPicURL string `json:"thumb_pic_url"` //thumbPicUrl
	} `json:"wireless_share_tpwd_query_response"`
}

// 查询解析淘口令
// https://open.taobao.com/api.htm?docId=32461&docType=2&scopeId=11998
type TPwdQueryParams struct {
	PasswordContent string `json:"password_content,omitempty"` //淘口令
}

// 获取淘宝系统当前时间
// https://open.taobao.com/api.htm?docId=120&docType=2&scopeId=381
type TaoBaoTimeGetResp struct {
	TimeGetResponse struct {
		Time string `json:"time"` //淘宝系统当前时间。格式:yyyy-MM-dd HH:mm:ss
	} `json:"time_get_response"`
}

// 关键词过滤匹配
type KeywordSearchResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	KfcKeywordSearchResponse struct {
		KfcSearchResult struct {
			Matched bool `json:"matched"` //是否匹配到关键词,匹配到则为true.
			//匹配到的关键词的等级，值为null，或为A、B、C、D。
			//当匹配不到关键词时，值为null，否则值为A、B、C、D中的一个。
			//A、B、C、D等级按严重程度从高至低排列。
			Level string `json:"level"`
			//过滤后的文本：
			//当匹配到B等级的词时，文本中的关键词被替换为*号，content即为关键词替换后的文本；
			//其他情况，content始终为null
			Content string `json:"content"`
		} `json:"kfc_search_result"`
	} `json:"kfc_keyword_search_response"`
}

// 关键词过滤匹配
// https://open.taobao.com/api.htm?docId=10385&docType=2&scopeId=381
type KeywordSearchParams struct {
	Nick    string `json:"nick,omitempty"`    //		发布信息的淘宝会员名，可以不传
	Content string `json:"content,omitempty"` //		需要过滤的文本信息
	//应用点，分为一级应用点、二级应用点。其中一级应用点通常是指某一个系统或产品，比如淘宝的商品应用（taobao_auction）；
	//二级应用点，是指一级应用点下的具体的分类，比如商品标题(title)、商品描述(content)。不同的二级应用可以设置不同关键词。
	//这里的apply参数是由一级应用点与二级应用点合起来的字符（一级应用点+"."+二级应用点），如taobao_auction.title。
	//通常apply参数是不需要传递的。如有特殊需求（比如特殊的过滤需求，需要自己维护一套自己词库），需传递此参数。
	Apply string `json:"apply,omitempty"`
}

// 获取ISV发起请求服务器IP
// https://open.taobao.com/api.htm?docId=21784&docType=2&scopeId=381
type AppIPGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	AppIPGetResponse struct {
		IP string `json:"ip"` //ISV发起请求服务器IP
	} `json:"appip_get_response"`
}

// 淘宝openUid 转换
type OpenUIDChangeResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	OpenUIDChangeResponse struct {
		NewOpenUID string `json:"new_open_uid"` //转换到新的openUId
	} `json:"openuid_change_response"`
}

// 淘宝openUid 转换
// https://open.taobao.com/api.htm?docId=23831&docType=2&scopeId=381
type OpenUIDChangeParams struct {
	OpenUID      string `json:"open_uid,omitempty"`       //	openUid
	TargetAppKey string `json:"target_app_key,omitempty"` //	转换到的appkey
}

// 刷新Access Token
type AuthTokenRefreshResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TopAuthTokenRefreshResponse struct {
		TokenResult string `json:"token_result"` //返回的是json信息
	} `json:"top_auth_token_refresh_response"`
}

// 刷新Access Token
// https://open.taobao.com/api.htm?docId=25387&docType=2&scopeId=381
type AuthTokenRefreshParams struct {
	RefreshToken string `json:"refresh_token,omitempty"` //grantType==refresh_token 时需要
}

// 获取Access Token
type AuthTokenCreateResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TopAuthTokenCreateResponse struct {
		TokenResult string `json:"token_result"` //返回的是json信息
	} `json:"top_auth_token_create_response"`
}

// 获取Access Token
// https://open.taobao.com/api.htm?docId=25388&docType=2&scopeId=381
type AuthTokenCreateParams struct {
	Code string `json:"code,omitempty"` //授权code，grantType==authorization_code 时需要
	Uuid string `json:"uuid,omitempty"` // 与生成code的uuid配对
}

// TOPDNS配置
// https://open.taobao.com/api.htm?docId=25414&docType=2&scopeId=381
type HttpDnsGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	HttpDnsGetResponse struct {
		Result string `json:"result"` //HTTP DNS配置信息
	} `json:"httpdns_get_response"`
}

// 获取开放平台出口IP段
// https://open.taobao.com/api.htm?docId=25441&docType=2&scopeId=381
type TopIPOutGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TopIPOutGetResponse struct {
		IPList string `json:"ip_list"` //TOP网关出口IP列表
	} `json:"top_ipout_get_response"`
}

// 获取TOP通道解密秘钥
type TopSecretGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TopSecretGetResponse struct {
		Interval      int    `json:"interval"`       //下次更新秘钥间隔，单位（秒）
		Secret        string `json:"secret"`         //秘钥值
		SecretVersion int    `json:"secret_version"` //秘钥版本号
		MaxInterval   int    `json:"max_interval"`   //最长有效期，容灾使用，单位（秒）
		AppConfig     string `json:"app_config"`     //app配置信息
	} `json:"top_secret_get_response"`
}

// 获取TOP通道解密秘钥
// https://open.taobao.com/api.htm?docId=26567&docType=2&scopeId=381
type TopSecretGetParams struct {
	SecretVersion  int    `json:"secret_version,omitempty"`   //	秘钥版本号
	RandomNum      string `json:"random_num,omitempty"`       //	伪随机数
	CustomerUserID int    `json:"customer_user_id,omitempty"` //	自定义用户id
}

// 注册加密账号
type TopSecretRegisterResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TopSecretRegisterResponse struct {
		Result bool `json:"result"` //返回操作是否成功
	} `json:"top_secret_register_response"`
}

// 注册加密账号
// https://open.taobao.com/api.htm?docId=27385&docType=2&scopeId=381
type TopSecretRegisterParams struct {
	UserID int `json:"user_id,omitempty"` //用户id，保证唯一
}

// sdk信息回调
type SdkFeedbackUploadResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TopSdkFeedbackUploadResponse struct {
		UploadInterval int `json:"upload_interval"` //控制回传间隔（单位：秒）
	} `json:"top_sdk_feedback_upload_response"`
}

// sdk信息回调
// https://open.taobao.com/api.htm?docId=27512&docType=2&scopeId=381
type SdkFeedbackUploadParams struct {
	Type    string `json:"type,omitempty"`    //	1、回传加密信息
	Content string `json:"content,omitempty"` //	具体内容，json形式
}

// 业务文件获取
type FilesGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	FilesGetResponse struct {
		Results struct {
			TopDownloadRecordDo []struct {
				URL     string `json:"url"`     //下载链接
				Created string `json:"created"` //文件创建时间
				Status  int    `json:"status"`  //下载链接状态。1:未下载。2:已下载
			} `json:"top_download_record_do"`
		} `json:"results"`
	} `json:"files_get_response"`
}

// 业务文件获取
// https://open.taobao.com/api.htm?docId=32298&docType=2&scopeId=381
type FilesGetParams struct {
	StartDate string `json:"start_date,omitempty"` //2017-11-11 00:00:00	搜索开始时间
	EndDate   string `json:"end_date,omitempty"`   //2017-11-12 00:00:00	搜索结束时间
	Status    int    `json:"status,omitempty"`     //下载链接状态。1:未下载。2:已下载
}

// 获取授权账号对应的OpenUid
// https://open.taobao.com/api.htm?docId=33220&docType=2&scopeId=381
type OpenUIDGetResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	OpenUIDGetResponse struct {
		OpenUID string `json:"open_uid"` //OpenUID
	} `json:"openuid_get_response"`
}

// 过订单获取对应买家的openUID
type OpenUIDGetByTradeResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	OpenUIDGetByTradeResponse struct {
		OpenUID string `json:"open_uid"` //当前交易tid对应买家的openuid
	} `json:"openuid_get_bytrade_response"`
}

// 通过订单获取对应买家的openUID
// https://open.taobao.com/api.htm?docId=33221&docType=2&scopeId=381
type OpenUIDGetByTradeParams struct {
	Tid int `json:"tid,omitempty"` //订单ID
}

// 通过MixNick转换openUID
type OpenUIDGetByMixNickResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	OpenUIDGetByMixNickResponse struct {
		OpenUID string `json:"open_uid"` //OpenUID
	} `json:"openuid_get_bymixnick_response"`
}

// 通过MixNick转换openUID
// https://open.taobao.com/api.htm?docId=33223&docType=2&scopeId=381
type OpenUIDGetByMixNickParams struct {
	MixNick string `json:"mix_nick"` //无线类应用获取到的混淆的nick
}

// 淘宝客-权益物料精选
type OptimusPromotionResp struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	TbkDgOptimusPromotionResponse struct {
		ResultList struct {
			MapData []struct {
				PromotionType    string `json:"promotion_type"`     //权益类型。1 有价券（需要购买的店铺券或单品券） 2 公开券（直接领取的店铺券或单品券） 3 妈妈券（妈妈渠道领取的店铺券或单品券） 4.品类券 （跨店可用券，可与1，2，3叠加）
				ConditionType    string `json:"condition_type"`     //优惠门槛类型： 1 满元 2 满件
				DiscountType     string `json:"discount_type"`      //优惠类型： 1 减钱 2 打折
				TotalCount       int    `json:"total_count"`        //权益信息-总量（权益初始库存量）
				RemainCount      int    `json:"remain_count"`       //权益信息-剩余库存（权益剩余库存量）
				DisplayStartTime string `json:"display_start_time"` //权益信息展示开始时间，精确到毫秒时间戳
				DisplayEndTime   string `json:"display_end_time"`   //权益信息展示结束时间，精确到毫秒时间戳
				PromotionList    struct {
					PromotionList []struct {
						EntryUsedStartTime string `json:"entry_used_start_time"` //权益开始时间，精确到毫秒时间戳
						EntryUsedEndTime   string `json:"entry_used_end_time"`   //权益结束时间，精确到毫秒时间戳
						EntryCondition     string `json:"entry_condition"`       //权益起用门槛，满X元可用，券场景为满元，精确到分，如满100元可用
						EntryDiscount      string `json:"entry_discount"`        //权益面额，券场景为减钱，精确到分
					} `json:"promotion_list"`
				} `json:"promotion_list"` //权益信息
				PromotionExtend struct {
					RecommendItemList struct {
						RecommendItemList []struct {
							ItemID int    `json:"item_id"` //权益推荐商品id
							URL    string `json:"url"`     //商品链接
						} `json:"recommend_item_list"`
					} `json:"recommend_item_list"` //权益推荐商品
					YoujiaCouponInfo struct {
						ItemID string `json:"item_id"` //有价券商品id
						URL    string `json:"url"`     //商品链接
					} `json:"youjia_coupon_info"` //有价券信息
					PromotionURL string `json:"promotion_url"` //权益链接
				} `json:"promotion_extend"` //权益扩展信息
				SellerID       string `json:"seller_id"`        //店铺信息-卖家ID
				Nick           string `json:"nick"`             //店铺信息-卖家昵称
				ShopTitle      string `json:"shop_title"`       //店铺信息-店铺名称
				ShopPictureURL string `json:"shop_picture_url"` //店铺信息-店铺logo
			} `json:"map_data"`
		} `json:"result_list"`
	} `json:"tbk_dg_optimus_promotion_response"` //推广者接口数据
	TbkScOptimusPromotionResponse struct {
		ResultList struct {
			MapData []struct {
				PromotionType    string `json:"promotion_type"`     //权益类型。1 有价券（需要购买的店铺券或单品券） 2 公开券（直接领取的店铺券或单品券） 3 妈妈券（妈妈渠道领取的店铺券或单品券） 4.品类券 （跨店可用券，可与1，2，3叠加）
				ConditionType    string `json:"condition_type"`     //优惠门槛类型： 1 满元 2 满件
				DiscountType     string `json:"discount_type"`      //优惠类型： 1 减钱 2 打折
				TotalCount       int    `json:"total_count"`        //权益信息-总量（权益初始库存量）
				RemainCount      int    `json:"remain_count"`       //权益信息-剩余库存（权益剩余库存量）
				DisplayStartTime string `json:"display_start_time"` //权益信息展示开始时间，精确到毫秒时间戳
				DisplayEndTime   string `json:"display_end_time"`   //权益信息展示结束时间，精确到毫秒时间戳
				PromotionList    struct {
					PromotionList []struct {
						EntryUsedStartTime string `json:"entry_used_start_time"` //权益开始时间，精确到毫秒时间戳
						EntryUsedEndTime   string `json:"entry_used_end_time"`   //权益结束时间，精确到毫秒时间戳
						EntryCondition     string `json:"entry_condition"`       //权益起用门槛，满X元可用，券场景为满元，精确到分，如满100元可用
						EntryDiscount      string `json:"entry_discount"`        //权益面额，券场景为减钱，精确到分
					} `json:"promotion_list"`
				} `json:"promotion_list"` //权益信息
				PromotionExtend struct {
					RecommendItemList struct {
						RecommendItemList []struct {
							ItemID int    `json:"item_id"` //权益推荐商品id
							URL    string `json:"url"`     //商品链接
						} `json:"recommend_item_list"`
					} `json:"recommend_item_list"` //权益推荐商品
					YoujiaCouponInfo struct {
						ItemID string `json:"item_id"` //有价券商品id
						URL    string `json:"url"`     //商品链接
					} `json:"youjia_coupon_info"` //有价券信息
					PromotionURL string `json:"promotion_url"` //权益链接
				} `json:"promotion_extend"` //权益扩展信息
				SellerID       string `json:"seller_id"`        //店铺信息-卖家ID
				Nick           string `json:"nick"`             //店铺信息-卖家昵称
				ShopTitle      string `json:"shop_title"`       //店铺信息-店铺名称
				ShopPictureURL string `json:"shop_picture_url"` //店铺信息-店铺logo
			} `json:"map_data"`
		} `json:"result_list"`
	} `json:"tbk_sc_optimus_promotion_response"` //服务商接口数据
}

// 淘宝客-服务商-权益物料精选
// https://open.taobao.com/api.htm?docId=52701&docType=2&scopeId=16287
// 淘宝客-推广者-权益物料精选
// https://open.taobao.com/api.htm?docId=52700&docType=2&scopeId=16518
type OptimusPromotionParams struct {
	PageSize    int    `json:"page_size,omitempty"`    //   页大小，每次请求限制10以内
	PageNum     int    `json:"page_num,omitempty"`     //	第几页，默认：1
	AdZoneID    string `json:"ad_zone_id,omitempty"`   //	来源广告位ID(pid中mm_1_2_3)中第3位
	PromotionID int    `json:"promotion_id,omitempty"` //	官方提供的权益物料Id。有价券-37104、大额店铺券-37116，更多权益物料id敬请期待！
	SiteID      string `json:"site_id,omitempty"`      //	来源广告位ID(pid中mm_1_2_3)中第2位(服务商必须传入)
}
