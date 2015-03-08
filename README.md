

### Crontab

    0,15,30,45 * * * * /bin/bash -l -c 'cd /data/Dropbox/Camera\ Uploads && photorg >> ~/.photorg.log 2>&1'

