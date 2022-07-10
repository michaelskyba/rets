# rets
RETS Monitoring of Toronto Real Estate Board

## Installation
```sh
git clone https://github.com/michaelskyba/rets.git
cd rets
```

### Set up credentials
No credentials are included in this repo by default. To add them, you will
need to set up a wrapper script using ``./mgt/creds``:
```
$ ./mgt/creds
This will overwrite any previously-stored credentials.
So, press Ctrl+c if you want to cancel.
Enter your RETS username.
> {type the username}
Enter your RETS password.
> {type the password}
```
This creates ``pr``, a password wrapper around the ``sh/rets`` script.

## Usage
### 1. Login
Use ``./pr login``, which creates ``data/cookies.txt``, ``output/login.xml``,
and ``headers/login.txt``. The first two are used for further requests. This
will use the credentials you created using the ``creds`` script earlier.

Look at ``output/login.xml`` to confirm that your login was successful - i.e.
you entered the correct username and password. An unsuccessful login looks like
```xml
<RETS ReplyCode="20036" ReplyText="Login failed." >
</RETS>
```

### 2. Download metadata
Use ``./pr getmetadata``, which saves the metadata in ``output/metadata.xml``
and the ``curl`` headers in ``headers/metadata.txt``.

### 3. Search
Use ``./pr search <search_type> <class> "<query>"``.
Working examples:
- ``./pr search Property CondoProperty "(Ld=2022-07-10)"``
- ``./pr search Property ResidentialProperty "(Ml_num=N5671526)"``

The search results are saved in ``output/search.xml`` and the headers in
``headers/search.txt``.

Not working but provided query: ``"(MediaKey=),(MediaType=IMAGE)"``
