# go-gcs_csv-to-bigquery-sample

このリポジトリは、
GolangでGCSにおいたCSVからBigQueryにデータをinsertするサンプルです。

詳しくは以下の記事をご参照ください。

## setup
Please set credential file.  
```sh
$ export GOOGLE_APPLICATION_CREDENTIALS=./key.json
```

## exec
Please execute follows command.
```sh
# confirm variables definition
$ make def

# table create and data insert
$ make run

# bq show table
$ make show

# bq select
$ make select

# bq rm table
$ make del
```
