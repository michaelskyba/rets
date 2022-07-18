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

### Initial MariaDB setup
First, install your distribution's ``mariadb`` package. Then, initialize with
```sh
su # Enter root user
mariadb-install-db --user=mysql --basedir=/usr --datadir=/var/lib/mysql
systemctl start mariadb
mysql -u root -p # Just press Enter; no password
```
Now you're inside MariaDB:
```mysql
create database rets_db;
use rets_db;
create user 'rets'@'localhost' identified by 'password';
grant all privileges on rets_db.* to 'rets'@'localhost';
flush privileges;
quit
exit
```
Here, the literal string "password" will be used as the password. This is only a
problem if a malicious agent has direct access to the database, which would
never happen.

https://wiki.archlinux.org/title/MariaDB

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

### 4. Logout
Once you are finished your session, you are supposed to logout with ``./pr
logout``.

```
Once the RETS session is complete, a client application should ALWAYS use the
Logout transaction to  close the session on the server. This frees session
resources on the server. Users with too many open sessions could be labeled as a
network threat or bad behaving client and get disabled or blocked.
```

#### "The session has expired"
If you find something like
```xml
<RETS ReplyCode="20052" ReplyText="The session has expired" >
</RETS>
```
in your ``output/search.xml``, you'll have to ``./pr login`` again.

## Compiled protocol documentation
See doc/RETS.md
