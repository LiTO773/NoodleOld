# Noodle-core
A Moodle Content downloader written in Go with no frontend

The project is still at it's beginning stage so there is still much to be done.
Here is a brief description of the project's goal:
```
Noodle-core is a daemon/service that runs in the background and periodicly 
downloads new contents from a user's Moodle. There is no UI, so it uses a json 
file with all the data necessary required for the software to work (such as the
Moodle's domain, download interval, expected files to download, etc.).
```