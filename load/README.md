# load
This directory is for the loading and management of RETS records in the context
of the SQL database.

## Usage
### Adding data
``./load add <table> <input xml> <ignore file> <query date>``

The table should either be ``residential_records`` or ``condo_records``, and the
query date should be in ``YYYY-MM-DD`` format.

The ignore file should have one name per line, where each line lists a column
name to ignore when adding to the database. For example, you might want to
ignore the ``Shore_allow`` RETS field, which seems to always be empty.

Try the sample:
```sh
mysql -u rets -p rets_db --password=password
source sql/test.sql
select * from residential_records;
quit

./load add residential_records sample.xml ignore/sample 2022-02-22

mysql -u rets -p rets_db --password=password
select * from residential_records;
quit
```
