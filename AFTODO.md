# aftodo


#### legend

**test**: write a test.

**impl**: implement a feature.

**cont**: continously do.

**note**: a note to self.

--

**cont**

+ golint, go test, go vet and gocyclo everytime.

**v.0.1 | 10/01/16 - 24/01/16**:

* [x] **test**: add tests first.
* [x] **impl**: implement cli.
* [x] **test**: check for index.html and required files.
* [x] **impl**: add logging.
* [x] **impl**: add control file option.
* [x] **impl**: embed dpkg- scripts to script.go.
* [x] **impl**: implement release file generator.
* [x] **impl**: generate html file, cydia icon etc.
* [x] **impl**: provide basic control file parser.
* [x] **note**: It should check for dpkg-scanpackages.

**v.0.2 24/01/16-**:

* [ ] **impl**: add a file watcher.
* [ ] **impl**: basic server. (Not file server)
* [x] **impl**: implement optional Release file signing `afto -s <repo>` -> Release.gpg
* [ ] **impl**: draft and generate docs (manpage/markdown/etc.)
* [ ] **impl**: afto should walk over generated repo. (./afto)

