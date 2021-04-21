package main

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
)

const (
	projectID = "your-project"
	dataset   = "sample_csv_dataset"
	table     = "sample_csv_table"
	GcsURI    = "gs://csv_data_sample-1/sample data.csv"
)

func main() {
	// client
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("err: %+v", err)
	}
	defer client.Close()

	// create table from gcs csv
	gcsRef := bigquery.NewGCSReference(GcsURI)
	gcsRef.SkipLeadingRows = 1
	gcsRef.Schema = bigquery.Schema{
		{Name: "id", Required: true, Type: bigquery.StringFieldType},
		{Name: "name", Required: true, Type: bigquery.StringFieldType},
		{Name: "age", Required: true, Type: bigquery.IntegerFieldType},
	}
	loader := client.Dataset(dataset).Table(table).LoaderFrom(gcsRef)
	loader.WriteDisposition = bigquery.WriteEmpty

	job, err := loader.Run(ctx)
	if err != nil {
		log.Printf("err: %+v", err)
	}

	status, err := job.Wait(ctx)
	if err != nil {
		log.Printf("err: %+v", err)
	}

	if status.Err() != nil {
		log.Printf("job completed with error: %v", status.Err())
	}
}
