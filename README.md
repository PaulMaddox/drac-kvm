Dell DRAC KVM Launcher
=========

Fed up of logging into the DRAC web interface just to launch a KVM session?
This simple Go program should help ease the pain.

```bash
$ drac-kvm --help
Usage of drac
  -h, --host="some.hostname.com": The DRAC host (or IP)
  -j, --javaws="/usr/bin/javaws": The path to javaws binary
  -p, --password="calvin": The DRAC password
  -u, --username="root": The DRAC username

$ drac-kvm -h 10.25.1.100 
2014/06/26 14:17:02 Launching DRAC KVM session to 10.25.1.100
```

It has been tested on Dell's 11th gen servers (eg: PowerEdge R720 etc).

It requires that you have java installed on your machine (specifically the `javaws` binary).

Installing (via Go)
----

If you already have Go configured on your system then you can just run the following to quickly install it:
```bash
$ go install github.com/paulmaddox/drac-kvm
```

Installing (pre built binaries) 
----

If you don't have Go installed already on your system, then included in this repository are some prebuilt binaries:

* [Linux 64bit](https://github.com/PaulMaddox/drac-kvm/blob/master/binaries/drac.linux_64bit?raw=true)

* [Mac OSX 64bit](https://github.com/PaulMaddox/drac-kvm/blob/master/binaries/drac.osx_64bit?raw=true)
