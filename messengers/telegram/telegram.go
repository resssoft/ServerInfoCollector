package telegram

import (
	"../../credentials"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var BotTelegram *tgbotapi.BotAPI

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

type UserPromise struct {
	userID      int64
	toKeepOne   bool
	promiceType string
}

func Initialize() bool {
	log.Printf("init tg")
	botTelegram, err := tgbotapi.NewBotAPI(credentials.TelegramToken)
	BotTelegram = botTelegram
	if err != nil {
		log.Printf("\ntgbotapi problem " + err.Error())
		return false
	}
	//log.Printf("Authorized on account %s", BotTelegram.Self.UserName)
	//log.Printf("\nAuthorized on account %s", BotTelegram.Self.UserName)
	return true
}

func Spy() {
	log.Printf("\nStart spy tg messages")

	if BotTelegram == nil {
		log.Printf("\nBotTelegram pointer is nil")
		return
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := BotTelegram.GetUpdatesChan(u)
	time.Sleep(time.Millisecond * 500)

	//c := new(UserPromise)
	//var UserPromises map[int]UserPromise
	UserPromises := make(map[int64]UserPromise)

	for update := range updates {
		if update.Message == nil && update.InlineQuery != nil {
			continue
		}
		//query := update.InlineQuery.Query

		//command := update.Message.Command()

		if update.CallbackQuery != nil {
			class := update.CallbackQuery.Data
			switch class {
			case "project_ver":
				BotTelegram.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Enter project key (example POPSCONTO): ")) // + class
				//update.CallbackQuery.Message.From.ID
				log.Printf("Input from user %v", update.CallbackQuery.Message.Chat.ID)
				//From
				localPromise := UserPromise{userID: update.CallbackQuery.Message.Chat.ID, toKeepOne: false, promiceType: "project_ver"}

				UserPromises[update.CallbackQuery.Message.Chat.ID] = localPromise
				//msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Ok")
				//BotTelegram.Send(msg)
			case "project_ver_unreleased":
				BotTelegram.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Enter project key (example POPSCONTO): "))
				log.Printf("Input from user %v", update.CallbackQuery.Message.Chat.ID)
				localPromise := UserPromise{userID: update.CallbackQuery.Message.Chat.ID, toKeepOne: false, promiceType: "project_ver_unreleased"}
				UserPromises[update.CallbackQuery.Message.Chat.ID] = localPromise
			case "help":
				BotTelegram.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "test commands: \n/start - start info \n/help - comands list  \n/vote - show keybord  \n/closevote - close keybord \n/rep - reply example \n/buttons - show buttons \n/info"))
			}

		} else {
			if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
				log.Printf("[%s, %n] %s", update.Message.Chat.UserName, update.Message.Chat.ID, update.Message.Text)
				log.Printf("%#v", UserPromises)

				if _, ok := UserPromises[update.Message.Chat.ID]; ok {
					var promisedProjectName string
					promisedProjectName = strings.ToUpper(update.Message.Text)
					log.Printf("Debug: src ["+update.Message.Text+"] and normal name ["+promisedProjectName+"]", UserPromises)
					//?status=unreleased&expand=issuesstatus,operations,remotelinks
					//startAt, maxResults, orderBy - description name releaseDate  sequence  startDate

					/// USE CASE !!!

					if UserPromises[update.Message.Chat.ID].promiceType == "project_ver" || UserPromises[update.Message.Chat.ID].promiceType == "project_ver_unreleased" {

						delete(UserPromises, update.Message.Chat.ID)

						if false {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I have no information on the "+promisedProjectName+" project, maybe it does not exist")
							BotTelegram.Send(msg)
							continue
						}

						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Releases list: ")
						BotTelegram.Send(msg)

						continue
					}

					if UserPromises[update.Message.Chat.ID].promiceType == "setUserRelationPromise" {
						message := ""
						chatID := strconv.FormatInt(update.Message.Chat.ID, 10)
						//jiraUser, result := jiraSystem.GetJiraUserByKey(update.Message.Text)
						log.Printf("chatID %#v", chatID)
						log.Printf("update.Message %#v", update.Message)
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
						BotTelegram.Send(msg)
						delete(UserPromises, update.Message.Chat.ID)
						continue
					}
				}

				switch update.Message.Text {
				case "/info":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm a softomate jira bot, i can search information about releases and tasks.")
					BotTelegram.Send(msg)

				case "/task":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "ok")
					BotTelegram.Send(msg)

				case "/help":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "commands: \n/start - start info \n/help - comands list  \n/vote - show keybord  \n/closevote - close keybord \n/rep - reply example \n/buttons - show buttons \n/releases")
					BotTelegram.Send(msg)

				case "/vote":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please, choise rating")
					msg.ReplyMarkup = numericKeyboard
					BotTelegram.Send(msg)

				case "/tasks":

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "user list: ")
					BotTelegram.Send(msg)

					//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please, choise rating")
					//msg.ReplyMarkup = numericKeyboard
					//BotTelegram.Send(msg)

				case "/closevote":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "ok")
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					BotTelegram.Send(msg)

				case "/features":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, features())
					BotTelegram.Send(msg)

				case "/setrelation":
				case "/start":
					/*
						keyboard := tgbotapi.InlineKeyboardMarkup{}

						var row1 []tgbotapi.InlineKeyboardButton
						btn := tgbotapi.NewInlineKeyboardButtonData("Get project unreleased versions", "project_ver_unreleased")
						row1 = append(row1, btn)
						keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row1)

						var row2 []tgbotapi.InlineKeyboardButton
						btn2 := tgbotapi.NewInlineKeyboardButtonData("Get project all versions", "project_ver")
						row2 = append(row2, btn2)
						keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row2)

						var row3 []tgbotapi.InlineKeyboardButton
						btn3 := tgbotapi.NewInlineKeyboardButtonData("Help", "help")
						row3 = append(row3, btn3)
						keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row3)

						msg.ReplyMarkup = keyboard
					*/
					message := ""
					if update.Message.Chat.IsPrivate() {
						chatID := strconv.FormatInt(update.Message.Chat.ID, 10)

						log.Printf("chatID %#v", chatID)
						log.Printf("update.Message %#v", update.Message)
						if false {
							message = "Hi, . You already connected from this account"
						} else {
							localPromise := UserPromise{userID: update.Message.Chat.ID, toKeepOne: false, promiceType: "setUserRelationPromise"}
							UserPromises[update.Message.Chat.ID] = localPromise
							message = "Please, enter you Jira name"
						}
					} else {
						message = "Use this command in private chat"
					}
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
					BotTelegram.Send(msg)

				case "/rep":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					msg.ReplyToMessageID = update.Message.MessageID
					BotTelegram.Send(msg)

				default:
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "This is unsupport command.")
					msg.ReplyToMessageID = update.Message.MessageID
					BotTelegram.Send(msg)
				}
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Use text only!!!")
				BotTelegram.Send(msg)
			}
		}
	}
}

func features() string {
	featureList := `
		add messages pull
		add user services relations collection
		send changes about versions to special role version manager
		set role manager from chat
	`
	return featureList
}
