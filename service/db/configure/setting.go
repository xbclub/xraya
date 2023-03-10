package configure

import (
	"github.com/xbclub/xraya/common"
	"github.com/xbclub/xraya/core/ipforward"
	"github.com/xbclub/xraya/pkg/util/log"
)

type Setting struct {
	RulePortMode                       RulePortMode    `json:"pacMode"`
	ProxyModeWhenSubscribe             ProxyMode       `json:"proxyModeWhenSubscribe"`
	GFWListAutoUpdateMode              AutoUpdateMode  `json:"pacAutoUpdateMode"`
	GFWListAutoUpdateIntervalHour      int             `json:"pacAutoUpdateIntervalHour"`
	SubscriptionAutoUpdateMode         AutoUpdateMode  `json:"subscriptionAutoUpdateMode"`
	SubscriptionAutoUpdateIntervalHour int             `json:"subscriptionAutoUpdateIntervalHour"`
	TcpFastOpen                        DefaultYesNo    `json:"tcpFastOpen"`
	MuxOn                              DefaultYesNo    `json:"muxOn"`
	Mux                                int             `json:"mux"`
	Transparent                        TransparentMode `json:"transparent"`
	IpForward                          bool            `json:"ipforward"`
	PortSharing                        bool            `json:"portSharing"`
	SpecialMode                        SpecialMode     `json:"specialMode"`
	TransparentType                    TransparentType `json:"transparentType"`
	AntiPollution                      Antipollution   `json:"antipollution"`
}

func NewSetting() (setting *Setting) {
	return &Setting{
		RulePortMode:                       WhitelistMode,
		ProxyModeWhenSubscribe:             ProxyModeDirect,
		GFWListAutoUpdateMode:              NotAutoUpdate,
		GFWListAutoUpdateIntervalHour:      0,
		SubscriptionAutoUpdateMode:         NotAutoUpdate,
		SubscriptionAutoUpdateIntervalHour: 0,
		TcpFastOpen:                        Default,
		MuxOn:                              No,
		Mux:                                8,
		Transparent:                        TransparentClose,
		IpForward:                          ipforward.IsIpForwardOn(),
		PortSharing:                        false,
		SpecialMode:                        SpecialModeNone,
		TransparentType:                    TransparentRedirect,
		AntiPollution:                      AntipollutionClosed,
	}
}

func (s *Setting) FillEmpty() {
	if err := common.FillEmpty(s, GetSettingNotNil()); err != nil {
		log.Warn("FillEmpty: %v:", err)
	}
}

type CustomPac struct {
	DefaultProxyMode RoutingDefaultProxyMode `json:"defaultProxyMode"` //??????????????????, proxy??????direct
	RoutingRules     []RoutingRule           `json:"routingRules"`
}

// v2rayTmpl.RoutingRule?????????????????????
type RoutingRule struct {
	Filename  string       `json:"filename"`  //SiteDAT?????????
	Tags      []string     `json:"tags"`      //SiteDAT???????????????
	MatchType PacMatchType `json:"matchType"` //???domain????????????ip??????
	RuleType  PacRuleType  `json:"ruleType"`  //???????????????????????????????????????????????????
}
