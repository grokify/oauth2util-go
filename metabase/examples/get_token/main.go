package main

import (
	"fmt"
	"log"
	"os"

	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/type/stringsutil"
	"github.com/grokify/oauth2more/metabase"
)

func main() {
	loaded, err := config.LoadDotEnvSkipEmptyInfo(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(loaded)

	if len(os.Getenv(metabase.EnvMetabaseUsername)) == 0 {
		log.Fatal("E_NO_METABASE_USERNAME")
	}

	cfg := metabase.Config{
		BaseUrl:       os.Getenv(metabase.EnvMetabaseBaseUrl),
		SessionId:     os.Getenv(metabase.EnvMetabaseSessionId),
		Username:      os.Getenv(metabase.EnvMetabaseUsername),
		Password:      os.Getenv(metabase.EnvMetabasePassword),
		TlsSkipVerify: stringsutil.ToBool(os.Getenv(metabase.EnvMetabaseTlsSkipVerify))}
	fmtutil.PrintJSON(cfg)

	_, authResponse, err := metabase.NewClient(cfg)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AUTH_RESPONSE:")
	fmtutil.PrintJSON(authResponse)

	fmt.Println("DONE")
}