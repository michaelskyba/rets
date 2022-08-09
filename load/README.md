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

### load_instance
This script is a wrapper around the main ``load`` program that makes the
necessary RETS queries before loading the downloaded data into both respective
databases.

It's assumed that you have a normal init.sql setup with
``ignore/{condo,residential}`` and ``sql/{condo,residential}_records.sql``.

```
cd /path/to/rets
cd load
./load_instance <YYYY-MM-DD>
```
