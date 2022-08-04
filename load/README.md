# load
This directory is for the loading and management of RETS records in the context
of the SQL database.

## Usage
### Adding data
``./load add <table> <input xml> <query date>``

The table should either be ``residential_records`` or ``condo_records``, and the
query date should be in ``YYYY-MM-DD`` format.

Try the sample:
```sh
mysql -u rets -p rets_db --password=password
select * from residential_records;
quit

./load add residential_records sample.xml 2022-02-22

mysql -u rets -p rets_db --password=password
select * from residential_records;
quit
```
