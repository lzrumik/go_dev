package model

type Profile struct {
	Name          string
	Gender        string
	Age           int
	Height        string
	Weight        string
	Income        string
	Marriage      string
	Education     string
	Constellation string //星座
	Residence     string //户口
	House         string
	Car           string
}

type ZhenaiProfile struct {
	ObjectInfo struct {
		Age                      int      `json:"age"`
		AvatarPhotoID            int      `json:"avatarPhotoID"`
		AvatarPraiseCount        int      `json:"avatarPraiseCount"`
		AvatarPraised            bool     `json:"avatarPraised"`
		AvatarURL                string   `json:"avatarURL"`
		BasicInfo                []string `json:"basicInfo"`
		DetailInfo               []string `json:"detailInfo"`
		EducationString          string   `json:"educationString"`
		EmotionStatus            int      `json:"emotionStatus"`
		Gender                   int      `json:"gender"`
		GenderString             string   `json:"genderString"`
		HasIntroduce             bool     `json:"hasIntroduce"`
		HeightString             string   `json:"heightString"`
		HideVerifyModule         bool     `json:"hideVerifyModule"`
		IntroduceContent         string   `json:"introduceContent"`
		IntroducePraiseCount     int      `json:"introducePraiseCount"`
		IsActive                 bool     `json:"isActive"`
		IsFollowing              bool     `json:"isFollowing"`
		IsInBlackList            bool     `json:"isInBlackList"`
		IsStar                   bool     `json:"isStar"`
		IsZhenaiMail             bool     `json:"isZhenaiMail"`
		LastLoginTimeString      string   `json:"lastLoginTimeString"`
		LiveAudienceCount        int      `json:"liveAudienceCount"`
		LiveType                 int      `json:"liveType"`
		MarriageString           string   `json:"marriageString"`
		MemberID                 int      `json:"memberID"`
		MomentCount              int      `json:"momentCount"`
		Nickname                 string   `json:"nickname"`
		ObjectAgeString          string   `json:"objectAgeString"`
		ObjectBodyString         string   `json:"objectBodyString"`
		ObjectChildrenString     string   `json:"objectChildrenString"`
		ObjectEducationString    string   `json:"objectEducationString"`
		ObjectHeightString       string   `json:"objectHeightString"`
		ObjectInfo               []string `json:"objectInfo"`
		ObjectMarriageString     string   `json:"objectMarriageString"`
		ObjectSalaryString       string   `json:"objectSalaryString"`
		ObjectWantChildrenString string   `json:"objectWantChildrenString"`
		ObjectWorkCityString     string   `json:"objectWorkCityString"`
		Onlive                   int      `json:"onlive"`
		PhotoCount               int      `json:"photoCount"`
		Photos                   []struct {
			CreateTime  string `json:"createTime"`
			IsAvatar    bool   `json:"isAvatar"`
			PhotoID     int    `json:"photoID"`
			PhotoType   int    `json:"photoType"`
			PhotoURL    string `json:"photoURL"`
			PraiseCount int    `json:"praiseCount"`
			Praised     bool   `json:"praised"`
			Verified    bool   `json:"verified"`
		} `json:"photos"`
		PraisedIntroduce       bool   `json:"praisedIntroduce"`
		PreviewPhotoURL        string `json:"previewPhotoURL"`
		SalaryString           string `json:"salaryString"`
		ShowValidateIDCardFlag bool   `json:"showValidateIDCardFlag"`
		TotalPhotoCount        int    `json:"totalPhotoCount"`
		ValidateEducation      bool   `json:"validateEducation"`
		ValidateFace           bool   `json:"validateFace"`
		ValidateIDCard         bool   `json:"validateIDCard"`
		VideoCount             int    `json:"videoCount"`
		VideoID                int    `json:"videoID"`
		WorkCity               int    `json:"workCity"`
		WorkCityString         string `json:"workCityString"`
		WorkProvinceCityString string `json:"workProvinceCityString"`
	} `json:"objectInfo"`
	Interest []struct {
		AnswerContent              string `json:"answerContent"`
		AnswerContentDetail        string `json:"answerContentDetail"`
		AnswerContentDetailStatus  int    `json:"answerContentDetailStatus"`
		AnswerGuideWord            string `json:"answerGuideWord"`
		AnswerID                   int    `json:"answerID"`
		AnswerOrder                int    `json:"answerOrder"`
		AnswerRecordID             int    `json:"answerRecordID"`
		AnswerWriteRule            int    `json:"answerWriteRule"`
		IconURL                    string `json:"iconURL"`
		NewInterest                bool   `json:"newInterest"`
		PraiseCount                int    `json:"praiseCount"`
		QuestionGuideWord          string `json:"questionGuideWord"`
		QuestionID                 int    `json:"questionID"`
		QuestionName               string `json:"questionName"`
		QuestionType               int    `json:"questionType"`
		VerifyStatus               int    `json:"verifyStatus"`
		AnswerContentDetailRecords []struct {
			AnswerContentDetail string `json:"answerContentDetail"`
			Status              int    `json:"status"`
			TagID               int    `json:"tagId"`
		} `json:"answerContentDetailRecords,omitempty"`
	} `json:"interest"`
	MemberList []interface{} `json:"memberList"`
	SeoInfo    struct {
		Location []struct {
			CurrLocation bool   `json:"currLocation"`
			LinkContent  string `json:"linkContent"`
			LinkURL      string `json:"linkURL"`
		} `json:"location"`
		NearbyCitys []struct {
			LinkContent string `json:"linkContent"`
			LinkURL     string `json:"linkURL"`
		} `json:"nearbyCitys"`
	} `json:"seoInfo"`
	IsSearchEntryFlag bool `json:"isSearchEntryFlag"`
}
