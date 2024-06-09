import {createActionHelper} from "../../services/rxjs-helper/actions-helper";

export const StoreConversationError = createActionHelper('StoreConversationError')
export const GetConversationPrivateMessages = createActionHelper('GetConversationPrivateMessages')
export const StoreGetConversationPrivateMessages = createActionHelper('StoreGetConversationPrivateMessages')
export const SendConversationPrivateMessage = createActionHelper('SendConversationPrivateMessage')
export const StoreSendConversationPrivateMessage = createActionHelper('StoreSendConversationPrivateMessage')
export const GetNewConversationMessage = createActionHelper('NewMessage')
export const StoreGetNewConversationMessage = createActionHelper('StoreGetNewConversationMessage')
export const StoreCurrentUser = createActionHelper('StoreCurrentUser')
export const SendConversationPrivateCall = createActionHelper('SendConversationPrivateCall')
export const StoreSendConversationPrivateCall = createActionHelper('StoreSendConversationPrivateCall')
export const GetConversationPrivateCalls = createActionHelper('GetConversationPrivateCalls')
export const StoreGetConversationPrivateCalls = createActionHelper('StoreGetConversationPrivateCalls')
