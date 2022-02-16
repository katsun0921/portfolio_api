package google_sheets_api

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"os"
)

type SheetClient struct {
	srv           *sheets.Service
	spreadsheetID string
}

func NewSheetClient(spreadsheetID string) (*SheetClient, error) {

	ctx := context.Background()

	fileName := "secret.json"

	_, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(fileName), option.WithScopes(sheets.SpreadsheetsScope))

	if err != nil {
		return nil, err
	}
	return &SheetClient{
		srv:           srv,
		spreadsheetID: spreadsheetID,
	}, nil
}

// https://developers.google.com/sheets/api/guides/values#reading_a_single_range
func (s *SheetClient) Get(range_ string) ([][]interface{}, error) {
	resp, err := s.srv.Spreadsheets.Values.Get(s.spreadsheetID, range_).Do()
	if err != nil {
		return nil, err
	}
	return resp.Values, nil
}

func Main() [][]interface{} {
	client, err := NewSheetClient(os.Getenv("SPREAD_SHEET_ID"))
	if err != nil {
		panic(err)
		return nil
	}
	res, err := client.Get("'Skill'!A2:D18")

	if err != nil {
		panic(err)
		return nil
	}

	return res
}
