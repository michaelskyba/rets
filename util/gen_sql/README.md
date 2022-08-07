# gen_sql
This script takes a (stripped) property metadata table and a property name
(specifically, a "SystemName") as input and generates a SQL column statement as
ouptut. These SQL outputs can then be put to use inside a SQL table
specification file (``../../load/sql``).

## Usage
``./gen <metadata xml file> <property name>``

For example, try the sample xml: ``./detect sample.xml Property_foo``.

## Metadata file
By "stripped", I mean removing isolating the description of your property type.
The first line would start with ``<COLUMNS>``, describing the property
descriptions (metaproperties?), and the rest of the file would be ``<DATA>``
values, filling the columns in.

Only the ``SystemName``, ``MaximumLength``, and ``DataType`` columns affect the
SQL description and are thus required.

## gen_all
This script takes the same metadata table that you would give to ``gen_sql``,
but also takes an property list file instead of a single property. Then, it
iterates through the properties given and prints the SQL for each line.

Try the example: ``./gen_all sample.xml properties.sample``.
