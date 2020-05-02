package types

// djzh module event types
const (
	// TODO: Create your event types
	// EventType<Action>    		= "action"

	// TODO: Create keys fo your events, the values will be derivided from the msg
	// AttributeKeyAddress  		= "address"

	// TODO: Some events may not have values for that reason you want to emit that something happened.
	// AttributeValueDoubleSign = "double_sign"
	EventTypeCreateArticle       = "CreateArticle"
	EventTypeCreateComment       = "CreateComment"
	EventTypeCreateReturnVisit   = "CreateReturnVisit"
	EventTypeCreateAVote         = "CreateAVote"
	EventTypeCreateCVote         = "CreateCVote"



	AttributeText                = "a_text" 	
	AttributeA_title             = "a_title"
	AttributeTag                 = "tag"
	AttributeArticle_id          = "article_id"
	AttributeTid                 = "tid"
	AttributeUid                 = "uid"
	AttributeA_timestamp         = "a_timestamp"
	AttributeReward              = "reward"
	
	AttributeValueCategory = ModuleName
)
