# gen_fields
This script takes a (stripped) RETS search result as input and produces a
readable output, where columns and values are clearly associated. This makes it
easier to properly describe in init.sql.

## Usage
``./gen <input xml file> <output txt>``

For example, try the sample xml: ``./gen sample.xml parsed.txt``.

The input should only have two lines, where the first line contains the columns
separated by tabs, and the second line contains the values separated by tabs.
