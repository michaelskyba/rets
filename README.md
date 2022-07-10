# rets
RETS Monitoring of Toronto Real Estate Board

## Installation
``git clone https://github.com/michaelskyba/rets.git``

### Set up credentials
No credentials are included in this repo by default. To add them, you will
need to set up a wrapper script using ``mgt/creds``:
```sh
./mgt/creds
Enter your username.
> {username}
Enter your password.
> {password}
```
This creates ``pr``, a password wrapper around the ``sh/rets`` script.

## Usage
### 1. Login
Use ``./pr login``, which creates ``data/cookies.txt``, ``output/login.xml``,
and ``headers/login.txt``. The cookie is used for further requests. This will
use the credentials you created using the ``creds`` script earlier.

### 2. Download metadata
Use ``./pr getmetadata``, which saves the metadata in ``ouptut/metadata.xml``
and the ``curl`` headers in ``headers/metadata.txt``.

### 3. Search
Use ``./pr search <search_type> <class> "<query>"``.
Working examples:
- ``./pr search Property CondoProperty "(Ld=2022-07-10)"``
- ``./pr search Property ResidentialProperty "(Ml_num=N5671526)"``
The search results are saved in ``output/search.xml`` and the headers in
``headers/search.txt``.
