package content

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"net/http"

	jsontree "github.com/bmatsuo/go-jsontree"
	"github.com/labstack/echo"
)

type Event struct {
	Owner      string // The username of the owner of the repository
	Repo       string // The name of the repository
	Branch     string // The branch the event took place on
	Commit     string // The head commit hash attached to the event
	Type       string // Can be either "pull_request" or "push"
	Action     string // For Pull Requests, contains "assigned", "unassigned", "labeled", "unlabeled", "opened", "closed", "reopened", or "synchronize".
	BaseOwner  string // For Pull Requests, contains the base owner
	BaseRepo   string // For Pull Requests, contains the base repo
	BaseBranch string // For Pull Requests, contains the base branch
}

// Checks if the given ref should be ignored
func (s *Server) ignoreRef(rawRef string) bool {
	if rawRef[:10] == "refs/tags/" && !s.IgnoreTags {
		return false
	}
	return rawRef[:11] != "refs/heads/"
}

func (s *Server) webhookHandler(ctx *echo.Context) error {
	req := ctx.Request()
	defer req.Body.Close()

	eventType := ctx.Request().Header.Get("X-GitHub-Event")
	if eventType == "" {
		return ctx.JSON(http.StatusBadRequest, &ErrorResponse{Message: "400 Bad Request - Missing X-GitHub-Event Header"})
	}
	if eventType != "push" && eventType != "pull_request" {
		return ctx.JSON(http.StatusBadRequest, &ErrorResponse{Message: "400 Bad Request - Unknown Event Type " + eventType})
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	// If we have a Secret set, we should check the MAC
	if s.Secret != "" {
		sig := req.Header.Get("X-Hub-Signature")

		if sig == "" {
			return ctx.JSON(http.StatusForbidden, &ErrorResponse{Message: "403 Forbidden - Missing X-Hub-Signature required for HMAC verification"})
		}

		mac := hmac.New(sha1.New, []byte(s.Secret))
		mac.Write(body)
		expectedMAC := mac.Sum(nil)
		expectedSig := "sha1=" + hex.EncodeToString(expectedMAC)
		if !hmac.Equal([]byte(expectedSig), []byte(sig)) {
			return ctx.JSON(http.StatusForbidden, &ErrorResponse{Message: "403 Forbidden - HMAC verification failed"})
		}
	}

	request := jsontree.New()
	err = request.UnmarshalJSON(body)
	if err != nil {
		return err
	}

	// Parse the request and build the Event
	event := Event{}

	if eventType == "push" {
		rawRef, err := request.Get("ref").String()
		if err != nil {
			return err
		}
		// If the ref is not a branch, we don't care about it
		if s.ignoreRef(rawRef) || request.Get("head_commit").IsNull() {
			return ctx.JSON(http.StatusBadRequest, &ErrorResponse{Message: "400 BadRequest - Ref is not branch"})
		}

		// Fill in values
		event.Type = eventType
		event.Branch = rawRef[11:]
		event.Repo, err = request.Get("repository").Get("name").String()
		if err != nil {
			return err
		}
		event.Commit, err = request.Get("head_commit").Get("id").String()
		if err != nil {
			return err
		}
		event.Owner, err = request.Get("repository").Get("owner").Get("name").String()
		if err != nil {
			return err
		}
	} else if eventType == "pull_request" {
		event.Action, err = request.Get("action").String()
		if err != nil {
			return err
		}
		// Fill in values
		event.Type = eventType
		event.Owner, err = request.Get("pull_request").Get("head").Get("repo").Get("owner").Get("login").String()
		if err != nil {
			return err
		}
		event.Repo, err = request.Get("pull_request").Get("head").Get("repo").Get("name").String()
		if err != nil {
			return err
		}
		event.Branch, err = request.Get("pull_request").Get("head").Get("ref").String()
		if err != nil {
			return err
		}
		event.Commit, err = request.Get("pull_request").Get("head").Get("sha").String()
		if err != nil {
			return err
		}
		event.BaseOwner, err = request.Get("pull_request").Get("base").Get("repo").Get("owner").Get("login").String()
		if err != nil {
			return err
		}
		event.BaseRepo, err = request.Get("pull_request").Get("base").Get("repo").Get("name").String()
		if err != nil {
			return err
		}
		event.BaseBranch, err = request.Get("pull_request").Get("base").Get("ref").String()
		if err != nil {
			return err
		}
	} else {
		return ctx.JSON(http.StatusInternalServerError, &ErrorResponse{Message: "Unknown Event Type " + eventType})
	}

	// We've built our Event - put it into the channel and we're done
	go func() {
		s.Events <- event
	}()

	return ctx.JSON(http.StatusOK, &Response{Data: event, Type: 0})
}
