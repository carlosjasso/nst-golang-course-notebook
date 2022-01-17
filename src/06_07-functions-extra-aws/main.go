package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	common "nst-golang-course/00-common"
	"os"
	"strings"

	// imports from https://aws.github.io/aws-sdk-go-v2/docs/sdk-utilities/s3/
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

const (
	maxContentSize int64 = 1024 * 1024 * 10 // 10 MB
	maxPixels      int   = 500
)

var (
	client *s3.Client
	bucket = os.Getenv("BUCKET")
)

func ImageHandler(w http.ResponseWriter, r http.Request) {
	ctx := r.Context()

	switch r.Method {
	case http.MethodPost:
		w.Header().Set("Accept", "application/octet-stream")
		w.Header().Set("Content-Type", "aplication/json")

		contentType := r.Header.Get("Content-Type")

		if !strings.HasPrefix(contentType, "image/") {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
		if r.ContentLength > maxContentSize {
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			return
		}

		contentTypeChunks := strings.Split(contentType, "/")
		keyExtension := contentTypeChunks[len(contentTypeChunks)-1]
		key := uuid.New().String() + "." + keyExtension

		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("!ERROR: ", err)
			return
		}

		input := &s3.PutObjectInput{
			Bucket:      aws.String(bucket),
			Key:         aws.String(key),
			Body:        bytes.NewReader(content),
			ContentType: aws.String(contentType),
		}

		putObjectOutput, err := client.PutObject(ctx, input)
		if err != nil {
			fmt.Println("!ERROR: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println("AwsPutObjectOutput: ", putObjectOutput)

		conn, err := amqp.Dial("amqp://<user>:<pass>@<host>:<port>")
		if err != nil {
			fmt.Println("!ERROR: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		channel, err := conn.Channel()
		if err != nil {
			fmt.Println("!ERROR: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		queue, err := channel.QueueDeclare("image", true, false, false, false, nil)
		if err != nil {
			fmt.Println("!ERROR: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		msg := struct {
			Key string "json:key"
		}{
			Key: key,
		}

		jsonBytes, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("!ERROR: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		publishing := amqp.Publishing{
			Body: jsonBytes,
		}

		err = channel.Publish("", queue.Name, false, false, publishing)
		if err != nil {
			fmt.Println("!ERROR: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

//creds @ s06min5

func init() {
	// executes before main
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("!FALAT ERROR: ", err)
	}
	client = s3.NewFromConfig(cfg)
}

func main() {
	common.LabelMessage("Session 06 & 07 - AWS")
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/image", nil)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("!ERROR: ", err)
	}
}
