PID=your-project
DATASET=sample_csv_dataset
TABLE=sample_csv_table

def:
	@echo $(PID) $(DATASET) $(TABLE)

run:
	go run main.go

show:
	bq show --format=prettyjson $(PID):$(DATASET).$(TABLE)

del:
	bq rm -f -t $(PID):$(DATASET).$(TABLE)

select:
	bq query --use_legacy_sql=false 'SELECT id, name, age FROM $(PID).$(DATASET).$(TABLE)'
