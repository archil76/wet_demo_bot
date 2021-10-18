package singleSubscription

import "fmt"

var (
	CommandHelp   = "help__subscription__singleSubscription"
	CommandGet    = "get__subscription__singleSubscription"
	CommandList   = "list__subscription__singleSubscription"
	CommandDelete = "delete__subscription__singleSubscription"
	CommandNew    = "new__subscription__singleSubscription"
	CommandEdit   = "edit__subscription__singleSubscription"

	UsageGet  = fmt.Sprintf("/%s <singleSubscriptionID>\n\nsingleSubscriptionID - id элемента (строго больше 0)", CommandGet)
	UsageList = fmt.Sprintf(
		"/%s <cursor> <limit>\n\n"+
			"cursor - курсор в базе (строго больше 0)\n"+
			"limit - кол-во элементов в странице (строго больше 0)", CommandList)
	UsageDelete = fmt.Sprintf("/%s <singleSubscriptionID>\n\nsingleSubscriptionID - id элемента (строго больше 0)", CommandDelete)
	UsageNew    = fmt.Sprintf("/%s <singleSubscriptionJSON>\n\nExample: `{\"name\": \"Batman\"}`", CommandNew)
	UsageEdit   = fmt.Sprintf(
		"/%s <singleSubscriptionJSON>\n\n"+
			"Edits by specified id\n"+
			"Example: `{\"id\": 1, \"name\": \"Batman\"}`", CommandEdit)

	ErrNotFound  = "Не найдено"
	ErrOnDelete  = "Не удалось удалить"
	ErrOnEdit    = "Не удалось изменить"
	ErrEmptyList = "Пустой список 😢. Добавьте элементы."

	SuccessDelete = "Успешно удалено!"
	SuccessNew    = "Успешно добавлено!"
	SuccessEdit   = "Успешно изменено!"

	CallbackNameList = "list"
)
