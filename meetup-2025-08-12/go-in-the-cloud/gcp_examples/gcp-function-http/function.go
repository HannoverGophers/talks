package function

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
)

type Configuration struct {
	BucketName string
}

const msgprefix = "<<HTTP execution>>: "

func init() {
	functions.HTTP("HandleRequest", HandleRequest)
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	resp, err := run(r.Context())
	if err != nil {
		log.Printf("%sERROR: %v", msgprefix, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Printf("%sERROR: %v", msgprefix, err)
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Printf("%sERROR: %v", msgprefix, err)
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	log.Printf("%sResponse sent successfully", msgprefix)
}

func run(ctx context.Context) ([]string, error) {
	cfg := config()

	log.Printf("%sBucket-Name is: %s", msgprefix, cfg.BucketName)

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("create storage client: %w", err)
	}
	defer client.Close()

	// Delimiter "/" -> „pseudo directories“ ähnlich wie bei S3
	q := &storage.Query{
		Delimiter: "/",
	}

	items := make([]string, 0)

	it := client.Bucket(cfg.BucketName).Objects(ctx, q)
	for {
		obj, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("list objects: %w", err)
		}

		if obj == nil || obj.Name == "" {
			return nil, fmt.Errorf("object name is empty (cannot process)")
		}

		log.Printf("%sLIST OBJECT: Name %q", msgprefix, obj.Name)

		items = append(items, obj.Name)
	}

	return items, nil
}

func config() *Configuration {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if errors.As(err, &notFound) {
			fmt.Println("No config file found. Using environment variables.")
		}
	}

	return &Configuration{
		BucketName: viper.GetString("BUCKET_NAME"),
	}
}
