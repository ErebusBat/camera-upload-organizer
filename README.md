# Camera Upload Organizer
This is a quick and dirty little utility I wrote to mess around with Go and solve a problem.

## The Problem
I have the dropbox app on my wife and I's (is that correct?  English is hard) phone set to auto upload all of our photos.

Dropbox puts these photos into the `/Camera Uploads` folder, which is great for backup but not ideal for sharing as Dropbox does not allow you to create a public album from this folder.

I have a public *Photostream* photo album that  I share with my family.  Given that I am not uploading any TMZ worthy pictures, by a matter of practice I want all my photos that both my wife and I take to upload into this album in a sane, sortable manner.

## Installation

~~~ sh
$ go get -v github.com/ErebusBat/camera-upload-organizer/
$ go install github.com/ErebusBat/camera-upload-organizer/app/photorg
# Should now have $GOPATH/bin/photorg
~~~

## Config
The `photorg` (Photo Organizer) tool uses a very simple [TOML](https://github.com/toml-lang/toml) config file:

~~~ toml
SourcePath = "/data/Dropbox/Camera Uploads"
DestRoot = "/data/Dropbox/Photos/Photostream"
~~~

You can specifc the name of the config file as the first argument:

    photorg /path/to/config_file.toml

If no file is specified then the tool will look for a file named `config.toml` in the current working directory.

## Building
You can build from source:

~~~ sh
> ./build [platform]
~~~

## Operation
The `photorg` tool will move and all photos from the `SourcePath` into the `DestRoot`, creating a tree like:

    .
    ├── 2014
    │   ├── 2014-01 January
    │   ├── 2014-02 February
    │   ├── 2014-03 March
    │   ├── 2014-04 April
    │   ├── 2014-05 May
    │   ├── 2014-06 June
    │   ├── 2014-07 July
    │   ├── 2014-08 August
    │   ├── 2014-09 September
    │   ├── 2014-10 October
    │   ├── 2014-11 November
    │   └── 2014-12 December
    └── 2015
        ├── 2015-01 January
        ├── 2015-02 February
        └── 2015-03 March

It will only create folders for photos with the given dates. If you didn't have any pictures in May 2014 then that folder would not be created.

The `photorg` tool will attempt to first read the photo date from the EXIF data, then from the file name (formatted in the way the iOS Dropbox app uploads it), and then from the date (lstat) of the file.

If one were so inclined they could write a custom decoder (See `photorg.RegisterDecoder`).  In fact this is how the lstat encoder is implemented (See `main.dumpDecoderInfo`).

## Crontab
I set this up on my media server to run and move my pictures, always ensuring they are in the correct date folder.

    0,15,30,45 * * * * /bin/bash -l -c 'cd /data/Dropbox/Camera\ Uploads && photorg >> ~/.photorg.log 2>&1'


## Contributing
1. [Fork it!](https://github.com/ErebusBat/camera-upload-organizer/fork)
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## But... Why Go?
Go is far more strict and perhaps verbose than Ruby; however it provides a much better deploy time story.

I had a much uglier and simple version of this tool implemented as a [maid](https://github.com/benjaminoakes/maid) script.  The problem with that is that it ran on my laptop, because that is where my ruby environment was setup.  I could have installed ruby on my media server and all that jazz... I just never got around to it.

With go, I just compile and have a static binary that I scp to my server and bob is your uncle.

Also it is fast... like lightning fast.

You could even use something like [golang-geo](https://github.com/kellydunn/golang-geo) to perform actions on where the picture was taken at.
