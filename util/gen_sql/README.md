# gen_sql
This script takes a (stripped) property metadata table and a property name
(specifically, a "SystemName") as input and generates a SQL column statement as
ouptut. These SQL outputs can then be put to use inside a SQL table
specification file (``../../load/sql``).

## Usage
``./gen <metadata xml file> <property name>``

For example, try the sample xml: ``./detect sample.xml Property_foo``.

By "stripped", I mean removing isolating the description of your property type.
The first line would start with ``<COLUMNS>``, describing the property
descriptions (metaproperties?), and the rest of the file would be ``<DATA>``
values, filling the columns in.
