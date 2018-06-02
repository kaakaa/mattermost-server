package app

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/utils"
)

func TestSendChangeUsernameEmail(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	var oldUsername = "old_username"
	var newUsername = "new_username"
	var email = "test@example.com"
	var locale = ""
	var siteURL = "site_url"

	//Delete all the messages before check the sample email
	utils.DeleteMailBox(email)

	th.App.SendChangeUsernameEmail(oldUsername, newUsername, email, locale, siteURL)

	var resultsMailbox utils.JSONMessageHeaderInbucket
	err := utils.RetryInbucket(5, func() error {
		var err error
		resultsMailbox, err = utils.GetMailBox(email)
		return err
	})
	if err != nil {
		t.Log(err)
		t.Fatal("No email was received, maybe due load on the server. Disabling this verification")
	}
	if err == nil && len(resultsMailbox) > 0 {
		if !strings.ContainsAny(resultsMailbox[0].To[0], email) {
			t.Fatal("Wrong To recipient")
		} else {
			if resultsEmail, err := utils.GetMessageFromMailbox(email, resultsMailbox[0].ID); err == nil {
				b, err := ioutil.ReadFile("../tests/test-email-changeUsernameEmail.html")
				require.NoError(t, err)
				expected := strings.Split(string(b), "\n")			
				for i, a := range strings.Split(resultsEmail.Body.HTML, "\r\n",){
					assert.Equal(t, expected[i], a, fmt.Sprintf("Line %d is not match", i + 1))
				}
			}
		}
	}
}

func TestSendEmailChangeEmail(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	var oldEmail = "test@example.com"
	var newEmail = "new_email"
	var locale = ""
	var siteURL = "site_url"

	//Delete all the messages before check the sample email
	utils.DeleteMailBox(oldEmail)

	th.App.SendEmailChangeEmail(oldEmail, newEmail, locale, siteURL)

	var resultsMailbox utils.JSONMessageHeaderInbucket
	err := utils.RetryInbucket(5, func() error {
		var err error
		resultsMailbox, err = utils.GetMailBox(oldEmail)
		return err
	})
	if err != nil {
		t.Log(err)
		t.Fatal("No email was received, maybe due load on the server. Disabling this verification")
	}
	if err == nil && len(resultsMailbox) > 0 {
		if !strings.ContainsAny(resultsMailbox[0].To[0], oldEmail) {
			t.Fatal("Wrong To recipient")
		} else {
			if resultsEmail, err := utils.GetMessageFromMailbox(oldEmail, resultsMailbox[0].ID); err == nil {
				b, err := ioutil.ReadFile("../tests/test-email-emailChangeEmail.html")
				require.NoError(t, err)
				expected := strings.Split(string(b), "\n")			
				for i, a := range strings.Split(resultsEmail.Body.HTML, "\r\n",){
					assert.Equal(t, expected[i], a, fmt.Sprintf("Line %d is not match", i + 1))
				}
			}
		}
	}
}


func TestSendSignInChangeEmail(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	var email = "test@example.com"
	var method = "method"
	var locale = ""
	var siteURL = "site_url"

	//Delete all the messages before check the sample email
	utils.DeleteMailBox(email)

	th.App.SendSignInChangeEmail(email, method, locale, siteURL)

	var resultsMailbox utils.JSONMessageHeaderInbucket
	err := utils.RetryInbucket(5, func() error {
		var err error
		resultsMailbox, err = utils.GetMailBox(email)
		return err
	})
	if err != nil {
		t.Log(err)
		t.Fatal("No email was received, maybe due load on the server. Disabling this verification")
	}
	if err == nil && len(resultsMailbox) > 0 {
		if !strings.ContainsAny(resultsMailbox[0].To[0], email) {
			t.Fatal("Wrong To recipient")
		} else {
			if resultsEmail, err := utils.GetMessageFromMailbox(email, resultsMailbox[0].ID); err == nil {
				b, err := ioutil.ReadFile("../tests/test-email-signinChangeEmail.html")
				require.NoError(t, err)
				expected := strings.Split(string(b), "\n")			
				for i, a := range strings.Split(resultsEmail.Body.HTML, "\r\n",){
					assert.Equal(t, expected[i], a, fmt.Sprintf("Line %d is not match", i + 1))
				}
			}
		}
	}
}

func TestSendPasswordChangeEmail(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	var email = "test@example.com"
	var method = "method"
	var locale = ""
	var siteURL = "site_url"

	//Delete all the messages before check the sample email
	utils.DeleteMailBox(email)

	th.App.SendPasswordChangeEmail(email, method, locale, siteURL)

	var resultsMailbox utils.JSONMessageHeaderInbucket
	err := utils.RetryInbucket(5, func() error {
		var err error
		resultsMailbox, err = utils.GetMailBox(email)
		return err
	})
	if err != nil {
		t.Log(err)
		t.Log("No email was received, maybe due load on the server. Disabling this verification")
	}
	if err == nil && len(resultsMailbox) > 0 {
		if !strings.ContainsAny(resultsMailbox[0].To[0], email) {
			t.Fatal("Wrong To recipient")
		} else {
			if resultsEmail, err := utils.GetMessageFromMailbox(email, resultsMailbox[0].ID); err == nil {
				b, err := ioutil.ReadFile("../tests/test-email-passwordChangeEmail.html")
				require.NoError(t, err)
				expected := strings.Split(string(b), "\n")			
				for i, a := range strings.Split(resultsEmail.Body.HTML, "\r\n",){
					assert.Equal(t, expected[i], a, fmt.Sprintf("Line %d is not match", i + 1))
				}
			}
		}
	}
}

func TestSendUserAccessTokenAddedEmail(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	var email = "test@example.com"
	var locale = ""
	var siteURL = "site_url"

	//Delete all the messages before check the sample email
	utils.DeleteMailBox(email)

	th.App.SendUserAccessTokenAddedEmail(email, locale, siteURL)

	var resultsMailbox utils.JSONMessageHeaderInbucket
	err := utils.RetryInbucket(5, func() error {
		var err error
		resultsMailbox, err = utils.GetMailBox(email)
		return err
	})
	if err != nil {
		t.Log(err)
		t.Fatal("No email was received, maybe due load on the server. Disabling this verification")
	}
	if err == nil && len(resultsMailbox) > 0 {
		if !strings.ContainsAny(resultsMailbox[0].To[0], email) {
			t.Fatal("Wrong To recipient")
		} else {
			if resultsEmail, err := utils.GetMessageFromMailbox(email, resultsMailbox[0].ID); err == nil {
				b, err := ioutil.ReadFile("../tests/test-email-userAccessTokenAddedEmail.html")
				require.NoError(t, err)
				expected := strings.Split(string(b), "\n")			
				for i, a := range strings.Split(resultsEmail.Body.HTML, "\r\n",){
					assert.Equal(t, expected[i], a, fmt.Sprintf("Line %d is not match", i + 1))
				}
			}
		}
	}
}

func TestSendPasswordResetEmail(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	var email = "test@example.com"
	var token = &model.Token{Token: "sample_token"}
	var locale = ""
	var siteURL = "site_url"

	//Delete all the messages before check the sample email
	utils.DeleteMailBox(email)

	th.App.SendPasswordResetEmail(email, token, locale, siteURL)

	var resultsMailbox utils.JSONMessageHeaderInbucket
	err := utils.RetryInbucket(5, func() error {
		var err error
		resultsMailbox, err = utils.GetMailBox(email)
		return err
	})
	if err != nil {
		t.Log(err)
		t.Fatal("No email was received, maybe due load on the server. Disabling this verification")
	}
	if err == nil && len(resultsMailbox) > 0 {
		if !strings.ContainsAny(resultsMailbox[0].To[0], email) {
			t.Fatal("Wrong To recipient")
		} else {
			if resultsEmail, err := utils.GetMessageFromMailbox(email, resultsMailbox[0].ID); err == nil {
				b, err := ioutil.ReadFile("../tests/test-email-passwordResetEmail.html")
				require.NoError(t, err)
				expected := strings.Split(string(b), "\n")			
				for i, a := range strings.Split(resultsEmail.Body.HTML, "\r\n",){
					assert.Equal(t, expected[i], a, fmt.Sprintf("Line %d is not match", i + 1))
				}
			}
		}
	}
}

func TestSendMfaActivateEmail(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	var emailTo = "test@example.com"
	var activated = true
	var locale = ""
	var siteURL = "site_url"

	//Delete all the messages before check the sample email
	utils.DeleteMailBox(emailTo)

	th.App.SendMfaChangeEmail(emailTo, activated, locale, siteURL)

	var resultsMailbox utils.JSONMessageHeaderInbucket
	err := utils.RetryInbucket(5, func() error {
		var err error
		resultsMailbox, err = utils.GetMailBox(emailTo)
		return err
	})
	if err != nil {
		t.Log(err)
		t.Log("No email was received, maybe due load on the server. Disabling this verification")
	}
	if err == nil && len(resultsMailbox) > 0 {
		if !strings.ContainsAny(resultsMailbox[0].To[0], emailTo) {
			t.Fatal("Wrong To recipient")
		} else {
			if resultsEmail, err := utils.GetMessageFromMailbox(emailTo, resultsMailbox[0].ID); err == nil {
				b, err := ioutil.ReadFile("../tests/test-email-mfaActivateEmail.html")
				require.NoError(t, err)
				expected := strings.Split(string(b), "\n")			
				for i, a := range strings.Split(resultsEmail.Body.HTML, "\r\n",){
					assert.Equal(t, expected[i], a, fmt.Sprintf("Line %d is not match", i + 1))
				}
			}
		}
	}
}

func TestSendMfaDeactivateEmail(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	var emailTo = "test@example.com"
	var activated = false
	var locale = ""
	var siteURL = "site_url"

	//Delete all the messages before check the sample email
	utils.DeleteMailBox(emailTo)

	th.App.SendMfaChangeEmail(emailTo, activated, locale, siteURL)

	var resultsMailbox utils.JSONMessageHeaderInbucket
	err := utils.RetryInbucket(5, func() error {
		var err error
		resultsMailbox, err = utils.GetMailBox(emailTo)
		return err
	})
	if err != nil {
		t.Log(err)
		t.Log("No email was received, maybe due load on the server. Disabling this verification")
	}
	if err == nil && len(resultsMailbox) > 0 {
		if !strings.ContainsAny(resultsMailbox[0].To[0], emailTo) {
			t.Fatal("Wrong To recipient")
		} else {
			if resultsEmail, err := utils.GetMessageFromMailbox(emailTo, resultsMailbox[0].ID); err == nil {
				b, err := ioutil.ReadFile("../tests/test-email-mfaDeactivateEmail.html")
				require.NoError(t, err)
				expected := strings.Split(string(b), "\n")			
				for i, a := range strings.Split(resultsEmail.Body.HTML, "\r\n",){
					assert.Equal(t, expected[i], a, fmt.Sprintf("Line %d is not match", i + 1))
				}
			}
		}
	}
}
