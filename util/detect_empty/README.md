# detect_empty
This script takes a (stripped) RETS search result as input and produces data
describing the emptiness statistics for each field. This makes it easier to
decide which ones can be safely can be ignored.

## Usage
``./detect <input xml file> <output txt>``

For example, try the sample xml: ``./detect sample.xml parsed.txt``.

The input's first line should contain the column names, and the rest should be
data values which fill (or leave blank!) those initial fields.
