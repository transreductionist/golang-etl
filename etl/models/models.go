package models

import (
    "USERetl/common/processors"
    "time"
)

type RcvdUltsysData struct {
    ID                    int       `json:"ID,omitempty"`
    UID                   int       `json:"UID,omitempty"`
    UserName              string    `json:"UserName,omitempty"`
    Honorific             string    `json:"Honorific,omitempty"`
    FirstName             string    `json:"FirstName,omitempty"`
    LastName              string    `json:"LastName,omitempty"`
    Suffix                string    `json:"Suffix,omitempty"`
    Address               string    `json:"Address,omitempty"`
    City                  string    `json:"City,omitempty"`
    State                 string    `json:"State,omitempty"`
    Zip                   int       `json:"Zip,omitempty"`
    ZipFour               int       `json:"ZipFour,omitempty"`
    Phone                 string    `json:"Phone,omitempty"`
    Fax                   string    `json:"Fax,omitempty"`
    Email                 string    `json:"Email,omitempty"`
    DateJoined            time.Time `json:"DateJoined,omitempty"`
    Password              string    `json:"Password,omitempty"`
    LastVisit             int       `json:"LastVisit,omitempty"`
    LastFax               time.Time `json:"LastFax,omitempty"`
    LastAction            time.Time `json:"LastAction,omitempty"`
    InvalidEmail          int       `json:"InvalidEmail,omitempty"`
    Alive                 int       `json:"Alive,omitempty"`
    Title                 string    `json:"Title,omitempty"`
    FaxStyle              int       `json:"FaxStyle,omitempty"`
    Triggers              int       `json:"Triggers,omitempty"`
    Origin                string    `json:"Origin,omitempty"`
    SubmitCount           int       `json:"SubmitCount,omitempty"`
    SigStatus             string    `json:"SigStatus,omitempty"`  // TODO: This is an enumeration: https://play.golang.org/p/1BvOakvbj2
    IP                    string    `json:"IP,omitempty"`
    IPDate                time.Time `json:"IPDate,omitempty"`
    FaxCount              int       `json:"FaxCount,omitempty"`
    PetitionCount         int       `json:"PetitionCount,omitempty"`
    CallCount             int       `json:"CallCount,omitempty"`
    MeetingCount          int       `json:"MeetingCount,omitempty"`
    ClickCount            int       `json:"ClickCount,omitempty"`
    OpenCount             int       `json:"OpenCountD,omitempty"`
    DonationCount         int       `json:"DonationCount,omitempty"`
    DonationPriorAmount   int       `json:"DonationPriorAmount,omitempty"`
    DonationSum           int       `json:"DonationSum,omitempty"`
    LastDonation          time.Time `json:"LastDonation,omitempty"`
    LastOptIn             time.Time `json:"LastOptIn,omitempty"`
    LastOpOut             time.Time `json:"LastOpOut,omitempty"`
    LastBounce            time.Time `json:"LastBounce,omitempty"`
    LastOpen              time.Time `json:"LastOpen,omitempty"`
    LastClick             time.Time `json:"LastClick,omitempty"`
    LastActionOpen        time.Time `json:"LastActionOpen,omitempty"`
    LastActionClick       time.Time `json:"LastActionClick,omitempty"`
    LastNewsOpen          time.Time `json:"LastNewsOpen,omitempty"`
    LastNewsClick         time.Time `json:"LastNewsClick,omitempty"`
    LastWebcastOpen       time.Time `json:"LastWebcastOpen,omitempty"`
    LastWebcastClick      time.Time `json:"LastWebcastClick,omitempty"`
    LastMeetingOpen       time.Time `json:"LastMeetingOpen,omitempty"`
    LastMeetingClick      time.Time `json:"LastMeetingClick,omitempty"`
    LastFundraisingOpen   time.Time `json:"LastFundraisingOpen,omitempty"`
    LastFundraisingClick  time.Time `json:"LastFundraisingClick,omitempty"`
    LastSpecialOpen       time.Time `json:"LastSpecialOpen,omitempty"`
    LastSpecialClick      time.Time `json:"LastSpecialClick,omitempty"`
    LastMediaOpen         time.Time `json:"LastMediaOpen,omitempty"`
    LastMediaClick        time.Time `json:"LastMediaClick,omitempty"`
}


type Transformed struct {
    SQLWriteAccts processors.SQLWriterData  `json:"account"`
    SQLWriteEmails processors.SQLWriterData `json:"email"`
    SQLWriteFlags processors.SQLWriterData  `json:"flag"`
}

type Account struct {
    ID                    int       `json:"id,omitempty"`
    EmailID               int       `json:"email_id,omitempty"`
    UserName              string    `json:"username,omitempty"`
    Password              string    `json:"password,omitempty"`
    Type                  string    `json:"type,omitempty"`
}

type Flag struct {
    ID                    int       `json:"id,omitempty"`
    Name                  string    `json:"name,omitempty"`
    Type                  string    `json:"type,omitempty"`
    AgentID               int       `json:"agent_id,omitempty"`
    Date                  string    `json:"datetime_in_utc,omitempty"`
}

type Email struct {
    ID                    int       `json:"id,omitempty"`
    AccountID             int       `json:"account_id,omitempty"`
    Email                 string    `json:"email,omitempty"`
    LastOpenInUTC         string    `json:"last_open_in_utc,omitempty"`
    LastClickInUTC        string    `json:"last_click_in_utc,omitempty"`
    LastBounceInUTC       string    `json:"last_bounce_in_utc,omitempty"`
    TotalOpens            int       `json:"total_opens,omitempty"`
    TotalClicks           int       `json:"total_clicks,omitempty"`
    TotalBounces          int       `json:"total_bounces,omitempty"`
    CreatedInUTC          string    `json:"created_in_utc,omitempty"`
    LastVerificationInUTC string    `json:"last_verification_in_utc,omitempty"`
}

type Enum struct {
    items []EnumItem
}

type EnumItem struct {
    index int
    name  string
}

func (enum Enum) Name(findIndex int) string {
    for _, item := range enum.items {
        if item.index == findIndex {
            return item.name
        }
    }
    return "ID not found"
}

func (enum Enum) Index(findName string) int {
    for idx, item := range enum.items {
        if findName == item.name {
            return idx
        }
    }
    return -1
}

func (enum Enum) Last() (int, string) {
    n := len(enum.items)
    return n - 1, enum.items[n-1].name
}

var AgentTypes = Enum{[]EnumItem{{0, "StaffMember"}, {1, "Organization"}, {1, "Automated"}}}
var AccountTypes = Enum{[]EnumItem{{0, "Basic"}, {1, "Advanced"}}}
var FlagTypes = Enum{[]EnumItem{{0, "Custom"}, {1, "System"}}}
