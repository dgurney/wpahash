# wpahash
Windows EulaHash calculator in Go. Useful if you want to know the hash of a public build for whatever reason.
# Information about the hash
For those unfamiliar with confidential Windows pre-release builds, here's all the info you need:
* Looks like this: [rest of build string].*c0e1af437cb0d038*, visible on the desktop watermark as well anything using ShellAbout (such as winver) on confidential builds.
* Calculated from the default value of \WPA\478C035F-04BC-48C7-B324-2462D786DAD7-5P-9 in the registry.
* Present from ~79xx Windows 8 builds until Windows 10 1511 (yes, even in public/RTM builds).
# Credits
Based on research and code from [Lucas](http://twitter.com/thebookisclosed).
