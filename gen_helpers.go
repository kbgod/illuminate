// THIS FILE IS AUTOGENERATED. DO NOT EDIT.
// Regen by running 'go generate' in the repo root.

package illuminate

// Get Helper method for Bot.GetBusinessConnection.
func (bc BusinessConnection) Get(b *Bot, opts *GetBusinessConnectionOpts) (*BusinessConnection, error) {
	return b.GetBusinessConnection(bc.Id, opts)
}

// Answer Helper method for Bot.AnswerCallbackQuery.
func (cq CallbackQuery) Answer(b *Bot, opts *AnswerCallbackQueryOpts) (bool, error) {
	return b.AnswerCallbackQuery(cq.Id, opts)
}

// ApproveJoinRequest Helper method for Bot.ApproveChatJoinRequest.
func (c Chat) ApproveJoinRequest(b *Bot, userId int64, opts *ApproveChatJoinRequestOpts) (bool, error) {
	return b.ApproveChatJoinRequest(c.Id, userId, opts)
}

// BanMember Helper method for Bot.BanChatMember.
func (c Chat) BanMember(b *Bot, userId int64, opts *BanChatMemberOpts) (bool, error) {
	return b.BanChatMember(c.Id, userId, opts)
}

// BanSenderChat Helper method for Bot.BanChatSenderChat.
func (c Chat) BanSenderChat(b *Bot, senderChatId int64, opts *BanChatSenderChatOpts) (bool, error) {
	return b.BanChatSenderChat(c.Id, senderChatId, opts)
}

// CreateInviteLink Helper method for Bot.CreateChatInviteLink.
func (c Chat) CreateInviteLink(b *Bot, opts *CreateChatInviteLinkOpts) (*ChatInviteLink, error) {
	return b.CreateChatInviteLink(c.Id, opts)
}

// CreateSubscriptionInviteLink Helper method for Bot.CreateChatSubscriptionInviteLink.
func (c Chat) CreateSubscriptionInviteLink(b *Bot, subscriptionPeriod int64, subscriptionPrice int64, opts *CreateChatSubscriptionInviteLinkOpts) (*ChatInviteLink, error) {
	return b.CreateChatSubscriptionInviteLink(c.Id, subscriptionPeriod, subscriptionPrice, opts)
}

// DeclineJoinRequest Helper method for Bot.DeclineChatJoinRequest.
func (c Chat) DeclineJoinRequest(b *Bot, userId int64, opts *DeclineChatJoinRequestOpts) (bool, error) {
	return b.DeclineChatJoinRequest(c.Id, userId, opts)
}

// DeletePhoto Helper method for Bot.DeleteChatPhoto.
func (c Chat) DeletePhoto(b *Bot, opts *DeleteChatPhotoOpts) (bool, error) {
	return b.DeleteChatPhoto(c.Id, opts)
}

// DeleteStickerSet Helper method for Bot.DeleteChatStickerSet.
func (c Chat) DeleteStickerSet(b *Bot, opts *DeleteChatStickerSetOpts) (bool, error) {
	return b.DeleteChatStickerSet(c.Id, opts)
}

// EditInviteLink Helper method for Bot.EditChatInviteLink.
func (c Chat) EditInviteLink(b *Bot, inviteLink string, opts *EditChatInviteLinkOpts) (*ChatInviteLink, error) {
	return b.EditChatInviteLink(c.Id, inviteLink, opts)
}

// EditSubscriptionInviteLink Helper method for Bot.EditChatSubscriptionInviteLink.
func (c Chat) EditSubscriptionInviteLink(b *Bot, inviteLink string, opts *EditChatSubscriptionInviteLinkOpts) (*ChatInviteLink, error) {
	return b.EditChatSubscriptionInviteLink(c.Id, inviteLink, opts)
}

// ExportInviteLink Helper method for Bot.ExportChatInviteLink.
func (c Chat) ExportInviteLink(b *Bot, opts *ExportChatInviteLinkOpts) (string, error) {
	return b.ExportChatInviteLink(c.Id, opts)
}

// Get Helper method for Bot.GetChat.
func (c Chat) Get(b *Bot, opts *GetChatOpts) (*ChatFullInfo, error) {
	return b.GetChat(c.Id, opts)
}

// GetAdministrators Helper method for Bot.GetChatAdministrators.
func (c Chat) GetAdministrators(b *Bot, opts *GetChatAdministratorsOpts) ([]ChatMember, error) {
	return b.GetChatAdministrators(c.Id, opts)
}

// GetMember Helper method for Bot.GetChatMember.
func (c Chat) GetMember(b *Bot, userId int64, opts *GetChatMemberOpts) (ChatMember, error) {
	return b.GetChatMember(c.Id, userId, opts)
}

// GetMemberCount Helper method for Bot.GetChatMemberCount.
func (c Chat) GetMemberCount(b *Bot, opts *GetChatMemberCountOpts) (int64, error) {
	return b.GetChatMemberCount(c.Id, opts)
}

// GetMenuButton Helper method for Bot.GetChatMenuButton.
func (c Chat) GetMenuButton(b *Bot, opts *GetChatMenuButtonOpts) (MenuButton, error) {
	if opts == nil {
		opts = &GetChatMenuButtonOpts{}
	}

	if opts.ChatId == nil {
		opts.ChatId = &c.Id
	}

	return b.GetChatMenuButton(opts)
}

// GetUserBoosts Helper method for Bot.GetUserChatBoosts.
func (c Chat) GetUserBoosts(b *Bot, userId int64, opts *GetUserChatBoostsOpts) (*UserChatBoosts, error) {
	return b.GetUserChatBoosts(c.Id, userId, opts)
}

// Leave Helper method for Bot.LeaveChat.
func (c Chat) Leave(b *Bot, opts *LeaveChatOpts) (bool, error) {
	return b.LeaveChat(c.Id, opts)
}

// PinMessage Helper method for Bot.PinChatMessage.
func (c Chat) PinMessage(b *Bot, messageId int64, opts *PinChatMessageOpts) (bool, error) {
	return b.PinChatMessage(c.Id, messageId, opts)
}

// PromoteMember Helper method for Bot.PromoteChatMember.
func (c Chat) PromoteMember(b *Bot, userId int64, opts *PromoteChatMemberOpts) (bool, error) {
	return b.PromoteChatMember(c.Id, userId, opts)
}

// RestrictMember Helper method for Bot.RestrictChatMember.
func (c Chat) RestrictMember(b *Bot, userId int64, permissions ChatPermissions, opts *RestrictChatMemberOpts) (bool, error) {
	return b.RestrictChatMember(c.Id, userId, permissions, opts)
}

// RevokeInviteLink Helper method for Bot.RevokeChatInviteLink.
func (c Chat) RevokeInviteLink(b *Bot, inviteLink string, opts *RevokeChatInviteLinkOpts) (*ChatInviteLink, error) {
	return b.RevokeChatInviteLink(c.Id, inviteLink, opts)
}

// SendAction Helper method for Bot.SendChatAction.
func (c Chat) SendAction(b *Bot, action string, opts *SendChatActionOpts) (bool, error) {
	return b.SendChatAction(c.Id, action, opts)
}

// SetAdministratorCustomTitle Helper method for Bot.SetChatAdministratorCustomTitle.
func (c Chat) SetAdministratorCustomTitle(b *Bot, userId int64, customTitle string, opts *SetChatAdministratorCustomTitleOpts) (bool, error) {
	return b.SetChatAdministratorCustomTitle(c.Id, userId, customTitle, opts)
}

// SetDescription Helper method for Bot.SetChatDescription.
func (c Chat) SetDescription(b *Bot, opts *SetChatDescriptionOpts) (bool, error) {
	return b.SetChatDescription(c.Id, opts)
}

// SetMenuButton Helper method for Bot.SetChatMenuButton.
func (c Chat) SetMenuButton(b *Bot, opts *SetChatMenuButtonOpts) (bool, error) {
	if opts == nil {
		opts = &SetChatMenuButtonOpts{}
	}

	if opts.ChatId == nil {
		opts.ChatId = &c.Id
	}

	return b.SetChatMenuButton(opts)
}

// SetPermissions Helper method for Bot.SetChatPermissions.
func (c Chat) SetPermissions(b *Bot, permissions ChatPermissions, opts *SetChatPermissionsOpts) (bool, error) {
	return b.SetChatPermissions(c.Id, permissions, opts)
}

// SetPhoto Helper method for Bot.SetChatPhoto.
func (c Chat) SetPhoto(b *Bot, photo InputFile, opts *SetChatPhotoOpts) (bool, error) {
	return b.SetChatPhoto(c.Id, photo, opts)
}

// SetStickerSet Helper method for Bot.SetChatStickerSet.
func (c Chat) SetStickerSet(b *Bot, stickerSetName string, opts *SetChatStickerSetOpts) (bool, error) {
	return b.SetChatStickerSet(c.Id, stickerSetName, opts)
}

// SetTitle Helper method for Bot.SetChatTitle.
func (c Chat) SetTitle(b *Bot, title string, opts *SetChatTitleOpts) (bool, error) {
	return b.SetChatTitle(c.Id, title, opts)
}

// UnbanMember Helper method for Bot.UnbanChatMember.
func (c Chat) UnbanMember(b *Bot, userId int64, opts *UnbanChatMemberOpts) (bool, error) {
	return b.UnbanChatMember(c.Id, userId, opts)
}

// UnbanSenderChat Helper method for Bot.UnbanChatSenderChat.
func (c Chat) UnbanSenderChat(b *Bot, senderChatId int64, opts *UnbanChatSenderChatOpts) (bool, error) {
	return b.UnbanChatSenderChat(c.Id, senderChatId, opts)
}

// UnpinAllMessages Helper method for Bot.UnpinAllChatMessages.
func (c Chat) UnpinAllMessages(b *Bot, opts *UnpinAllChatMessagesOpts) (bool, error) {
	return b.UnpinAllChatMessages(c.Id, opts)
}

// UnpinMessage Helper method for Bot.UnpinChatMessage.
func (c Chat) UnpinMessage(b *Bot, opts *UnpinChatMessageOpts) (bool, error) {
	return b.UnpinChatMessage(c.Id, opts)
}

// Send Helper method for Bot.SendGift.
func (g Gift) Send(b *Bot, userId int64, opts *SendGiftOpts) (bool, error) {
	return b.SendGift(userId, g.Id, opts)
}

// Copy Helper method for Bot.CopyMessage.
func (im InaccessibleMessage) Copy(b *Bot, chatId int64, opts *CopyMessageOpts) (*MessageId, error) {
	return b.CopyMessage(chatId, im.Chat.Id, im.MessageId, opts)
}

// Delete Helper method for Bot.DeleteMessage.
func (im InaccessibleMessage) Delete(b *Bot, opts *DeleteMessageOpts) (bool, error) {
	return b.DeleteMessage(im.Chat.Id, im.MessageId, opts)
}

// EditCaption Helper method for Bot.EditMessageCaption.
func (im InaccessibleMessage) EditCaption(b *Bot, opts *EditMessageCaptionOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &EditMessageCaptionOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = im.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = im.MessageId
	}

	return b.EditMessageCaption(opts)
}

// EditLiveLocation Helper method for Bot.EditMessageLiveLocation.
func (im InaccessibleMessage) EditLiveLocation(b *Bot, latitude float64, longitude float64, opts *EditMessageLiveLocationOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &EditMessageLiveLocationOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = im.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = im.MessageId
	}

	return b.EditMessageLiveLocation(latitude, longitude, opts)
}

// EditMedia Helper method for Bot.EditMessageMedia.
func (im InaccessibleMessage) EditMedia(b *Bot, media InputMedia, opts *EditMessageMediaOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &EditMessageMediaOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = im.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = im.MessageId
	}

	return b.EditMessageMedia(media, opts)
}

// EditReplyMarkup Helper method for Bot.EditMessageReplyMarkup.
func (im InaccessibleMessage) EditReplyMarkup(b *Bot, opts *EditMessageReplyMarkupOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &EditMessageReplyMarkupOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = im.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = im.MessageId
	}

	return b.EditMessageReplyMarkup(opts)
}

// EditText Helper method for Bot.EditMessageText.
func (im InaccessibleMessage) EditText(b *Bot, text string, opts *EditMessageTextOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &EditMessageTextOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = im.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = im.MessageId
	}

	return b.EditMessageText(text, opts)
}

// Forward Helper method for Bot.ForwardMessage.
func (im InaccessibleMessage) Forward(b *Bot, chatId int64, opts *ForwardMessageOpts) (*Message, error) {
	return b.ForwardMessage(chatId, im.Chat.Id, im.MessageId, opts)
}

// Pin Helper method for Bot.PinChatMessage.
func (im InaccessibleMessage) Pin(b *Bot, opts *PinChatMessageOpts) (bool, error) {
	return b.PinChatMessage(im.Chat.Id, im.MessageId, opts)
}

// SetReaction Helper method for Bot.SetMessageReaction.
func (im InaccessibleMessage) SetReaction(b *Bot, opts *SetMessageReactionOpts) (bool, error) {
	return b.SetMessageReaction(im.Chat.Id, im.MessageId, opts)
}

// StopLiveLocation Helper method for Bot.StopMessageLiveLocation.
func (im InaccessibleMessage) StopLiveLocation(b *Bot, opts *StopMessageLiveLocationOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &StopMessageLiveLocationOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = im.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = im.MessageId
	}

	return b.StopMessageLiveLocation(opts)
}

// Unpin Helper method for Bot.UnpinChatMessage.
func (im InaccessibleMessage) Unpin(b *Bot, opts *UnpinChatMessageOpts) (bool, error) {
	if opts == nil {
		opts = &UnpinChatMessageOpts{}
	}

	if opts.MessageId == nil {
		opts.MessageId = &im.MessageId
	}

	return b.UnpinChatMessage(im.Chat.Id, opts)
}

// Answer Helper method for Bot.AnswerInlineQuery.
func (iq InlineQuery) Answer(b *Bot, results []InlineQueryResult, opts *AnswerInlineQueryOpts) (bool, error) {
	return b.AnswerInlineQuery(iq.Id, results, opts)
}

// Copy Helper method for Bot.CopyMessage.
func (m Message) Copy(b *Bot, chatId int64, opts *CopyMessageOpts) (*MessageId, error) {
	return b.CopyMessage(chatId, m.Chat.Id, m.MessageId, opts)
}

// Delete Helper method for Bot.DeleteMessage.
func (m Message) Delete(b *Bot, opts *DeleteMessageOpts) (bool, error) {
	return b.DeleteMessage(m.Chat.Id, m.MessageId, opts)
}

// EditCaption Helper method for Bot.EditMessageCaption.
func (m Message) EditCaption(b *Bot, opts *EditMessageCaptionOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &EditMessageCaptionOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = m.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = m.MessageId
	}

	return b.EditMessageCaption(opts)
}

// EditLiveLocation Helper method for Bot.EditMessageLiveLocation.
func (m Message) EditLiveLocation(b *Bot, latitude float64, longitude float64, opts *EditMessageLiveLocationOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &EditMessageLiveLocationOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = m.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = m.MessageId
	}

	return b.EditMessageLiveLocation(latitude, longitude, opts)
}

// EditMedia Helper method for Bot.EditMessageMedia.
func (m Message) EditMedia(b *Bot, media InputMedia, opts *EditMessageMediaOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &EditMessageMediaOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = m.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = m.MessageId
	}

	return b.EditMessageMedia(media, opts)
}

// EditReplyMarkup Helper method for Bot.EditMessageReplyMarkup.
func (m Message) EditReplyMarkup(b *Bot, opts *EditMessageReplyMarkupOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &EditMessageReplyMarkupOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = m.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = m.MessageId
	}

	return b.EditMessageReplyMarkup(opts)
}

// EditText Helper method for Bot.EditMessageText.
func (m Message) EditText(b *Bot, text string, opts *EditMessageTextOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &EditMessageTextOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = m.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = m.MessageId
	}

	return b.EditMessageText(text, opts)
}

// Forward Helper method for Bot.ForwardMessage.
func (m Message) Forward(b *Bot, chatId int64, opts *ForwardMessageOpts) (*Message, error) {
	return b.ForwardMessage(chatId, m.Chat.Id, m.MessageId, opts)
}

// Pin Helper method for Bot.PinChatMessage.
func (m Message) Pin(b *Bot, opts *PinChatMessageOpts) (bool, error) {
	return b.PinChatMessage(m.Chat.Id, m.MessageId, opts)
}

// SetReaction Helper method for Bot.SetMessageReaction.
func (m Message) SetReaction(b *Bot, opts *SetMessageReactionOpts) (bool, error) {
	return b.SetMessageReaction(m.Chat.Id, m.MessageId, opts)
}

// StopLiveLocation Helper method for Bot.StopMessageLiveLocation.
func (m Message) StopLiveLocation(b *Bot, opts *StopMessageLiveLocationOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &StopMessageLiveLocationOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = m.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = m.MessageId
	}

	return b.StopMessageLiveLocation(opts)
}

// Unpin Helper method for Bot.UnpinChatMessage.
func (m Message) Unpin(b *Bot, opts *UnpinChatMessageOpts) (bool, error) {
	if opts == nil {
		opts = &UnpinChatMessageOpts{}
	}

	if opts.MessageId == nil {
		opts.MessageId = &m.MessageId
	}

	return b.UnpinChatMessage(m.Chat.Id, opts)
}

// Answer Helper method for Bot.AnswerPreCheckoutQuery.
func (pcq PreCheckoutQuery) Answer(b *Bot, ok bool, opts *AnswerPreCheckoutQueryOpts) (bool, error) {
	return b.AnswerPreCheckoutQuery(pcq.Id, ok, opts)
}

// Answer Helper method for Bot.AnswerShippingQuery.
func (sq ShippingQuery) Answer(b *Bot, ok bool, opts *AnswerShippingQueryOpts) (bool, error) {
	return b.AnswerShippingQuery(sq.Id, ok, opts)
}

// EditStarSubscription Helper method for Bot.EditUserStarSubscription.
func (u User) EditStarSubscription(b *Bot, telegramPaymentChargeId string, isCanceled bool, opts *EditUserStarSubscriptionOpts) (bool, error) {
	return b.EditUserStarSubscription(u.Id, telegramPaymentChargeId, isCanceled, opts)
}

// GetChatBoosts Helper method for Bot.GetUserChatBoosts.
func (u User) GetChatBoosts(b *Bot, chatId int64, opts *GetUserChatBoostsOpts) (*UserChatBoosts, error) {
	return b.GetUserChatBoosts(chatId, u.Id, opts)
}

// GetProfilePhotos Helper method for Bot.GetUserProfilePhotos.
func (u User) GetProfilePhotos(b *Bot, opts *GetUserProfilePhotosOpts) (*UserProfilePhotos, error) {
	return b.GetUserProfilePhotos(u.Id, opts)
}

// SetEmojiStatus Helper method for Bot.SetUserEmojiStatus.
func (u User) SetEmojiStatus(b *Bot, opts *SetUserEmojiStatusOpts) (bool, error) {
	return b.SetUserEmojiStatus(u.Id, opts)
}
