package publ

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

const loginURL = "https://teams.microsoft.com/"
const tokenURLFragment = "oauth2/v2.0/token"

type TokenResponse struct {
	TokenType string `json:"token_type"`

	ExpiresIn             int `json:"expires_in"`
	ExtExpiresIn          int `json:"ext_expires_in"`
	RefreshTokenExpiresIn int `json:"refresh_token_expires_in"`

	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`

	ClientInfo string `json:"client_info"`
}

func DemandTokenResponse() (TokenResponse, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-popup-blocking", true),
		chromedp.Flag("app", loginURL),

		chromedp.Flag("enable-automation", false),
		//chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("disable-extensions", true),
		chromedp.WindowSize(400, 512),
	)

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelAlloc()

	ctx, cancelCtx := chromedp.NewContext(allocCtx)
	defer cancelCtx()

	tokenCh := make(chan TokenResponse, 1)
	errCh := make(chan error, 1)

	chromedp.ListenTarget(ctx, func(ev any) {
		e, ok := ev.(*network.EventResponseReceived)
		if !ok {
			return
		}
		if !strings.Contains(e.Response.URL, tokenURLFragment) {
			return
		}

		go func(requestID network.RequestID) {
			c := chromedp.FromContext(ctx)
			bodyCtx := cdp.WithExecutor(ctx, c.Target)

			body, err := network.GetResponseBody(requestID).Do(bodyCtx)
			if err != nil {
				errCh <- fmt.Errorf("get response body: %w", err)
				return
			}

			var tr TokenResponse
			if err := json.Unmarshal(body, &tr); err != nil {
				errCh <- fmt.Errorf("unmarshal token: %w", err)
				return
			}

			tokenCh <- tr
		}(e.RequestID)
	})

	if err := chromedp.Run(ctx,
		network.Enable(),
	); err != nil {
		return TokenResponse{}, fmt.Errorf("chromedp run: %w", err)
	}

	log.Println("(emulated-login) waiting for login")

	select {
	case tr := <-tokenCh:
		return tr, nil
	case err := <-errCh:
		return TokenResponse{}, err
	case <-ctx.Done():
		return TokenResponse{}, fmt.Errorf("context cancelled before token received")
	}
}
